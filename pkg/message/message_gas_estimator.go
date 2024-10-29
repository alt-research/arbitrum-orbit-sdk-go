package message

import (
	"context"
	"encoding/hex"
	"errors"
	"math/big"

	"github.com/alt-research/arbitrum-orbit-sdk-go/pkg/bindings"
	"github.com/alt-research/arbitrum-orbit-sdk-go/pkg/bindings/rollupgen"
	"github.com/alt-research/arbitrum-orbit-sdk-go/pkg/types"
	"github.com/alt-research/arbitrum-orbit-sdk-go/pkg/utils"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

var (
	DEFAULT_GAS_PRICE_PERCENT_INCREASE      = big.NewInt(500)
	DEFAULT_SUBMISSION_FEE_PERCENT_INCREASE = big.NewInt(300)
)

type L1ToL2MessageGasEstimator struct {
	BaseChainClient  *ethclient.Client
	OrbitChainClient *ethclient.Client
}

func NewL1ToL2MessageGasEstimator(baseChainRpc, orbitChainRpc string) (*L1ToL2MessageGasEstimator, error) {
	baseChainClient, err := ethclient.Dial(baseChainRpc)
	if err != nil {
		return nil, err
	}
	orbitChainClient, err := ethclient.Dial(orbitChainRpc)
	if err != nil {
		return nil, err
	}
	return &L1ToL2MessageGasEstimator{
		BaseChainClient:  baseChainClient,
		OrbitChainClient: orbitChainClient,
	}, nil
}

// Get gas limit, gas price and submission price estimates for sending an L1->L2 message
func (e *L1ToL2MessageGasEstimator) EstimateAll(
	ctx context.Context,
	retryableData *types.RetryableData,
	l1BaseFee *big.Int,
	inbox common.Address,
	gasOverrides *types.GasOverrides,
) (uint64, *big.Int, *big.Int, *big.Int, error) {
	// estimate l2 maxFeePerGas
	maxFeePerGas, err := e.EstimateMaxFeePerGas(ctx, gasOverrides.MaxFeePerGas)
	if err != nil {
		return 0, nil, nil, nil, utils.ConcatError("EstimateMaxFeePerGas: ", err)
	}
	data, err := hex.DecodeString(retryableData.Data)
	if err != nil {
		return 0, nil, nil, nil, utils.ConcatError("DecodeString for retryableData.Data: ", err)
	}

	// estimate the l2 gas limit
	gasLimit, err := e.EstimateRetryableTicketGasLimit(
		ctx,
		common.HexToAddress(retryableData.From),
		retryableData.Deposit,
		common.HexToAddress(retryableData.To),
		retryableData.L2CallValue,
		common.HexToAddress(retryableData.From),
		common.HexToAddress(retryableData.From),
		data,
	)

	if err != nil {
		return 0, nil, nil, nil, utils.ConcatError("EstimateRetryableTicketGasLimit: ", err)
	}

	// estimate the submission fee
	dataLength := len(data)
	submissionFee, err := e.EstimateSubmissionFee(ctx, l1BaseFee, big.NewInt(int64(dataLength)), inbox, gasOverrides.MaxSubmissionFee)
	if err != nil {
		return 0, nil, nil, nil, utils.ConcatError("EstimateSubmissionFee: ", err)
	}
	deposit := new(big.Int).Mul(big.NewInt(int64(gasLimit)), maxFeePerGas)
	deposit = new(big.Int).Add(deposit, submissionFee)
	deposit = new(big.Int).Add(deposit, retryableData.L2CallValue)
	return gasLimit, submissionFee, maxFeePerGas, deposit, nil

}

// Estimate the amount of L2 gas required for putting the transaction in the L2 inbox, and executing it.
func (e *L1ToL2MessageGasEstimator) EstimateRetryableTicketGasLimit(
	ctx context.Context,
	sender common.Address,
	deposit *big.Int,
	to common.Address,
	l2CallValue *big.Int,
	excessFeeRefundAddress common.Address,
	callValueRefundAddress common.Address,
	data []byte,
) (uint64, error) {
	abi, err := bindings.NodeInterfaceMetaData.GetAbi()
	if err != nil {
		return 0, utils.ConcatError("GetAbi: ", err)
	}
	deposit = new(big.Int).Add(big.NewInt(1000000000000000000), l2CallValue)
	callData, err := abi.Pack(
		"estimateRetryableTicket",
		sender,
		deposit,
		to,
		l2CallValue,
		excessFeeRefundAddress,
		callValueRefundAddress,
		data,
	)
	if err != nil {
		return 0, utils.ConcatError("Pack: ", err)
	}
	callMsg := ethereum.CallMsg{
		To:   &types.NODE_INTERFACE_ADDRESS,
		Data: callData,
	}
	return e.OrbitChainClient.EstimateGas(context.Background(), callMsg)

}

// Return the fee, in wei, of submitting a new retryable tx with a give calldata size.
func (e *L1ToL2MessageGasEstimator) EstimateSubmissionFee(
	ctx context.Context,
	l1BaseFee *big.Int,
	callDataSize *big.Int,
	inbox common.Address,
	option *types.PercentIncrease,
) (*big.Int, error) {
	if option == nil {
		return nil, errors.New("maxFeePerGasOptions cannot be nil")
	}
	pi := &types.PercentIncrease{}
	if option.Base == nil {
		inboxInstance, err := rollupgen.NewInbox(inbox, e.BaseChainClient)
		if err != nil {
			return nil, err
		}
		submissionFee, err := inboxInstance.CalculateRetryableSubmissionFee(&bind.CallOpts{}, callDataSize, l1BaseFee)
		if err != nil {
			return nil, err
		}
		pi.Base = submissionFee
	} else {
		pi.Base = option.Base
	}
	if option.PercentIncrease != nil {
		pi.PercentIncrease = option.PercentIncrease
	} else {
		pi.PercentIncrease = DEFAULT_SUBMISSION_FEE_PERCENT_INCREASE
	}
	return e.percentIncrease(pi.Base, pi.PercentIncrease), nil
}

// Provides an estimate for the L2 maxFeePerGas, adding some margin to allow for gas price variation
func (e *L1ToL2MessageGasEstimator) EstimateMaxFeePerGas(
	ctx context.Context,
	maxFeePerGasOptions *types.PercentIncrease,
) (*big.Int, error) {
	if maxFeePerGasOptions == nil {
		return nil, errors.New("maxFeePerGasOptions cannot be nil")
	}
	pi := &types.PercentIncrease{}
	if maxFeePerGasOptions.Base != nil {
		pi.Base = maxFeePerGasOptions.Base
	} else {
		suggestedGasPrice, err := e.OrbitChainClient.SuggestGasPrice(ctx)
		if err != nil {
			return nil, err
		}
		pi.Base = suggestedGasPrice
	}
	if maxFeePerGasOptions.PercentIncrease != nil {
		pi.PercentIncrease = maxFeePerGasOptions.PercentIncrease
	} else {
		pi.PercentIncrease = DEFAULT_GAS_PRICE_PERCENT_INCREASE
	}

	return e.percentIncrease(pi.Base, pi.PercentIncrease), nil
}

func (e *L1ToL2MessageGasEstimator) percentIncrease(num, increase *big.Int) *big.Int {
	a := new(big.Int).Mul(num, increase)
	b := new(big.Int).Div(a, big.NewInt(100))
	return new(big.Int).Add(num, b)
}

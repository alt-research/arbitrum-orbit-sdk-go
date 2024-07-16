package message

import (
	"context"
	"errors"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/renlulu/arbitrum-orbit-sdk-go/pkg/bindings"
	"github.com/renlulu/arbitrum-orbit-sdk-go/pkg/types"
)

var (
	DEFAULT_GAS_PRICE_PERCENT_INCREASE      = big.NewInt(500)
	DEFAULT_SUBMISSION_FEE_PERCENT_INCREASE = big.NewInt(300)
)

type L1ToL2MessageGasEstimator struct {
	BaseChainClient  *ethclient.Client
	OrbitChainClient *ethclient.Client
}

func (e *L1ToL2MessageGasEstimator) estimateAll() {

}

// Estimate the amount of L2 gas required for putting the transaction in the L2 inbox, and executing it.
func (e *L1ToL2MessageGasEstimator) estimateRetryableTicketGasLimit() {

}

// Return the fee, in wei, of submitting a new retryable tx with a give calldata size.
func (e *L1ToL2MessageGasEstimator) estimateSubmissionFee(
	ctx context.Context,
	l1BaseFee *big.Int,
	callDataSize *big.Int,
	inbox common.Address,
	option *types.PercentIncrease,
) (*types.PercentIncrease, error) {
	if option == nil {
		return nil, errors.New("maxFeePerGasOptions cannot be nil")
	}
	pi := &types.PercentIncrease{}
	if option.Base == nil {
		inboxInstance, err := bindings.NewInbox(inbox, e.BaseChainClient)
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
	return pi, nil
}

func (e *L1ToL2MessageGasEstimator) estimateMaxFeePerGas(
	ctx context.Context,
	maxFeePerGasOptions *types.PercentIncrease,
) (*types.PercentIncrease, error) {
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
	return pi, nil
}

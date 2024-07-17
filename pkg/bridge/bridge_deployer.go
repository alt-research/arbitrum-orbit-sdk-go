package bridge

import (
	"context"
	"errors"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	ethtypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/renlulu/arbitrum-orbit-sdk-go/pkg/bindings"
	"github.com/renlulu/arbitrum-orbit-sdk-go/pkg/bindings/bridgegen"
	"github.com/renlulu/arbitrum-orbit-sdk-go/pkg/message"
	"github.com/renlulu/arbitrum-orbit-sdk-go/pkg/types"
	"github.com/renlulu/arbitrum-orbit-sdk-go/pkg/utils"
)

// 0.02 ether
var createTokenBridgeDefaultRetryableFees = big.NewInt(20000000000000000)
var zeroAddress = common.Address{}

type BridgeDeployer struct {
	BaseChainID           *big.Int
	BaseChainRpc          string
	BaseChainDeployerKey  string
	BaseChainClient       *ethclient.Client
	BridgeCreatorAddress  string
	BaseChainTransactOpts *bind.TransactOpts
	ChildChainRpc         string
	ChildChainClient      *ethclient.Client
	RollupAddress         string
	NativeToken           common.Address
}

func NewBridgeDeployer(
	privateKey string,
	baseChainRpc string,
	childChainRpc string,
	rollupAddress string,
	nativeToken common.Address,
) (*BridgeDeployer, error) {
	baseChainClient, err := ethclient.Dial(baseChainRpc)
	if err != nil {
		return nil, err
	}
	childChainClient, err := ethclient.Dial(childChainRpc)
	if err != nil {
		return nil, err
	}
	chainID, err := baseChainClient.ChainID(context.Background())
	if err != nil {
		return nil, err
	}
	bridgeCreatorAddress := types.TokenBridgeCreatorAddress[int(chainID.Int64())]
	if bridgeCreatorAddress == "" {
		return nil, errors.New("invalid parent chain id")
	}
	auth, err := utils.GetBasicAuth(baseChainClient, privateKey)
	if err != nil {
		return nil, err
	}
	return &BridgeDeployer{
		BaseChainID:           chainID,
		BaseChainRpc:          baseChainRpc,
		BaseChainClient:       baseChainClient,
		BaseChainTransactOpts: auth,
		BaseChainDeployerKey:  privateKey,
		BridgeCreatorAddress:  bridgeCreatorAddress,
		ChildChainRpc:         childChainRpc,
		ChildChainClient:      childChainClient,
		RollupAddress:         rollupAddress,
		NativeToken:           nativeToken,
	}, nil
}

// Create token bridge
// Steps:
// - create token bridge
//   - validate parent chain
//   - get inputs: maxGasForContracts, gasPrice, retryableFee etc
//   - send transaction
//
// - wait for retryables to execute
// - set weth gateway (only for eth-based chains)
// - wait for retryables to execute
// - fetch core contracts and parse them
func (b *BridgeDeployer) CreateNewTokenBridge(
	ctx context.Context,
	inbox common.Address,
	l1TokenBridgeCreatorAddress common.Address,
	l1BaseFee *big.Int,
	retryableGasOverrides *types.TransactionRequestRetryableGasOverrides,
) error {
	// 1. validate parent chain
	// 2. create token bridge get inputs
	maxGasForContracts, gasPrice, retryableFee, err := b.CreateTokenBridgeGetInputs(ctx, inbox, l1TokenBridgeCreatorAddress, l1BaseFee, retryableGasOverrides)
	if err != nil {
		return err
	}

	fmt.Println("maxGasForContracts: ", maxGasForContracts)
	fmt.Println("gasPrice: ", gasPrice)
	fmt.Println("retryableFee: ", retryableFee)
	return nil
}

func (b *BridgeDeployer) CreateTokenBridgeGetInputs(
	ctx context.Context,
	inbox common.Address,
	l1TokenBridgeCreatorAddress common.Address,
	l1BaseFee *big.Int,
	retryableGasOverrides *types.TransactionRequestRetryableGasOverrides,
) (uint64, *big.Int, *big.Int, error) {
	_, deployerAddress, err := utils.GetAddressFromPrivateKey(b.BaseChainDeployerKey)
	if err != nil {
		return 0, nil, nil, err
	}
	l1AtomicTokenBridgeFactory, err := bindings.NewL1AtomicTokenBridgeFactory(l1TokenBridgeCreatorAddress, b.BaseChainClient)
	if err != nil {
		return 0, nil, nil, err
	}
	gasPrice, err := b.ChildChainClient.SuggestGasPrice(ctx)
	if err != nil {
		return 0, nil, nil, err
	}

	// run retryable estimate for deploying L2 factory
	_, submissionFee, _, _, err := b.GetEstimateForDeployingFactory(ctx, inbox, deployerAddress.Hex(), l1BaseFee)
	if err != nil {
		return 0, nil, nil, err
	}
	maxSubmissionCostForFactoryEstimation := new(big.Int).Mul(submissionFee, big.NewInt(2))
	maxGasForFactoryEstimation, err := l1AtomicTokenBridgeFactory.GasLimitForL2FactoryDeployment(&bind.CallOpts{})
	if err != nil {
		return 0, nil, nil, err
	}

	// run retryable estimate for deploying L2 contracts
	// we do this estimate using L2 factory template on L1 because on L2 factory does not yet exist
	gasEstimateToDeployContracts, err := b.GetEstimateToDeployContracts(ctx)
	if err != nil {
		return 0, nil, nil, err
	}
	maxSubmissionCostForContractsEstimation := new(big.Int).Mul(submissionFee, big.NewInt(2))
	retryableFee := new(big.Int).Add(maxSubmissionCostForFactoryEstimation, maxSubmissionCostForContractsEstimation)
	retryableFee = new(big.Int).Add(retryableFee, new(big.Int).Mul(maxGasForFactoryEstimation, gasPrice))
	retryableFee = new(big.Int).Add(retryableFee, new(big.Int).Mul(big.NewInt(int64(gasEstimateToDeployContracts)), gasPrice))

	return gasEstimateToDeployContracts, gasPrice, retryableFee, nil

}

func (b *BridgeDeployer) GetEstimateForDeployingFactory(
	ctx context.Context,
	inbox common.Address,
	deployerAddr string,
	l1baseFee *big.Int,
) (uint64, *big.Int, *big.Int, *big.Int, error) {
	// init l1 -> l2 message gas estimator
	l1ToL2MessageGasEstimator, err := message.NewL1ToL2MessageGasEstimator(b.BaseChainRpc, b.ChildChainRpc)
	if err != nil {
		return 0, nil, nil, nil, utils.ConcatError("NewL1ToL2MessageGasEstimator: ", err)
	}
	retryableData := &types.RetryableData{
		From:                   deployerAddr,
		To:                     "0x0000000000000000000000000000000000000000",
		L2CallValue:            big.NewInt(0),
		ExcessFeeRefundAddress: deployerAddr,
		CallValueRefundAddress: deployerAddr,
		Deposit:                big.NewInt(0),
		Data:                   types.L2_BRIDGE_FACTORY_BYTECODE,
	}
	gasOverrides := &types.GasOverrides{
		GasLimit:         &types.PercentIncrease{},
		MaxSubmissionFee: &types.PercentIncrease{},
		MaxFeePerGas:     &types.PercentIncrease{},
	}

	return l1ToL2MessageGasEstimator.EstimateAll(ctx, retryableData, l1baseFee, inbox, gasOverrides)
}

func (b *BridgeDeployer) CreateERC20Bridge() {

}

func (b *BridgeDeployer) createNewTokenBridge(
	parentChainId int,
	inbox common.Address,
	rollupOwner common.Address,
	maxGasForContracts *big.Int,
	gasPriceBid *big.Int,
) (*ethtypes.Transaction, error) {
	if b.NativeToken == zeroAddress {

	} else {
		// handle allowance
	}
	bridgeCreatorTransactor, err := bridgegen.NewBridgeCreatorTransactor(common.HexToAddress(b.BridgeCreatorAddress), b.BaseChainClient)
	if err != nil {
		return nil, err
	}
	return bridgeCreatorTransactor.CreateTokenBridge(b.BaseChainTransactOpts, inbox, rollupOwner, maxGasForContracts, gasPriceBid)

}

func (b *BridgeDeployer) GetL2TokenBridgeFactoryTemplate() (common.Address, error) {
	bridgeCreator, err := bridgegen.NewBridgeCreator(common.HexToAddress(b.BridgeCreatorAddress), b.BaseChainClient)
	if err != nil {
		return common.Address{}, err
	}
	return bridgeCreator.L2TokenBridgeFactoryTemplate(&bind.CallOpts{})
}

func (b *BridgeDeployer) GetMaxGasForContracts(ctx context.Context) {

}

func (b *BridgeDeployer) GetEstimateToDeployContracts(ctx context.Context) (uint64, error) {
	bridgeCreator, err := bindings.NewL1AtomicTokenBridgeFactory(common.HexToAddress(b.BridgeCreatorAddress), b.BaseChainClient)
	if err != nil {
		return 0, err
	}
	callOpts := &bind.CallOpts{}
	routerAddress, err := bridgeCreator.L2RouterTemplate(callOpts)
	if err != nil {
		return 0, err
	}
	router, err := b.BaseChainClient.CodeAt(ctx, routerAddress, nil)
	if err != nil {
		return 0, err
	}
	standardGatewayAddress, err := bridgeCreator.L2StandardGatewayTemplate(callOpts)
	if err != nil {
		return 0, err
	}
	standardGateway, err := b.BaseChainClient.CodeAt(ctx, standardGatewayAddress, nil)
	if err != nil {
		return 0, err
	}
	customGatewayAddress, err := bridgeCreator.L2CustomGatewayTemplate(callOpts)
	if err != nil {
		return 0, err
	}
	customGateway, err := b.BaseChainClient.CodeAt(ctx, customGatewayAddress, nil)
	if err != nil {
		return 0, err
	}
	wethGatewayAddress, err := bridgeCreator.L2WethGatewayTemplate(callOpts)
	if err != nil {
		return 0, err
	}
	wethGateway, err := b.BaseChainClient.CodeAt(ctx, wethGatewayAddress, nil)
	if err != nil {
		return 0, err
	}
	aeWethAddress, err := bridgeCreator.L2WethTemplate(callOpts)
	if err != nil {
		return 0, err
	}
	aeWeth, err := b.BaseChainClient.CodeAt(ctx, aeWethAddress, nil)
	if err != nil {
		return 0, err
	}
	l1Templates, err := bridgeCreator.L1Templates(callOpts)
	if err != nil {
		return 0, err
	}
	upgradeExecutorAddress := l1Templates.UpgradeExecutor
	upgradeExecutor, err := b.BaseChainClient.CodeAt(ctx, upgradeExecutorAddress, nil)
	if err != nil {
		return 0, err
	}
	multicallAddress, err := bridgeCreator.L2MulticallTemplate(callOpts)
	if err != nil {
		return 0, err
	}
	multicall, err := b.BaseChainClient.CodeAt(ctx, multicallAddress, nil)
	if err != nil {
		return 0, err
	}
	l2RunTimeCode := bindings.L2RuntimeCode{
		Router:          router,
		StandardGateway: standardGateway,
		CustomGateway:   customGateway,
		WethGateway:     wethGateway,
		AeWeth:          aeWeth,
		UpgradeExecutor: upgradeExecutor,
		Multicall:       multicall,
	}
	abi, err := bindings.L2AtomicTokenBridgeFactoryMetaData.GetAbi()
	if err != nil {
		return 0, err
	}
	l2TokenBridgeFactoryTemplate, err := b.GetL2TokenBridgeFactoryTemplate()
	if err != nil {
		return 0, err
	}

	_, _, addr, err := utils.GenerateECDSAKeys()
	dummyAddr := common.HexToAddress(addr)
	data, err := abi.Pack(
		"deployL2Contracts",
		l2RunTimeCode,
		dummyAddr,
		dummyAddr,
		dummyAddr,
		dummyAddr,
		dummyAddr,
		dummyAddr,
		dummyAddr,
		dummyAddr,
	)
	if err != nil {
		return 0, err
	}

	callMsg := ethereum.CallMsg{
		To:   &l2TokenBridgeFactoryTemplate,
		Data: data,
	}
	gasLimit, err := b.BaseChainClient.EstimateGas(context.Background(), callMsg)
	if err != nil {
		return 0, err
	}

	fmt.Println("gas limit: ", gasLimit)
	return gasLimit, nil

}
func (b *BridgeDeployer) fetchAllowance(owner, spender common.Address) (*big.Int, error) {
	instance, err := bindings.NewERC20(b.NativeToken, b.BaseChainClient)
	if err != nil {
		return nil, err
	}
	allowance, err := instance.Allowance(&bind.CallOpts{}, owner, spender)
	if err != nil {
		return nil, err
	}
	return allowance, nil

}

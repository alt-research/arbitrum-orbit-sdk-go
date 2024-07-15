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
	chainID, err := baseChainClient.ChainID(context.Background())
	if err != nil {
		return nil, err
	}
	bridgeCreatorAddress := types.TokenBridgeCreatorAddress[int(chainID.Int64())]
	if bridgeCreatorAddress == "" {
		return nil, errors.New("invalid parent chain id")
	}
	auth, err := utils.GetBasicAuth(baseChainClient, privateKey)

	return &BridgeDeployer{
		BaseChainID:           chainID,
		BaseChainRpc:          baseChainRpc,
		BaseChainClient:       baseChainClient,
		BaseChainTransactOpts: auth,
		BaseChainDeployerKey:  privateKey,
		BridgeCreatorAddress:  bridgeCreatorAddress,
		ChildChainRpc:         childChainRpc,
		RollupAddress:         rollupAddress,
		NativeToken:           nativeToken,
	}, nil

}

func (b *BridgeDeployer) CreateERC20Bridge() {

}

// steps:
// 1. read network info
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
	bridgeCreatorTransactor, err := bindings.NewBridgeCreatorTransactor(common.HexToAddress(b.BridgeCreatorAddress), b.BaseChainClient)
	if err != nil {
		return nil, err
	}
	return bridgeCreatorTransactor.CreateTokenBridge(b.BaseChainTransactOpts, inbox, rollupOwner, maxGasForContracts, gasPriceBid)

}

func (b *BridgeDeployer) GetL2TokenBridgeFactoryTemplate() (common.Address, error) {
	bridgeCreator, err := bindings.NewBridgeCreator(common.HexToAddress(b.BridgeCreatorAddress), b.BaseChainClient)
	if err != nil {
		return common.Address{}, err
	}
	return bridgeCreator.L2TokenBridgeFactoryTemplate(&bind.CallOpts{})
}

func (b *BridgeDeployer) GetEstimateForDeployingFactory(ctx context.Context) error {
	bridgeCreator, err := bindings.NewBridgeCreator(common.HexToAddress(b.BridgeCreatorAddress), b.BaseChainClient)
	if err != nil {
		return err
	}
	callOpts := &bind.CallOpts{}
	routerAddress, err := bridgeCreator.L2RouterTemplate(callOpts)
	if err != nil {
		return err
	}
	router, err := b.BaseChainClient.CodeAt(ctx, routerAddress, nil)
	if err != nil {
		return err
	}
	standardGatewayAddress, err := bridgeCreator.L2StandardGatewayTemplate(callOpts)
	if err != nil {
		return nil
	}
	standardGateway, err := b.BaseChainClient.CodeAt(ctx, standardGatewayAddress, nil)
	customGatewayAddress, err := bridgeCreator.L2CustomGatewayTemplate(callOpts)
	if err != nil {
		return err
	}
	customGateway, err := b.BaseChainClient.CodeAt(ctx, customGatewayAddress, nil)
	if err != nil {
		return err
	}
	wethGatewayAddress, err := bridgeCreator.L2WethGatewayTemplate(callOpts)
	if err != nil {
		return err
	}
	wethGateway, err := b.BaseChainClient.CodeAt(ctx, wethGatewayAddress, nil)
	if err != nil {
		return err
	}
	aeWethAddress, err := bridgeCreator.L2WethTemplate(callOpts)
	if err != nil {
		return err
	}
	aeWeth, err := b.BaseChainClient.CodeAt(ctx, aeWethAddress, nil)
	if err != nil {
		return err
	}
	l1Templates, err := bridgeCreator.L1Templates(callOpts)
	if err != nil {
		return err
	}
	upgradeExecutorAddress := l1Templates.UpgradeExecutor
	upgradeExecutor, err := b.BaseChainClient.CodeAt(ctx, upgradeExecutorAddress, nil)
	if err != nil {
		return err
	}
	multicallAddress, err := bridgeCreator.L2MulticallTemplate(callOpts)
	if err != nil {
		return err
	}
	multicall, err := b.BaseChainClient.CodeAt(ctx, multicallAddress, nil)
	l2RunTimeCode := bindings.L2RuntimeCode{
		Router:          router,
		StandardGateway: standardGateway,
		CustomGateway:   customGateway,
		WethGateway:     wethGateway,
		AeWeth:          aeWeth,
		UpgradeExecutor: upgradeExecutor,
		Multicall:       multicall,
	}
	abi, err := bindings.L2TokenBridgeFactoryTemplateMetaData.GetAbi()
	if err != nil {
		return err
	}
	l2TokenBridgeFactoryTemplate, err := b.GetL2TokenBridgeFactoryTemplate()
	if err != nil {
		return err
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
		return err
	}

	callMsg := ethereum.CallMsg{
		To:   &l2TokenBridgeFactoryTemplate,
		Data: data,
	}
	gasLimit, err := b.BaseChainClient.EstimateGas(context.Background(), callMsg)
	if err != nil {
		return err
	}

	fmt.Println("gas limit: ", gasLimit)
	return nil

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

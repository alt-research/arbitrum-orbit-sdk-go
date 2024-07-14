package bridge

import (
	"context"
	"errors"
	"math/big"

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
	bridgeCreatror, err := bindings.NewBridgeCreator(common.HexToAddress(b.BridgeCreatorAddress), b.BaseChainClient)
	if err != nil {
		return common.Address{}, err
	}
	return bridgeCreatror.L2TokenBridgeFactoryTemplate(&bind.CallOpts{})
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

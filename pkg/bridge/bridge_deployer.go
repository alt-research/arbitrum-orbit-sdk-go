package bridge

import (
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/renlulu/arbitrum-orbit-sdk-go/pkg/bindings"
)

// 0.02 ether
var createTokenBridgeDefaultRetryableFees = big.NewInt(20000000000000000)
var zeroAddress = common.Address{}

type BridgeDeployer struct {
	BaseChainRpc         string
	BaseChainDeployerKey string
	ChildChainRpc        string
	RollupAddress        string
	NativeToken          common.Address
}

func (b *BridgeDeployer) CreateERC20Bridge() {

}

// steps:
// 1. read network info
func (b *BridgeDeployer) createNewTokenBridge() {
	if b.NativeToken == zeroAddress {
		// check allowance

	} else {

	}
}

func (b *BridgeDeployer) fetchAllowance(owner, spender common.Address) (*big.Int, error) {
	baseChainClient, err := ethclient.Dial(b.BaseChainRpc)
	if err != nil {
		return nil, err
	}
	instance, err := bindings.NewERC20(b.NativeToken, baseChainClient)
	if err != nil {
		return nil, err
	}
	allowance, err := instance.Allowance(&bind.CallOpts{}, owner, spender)
	if err != nil {
		return nil, err
	}
	return allowance, nil

}

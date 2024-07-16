package bridge

import (
	"context"
	"fmt"
	"testing"

	"github.com/ethereum/go-ethereum/common"
)

func TestGetL2TokenBridgeFactoryTemplate(t *testing.T) {
	// dummy private key
	privatekey := "afe2bdd2c6bc8a7a87f1aee195a8f5a45de1007df742c8d85608ef6c85e3fb7c"
	baseChainRpc := "https://arb1.arbitrum.io/rpc"
	chainChainRpc := ""
	rollupAddress := ""
	nativeToken := common.Address{}
	bridgeDeployer, err := NewBridgeDeployer(privatekey, baseChainRpc, chainChainRpc, rollupAddress, nativeToken)
	if err != nil {
		t.Fail()
	}
	add, err := bridgeDeployer.GetL2TokenBridgeFactoryTemplate()
	if err != nil {
		fmt.Println(err.Error())
		t.Fail()
	}
	fmt.Println(add)
}

func TestGetEstimateToDeployContracts(t *testing.T) {
	// dummy private key
	privatekey := "afe2bdd2c6bc8a7a87f1aee195a8f5a45de1007df742c8d85608ef6c85e3fb7c"
	baseChainRpc := "https://arb1.arbitrum.io/rpc"
	chainChainRpc := ""
	rollupAddress := ""
	nativeToken := common.Address{}
	bridgeDeployer, err := NewBridgeDeployer(privatekey, baseChainRpc, chainChainRpc, rollupAddress, nativeToken)
	if err != nil {
		t.Fail()
	}
	gasEstimateToDeployContracts, err := bridgeDeployer.GetEstimateToDeployContracts(context.Background())
	if err != nil {
		fmt.Println(err.Error())
		t.Fail()
	}
	fmt.Println(gasEstimateToDeployContracts)
}

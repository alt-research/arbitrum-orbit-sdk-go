package bridge

import (
	"context"
	"fmt"
	"math/big"
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

func TestGetEstimateForDeployingFactory(t *testing.T) {
	// dummy private key
	privatekey := "afe2bdd2c6bc8a7a87f1aee195a8f5a45de1007df742c8d85608ef6c85e3fb7c"
	baseChainRpc := "https://arbitrum-sepolia.blockpi.network/v1/rpc/public"
	chainChainRpc := "https://orbit-demo.alt.technology"
	rollupAddress := "0x31ec695619dAba46Bcd1A0080af5C58594BF5843"
	inboxAddress := "0xB9892e41ca8Ead0a9968d4dA5c1d08A833a07a36"
	nativeToken := common.Address{}
	bridgeDeployer, err := NewBridgeDeployer(privatekey, baseChainRpc, chainChainRpc, rollupAddress, nativeToken)
	if err != nil {
		fmt.Println(err.Error())
		t.Fail()
	}
	gasLimit, submissionFee, maxFeePerGas, deposit, err := bridgeDeployer.GetEstimateForDeployingFactory(
		context.Background(),
		common.HexToAddress(inboxAddress),
		"0x66530799037b46913e52e9e0144D15ab6ed954f5",
		big.NewInt(100000000),
	)
	if err != nil {
		fmt.Println(err.Error())
		t.Fail()
	}
	fmt.Println("l2 gas price: ", maxFeePerGas)
	fmt.Println("l2 gas limit: ", gasLimit)
	fmt.Println("submission fee: ", submissionFee)
	fmt.Println("deposit: ", deposit)
}

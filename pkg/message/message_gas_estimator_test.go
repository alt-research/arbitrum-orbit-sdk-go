package message

import (
	"context"
	"fmt"
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/common"
)

func TestEstimateRetryableTicketGasLimit(t *testing.T) {
	baseChainRpc := "https://arbitrum-sepolia.blockpi.network/v1/rpc/public"
	orbitChainRpc := "https://orbit-demo.alt.technology"
	sender := common.HexToAddress("0x66530799037b46913e52e9e0144D15ab6ed954f5")
	deposit := big.NewInt(10000000000000000)
	l2CallValue := big.NewInt(0)
	gasEstimator, err := NewL1ToL2MessageGasEstimator(baseChainRpc, orbitChainRpc)
	data := []byte{}
	if err != nil {
		fmt.Println(err.Error())
		t.Fail()
	}
	gasLimit, err := gasEstimator.EstimateRetryableTicketGasLimit(
		context.Background(),
		sender,
		deposit,
		sender,
		l2CallValue,
		sender,
		sender,
		data,
	)
	if err != nil {
		fmt.Println(err.Error())
		t.Fail()
	}
	fmt.Println(gasLimit)
}

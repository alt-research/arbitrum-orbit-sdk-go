package message

import (
	"fmt"
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/renlulu/arbitrum-orbit-sdk-go/pkg/utils"
)

func TestDecodeMessageData(t *testing.T) {
	dummyValue := big.NewInt(1234)
	dummyAddress := utils.AddressToBigInt(common.HexToAddress("0x547f58583cAdC65e39Fd241Bc247eB321EA7bD2d"))
	fmt.Println("address: ", dummyAddress)
	md := &MessageData{
		Dest:                   dummyAddress,
		L2CallValue:            dummyValue,
		MaxSubmissionCost:      dummyValue,
		ExcessFeeRefundAddress: dummyAddress,
		CallValueRefundAddress: dummyAddress,
		GasLimit:               dummyValue,
		MaxFeePerGas:           dummyValue,
		DataLength:             dummyValue,
	}
	encoded, err := typ.Encode(md)
	if err != nil {
		t.Fatal(err)
	}
	m, err := DecodeMessageData(encoded)
	if err != nil {
		t.Fatal(err)
	}

	fmt.Println(m)
}

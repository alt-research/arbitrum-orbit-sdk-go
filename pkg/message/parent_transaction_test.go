package message

import (
	"context"
	"encoding/hex"
	"fmt"
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

func TestGetParentToChildMessages(t *testing.T) {
	err := hepler()
	if err != nil {
		t.Fatal(err.Error())
	}
}

func hepler() error {
	baseChainRpc := "https://arbitrum-sepolia.blockpi.network/v1/rpc/public"
	baseChainClient, err := ethclient.Dial(baseChainRpc)
	if err != nil {
		return err
	}

	txn, _, err := baseChainClient.TransactionByHash(context.Background(), common.HexToHash("0xe2da6b5225e7194d6bcd782e1c3ebe016117e9d5e91fe282c7ad5a2393379e08"))
	if err != nil {
		return err
	}
	fmt.Println("txn: ", txn)
	receipt, err := bind.WaitMined(context.Background(), baseChainClient, txn)
	if err != nil {
		return err
	}
	fmt.Println("receipt: ", receipt.Logs)
	parentChainReceipt := NewParentTransactionReceipt(common.HexToAddress("0x66530799037b46913e52e9e0144d15ab6ed954f5"), common.HexToAddress("0x56c486d3786fa26cc61473c499a36eb9cc1fbd8e"), receipt)
	l1BaseFee := big.NewInt(100000000)
	inbox := "0x43696c385947Dc1f1F36b8D3C65589D2D70607C6"
	parentToChildMessages, err := parentChainReceipt.GetParentToChildMessages(big.NewInt(421614), txn.Value(), l1BaseFee, inbox)
	if err != nil {
		return err
	}
	for _, parentToChildMessage := range parentToChildMessages {
		fmt.Println(hex.EncodeToString(parentToChildMessage.RetryableCreationId))

	}

	return nil
}

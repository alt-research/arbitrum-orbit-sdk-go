package message

import (
	"context"
	"fmt"
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

	txn, _, err := baseChainClient.TransactionByHash(context.Background(), common.HexToHash("0xfa0cba9792ab55a3913413dde1b4f29461eb81362719b0a069593a3689b819b6"))
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
	err = parentChainReceipt.GetParentToChildMessages("0xB9892e41ca8Ead0a9968d4dA5c1d08A833a07a36")
	if err != nil {
		return err
	}
	return nil
}

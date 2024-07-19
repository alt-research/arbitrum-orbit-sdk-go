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

	txn, _, err := baseChainClient.TransactionByHash(context.Background(), common.HexToHash("0xfa0cba9792ab55a3913413dde1b4f29461eb81362719b0a069593a3689b819b6"))
	if err != nil {
		return err
	}
	receipt, err := bind.WaitMined(context.Background(), baseChainClient, txn)
	if err != nil {
		return err
	}
	fmt.Println("receipt: ", receipt.Logs)
	from := common.HexToAddress("0x66530799037b46913e52e9e0144d15ab6ed954f5")
	to := common.HexToAddress("0x56c486d3786fa26cc61473c499a36eb9cc1fbd8e")
	parentChainReceipt := NewParentTransactionReceipt(to, from, receipt)
	l1BaseFee := big.NewInt(100000000)
	inbox := "0xB9892e41ca8Ead0a9968d4dA5c1d08A833a07a36"
	parentToChildMessages, err := parentChainReceipt.GetParentToChildMessages(big.NewInt(20240328), l1BaseFee, inbox)
	if err != nil {
		return err
	}
	for _, parentToChildMessage := range parentToChildMessages {
		fmt.Println(hex.EncodeToString(parentToChildMessage.RetryableCreationId))

	}

	return nil
}

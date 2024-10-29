package main

import (
	"context"
	"fmt"
	"os"

	"github.com/alt-research/arbitrum-orbit-sdk-go/pkg/bindings"
	"github.com/alt-research/arbitrum-orbit-sdk-go/pkg/rollup"
	"github.com/alt-research/arbitrum-orbit-sdk-go/pkg/types"
	"github.com/ethereum/go-ethereum/common"
)

func main() {
	fmt.Println("start parsing rollup contracts")
	// dummy one
	privateKey := "00d7029b9124e18374a3f9b173d22b41a07c60e509a42fa64e4282378c4be171"
	l1conn := os.Getenv("L1CONN")
	rollupCreator, err := rollup.NewRollupCreator(privateKey, l1conn)
	if err != nil {
		fmt.Printf("create rollup creator failed: %s\n", err.Error())
		os.Exit(1)
	}
	receipt, err := rollupCreator.Client.TransactionReceipt(context.Background(), common.HexToHash("0x9b906154420b9bcb7cc7eaf25acebdc6380d1e3235c450094c958265dfc0a742"))
	if err != nil {
		fmt.Printf("get transaction receipt failed: %s\n", err.Error())
		os.Exit(1)
	}
	fmt.Println("receipt: ", receipt.Logs[len(receipt.Logs)-1])
	rollupCreatorAddr := types.RollupCreatorAddress[421614]
	rollupCreatorParser, err := bindings.NewRollupCreator(common.HexToAddress(rollupCreatorAddr), rollupCreator.Client)
	if err != nil {
		fmt.Printf("create rollup contract parser failed: %s\n", err.Error())
		os.Exit(1)
	}
	rollupCreatorRollupCreated, err := rollupCreatorParser.ParseRollupCreated(*receipt.Logs[len(receipt.Logs)-1])
	if err != nil {
		fmt.Printf("parse rollup contracts failed: %s\n", err.Error())
		os.Exit(1)
	}
	fmt.Println(rollupCreatorRollupCreated)

}

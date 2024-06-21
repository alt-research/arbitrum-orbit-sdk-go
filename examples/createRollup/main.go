package main

import (
	"context"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"math/big"
	"os"

	"github.com/ethereum/go-ethereum/common"
	"github.com/renlulu/arbitrum-orbit-sdk-go/pkg/rollup"
)

func main() {
	fmt.Println("start creating rollup")
	privateKey := os.Getenv("PRIVATE_KEY")
	owner := os.Getenv("OWNER")
	l1conn := os.Getenv("L1CONN")
	rollupCreator, err := rollup.NewRollupCreator(privateKey, l1conn)
	if err != nil {
		fmt.Printf("create rollup creator failed: %s\n", err.Error())
		os.Exit(1)
	}

	l2configtype := rollup.GenerateL2Config(20240621, owner)
	l2config, err := json.Marshal(l2configtype)
	if err != nil {
		fmt.Printf("create l2config failed: %s\n", err.Error())
		os.Exit(1)
	}

	// 0x8b104a2e80ac6165dc58b9048de12f301d70b02a0ab51396c22b4b4b802a16a4
	wmrb, err := hex.DecodeString("8b104a2e80ac6165dc58b9048de12f301d70b02a0ab51396c22b4b4b802a16a4")
	if err != nil {
		fmt.Printf("decode wasm module root failed: %s\n", err.Error())
		os.Exit(1)
	}
	var wasmModuleRoot [32]byte
	copy(wasmModuleRoot[:], wmrb)

	txn, err := rollupCreator.CreateRollup(
		context.Background(),
		6,
		common.HexToAddress(owner),
		big.NewInt(20240625),
		string(l2config),
		uint64(l2configtype.Arbitrum.GenesisBlockNum),
		common.HexToAddress(owner),
		wasmModuleRoot,
		common.HexToAddress("0x0B03bF93Ef8A8626E5d73DB4d9181E8c10568D7B"),
		[]common.Address{common.HexToAddress("0x7B5bCf696b6C9Ef9189fed66597A4aAC87957a08")},
		big.NewInt(125000000000000000),
	)

	if err != nil {
		fmt.Printf("invoke createRollup failed: %s\n", err.Error())
		os.Exit(1)
	}

	fmt.Println(txn.Hash().String())
	rollupContracts, err := rollupCreator.ParseRollupContracts(context.Background(), 6, txn)
	if err != nil {
		fmt.Printf("parse rollup contracts failed: %s\n", err.Error())
		os.Exit(1)
	}

	contracts, err := json.Marshal(rollupContracts)
	if err != nil {
		fmt.Printf("marshal rollup contracts failed: %s\n", err.Error())
		os.Exit(1)
	}
	fmt.Println(string(contracts))
}

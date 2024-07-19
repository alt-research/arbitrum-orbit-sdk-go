# Arbitrum Orbit SDK

A golang library and command line tool for client-side intereactions with Orbit chains.

## Overview

Arbitrum Orbit SDK simplifies the process of deploying nitro contracts, bridge setup, as well as interacting with orbit chains, offering a robust set of tools for asset bridging and cross-chain messaging.

## Installation

```bash
go get github.com/renlulu/arbitrum-orbit-sdk-go
```

## Key Features

### Rollup Creation/Nitro Contract Deployment

Rollup creation or nitro contract deployment is handled through `RollupCreator`. The creator accepts a set of parameters as rollup config, call `CreateRollup` on `RollupCreator` factory contract, parse rollup contracts' address from receipt logs.

### Token Bridge Setup

Bridge setup is handled through `BridgeDeployer`. The deployer accepts a setp of parameters including rollup address created by `RollupCreator`, simulate and estimate necessary transactions would be happening both on parent chain and child chain, then call `CreateTokenBridge` on `TokenBridgeFactory` contract. 

### Cross-Chain Messages

L1 -> L2 corss chain messages is handled by `ParentToChildMessage`, `ArbitrumSubmitRetryableTx` and `ParentTransactionReceipt`. Those structs parse logs from inbox and bridge contracts, get all necessary data for calculating retryable id, which can be used for querying redeem status on chain chains.


## Usage

Here's a basic example of using the SDK create rollup:

```go
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

	l2configtype := rollup.GenerateL2Config(20240621, owner, false)
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
		common.HexToAddress(owner),
		big.NewInt(20240621),
		421614,
		string(l2config),
		uint64(l2configtype.Arbitrum.GenesisBlockNum),
		common.HexToAddress(owner),
		wasmModuleRoot,
		common.HexToAddress("0x0B03bF93Ef8A8626E5d73DB4d9181E8c10568D7B"),
		[]common.Address{common.HexToAddress("0x7B5bCf696b6C9Ef9189fed66597A4aAC87957a08")},
		big.NewInt(125000000000000000),
		10266112,
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
```


Here's a basic example of using the SDK deploy:


```go
func main() {
	// dummy private key
	privatekey := ""
	baseChainRpc := ""
	chainChainRpc := ""
	rollupAddress := ""
	tokenBridgeCreator := ""
	inboxAddress := ""
	nativeToken := common.Address{}
	bridgeDeployer, err := NewBridgeDeployer(privatekey, baseChainRpc, chainChainRpc, rollupAddress, nativeToken)
	if err != nil {
		fmt.Println(err.Error())
		t.Fail()
	}
	txn, err := bridgeDeployer.CreateNewTokenBridge(
		context.Background(),
		common.HexToAddress(inboxAddress),
		common.HexToAddress(tokenBridgeCreator),
		big.NewInt(100000000),
		nil,
	)
	if err != nil {
		fmt.Println(err.Error())
		t.Fail()
	}
	receipt, err := bind.WaitMined(context.Background(), bridgeDeployer.BaseChainClient, txn)
	if err != nil {
		fmt.Println(err.Error())
		t.Fail()
	}
	fmt.Println(receipt.TxHash)
}
```

Here's a basic example of using the SDK calculate retryable id:

```go
func hepler() error {
	baseChainRpc := ""
	transactionHash := ""
	from := common.HexToAddress("")
	to := common.HexToAddress("")
	inbox := ""
	l1BaseFee := big.NewInt(100000000)
	baseChainClient, err := ethclient.Dial(baseChainRpc)
	if err != nil {
		return err
	}

	txn, _, err := baseChainClient.TransactionByHash(context.Background(), common.HexToHash(transactionHash))
	if err != nil {
		return err
	}
	receipt, err := bind.WaitMined(context.Background(), baseChainClient, txn)
	if err != nil {
		return err
	}

	parentChainReceipt := NewParentTransactionReceipt(to, from, receipt)
	parentToChildMessages, err := parentChainReceipt.GetParentToChildMessages(big.NewInt(20240328), l1BaseFee, inbox)
	if err != nil {
		return err
	}
	for _, parentToChildMessage := range parentToChildMessages {
		fmt.Println(hex.EncodeToString(parentToChildMessage.RetryableCreationId))
	}
	return nil
}
```
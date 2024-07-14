package bindings

//go:generate abigen --abi RollupCreator.abi --pkg bindings --type RollupCreator --out RollupCreator.go
//go:generate abigen --abi BridgeCreator.abi --pkg bindings --type BridgeCreator --out BridgeCreator.go
//go:generate abigen --abi UpgradeExecutor.abi --pkg bindings --type UpgradeExecutor --out UpgradeExecutor.go
//go:generate abigen --abi SequencerInbox.abi --pkg bindings --type SequencerInbox --out SequencerInbox.go
//go:generate abigen --abi ERC20.abi --pkg bindings --type ERC20 --out ERC20.go

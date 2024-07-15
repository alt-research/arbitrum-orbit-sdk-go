package bindings

//go:generate abigen --abi rollup_creator.abi --pkg bindings --type RollupCreator --out rollup_creator.go
//go:generate abigen --abi bridge_creator.abi --pkg bindings --type BridgeCreator --out bidge_creator.go
//go:generate abigen --abi upgrade_executor.abi --pkg bindings --type UpgradeExecutor --out upgrade_executor.go
//go:generate abigen --abi sequencer_inbox.abi --pkg bindings --type SequencerInbox --out sequencer_inbox.go
//go:generate abigen --abi erc20.abi --pkg bindings --type ERC20 --out erc20.go
//go:generate abigen --abi l2_token_bridge_factory_template.abi --pkg bindings --type L2TokenBridgeFactoryTemplate --out l2_token_bridge_factory_template.go

package bindings

//go:generate abigen --abi rollup_creator.abi --pkg bindings --type RollupCreator --out rollup_creator.go
//go:generate abigen --abi bridge_creator.abi --pkg bindings --type BridgeCreator --out bridge_creator.go
//go:generate abigen --abi upgrade_executor.abi --pkg bindings --type UpgradeExecutor --out upgrade_executor.go
//go:generate abigen --abi sequencer_inbox.abi --pkg bindings --type SequencerInbox --out sequencer_inbox.go
//go:generate abigen --abi inbox.abi --pkg bindings --type Inbox --out inbox.go
//go:generate abigen --abi erc20.abi --pkg bindings --type ERC20 --out erc20.go
//go:generate abigen --abi l2_token_bridge_factory_template.abi --pkg bindings --type L2TokenBridgeFactoryTemplate --out l2_token_bridge_factory_template.go
//go:generate abigen --abi l2_atomic_token_bridge_factory.abi --pkg bindings --type L2AtomicTokenBridgeFactory --out l2_atomic_token_bridge_factory.go
//go:generate abigen --abi node_interface.abi --pkg bindings --type NodeInterface --out node_interface.go

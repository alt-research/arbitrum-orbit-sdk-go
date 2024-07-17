package bindings

//go:generate abigen --abi rollup_creator.abi --pkg rollupgen --type RollupCreator --out rollupgen/rollup_creator.go
//go:generate abigen --abi inbox.abi --pkg rollupgen --type Inbox --out rollupgen/inbox.go
//go:generate abigen --abi sequencer_inbox.abi --pkg sequencerinboxgen --type SequencerInbox --out sequencerinboxgen/sequencer_inbox.go
//go:generate abigen --abi bridge_creator.abi --pkg bridgegen --type BridgeCreator --out bridgegen/bridge_creator.go
//go:generate abigen --abi upgrade_executor.abi --pkg upgradegen --type UpgradeExecutor --out upgradegen/upgrade_executor.go

//go:generate abigen --abi erc20.abi --pkg bindings --type ERC20 --out erc20.go
//go:generate abigen --abi l2_atomic_token_bridge_factory.abi --pkg bindings --type L2AtomicTokenBridgeFactory --out l2_atomic_token_bridge_factory.go
//go:generate abigen --abi l1_atomic_token_bridge_factory.abi --pkg bindings --type L1AtomicTokenBridgeFactory --out l1_atomic_token_bridge_factory.go
//go:generate abigen --abi node_interface.abi --pkg bindings --type NodeInterface --out node_interface.go

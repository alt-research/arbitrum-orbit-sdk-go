# Arbitrum Orbit SDK (in golang)

## Installation

```bash
go get github.com/renlulu/arbitrum-orbit-sdk-go
```

## Convention

1. Rollup Creator

```go
var RollupCreatorAddress = map[int]string{
	1:        "0x90d68b056c411015eae3ec0b98ad94e2c91419f1",
	42161:    "0x9CAd81628aB7D8e239F1A5B497313341578c5F71",
	42170:    "0x9CAd81628aB7D8e239F1A5B497313341578c5F71",
	8453:     "0x850F050C65B34966895AdA26a4D06923901916DB",
	11155111: "0xfbd0b034e6305788007f6e0123cc5eae701a5751",
	17000:    "0xB512078282F462Ba104231ad856464Ceb0a7747e",
	421614:   "0x06E341073b2749e0Bb9912461351f716DeCDa9b0",
	84532:    "0x1E0921818df948c338380e722C8aE91Bb285763C",
}
```

2. Bridge Creator

```go
var TokenBridgeCreatorAddress = map[int]string{
	1:        "0x60D9A46F24D5a35b95A78Dd3E793e55D94EE0660",
	42161:    "0x2f5624dc8800dfA0A82AC03509Ef8bb8E7Ac000e",
	42170:    "0x8B9D9490a68B1F16ac8A21DdAE5Fd7aB9d708c14",
	8453:     "0x4C240987d6fE4fa8C7a0004986e3db563150CA55",
	11155111: "0x7edb2dfBeEf9417e0454A80c51EE0C034e45a570",
	17000:    "0xac890ED9bC2494C053cE701F138958df95966d94",
	421614:   "0x38F35Af53bF913c439eaB06A367e09D6eb253492",
	84532:    "0x4C240987d6fE4fa8C7a0004986e3db563150CA55",
}
```

## Functions

### Rollup Creation

1. Create `RollupCreator`

```go
func NewRollupCreator(privateKey string, l1conn string) (*RollupCreator, error)
```

2. Create Rollup using $ETH as gas token

```go
func (r *RollupCreator) CreateRollup(
	ctx context.Context,
    chainIndex int,
	owner common.Address,
	chainId *big.Int,
	chainConfig string,
	genesisBlockNum uint64,
	loserStakeEscrow common.Address,
	wasmModuleRoot [32]byte,
	batchPoster common.Address,
	validators []common.Address,

) (*ethtypes.Transaction, error)
```

3. Create Rollup using custom gas token

4. Parse rollup contracts from transaction receipt

```go
func (r *RollupCreator) ParseRollupContracts(ctx context.Context, txn *ethtypes.Transaction) (*bindings.RollupCreatorRollupCreated, error)
```

## Bridge Setup

1. GetL2TokenBridgeFactoryTemplate
2. GetEstimateForDeployingFactory

## Interact with Rollup contracts (on parent chain)

```go
const (
	NODE_INTERFACE_ADDRESS    = "0x00000000000000000000000000000000000000C8"
	ARB_SYS_ADDRESS           = "0x0000000000000000000000000000000000000064"
	ARB_ADDRESS_TABLE_ADDRESS = "0x0000000000000000000000000000000000000066"
	ARB_OWNER_PUBLIC          = "0x000000000000000000000000000000000000006B"
	ARB_GAS_INFO              = "0x000000000000000000000000000000000000006C"
	ARB_STATISTICS            = "0x000000000000000000000000000000000000006F"
)
```

1. Whitelist DA Key

## Interact with Arbos (on orbit chain)

## Cross Chain Transactions/Message

1. L1ToL2MessageGasEstimator

## Utils

1. Wait transaction to pass the safe block
2. Generate BLS Keys
3. Prepare DA Keyset for whitelisting

```go
func WaitTx(ctx context.Context, client *ethclient.Client, txn *ethtypes.Transaction, waitForSafePoll bool) error
```

## Calldata generation

## Cli
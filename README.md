# Arbitrum Orbit SDK (in golang)

## Installation

```bash
go get github.com/renlulu/arbitrum-orbit-sdk-go
```

## Convention

1. chain index

```
0 => ethereum
1 => arbitrumOne
2 => arbitrumNova
3 => base
4 => sepolia
5 => holesky
6 => arbitrumSepolia
7 => baseSepolia
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

## Interact with Rollup contracts (on parent chain)

## Interact  with Arbos (on orbit chain)

## Cross Chain Transactions

## Utils

1. Wait transaction to pass the safe block

```go
func WaitTx(ctx context.Context, client *ethclient.Client, txn *ethtypes.Transaction, waitForSafePoll bool) error
```

## Calldata generation

## Cli
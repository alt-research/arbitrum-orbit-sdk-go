package rollup

import (
	"context"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	ethtypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/renlulu/arbitrum-orbit-sdk-go/pkg/bindings"
	"github.com/renlulu/arbitrum-orbit-sdk-go/pkg/types"
	"github.com/renlulu/arbitrum-orbit-sdk-go/pkg/utils"
)

type RollupCreator struct {
	RPC        string
	PrivateKey string
	Client     *ethclient.Client
	opts       *bind.TransactOpts
}

func NewRollupCreator(privateKey string, l1conn string) (*RollupCreator, error) {
	client, err := ethclient.Dial(l1conn)
	if err != nil {
		return nil, err
	}
	auth, err := utils.GetBasicAuth(client, privateKey)
	if err != nil {
		return nil, err
	}
	return &RollupCreator{
		RPC:    l1conn,
		Client: client,
		opts:   auth,
	}, nil
}

func (r *RollupCreator) CreateRollup(
	ctx context.Context,
	owner common.Address,
	chainId *big.Int,
	parentChainId int,
	chainConfig string,
	genesisBlockNum uint64,
	loserStakeEscrow common.Address,
	wasmModuleRoot [32]byte,
	batchPoster common.Address,
	validators []common.Address,
	value *big.Int,
	gasLimit uint64,
) (*ethtypes.Transaction, error) {
	var config bindings.Config
	config.ConfirmPeriodBlocks = types.DefaultConfig.ConfirmPeriodBlocks
	config.ExtraChallengeTimeBlocks = types.DefaultConfig.ExtraChallengeTimeBlocks
	config.StakeToken = types.DefaultConfig.StakeToken
	config.BaseStake = types.DefaultConfig.BaseStake
	config.WasmModuleRoot = wasmModuleRoot
	config.Owner = owner
	config.LoserStakeEscrow = loserStakeEscrow
	config.ChainConfig = chainConfig
	config.ChainId = chainId
	config.GenesisBlockNum = genesisBlockNum
	config.SequencerInboxMaxTimeVariation = types.DefaultSequencerInboxMaxTimeVariation

	var deploymentParams bindings.RollupCreatorRollupDeploymentParams
	deploymentParams.Config = config
	deploymentParams.BatchPoster = batchPoster
	deploymentParams.Validators = validators
	deploymentParams.MaxDataSize = types.DefaultRollupCreatorRollupDeploymentParams.MaxDataSize
	deploymentParams.NativeToken = types.DefaultRollupCreatorRollupDeploymentParams.NativeToken
	deploymentParams.DeployFactoriesToL2 = true
	deploymentParams.MaxFeePerGasForRetryables = types.DefaultRollupCreatorRollupDeploymentParams.MaxFeePerGasForRetryables
	rollupCreatorAddr := types.ContractConfig[parentChainId]
	rollupCreatorTransactor, err := bindings.NewRollupCreatorTransactor(common.HexToAddress(rollupCreatorAddr), r.Client)
	if err != nil {
		return nil, err
	}

	// setting nonce and gas price
	nonce, err := r.Client.PendingNonceAt(context.Background(), r.opts.From)
	if err != nil {
		return nil, err
	}
	r.opts.Nonce = big.NewInt(int64(nonce))
	r.opts.Value = value
	r.opts.GasLimit = gasLimit

	return rollupCreatorTransactor.CreateRollup(r.opts, deploymentParams)
}

func (r *RollupCreator) ParseRollupContracts(ctx context.Context, chainId int, txn *ethtypes.Transaction) (*bindings.RollupCreatorRollupCreated, error) {
	err := utils.WaitTx(ctx, r.Client, txn, false)
	if err != nil {
		return nil, err
	}
	receipt, err := r.Client.TransactionReceipt(ctx, txn.Hash())
	if err != nil {
		return nil, err
	}
	rollupCreatorAddr := types.ContractConfig[chainId]

	rollupCreatorParser, err := bindings.NewRollupCreator(common.HexToAddress(rollupCreatorAddr), r.Client)
	if err != nil {
		return nil, err
	}

	rollupCreatorRollupCreated, err := rollupCreatorParser.ParseRollupCreated(*receipt.Logs[len(receipt.Logs)-1])
	if err != nil {
		return nil, err
	}
	return rollupCreatorRollupCreated, nil
}

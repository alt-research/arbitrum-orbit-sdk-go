package types

import (
	"context"
	"encoding/hex"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	ethtypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/renlulu/arbitrum-orbit-sdk-go/pkg/bindings"
)

// source: https://github.com/OffchainLabs/arbitrum-orbit-sdk/blob/main/wagmi.config.ts#L94
const (
	ethereum        = "0x90d68b056c411015eae3ec0b98ad94e2c91419f1"
	arbitrumOne     = "0x9CAd81628aB7D8e239F1A5B497313341578c5F71"
	arbitrumNova    = "0x9CAd81628aB7D8e239F1A5B497313341578c5F71"
	base            = "0x850F050C65B34966895AdA26a4D06923901916DB"
	sepolia         = "0xfbd0b034e6305788007f6e0123cc5eae701a5751"
	holesky         = "0xB512078282F462Ba104231ad856464Ceb0a7747e"
	arbitrumSepolia = "0x06E341073b2749e0Bb9912461351f716DeCDa9b0"
	baseSepolia     = "0x1E0921818df948c338380e722C8aE91Bb285763C"
)

var (
	DefaultBaseStake                      = big.NewInt(100000000000000000)
	DefaultWasmModuleRoot, _              = hex.DecodeString("0x8b104a2e80ac6165dc58b9048de12f301d70b02a0ab51396c22b4b4b802a16a4")
	DefaultSequencerInboxMaxTimeVariation = bindings.ISequencerInboxMaxTimeVariation{
		DelayBlocks:   big.NewInt(5760),
		FutureBlocks:  big.NewInt(48),
		DelaySeconds:  big.NewInt(86400),
		FutureSeconds: big.NewInt(3600),
	}
	DefaultConfig = bindings.Config{
		ConfirmPeriodBlocks:            15,
		ExtraChallengeTimeBlocks:       0,
		StakeToken:                     common.HexToAddress("0x0000000000000000000000000000000000000000"),
		BaseStake:                      DefaultBaseStake,
		SequencerInboxMaxTimeVariation: DefaultSequencerInboxMaxTimeVariation,
	}
	DefaultRollupCreatorRollupDeploymentParams = bindings.RollupCreatorRollupDeploymentParams{
		Config:                    DefaultConfig,
		MaxDataSize:               big.NewInt(104857),
		NativeToken:               common.HexToAddress("0x0000000000000000000000000000000000000000"),
		DeployFactoriesToL2:       true,
		MaxFeePerGasForRetryables: big.NewInt(100000000),
	}
)

type RollupCreator struct {
	RPC                     string
	Address                 string
	opts                    *bind.TransactOpts
	RollupCreatorTransactor *bindings.RollupCreatorTransactor
}

func NewRollupCreator(rpc string, address string, opts *bind.TransactOpts) *RollupCreator {
	return &RollupCreator{
		RPC:     rpc,
		Address: address,
		opts:    opts,
	}
}

func (r *RollupCreator) CreateRollupWithNativeEther(
	ctx context.Context,
	owner common.Address,
	chainId *big.Int,
	chainConfig string,
	genesisBlockNum uint64,
	loserStakeEscrow common.Address,
	wasmModuleRoot [32]byte,
	batchPoster common.Address,
	validators []common.Address,

) (*ethtypes.Transaction, error) {
	var config bindings.Config
	config.ConfirmPeriodBlocks = DefaultConfig.ConfirmPeriodBlocks
	config.ExtraChallengeTimeBlocks = DefaultConfig.ExtraChallengeTimeBlocks
	config.StakeToken = DefaultConfig.StakeToken
	config.BaseStake = DefaultConfig.BaseStake
	config.WasmModuleRoot = wasmModuleRoot
	config.Owner = owner
	config.LoserStakeEscrow = loserStakeEscrow
	config.ChainConfig = chainConfig
	config.ChainConfig = chainConfig
	config.GenesisBlockNum = genesisBlockNum
	config.SequencerInboxMaxTimeVariation = DefaultSequencerInboxMaxTimeVariation

	var deploymentParams bindings.RollupCreatorRollupDeploymentParams
	deploymentParams.Config = config
	deploymentParams.BatchPoster = batchPoster
	deploymentParams.Validators = validators
	deploymentParams.MaxDataSize = DefaultRollupCreatorRollupDeploymentParams.MaxDataSize
	deploymentParams.NativeToken = DefaultRollupCreatorRollupDeploymentParams.NativeToken
	deploymentParams.DeployFactoriesToL2 = true
	deploymentParams.MaxFeePerGasForRetryables = DefaultRollupCreatorRollupDeploymentParams.MaxFeePerGasForRetryables
	return r.RollupCreatorTransactor.CreateRollup(r.opts, deploymentParams)
}

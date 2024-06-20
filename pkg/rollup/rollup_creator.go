package rollup

import (
	"context"
	"crypto/ecdsa"
	"errors"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	ethtypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/renlulu/arbitrum-orbit-sdk-go/pkg/bindings"
	"github.com/renlulu/arbitrum-orbit-sdk-go/pkg/types"
)

type RollupCreator struct {
	RPC                     string
	Address                 string
	PrivateKey              string
	opts                    *bind.TransactOpts
	RollupCreatorTransactor *bindings.RollupCreatorTransactor
}

func NewRollupCreator(chainIndex int, privateKey string, l1conn string, address string) (*RollupCreator, error) {
	key, err := crypto.HexToECDSA(privateKey)
	if err != nil {
		return nil, err
	}

	client, err := ethclient.Dial(l1conn)
	if err != nil {
		return nil, err
	}

	publicKey := key.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		return nil, errors.New("Falied to cast public key to ECDSA")
	}
	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)

	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		return nil, err
	}

	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		return nil, err
	}

	chainID, err := client.NetworkID(context.Background())
	if err != nil {
		return nil, err
	}

	auth, err := bind.NewKeyedTransactorWithChainID(key, chainID)
	if err != nil {
		return nil, err
	}
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)
	auth.GasLimit = uint64(300000) // in units
	auth.GasPrice = gasPrice

	rollupCreatorTransactor, err := bindings.NewRollupCreatorCaller(common.HexToAddress(types.RollupCreatorAddr[chainIndex]), client)
	if err != nil {
		return nil, err
	}
	return &RollupCreator{
		RPC:                     l1conn,
		Address:                 address,
		opts:                    auth,
		RollupCreatorTransactor: (*bindings.RollupCreatorTransactor)(rollupCreatorTransactor),
	}, nil
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
	config.ConfirmPeriodBlocks = types.DefaultConfig.ConfirmPeriodBlocks
	config.ExtraChallengeTimeBlocks = types.DefaultConfig.ExtraChallengeTimeBlocks
	config.StakeToken = types.DefaultConfig.StakeToken
	config.BaseStake = types.DefaultConfig.BaseStake
	config.WasmModuleRoot = wasmModuleRoot
	config.Owner = owner
	config.LoserStakeEscrow = loserStakeEscrow
	config.ChainConfig = chainConfig
	config.ChainConfig = chainConfig
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
	return r.RollupCreatorTransactor.CreateRollup(r.opts, deploymentParams)
}

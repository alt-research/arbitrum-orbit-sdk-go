package types

import (
	"encoding/hex"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/renlulu/arbitrum-orbit-sdk-go/pkg/bindings"
)

var (
	NODE_INTERFACE_ADDRESS    = common.HexToAddress("0x00000000000000000000000000000000000000C8")
	ARB_SYS_ADDRESS           = common.HexToAddress("0x0000000000000000000000000000000000000064")
	ARB_ADDRESS_TABLE_ADDRESS = common.HexToAddress("0x0000000000000000000000000000000000000066")
	ARB_OWNER_PUBLIC          = common.HexToAddress("0x000000000000000000000000000000000000006B")
	ARB_GAS_INFO              = common.HexToAddress("0x000000000000000000000000000000000000006C")
	ARB_STATISTICS            = common.HexToAddress("0x000000000000000000000000000000000000006F")
)

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

// source: https://github.com/OffchainLabs/arbitrum-orbit-sdk/blob/main/wagmi.config.ts#L94
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
		StakeToken:                     common.Address{},
		BaseStake:                      DefaultBaseStake,
		SequencerInboxMaxTimeVariation: DefaultSequencerInboxMaxTimeVariation,
	}
	DefaultRollupCreatorRollupDeploymentParams = bindings.RollupCreatorRollupDeploymentParams{
		Config:                    DefaultConfig,
		MaxDataSize:               big.NewInt(104857),
		NativeToken:               common.Address{},
		DeployFactoriesToL2:       true,
		MaxFeePerGasForRetryables: big.NewInt(100000000),
	}
)

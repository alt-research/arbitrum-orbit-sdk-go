package rollup

import "github.com/renlulu/arbitrum-orbit-sdk-go/pkg/types"

func GenerateL2Config(chainId int64, initialChainOwner string, enableAnyTrust bool) *types.L2Config {
	return &types.L2Config{
		ChainId:             chainId,
		HomesteadBlock:      0,
		DaoForkBlock:        nil,
		DaoForkSupport:      true,
		Eip150Block:         0,
		Eip150Hash:          "0x0000000000000000000000000000000000000000000000000000000000000000",
		Eip155Block:         0,
		Eip158Block:         0,
		ByzantiumBlock:      0,
		ConstantinopleBlock: 0,
		PetersburgBlock:     0,
		IstanbulBlock:       0,
		MuirGlacierBlock:    0,
		BerlinBlock:         0,
		LondonBlock:         0,
		Clique: types.Clique{
			Period: 0,
			Epoch:  0,
		},
		Arbitrum: types.Arbitrum{
			EnableArbOS:               true,
			AllowDebugPrecompiles:     false,
			DataAvailabilityCommittee: enableAnyTrust,
			InitialArbOSVersion:       20,
			EigenDA:                   false,
			InitialChainOwner:         initialChainOwner,
			GenesisBlockNum:           0,
		},
	}
}

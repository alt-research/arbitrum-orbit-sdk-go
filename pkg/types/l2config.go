package types

//	{
//		"chainId": 10241024,
//		"homesteadBlock": 0,
//		"daoForkBlock": null,
//		"daoForkSupport": true,
//		"eip150Block": 0,
//		"eip150Hash": "0x0000000000000000000000000000000000000000000000000000000000000000",
//		"eip155Block": 0,
//		"eip158Block": 0,
//		"byzantiumBlock": 0,
//		"constantinopleBlock": 0,
//		"petersburgBlock": 0,
//		"istanbulBlock": 0,
//		"muirGlacierBlock": 0,
//		"berlinBlock": 0,
//		"londonBlock": 0,
//		"clique": {
//		  "period": 0,
//		  "epoch": 0
//		},
//		"arbitrum": {
//		  "EnableArbOS": true,
//		  "AllowDebugPrecompiles": false,
//		  "DataAvailabilityCommittee": true,
//		  "InitialArbOSVersion": 20,
//		  "EigenDA": false,
//		  "InitialChainOwner": "0xa179bb7dba00815007866f20539d39116b36b052",
//		  "GenesisBlockNum": 0
//		}
//	}
type L2Config struct {
	ChainId             int64    `json:"chainId"`
	HomesteadBlock      int64    `json:"homesteadBlock"`
	DaoForkBlock        *int64   `json:"daoForkBlock"`
	DaoForkSupport      bool     `json:"daoForkSupport"`
	Eip150Block         int64    `json:"eip150Block"`
	Eip150Hash          string   `json:"eip150Hash"`
	Eip155Block         int64    `json:"eip155Block"`
	Eip158Block         int64    `json:"eip158Block"`
	ByzantiumBlock      int64    `json:"byzantiumBlock"`
	ConstantinopleBlock int64    `json:"constantinopleBlock"`
	PetersburgBlock     int64    `json:"petersburgBlock"`
	IstanbulBlock       int64    `json:"istanbulBlock"`
	MuirGlacierBlock    int64    `json:"muirGlacierBlock"`
	BerlinBlock         int64    `json:"berlinBlock"`
	LondonBlock         int64    `json:"londonBlock"`
	Clique              Clique   `json:"clique"`
	Arbitrum            Arbitrum `json:"arbitrum"`
}

type Clique struct {
	Period int64 `json:"period"`
	Epoch  int64 `json:"epoch"`
}

type Arbitrum struct {
	EnableArbOS               bool   `json:"EnableArbOS"`
	AllowDebugPrecompiles     bool   `json:"AllowDebugPrecompiles"`
	DataAvailabilityCommittee bool   `json:"DataAvailabilityCommittee"`
	InitialArbOSVersion       int64  `json:"InitialArbOSVersion"`
	EigenDA                   bool   `json:"EigenDA"`
	InitialChainOwner         string `json:"InitialChainOwner"`
	GenesisBlockNum           int64  `json:"GenesisBlockNum"`
}

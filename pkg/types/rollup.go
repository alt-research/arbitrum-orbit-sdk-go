package types

import (
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
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

type RollupCreator struct {
	RPC     string
	Address string
	opts    *bind.TransactOpts
}

func NewRollupCreator(rpc string, address string, opts *bind.TransactOpts) *RollupCreator {
	return &RollupCreator{
		RPC:     rpc,
		Address: address,
		opts:    opts,
	}
}

var DefaultConfig = bindings.Config{}

var DefaultRollupCreatorRollupDeploymentParams = bindings.RollupCreatorRollupDeploymentParams{}

func (r *RollupCreator) CreateRollup() {

}

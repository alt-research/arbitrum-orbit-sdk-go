package types

import "math/big"

// An optional big number percentage increase
type PercentIncrease struct {
	// If not nil, will override the estimated base
	Base *big.Int
	// How much to increase the base by
	PercentIncrease *big.Int
}

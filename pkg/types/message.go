package types

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
)

// An optional big number percentage increase
type PercentIncrease struct {
	// If not nil, will override the estimated base
	Base *big.Int
	// How much to increase the base by
	PercentIncrease *big.Int
	Min             *big.Int
}

type GasOverrides struct {
	GasLimit         *PercentIncrease
	MaxSubmissionFee *PercentIncrease
	MaxFeePerGas     *PercentIncrease
}

// type TransactionRequestRetryableGasOverrides struct {
// 	MaxSubmissionCostForFactory   GasOverrideOptions
// 	MaxGasForFactory              GasOverrideOptions
// 	MaxSubmissionCostForContracts GasOverrideOptions
// 	MaxGasForContracts            GasOverrideOptions
// 	maxGasPrice                   GasOverrideOptions
// }

type CreateTokenBridgeGetInputsResult struct {
	Inbox              common.Address
	MaxGasForContracts *big.Int
	GasPrice           *big.Int
	RetryableFee       *big.Int
}

type RetryableData struct {
	From string
	// The address to be called on L2
	To string
	// The value to call the L2 address with
	L2CallValue            *big.Int
	Deposit                *big.Int
	MaxSubmissionCost      *big.Int
	ExcessFeeRefundAddress string
	CallValueRefundAddress string
	GasLimit               *big.Int
	MaxFeePerGas           *big.Int
	Data                   string
}

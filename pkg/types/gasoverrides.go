package types

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
)

type GasOverrideOptions struct {
	Base            *big.Int
	PercentIncrease *big.Int
}

type TransactionRequestRetryableGasOverrides struct {
	MaxSubmissionCostForFactory   GasOverrideOptions
	MaxGasForFactory              GasOverrideOptions
	MaxSubmissionCostForContracts GasOverrideOptions
	MaxGasForContracts            GasOverrideOptions
	maxGasPrice                   GasOverrideOptions
}

type CreateTokenBridgeGetInputsResult struct {
	Inbox              common.Address
	MaxGasForContracts *big.Int
	GasPrice           *big.Int
	RetryableFee       *big.Int
}

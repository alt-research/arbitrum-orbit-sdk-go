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

type TransactionRequestRetryableGasOverrides struct {
	MaxSubmissionCostForFactory   PercentIncrease
	MaxGasForFactory              PercentIncrease
	MaxSubmissionCostForContracts PercentIncrease
	MaxGasForContracts            PercentIncrease
	maxGasPrice                   PercentIncrease
}

type CreateTokenBridgeGetInputsResult struct {
	GasEstimateToDeployContracts *big.Int
	GasPrice                     *big.Int
	RetryableFee                 *big.Int
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

/**
 * The components of a submit retryable message. Can be parsed from the
 * events emitted from the Inbox.
 */

type RetryableMessageParams struct {
	// Destination address for L2 message
	DestAddress common.Address
	// Call value in L2 message
	L2CallValue *big.Int
	// Value set at L1
	L1Value *big.Int
	// Max gas deducted from L2 balance to cover the base submission fee
	MaxSubmissionFee *big.Int
	// L2 address address to credit (gaslimit x gasprice - execution cost)
	ExcessFeeRefundAddress common.Address
	// Address to credit l2Callvalue on L2 if retryable txn times out or gets cancelled
	CallValueRefundAddress common.Address
	// Max gas deducted from user's L2 balance to cover L2 execution
	GasLimit *big.Int
	// maxFeePerGas
	MaxFeePerGas *big.Int
	// Calldata for of the L2 message
	Data []byte
}

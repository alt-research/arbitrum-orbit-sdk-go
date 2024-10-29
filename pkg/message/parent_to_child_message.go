package message

import (
	"bytes"
	"encoding/hex"
	"math/big"

	"github.com/alt-research/arbitrum-orbit-sdk-go/pkg/types"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/math"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/rlp"
)

type ParentToChildMessage struct {
	RetryableCreationId []byte
}

func NewParentToChildMessage(
	chainId *big.Int,
	sender common.Address,
	messageNumber big.Int,
	parentBaseFee *big.Int,
	messageData types.RetryableMessageParams,
) (*ParentToChildMessage, error) {
	arbitrumSubmitRetryableTx := NewArbitrumSubmitRetryableTx(chainId, sender, messageNumber, parentBaseFee, messageData)
	retryableId, err := arbitrumSubmitRetryableTx.CalculateSubmitRetryableId()
	if err != nil {
		return nil, err
	}
	return &ParentToChildMessage{
		RetryableCreationId: retryableId,
	}, nil

}

// port from nitro
type ArbitrumSubmitRetryableTx struct {
	ChainId          *big.Int
	RequestId        common.Hash
	From             common.Address
	L1BaseFee        *big.Int
	DepositValue     *big.Int
	GasFeeCap        *big.Int        // wei per gas
	Gas              uint64          // gas limit
	RetryTo          *common.Address `rlp:"nil"` // nil means contract creation
	RetryValue       *big.Int        // wei amount
	Beneficiary      common.Address
	MaxSubmissionFee *big.Int
	FeeRefundAddr    common.Address
	RetryData        []byte // contract invocation input data
}

// nitro code: for referring the naming
//
//	submitTx := &types.ArbitrumSubmitRetryableTx{
//		ChainId:          nil,
//		RequestId:        hash{},
//		From:             util.RemapL1Address(sender),
//		L1BaseFee:        l1BaseFee,
//		DepositValue:     deposit,
//		GasFeeCap:        n.sourceMessage.GasPrice,
//		Gas:              n.sourceMessage.GasLimit,
//		RetryTo:          pRetryTo,
//		RetryValue:       l2CallValue,
//		Beneficiary:      callValueRefundAddress,
//		MaxSubmissionFee: maxSubmissionFee,
//		FeeRefundAddr:    excessFeeRefundAddress,
//		RetryData:        data,
//	}
func NewArbitrumSubmitRetryableTx(
	chainId *big.Int,
	sender common.Address,
	messageNumber big.Int,
	parentBaseFee *big.Int,
	messageData types.RetryableMessageParams,
) *ArbitrumSubmitRetryableTx {
	retryTo := &messageData.DestAddress
	if retryTo.Hex() == "0x0000000000000000000000000000000000000000" {
		retryTo = nil
	}
	return &ArbitrumSubmitRetryableTx{
		ChainId:          chainId,
		RequestId:        common.BytesToHash(math.U256Bytes(&messageNumber)),
		From:             sender,
		L1BaseFee:        parentBaseFee,
		DepositValue:     messageData.L1Value,
		GasFeeCap:        messageData.MaxFeePerGas,
		Gas:              messageData.GasLimit.Uint64(),
		RetryTo:          retryTo,
		RetryValue:       messageData.L2CallValue,
		Beneficiary:      messageData.CallValueRefundAddress,
		MaxSubmissionFee: messageData.MaxSubmissionFee,
		FeeRefundAddr:    messageData.ExcessFeeRefundAddress,
		RetryData:        messageData.Data,
	}
}

func (a *ArbitrumSubmitRetryableTx) CalculateSubmitRetryableId() ([]byte, error) {
	var buf bytes.Buffer
	err := a.encode(&buf)
	if err != nil {
		return nil, err
	}
	prefix, err := hex.DecodeString("69")
	if err != nil {
		return nil, err
	}
	ret := append(prefix, buf.Bytes()...)
	return crypto.Keccak256(ret), nil

}

func (a *ArbitrumSubmitRetryableTx) encode(b *bytes.Buffer) error {
	return rlp.Encode(b, a)
}

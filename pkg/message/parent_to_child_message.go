package message

import (
	"bytes"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/rlp"
	"github.com/renlulu/arbitrum-orbit-sdk-go/pkg/types"
)

type ParentToChildMessage struct {
}

// port from nitro
type ArbitrumSubmitRetryableTx struct {
	ChainId   *big.Int
	RequestId common.Hash
	From      common.Address
	L1BaseFee *big.Int

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

func NewArbitrumSubmitRetryableTx(
	chainId *big.Int,
	sender common.Address,
	messageNumber big.Int,
	parentBaseFee *big.Int,
	messageData types.RetryableMessageParams,
) *ArbitrumSubmitRetryableTx {
	return &ArbitrumSubmitRetryableTx{}
}

func (r *ArbitrumSubmitRetryableTx) CalculateSubmitRetryableId() ([]byte, error) {
	// Buffer to store the encoded data
	var buf bytes.Buffer
	err := rlp.Encode(&buf, r)
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

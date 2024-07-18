package message

import (
	"bytes"
	"context"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
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

func Txx() error {
	baseChainRpc := "https://arbitrum-sepolia.blockpi.network/v1/rpc/public"
	baseChainClient, err := ethclient.Dial(baseChainRpc)
	if err != nil {
		return err
	}

	txn, _, err := baseChainClient.TransactionByHash(context.Background(), common.HexToHash("0xfa0cba9792ab55a3913413dde1b4f29461eb81362719b0a069593a3689b819b6"))
	if err != nil {
		return err
	}
	fmt.Println("txn: ", txn)
	receipt, err := bind.WaitMined(context.Background(), baseChainClient, txn)
	if err != nil {
		return err
	}
	fmt.Println("receipt: ", receipt.Logs)
	parentChainReceipt := NewParentTransactionReceipt(common.HexToAddress("0x66530799037b46913e52e9e0144d15ab6ed954f5"), common.HexToAddress("0x56c486d3786fa26cc61473c499a36eb9cc1fbd8e"), receipt)
	messsageEvents, err := parentChainReceipt.GetMessageEvents()
	if err != nil {
		return err
	}
	fmt.Println("length of messsageEvents: ", len(messsageEvents))
	return nil

}

package message

import (
	"errors"
	"math/big"

	"github.com/umbracle/ethgo/abi"
)

var (
	typ = abi.MustNewType("tuple(uint256 dest,uint256 l2callvalue,uint256 submissioncost,uint256 excess,uint256 callvaluerefund,uint256 gaslimit,uint256 maxfeebid,uint256 datalength)")
)

type MessageData struct {
	Dest                   *big.Int `abi:"dest"`
	L2CallValue            *big.Int `abi:"l2callvalue"`
	MaxSubmissionCost      *big.Int `abi:"submissioncost"`
	ExcessFeeRefundAddress *big.Int `abi:"excess"`
	CallValueRefundAddress *big.Int `abi:"callvaluerefund"`
	GasLimit               *big.Int `abi:"gaslimit"`
	MaxFeePerGas           *big.Int `abi:"maxfeebid"`
	DataLength             *big.Int `abi:"datalength"`
}

func DecodeMessageData(encoded []byte) (*MessageData, error) {
	md, err := typ.Decode(encoded)
	if err != nil {
		return nil, err
	}

	decoded, ok := md.(map[string]interface{})
	if !ok {
		return nil, errors.New("ArbSDK: cannot decode")
	}

	return &MessageData{
		Dest:                   decoded["dest"].(*big.Int),
		L2CallValue:            decoded["l2callvalue"].(*big.Int),
		MaxSubmissionCost:      decoded["submissioncost"].(*big.Int),
		ExcessFeeRefundAddress: decoded["excess"].(*big.Int),
		CallValueRefundAddress: decoded["callvaluerefund"].(*big.Int),
		GasLimit:               decoded["gaslimit"].(*big.Int),
		MaxFeePerGas:           decoded["maxfeebid"].(*big.Int),
		DataLength:             decoded["datalength"].(*big.Int),
	}, nil
}

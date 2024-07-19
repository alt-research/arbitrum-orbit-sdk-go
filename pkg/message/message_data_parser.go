package message

import (
	"errors"
	"math/big"

	"github.com/umbracle/ethgo/abi"
)

var (
	typ = abi.MustNewType("tuple(uint256 dest,uint256 l2callvalue,uint256 l1callvalue, uint256 submissioncost,uint256 excess,uint256 callvaluerefund,uint256 gaslimit,uint256 maxfeebid,uint256 datalength)")
)

type MessageDataParam struct {
	Dest                   *big.Int `abi:"dest"`
	L2CallValue            *big.Int `abi:"l2callvalue"`
	L1CallValue            *big.Int `abi:"l1callvalue"`
	MaxSubmissionCost      *big.Int `abi:"submissioncost"`
	ExcessFeeRefundAddress *big.Int `abi:"excess"`
	CallValueRefundAddress *big.Int `abi:"callvaluerefund"`
	GasLimit               *big.Int `abi:"gaslimit"`
	MaxFeePerGas           *big.Int `abi:"maxfeebid"`
	DataLength             *big.Int `abi:"datalength"`
}

type MessageData struct {
	Dest                   *big.Int
	L2CallValue            *big.Int
	L1CallValue            *big.Int
	MaxSubmissionCost      *big.Int
	ExcessFeeRefundAddress *big.Int
	CallValueRefundAddress *big.Int
	GasLimit               *big.Int
	MaxFeePerGas           *big.Int
	DataLength             *big.Int
	Data                   []byte
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

	calldataLength := decoded["datalength"].(*big.Int)
	eventDataLength := len(encoded)
	start := eventDataLength - int(calldataLength.Int64())
	data := encoded[start:]

	return &MessageData{
		Dest:                   decoded["dest"].(*big.Int),
		L2CallValue:            decoded["l2callvalue"].(*big.Int),
		L1CallValue:            decoded["l1callvalue"].(*big.Int),
		MaxSubmissionCost:      decoded["submissioncost"].(*big.Int),
		ExcessFeeRefundAddress: decoded["excess"].(*big.Int),
		CallValueRefundAddress: decoded["callvaluerefund"].(*big.Int),
		GasLimit:               decoded["gaslimit"].(*big.Int),
		MaxFeePerGas:           decoded["maxfeebid"].(*big.Int),
		DataLength:             decoded["datalength"].(*big.Int),
		Data:                   data,
	}, nil
}

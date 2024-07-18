package message

import (
	"errors"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/renlulu/arbitrum-orbit-sdk-go/pkg/bindings/rollupgen"
)

var (
	// Solidity: event MessageDelivered(uint256 indexed messageIndex, bytes32 indexed beforeInboxAcc, address inbox, uint8 kind, address sender, bytes32 messageDataHash, uint256 baseFeeL1, uint64 timestamp)
	MessageDeliveredSig     = []byte("MessageDelivered(uint256,bytes32,address,uint8,address,bytes32,uint256,uint64)")
	MessageDeliveredSigHash = common.BytesToHash(crypto.Keccak256(MessageDeliveredSig))
	MessageDeliveredABI, _  = rollupgen.BridgeMetaData.GetAbi()
	// Solidity: event InboxMessageDelivered(uint256 indexed messageNum, bytes data)
	InboxMessageDeliveredSig     = []byte("InboxMessageDelivered(uint256,bytes)")
	InboxMessageDeliveredSigHash = common.BytesToHash(crypto.Keccak256(InboxMessageDeliveredSig))
	InboxMessageDeliveredABI, _  = rollupgen.InboxMetaData.GetAbi()
)

type ParentTransaction struct {
}

// basically mirror from types.Receipt with a few more fields
type ParentTransactionReceipt struct {
	To                common.Address
	From              common.Address
	ContractAddress   common.Address
	TransactionIndex  uint
	GasUsed           uint64
	LogsBloom         types.Bloom
	BlockHash         common.Hash
	TransactionHash   common.Hash
	Logs              []*types.Log
	BlockNumber       *big.Int
	Confirmations     uint64
	CumulativeGasUsed uint64
	Type              uint8
	Status            uint64
	// todo
	Root              string
	EffectiveGasPrice *big.Int
	Byzantium         bool
}

func NewParentTransactionReceipt(to, from common.Address, receipt *types.Receipt) *ParentTransactionReceipt {
	return &ParentTransactionReceipt{
		To:                to,
		From:              from,
		ContractAddress:   receipt.ContractAddress,
		TransactionIndex:  receipt.TransactionIndex,
		GasUsed:           receipt.GasUsed,
		LogsBloom:         receipt.Bloom,
		BlockHash:         receipt.BlockHash,
		TransactionHash:   receipt.TxHash,
		Logs:              receipt.Logs,
		BlockNumber:       receipt.BlockNumber,
		CumulativeGasUsed: receipt.CumulativeGasUsed,
		Type:              receipt.Type,
		Status:            receipt.Status,
	}
}

func (p *ParentTransactionReceipt) GetMessageEvents() ([]*rollupgen.BridgeMessageDelivered, []*rollupgen.InboxInboxMessageDelivered, error) {
	messageDeliveredEvents, err := p.getMessageDeliveredEvents()
	if err != nil {
		return nil, nil, err
	}
	inboxMessageDeliveredEvents, err := p.getInboxMessageDeliveredEvents()
	if err != nil {
		return nil, nil, err
	}
	if len(messageDeliveredEvents) != len(inboxMessageDeliveredEvents) {
		return nil, nil, errors.New("ArbSDK: missing event for message index")
	}
	return messageDeliveredEvents, inboxMessageDeliveredEvents, err

}

func (p *ParentTransactionReceipt) getMessageDeliveredEvents() ([]*rollupgen.BridgeMessageDelivered, error) {
	var messageDeliveredEvents []*rollupgen.BridgeMessageDelivered
	for _, vLog := range p.Logs {
		if vLog.Topics[0] == MessageDeliveredSigHash {
			var messageDelivered rollupgen.BridgeMessageDelivered
			err := MessageDeliveredABI.UnpackIntoInterface(&messageDelivered, "MessageDelivered", vLog.Data)
			if err != nil {
				return nil, err
			}
			messageDeliveredEvents = append(messageDeliveredEvents, &messageDelivered)
		}
	}
	return messageDeliveredEvents, nil
}

func (p *ParentTransactionReceipt) getInboxMessageDeliveredEvents() ([]*rollupgen.InboxInboxMessageDelivered, error) {
	var inboxMessageDeliveredEvents []*rollupgen.InboxInboxMessageDelivered
	for _, vLog := range p.Logs {
		if vLog.Topics[0] == InboxMessageDeliveredSigHash {
			var inboxMessageDelivered rollupgen.InboxInboxMessageDelivered
			err := InboxMessageDeliveredABI.UnpackIntoInterface(&inboxMessageDelivered, "InboxMessageDelivered", vLog.Data)
			if err != nil {
				return nil, err
			}
			inboxMessageDeliveredEvents = append(inboxMessageDeliveredEvents, &inboxMessageDelivered)

		}
	}
	return inboxMessageDeliveredEvents, nil
}

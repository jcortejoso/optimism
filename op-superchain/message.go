package superchain

import (
	"bytes"
	"fmt"
	"math/big"

	"github.com/ethereum-optimism/optimism/op-service/solabi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
)

var (
	// TODO: use predeploy constant when available
	crossL2InboxAddr = common.Address{}

	inboxExecuteMessageSignature      = "executeMessage((address,uint256,uint256,uint256,uint256),address,bytes)"
	inboxExecuteMessageBytes4         = crypto.Keccak256([]byte(inboxExecuteMessageSignature))[:4]
	inboxExecuteMessagePayloadDataLoc = common.HexToHash("0xe0")
)

type MessageSafetyLabel int

const (
	MessageUnknown MessageSafetyLabel = iota - 1
	MessageInvalid
	MessageUnsafe
	MessageCrossUnsafe
	MessageSafe
	MessageFinalized
)

type MessageIdentifier struct {
	Origin      common.Address
	BlockNumber *big.Int
	LogIndex    uint64
	Timestamp   uint64
	ChainId     *big.Int
}

func MessagePayloadBytes(log *types.Log) []byte {
	msg := []byte{}
	for _, topic := range log.Topics {
		msg = append(msg, topic.Bytes()...)
	}
	return append(msg, log.Data...)
}

// Check if a transaction is an executing message to the CrossL2Inbox
func IsInboxExecutingMessageTx(tx *types.Transaction) bool {
	if tx.To() == nil || *tx.To() == crossL2InboxAddr {
		return false
	}

	txData := tx.Data()
	return len(txData) >= 4 && bytes.Equal(txData[:4], inboxExecuteMessageBytes4)
}

// Check the message id and payload against the fields of the log.
func CheckMessageLog(id MessageIdentifier, payload hexutil.Bytes, log *types.Log) error {
	if id.LogIndex != uint64(log.Index) {
		return fmt.Errorf("log index mismatch")
	}
	if !bytes.Equal(payload, MessagePayloadBytes(log)) {
		return fmt.Errorf("payload mismatch")
	}
	if id.Origin != log.Address {
		return fmt.Errorf("origin mismatch")
	}
	if id.BlockNumber.Uint64() != log.BlockNumber {
		return fmt.Errorf("block number mismatch")
	}
	return nil
}

// Parse the transaction data posted to the inbox `executeMessage` function, extracing it's parameters
func ParseInboxExecuteMessageTxData(txData []byte) (common.Address, MessageIdentifier, []byte, error) {
	var target common.Address
	var id MessageIdentifier

	r := bytes.NewReader(txData)

	// Validate Function Signature
	_, err := solabi.ReadAndValidateSignature(r, inboxExecuteMessageBytes4)
	if err != nil {
		return target, id, nil, err
	}

	// Read Identifier
	id.Origin, err = solabi.ReadAddress(r)
	if err != nil {
		return target, id, nil, fmt.Errorf("failed to read identifier origin: %w", err)
	}
	id.BlockNumber, err = solabi.ReadUint256(r)
	if err != nil {
		return target, id, nil, fmt.Errorf("failed to read identifier block number: %w", err)
	}
	id.LogIndex, err = solabi.ReadUint64(r)
	if err != nil {
		return target, id, nil, fmt.Errorf("failed to read identifier log index: %w", err)
	}
	id.Timestamp, err = solabi.ReadUint64(r)
	if err != nil {
		return target, id, nil, fmt.Errorf("failed to read identifier timestamp: %w", err)
	}
	id.ChainId, err = solabi.ReadUint256(r)
	if err != nil {
		return target, id, nil, fmt.Errorf("failed to read identifier chain id: %w", err)
	}

	// Read target
	target, err = solabi.ReadAddress(r)
	if err != nil {
		return target, id, nil, fmt.Errorf("failed to read target: %w", err)
	}

	// Read Message Bytes
	dataLoc, err := solabi.ReadHash(r)
	if err != nil {
		return target, id, nil, fmt.Errorf("failed to read message data loc: %w", err)
	}
	if dataLoc != inboxExecuteMessagePayloadDataLoc {
		return target, id, nil, fmt.Errorf("mismatched message data loc. Got %s, Expected: %s", dataLoc, inboxExecuteMessagePayloadDataLoc)
	}
	message, err := solabi.ReadBytes(r)
	if err != nil {
		return target, id, nil, fmt.Errorf("failed to read message: %w", err)
	}

	return target, id, message, nil
}

package pregod_etl

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/shopspring/decimal"
)

type GetLogsRequest struct {
	Address         string `url:"address,omitempty"`
	TopicFirst      string `url:"topic_first,omitempty"`
	TopicSecond     string `url:"topic_second,omitempty"`
	TopicThird      string `url:"topic_third,omitempty"`
	TopicFourth     string `url:"topic_fourth,omitempty"`
	TransactionHash string `url:"transaction_hash,omitempty"`
	BlockNumberFrom int64  `url:"block_number_from,omitempty"`
	BlockNumberTo   int64  `url:"block_number_to,omitempty"`
}

// EthereumLog is the model entity for the EthereumLog schema.
type EthereumLog struct {
	// ID of the ent.
	ID common.Hash `json:"id,omitempty"`
	// LogIndex holds the value of the "log_index" field.
	LogIndex uint8 `json:"log_index,omitempty"`
	// BlockNumber holds the value of the "block_number" field.
	BlockNumber decimal.Decimal `json:"block_number,omitempty"`
	// BlockHash holds the value of the "block_hash" field.
	BlockHash common.Hash `json:"block_hash,omitempty"`
	// TransactionHash holds the value of the "transaction_hash" field.
	TransactionHash common.Hash `json:"transaction_hash,omitempty"`
	// TransactionIndex holds the value of the "transaction_index" field.
	TransactionIndex uint8 `json:"transaction_index,omitempty"`
	// Address holds the value of the "address" field.
	Address common.Address `json:"address,omitempty"`
	// Removed holds the value of the "removed" field.
	Removed bool `json:"removed,omitempty"`
	// TopicFirst holds the value of the "topic_first" field.
	TopicFirst *common.Hash `json:"topic_first,omitempty"`
	// TopicSecond holds the value of the "topic_second" field.
	TopicSecond *common.Hash `json:"topic_second,omitempty"`
	// TopicThird holds the value of the "topic_third" field.
	TopicThird *common.Hash `json:"topic_third,omitempty"`
	// TopicFourth holds the value of the "topic_fourth" field.
	TopicFourth *common.Hash `json:"topic_fourth,omitempty"`
	// Data holds the value of the "data" field.
	Data hexutil.Bytes `json:"data,omitempty"`
}

package pregod_etl

import (
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
	Cursor          string `url:"cursor,omitempty"`
	Order           string `url:"order"`
	Limit           int    `url:"limit"`
}

type GetLogsResponse struct {
	Cursor string        `json:"cursor,omitempty"`
	Total  int64         `json:"total"`
	Result []EthereumLog `json:"result"`
}

// EthereumLog is the model entity for the EthereumLog schema.
type EthereumLog struct {
	// LogIndex holds the value of the "log_index" field.
	LogIndex decimal.Decimal `json:"log_index,omitempty"`
	// BlockNumber holds the value of the "block_number" field.
	BlockNumber decimal.Decimal `json:"block_number,omitempty"`
	// BlockHash holds the value of the "block_hash" field.
	BlockHash string `json:"block_hash,omitempty"`
	// TransactionHash holds the value of the "transaction_hash" field.
	TransactionHash string `json:"transaction_hash,omitempty"`
	// TransactionIndex holds the value of the "transaction_index" field.
	TransactionIndex decimal.Decimal `json:"transaction_index,omitempty"`
	// Address holds the value of the "address" field.
	Address string `json:"address,omitempty"`
	// TopicFirst holds the value of the "topic_first" field.
	TopicFirst string `json:"topic_first,omitempty"`
	// TopicSecond holds the value of the "topic_second" field.
	TopicSecond string `json:"topic_second,omitempty"`
	// TopicThird holds the value of the "topic_third" field.
	TopicThird string `json:"topic_third,omitempty"`
	// TopicFourth holds the value of the "topic_fourth" field.
	TopicFourth string `json:"topic_fourth,omitempty"`
	// Data holds the value of the "data" field.
	Data string `json:"data,omitempty"`
}

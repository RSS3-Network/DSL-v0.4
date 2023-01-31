package aptos

import (
	"github.com/shopspring/decimal"
)

type Event struct {
	Guid struct {
		CreationNumber string `json:"creation_number"`
		AccountAddress string `json:"account_address"`
	} `json:"guid"`
	SequenceNumber decimal.Decimal `json:"sequence_number"`
	Type           string          `json:"type"`
	Data           struct {
		Amount decimal.Decimal `json:"amount"`
	} `json:"data"`
}

type Payload struct {
	Function      string        `json:"function"`
	TypeArguments []interface{} `json:"type_arguments"`
	Arguments     []interface{} `json:"arguments"`
	Type          string        `json:"type"`
}

type GetAccountTransactionsParameter struct {
	Address string
	Limit   int             `url:"limit,omitempty"`
	Start   decimal.Decimal `url:"start,omitempty"`
}

type GetAccountTransactionsResult struct {
	Version        decimal.Decimal `json:"version"`
	SequenceNumber decimal.Decimal `json:"sequence_number"`
	Hash           string          `json:"hash"`
	Success        bool            `json:"success"`
	GasUsed        decimal.Decimal `json:"gas_used"`
	GasUnitPrice   decimal.Decimal `json:"gas_unit_price"`
	Sender         string          `json:"sender"`
	Timestamp      decimal.Decimal `json:"timestamp"`
	Type           string          `json:"type"`
	Payload        Payload         `json:"payload"`
	Events         []Event         `json:"events"`
}

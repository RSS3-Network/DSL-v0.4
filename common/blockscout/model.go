package blockscout

import (
	"encoding/json"

	"github.com/shopspring/decimal"
)

type Response struct {
	Message string          `json:"message"`
	Result  json.RawMessage `json:"result"`
	Status  decimal.Decimal `json:"status"`
}

type Transaction struct {
	BlockHash         string          `json:"blockHash"`
	BlockNumber       decimal.Decimal `json:"blockNumber"`
	Confirmations     decimal.Decimal `json:"confirmations"`
	ContractAddress   string          `json:"contractAddress"`
	CumulativeGasUsed string          `json:"cumulativeGasUsed"`
	From              string          `json:"from"`
	Gas               decimal.Decimal `json:"gas"`
	GasPrice          decimal.Decimal `json:"gasPrice"`
	GasUsed           decimal.Decimal `json:"gasUsed"`
	Hash              string          `json:"hash"`
	Input             string          `json:"input"`
	IsError           string          `json:"isError"`
	LogIndex          decimal.Decimal `json:"logIndex"`
	Nonce             string          `json:"nonce"`
	TimeStamp         decimal.Decimal `json:"timeStamp"`
	To                string          `json:"to"`
	TokenDecimal      string          `json:"tokenDecimal,omitempty"`
	TokenName         string          `json:"tokenName,omitempty"`
	TokenSymbol       string          `json:"tokenSymbol,omitempty"`
	TransactionIndex  decimal.Decimal `json:"transactionIndex"`
	TxReceiptStatus   decimal.Decimal `json:"txreceipt_status"`
	Value             decimal.Decimal `json:"value"`
}

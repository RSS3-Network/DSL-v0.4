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
	TokenID           decimal.Decimal `json:"tokenID,omitempty"`
	TokenDecimal      string          `json:"tokenDecimal,omitempty"`
	TokenName         string          `json:"tokenName,omitempty"`
	TokenSymbol       string          `json:"tokenSymbol,omitempty"`
	TransactionIndex  decimal.Decimal `json:"transactionIndex"`
	TxReceiptStatus   decimal.Decimal `json:"txreceipt_status"`
	Value             decimal.Decimal `json:"value"`
}

type TransactionInfo struct {
	RevertReason  string               `json:"revertReason"`
	BlockNumber   decimal.Decimal      `json:"blockNumber"`
	Confirmations decimal.Decimal      `json:"confirmations"`
	From          string               `json:"from"`
	GasLimit      decimal.Decimal      `json:"gasLimit"`
	GasPrice      decimal.Decimal      `json:"gasPrice"`
	GasUsed       decimal.Decimal      `json:"gasUsed"`
	Hash          string               `json:"hash"`
	Input         string               `json:"input"`
	Logs          []TransactionInfoLog `json:"logs"`
	Success       bool                 `json:"success"`
	TimeStamp     decimal.Decimal      `json:"timeStamp"`
	To            string               `json:"to"`
	Value         decimal.Decimal      `json:"value"`
}

type TransactionInfoLog struct {
	Address string   `json:"address"`
	Data    string   `json:"data"`
	Topics  []string `json:"topics"`
}
type GetTransactionListOption struct {
	Module          string `url:"module"`
	Action          string `url:"action"`
	Address         string `url:"address"`
	ContractAddress string `url:"contractaddress,omitempty"`
	Sort            string `url:"sort,omitempty"`
	StartBlock      int64  `url:"startblock,omitempty"`
	EndBlock        int64  `url:"endblock,omitempty"`
	Page            int64  `url:"page,omitempty"`
	Offset          int64  `url:"offset,omitempty"`
	FilterBy        string `url:"filterby,omitempty"`
	StartTimestamp  int64  `url:"starttimestamp,omitempty"`
	EndTimestamp    int64  `url:"endtimestamp,omitempty"`
}

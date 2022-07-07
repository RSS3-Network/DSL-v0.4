package zksync

import (
	"encoding/json"
	"time"
)

type Response struct {
	Status string          `json:"status"`
	Result json.RawMessage `json:"result,omitempty"`
	Error  any             `json:"error,omitempty"`
}

type Error struct {
	Type    string `json:"errorType"`
	Code    int64  `json:"code"`
	Message string `json:"message"`
}

type GetTransactionData struct {
	Transaction       GetTransactionDataTransaction `json:"tx"`
	EthereumSignature string                        `json:"ethSignature"`
}

type GetTransactionDataTransaction struct {
	TransactionHash string                      `json:"txHash"`
	BlockNumber     int                         `json:"blockNumber"`
	Operation       GetTransactionDataOperation `json:"op"`
	Status          string                      `json:"status"`
	FailReason      string                      `json:"failReason"`
	CreatedAt       string                      `json:"createdAt"`
	BatchID         int                         `json:"batchId"`
}

// https://docs.zksync.io/apiv02-docs/#transactions-api-v0.2-transactions-txhash-data
// only support `Transfer` now
type GetTransactionDataOperation struct {
	Type       string                               `json:"type"`
	To         string                               `json:"to"`
	Fee        string                               `json:"fee"`
	From       string                               `json:"from"`
	Nonce      int                                  `json:"nonce"`
	Token      int                                  `json:"transaction"`
	Amount     string                               `json:"amount"`
	AccountId  int                                  `json:"accountId"`
	ValidFrom  int                                  `json:"validFrom"`
	ValidUntil int                                  `json:"validUntil"`
	Signature  GetTransactionDataOperationSignature `json:"signature"`
}

type GetTransactionDataOperationSignature struct {
	PubKey    string `json:"pubKey"`
	Signature string `json:"signature"`
}

type GetAccountTransactionList struct {
	Pagination GetAccountTransactionListPagination `json:"pagination"`
	List       []GetAccountTransactionListItem     `json:"list"`
}

type GetAccountTransactionListPagination struct {
	From      string `json:"from"`
	Limit     int    `json:"limit"`
	Direction string `json:"direction"`
	Count     int    `json:"count"`
}

type GetAccountTransactionListItem struct {
	TransactionHash string                                 `json:"txHash"`
	BlockNumber     int64                                  `json:"blockNumber"`
	Operation       GetAccountTransactionListItemOperation `json:"op"`
	Status          string                                 `json:"status"`
	FailReason      string                                 `json:"failReason"`
	CreatedAt       time.Time                              `json:"createdAt"`
	BatchID         int64                                  `json:"batchId"`
}

type GetAccountTransactionListItemOperation struct {
	Type            string `json:"type"`
	From            string `json:"from"`
	TokenID         int    `json:"tokenId"`
	Amount          string `json:"amount"`
	To              string `json:"to"`
	AccountID       int    `json:"accountId"`
	EthereumHash    string `json:"ethHash"`
	ID              int    `json:"id"`
	TransactionHash string `json:"txHash"`
}

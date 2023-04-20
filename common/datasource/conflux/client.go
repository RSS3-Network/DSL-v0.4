package conflux

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"net/http"

	"github.com/ethereum/go-ethereum/common/hexutil"
)

const (
	Endpoint = "https://main.confluxrpc.com"
	MaxCount = 1000
)

type Client struct {
	httpClient *http.Client
}

func (c *Client) GetBlockTransactions(ctx context.Context, parameter GetBlockTransactionsParameter) (*ConfluxBlock, error) {
	number := hexutil.EncodeUint64(uint64(parameter.BlockNumber))
	jsonRequest := JsonRPCRequest{
		Jsonrpc: "2.0",
		Method:  "cfx_getBlockByEpochNumber",
		Params:  []any{number, true},
		Id:      67,
	}
	reqBytes, err := json.Marshal(jsonRequest)
	if err != nil {
		return nil, err
	}
	resp, err := c.httpClient.Post(Endpoint, "application/json", bytes.NewReader(reqBytes))
	if err != nil {
		return nil, err
	}
	respBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	_ = resp.Body.Close()
	var result BlockEpochResp
	err = json.Unmarshal(respBytes, &result)
	if err != nil {
		return nil, err
	}
	for i := range result.Result.Transactions {
		receiptInfo, err := c.GetTransactionReceipt(ctx, result.Result.Transactions[i].Hash)
		if err != nil {
			return nil, err
		}
		result.Result.Transactions[i].Receipt = *receiptInfo
	}

	return &result.Result, nil
}

func (c *Client) GetTransactionReceipt(ctx context.Context, txHash string) (*ConfluxTransactionReceipt, error) {
	jsonRequest := JsonRPCRequest{
		Jsonrpc: "2.0",
		Method:  "cfx_getTransactionReceipt",
		Params:  []any{txHash},
		Id:      67,
	}
	reqBytes, err := json.Marshal(jsonRequest)
	if err != nil {
		return nil, err
	}
	resp, err := c.httpClient.Post(Endpoint, "application/json", bytes.NewReader(reqBytes))
	if err != nil {
		return nil, err
	}
	respBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	_ = resp.Body.Close()
	var result ConfluxTransactionReceiptResp
	err = json.Unmarshal(respBytes, &result)
	if err != nil {
		return nil, err
	}

	return &result.Result, nil
}

func NewClient() *Client {
	return &Client{
		httpClient: http.DefaultClient,
	}
}

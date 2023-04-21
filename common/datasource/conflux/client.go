package conflux

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/google/go-querystring/query"
	"github.com/naturalselectionlabs/pregod/common/utils/httpx"
)

const (
	ScanEndpoint = "api.confluxscan.net"
	NodeEndpoint = "https://main.confluxrpc.com"
	MaxCount     = 100
)

type Client struct {
	httpClient *http.Client
}

func (c *Client) GetAccountTransactions(ctx context.Context, parameter GetAccountTxParameter) (*ConfluxScanAccountTxResp, error) {
	if parameter.Limit > MaxCount {
		parameter.Limit = MaxCount
	}
	parameter.Sort = "DESC"
	parameter.WithInput = false

	values, err := query.Values(parameter)
	if err != nil {
		return nil, err
	}

	url := url.URL{
		Scheme:   "https",
		Host:     ScanEndpoint,
		Path:     "/account/transactions",
		RawQuery: values.Encode(),
	}
	request, err := http.NewRequest(http.MethodGet, url.String(), nil)
	if err != nil {
		return nil, err
	}

	var result ConfluxScanAccountTxResp
	err = httpx.DoRequest(ctx, c.httpClient, request, &result)
	if err != nil {
		return nil, err
	}
	if result.Code != 0 || result.Message != "OK" {
		return nil, fmt.Errorf("confluxscan api error: %s", result.Message)
	}
	return &result, nil
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
	resp, err := c.httpClient.Post(NodeEndpoint, "application/json", bytes.NewReader(reqBytes))
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
	resp, err := c.httpClient.Post(NodeEndpoint, "application/json", bytes.NewReader(reqBytes))
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

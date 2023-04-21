package conflux

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"

	"github.com/naturalselectionlabs/pregod/common/utils/loggerx"
	"go.uber.org/zap"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/google/go-querystring/query"
	"github.com/naturalselectionlabs/pregod/common/utils/httpx"
)

const (
	ScanEndpoint = "api.confluxscan.net"
	NodeEndpoint = "https://main.confluxrpc.com"
	MaxCount     = 100
	// prevent rate limit
	DefaultCount = 10
)

// https://api.confluxscan.net/doc
const (
	TransferTypeTransaction = "transaction"
)

type Client struct {
	httpClient *http.Client
}

func (c *Client) GetAccountTransactions(ctx context.Context, parameter GetAccountParameter) (*ConfluxScanResp[ConfluxScanTxBrief], error) {
	if parameter.Limit > MaxCount {
		parameter.Limit = MaxCount
	}
	parameter.Sort = "DESC"
	parameter.TransferType = TransferTypeTransaction

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

	var result ConfluxScanResp[ConfluxScanTxBrief]
	err = httpx.DoRequest(ctx, c.httpClient, request, &result)
	if err != nil {
		return nil, err
	}
	if result.Code != 0 || result.Message != "OK" {
		return nil, fmt.Errorf("confluxscan api error: %s", result.Message)
	}
	return &result, nil
}

func (c *Client) GetAccountTransfers(ctx context.Context, parameter GetAccountParameter) (*ConfluxScanResp[ConfluxScanTransferBrief], error) {
	parameter.Sort = "DESC"
	parameter.TransferType = TransferTypeTransaction
	values, err := query.Values(parameter)
	if err != nil {
		return nil, err
	}

	url := url.URL{
		Scheme:   "https",
		Host:     ScanEndpoint,
		Path:     "/account/transfers",
		RawQuery: values.Encode(),
	}
	request, err := http.NewRequest(http.MethodGet, url.String(), nil)
	if err != nil {
		return nil, err
	}

	var result ConfluxScanResp[ConfluxScanTransferBrief]
	err = httpx.DoRequest(ctx, c.httpClient, request, &result)
	if err != nil {
		return nil, err
	}
	if result.Code != 0 || result.Message != "OK" {
		return nil, fmt.Errorf("confluxscan api error: %s", result.Message)
	}
	if len(result.Data.List) == 100 {
		// fetch 500 native token transfers
		result.Data.List = c.fetchAccountLatestTransfers(ctx, result.Data.List, url, parameter)
	}
	result.Data.List = c.filterKnownError(result.Data.List)
	return &result, nil
}

func (c *Client) filterKnownError(transfers []ConfluxScanTransferBrief) []ConfluxScanTransferBrief {
	result := make([]ConfluxScanTransferBrief, 0)
	for i := range transfers {
		if transfers[i].TransactionHash == "0xundefined" {
			continue
		}
		result = append(result, transfers[i])
	}
	return result
}

func (c *Client) fetchAccountLatestTransfers(ctx context.Context, result []ConfluxScanTransferBrief, url url.URL, parameter GetAccountParameter) []ConfluxScanTransferBrief {
	for i := 0; i < 4; i++ {
		// limit 100
		if len(result)%100 != 0 {
			return result
		}
		parameter.MinEpochNumber = 0
		parameter.MaxEpochNumber = result[len(result)-1].EpochNumber - 1
		parameter.Limit = 100
		parameter.Sort = "DESC"
		parameter.TransferType = TransferTypeTransaction

		values, err := query.Values(parameter)
		if err != nil {
			loggerx.Global().Error("fetchAccountLatestTransfers error", zap.Error(err), zap.Any("parameter", parameter))
			return result
		}

		url.RawQuery = values.Encode()
		var temp ConfluxScanResp[ConfluxScanTransferBrief]

		request, err := http.NewRequest(http.MethodGet, url.String(), nil)
		if err != nil {
			loggerx.Global().Error("fetchAccountLatestTransfers error", zap.Error(err), zap.Any("url", url.String()))
			return result
		}

		err = httpx.DoRequest(ctx, c.httpClient, request, &temp)
		if err != nil {
			loggerx.Global().Error("fetchAccountLatestTransfers error", zap.Error(err), zap.Any("url", url.String()))
			return result
		}

		if temp.Code != 0 || temp.Message != "OK" {
			loggerx.Global().Error("fetchAccountLatestTransfers error", zap.Error(err), zap.Any("message", temp))
			return result
		}
		result = append(result, temp.Data.List...)
	}
	return result
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

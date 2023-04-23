package conflux

import (
	"context"
	"fmt"
	"time"

	"github.com/go-resty/resty/v2"
	"github.com/naturalselectionlabs/pregod/common/utils/loggerx"
	"go.uber.org/zap"

	"github.com/ethereum/go-ethereum/common/hexutil"
)

const (
	ScanEndpoint = "https://api.confluxscan.net"
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
	restyClient *resty.Client
}

func (c *Client) GetAccountTransactions(ctx context.Context, parameter GetAccountParameter) (*ConfluxScanResp[ConfluxScanTxBrief], error) {
	var result ConfluxScanResp[ConfluxScanTxBrief]

	_, err := c.confluxScanRequest(ctx, parameter).SetResult(&result).Get(ScanEndpoint + "/account/transactions")
	if err != nil {
		return nil, err
	}

	if result.Code != 0 || result.Message != "OK" {
		return nil, fmt.Errorf("confluxscan api error: %s", result.Message)
	}

	return &result, nil
}

func (c *Client) GetAccountCfxTransfers(ctx context.Context, parameter GetAccountParameter) (*ConfluxScanResp[ConfluxScanTransferBrief], error) {
	var result ConfluxScanResp[ConfluxScanTransferBrief]

	_, err := c.confluxScanRequest(ctx, parameter).SetResult(&result).Get(ScanEndpoint + "/account/cfx/transfers")
	if err != nil {
		return nil, err
	}

	if result.Code != 0 || result.Message != "OK" {
		return nil, fmt.Errorf("confluxscan api error: %s", result.Message)
	}
	if len(result.Data.List) == 100 {
		// fetch 500 native token transfers
		result.Data.List = c.fetchAccountLatestTransfers(ctx, result.Data.List, parameter)
	}
	result.Data.List = c.filterKnownError(result.Data.List)
	return &result, nil
}

func (c *Client) confluxScanRequest(ctx context.Context, paramster GetAccountParameter) *resty.Request {
	r := c.restyClient.R().SetContext(ctx)
	if paramster.MaxEpochNumber != 0 {
		r.SetQueryParam("maxEpochNumber", fmt.Sprintf("%d", paramster.MaxEpochNumber))
	}
	if paramster.MinEpochNumber != 0 {
		r.SetQueryParam("minEpochNumber", fmt.Sprintf("%d", paramster.MinEpochNumber))
	}
	if paramster.Limit != 0 {
		r.SetQueryParam("limit", fmt.Sprintf("%d", paramster.Limit))
	}
	if paramster.Sort == "" {
		paramster.Sort = "DESC"
	}
	if paramster.TransferType == "" {
		paramster.TransferType = TransferTypeTransaction
	}
	r.SetQueryParams(map[string]string{
		"account":      paramster.Address,
		"transferType": paramster.TransferType,
		"sort":         paramster.Sort,
		"from":         paramster.Address,
	})
	return r
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

func (c *Client) fetchAccountLatestTransfers(ctx context.Context, result []ConfluxScanTransferBrief, parameter GetAccountParameter) []ConfluxScanTransferBrief {
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

		var temp ConfluxScanResp[ConfluxScanTransferBrief]

		_, err := c.confluxScanRequest(ctx, parameter).SetResult(&temp).Get(ScanEndpoint + "/account/cfx/transfers")
		if err != nil {
			loggerx.Global().Error("fetchAccountLatestTransfers error", zap.Error(err), zap.Any("parameter", parameter))
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
	var result BlockEpochResp

	_, err := c.restyClient.R().SetContext(ctx).SetBody(jsonRequest).SetResult(&result).Post(NodeEndpoint)
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
	var result ConfluxTransactionReceiptResp

	_, err := c.restyClient.R().SetContext(ctx).SetBody(jsonRequest).SetResult(&result).Post(NodeEndpoint)
	if err != nil {
		return nil, err
	}

	return &result.Result, nil
}

func NewClient() *Client {
	return &Client{
		restyClient: resty.New().SetRetryCount(10).SetRetryAfter(func(c *resty.Client, r *resty.Response) (time.Duration, error) {
			loggerx.Global().Warn("retry request conflux scan api", zap.String("url", r.String()))
			return time.Millisecond * 500, nil
		}),
	}
}

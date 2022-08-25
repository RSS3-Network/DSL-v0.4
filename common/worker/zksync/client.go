package zksync

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"

	"github.com/ethereum/go-ethereum/common"
	"github.com/google/go-querystring/query"
	"github.com/naturalselectionlabs/pregod/common/database"
	"github.com/naturalselectionlabs/pregod/common/database/model"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

const (
	Scheme   = "https"
	Endpoint = "api.zksync.io"

	FromLatest = "latest"

	DirectionNewer = "newer"
	DirectionOlder = "older"

	KindERC20 = "ERC20"
)

type Client struct {
	httpClient *http.Client
}

func (c *Client) DoRequest(_ context.Context, request *http.Request) (*Response, *http.Response, error) {
	httpResponse, err := c.httpClient.Do(request)
	if err != nil {
		return nil, nil, err
	}

	defer func() {
		_ = httpResponse.Body.Close()
	}()

	response := &Response{}

	if err := json.NewDecoder(httpResponse.Body).Decode(&response); err != nil {
		return nil, httpResponse, err
	}

	return response, httpResponse, nil
}

type GetTokenList []GetTokenListItem

type GetTokenListItem struct {
	ID       int    `json:"id"`
	Address  string `json:"address"`
	Symbol   string `json:"symbol"`
	Decimals uint8  `json:"decimals"`
	Kind     string `json:"kind"`
	IsNFT    bool   `json:"is_nft"`
}

func (c *Client) GetTokenList() (GetTokenList, error) {
	requestURL := &url.URL{
		Scheme: Scheme,
		Host:   Endpoint,
		Path:   "/api/v0.1/tokens",
	}

	httpResponse, err := c.httpClient.Get(requestURL.String())
	if err != nil {
		return nil, err
	}

	defer func() {
		_ = httpResponse.Body.Close()
	}()

	tokenList := GetTokenList{}

	if err := json.NewDecoder(httpResponse.Body).Decode(&tokenList); err != nil {
		return nil, err
	}

	return tokenList, nil
}

type GetAddressTransactionsOption struct {
	From      string `url:"from"`
	Limit     int64  `url:"limit"`
	Direction string `url:"direction"`
}

func (c *Client) GetAddressTransactions(ctx context.Context, address common.Address, from string, limit int64, direction string) (*GetAccountTransactionList, *Response, error) {
	values, err := query.Values(&GetAddressTransactionsOption{
		From:      from,
		Limit:     limit,
		Direction: direction,
	})
	if err != nil {
		return nil, nil, err
	}

	requestURL := &url.URL{
		Scheme:   Scheme,
		Host:     Endpoint,
		Path:     fmt.Sprintf("/api/v0.2/accounts/%s/transactions", address.String()),
		RawQuery: values.Encode(),
	}

	httpRequest, err := http.NewRequest(http.MethodGet, requestURL.String(), nil)
	if err != nil {
		return nil, nil, err
	}

	response, _, err := c.DoRequest(ctx, httpRequest)
	if err != nil {
		return nil, nil, err
	}

	transactionList := GetAccountTransactionList{}

	if err := json.Unmarshal(response.Result, &transactionList); err != nil {
		return nil, nil, err
	}

	return &transactionList, response, nil
}

// https://docs.zksync.io/apiv02-docs/#transactions-api-v0.2-transactions-txhash-data
func (c *Client) GetTransactionData(ctx context.Context, transactionHash common.Hash) (*GetTransactionData, *Response, error) {
	requestURL := &url.URL{
		Scheme: Scheme,
		Host:   Endpoint,
		Path:   fmt.Sprintf("/api/v0.2/transactions/%s/data", transactionHash.String()),
	}

	httpRequest, err := http.NewRequest(http.MethodGet, requestURL.String(), nil)
	if err != nil {
		return nil, nil, err
	}

	response, _, err := c.DoRequest(ctx, httpRequest)
	if err != nil {
		return nil, nil, err
	}

	transactionData := GetTransactionData{}

	if err := json.Unmarshal(response.Result, &transactionData); err != nil {
		return nil, nil, err
	}

	return &transactionData, response, nil
}

// https://docs.zksync.io/apiv02-docs/#tokens-api-v0.2-tokens-tokenlike
func (c *Client) GetToken(ctx context.Context, tokenID uint) (*model.GetTokenInfo, *Response, error) {
	tokenInfo := model.GetTokenInfo{}

	// first try to get from db
	if err := database.Global().Where(
		"id = ?", tokenID,
	).First(&tokenInfo).Error; err != nil && err != gorm.ErrRecordNotFound {
		logrus.Error(err)
		return nil, nil, err
	} else if err == nil { // exists
		return &tokenInfo, nil, nil
	} // else not found, continue

	requestURL := &url.URL{
		Scheme: Scheme,
		Host:   Endpoint,
		Path:   fmt.Sprintf("/api/v0.2/tokens/%v", tokenID),
	}

	httpRequest, err := http.NewRequest(http.MethodGet, requestURL.String(), nil)
	if err != nil {
		return nil, nil, err
	}

	response, _, err := c.DoRequest(ctx, httpRequest)
	if err != nil {
		return nil, nil, err
	}

	if err := json.Unmarshal(response.Result, &tokenInfo); err != nil {
		return nil, nil, err
	}

	// save
	if err := database.Global().Clauses(clause.OnConflict{
		UpdateAll: true,
	}).Create(&tokenInfo).Error; err != nil {
		logrus.Error(err)
		return nil, nil, err
	}

	return &tokenInfo, response, nil
}

// https://docs.zksync.io/apiv02-docs/#tokens-api-v0.2-tokens-nft-id
func (c *Client) GetNFTToken(ctx context.Context, tokenID uint) (*model.GetNFTTokenInfo, *Response, error) {
	tokenInfo := model.GetNFTTokenInfo{}

	// first try to get from db
	if err := database.Global().Where(
		"id = ?", tokenID,
	).First(&tokenInfo).Error; err != nil && err != gorm.ErrRecordNotFound {
		logrus.Error(err)
		return nil, nil, err
	} else if err == nil { // exists
		return &tokenInfo, nil, nil
	} // else not found, continue

	requestURL := &url.URL{
		Scheme: Scheme,
		Host:   Endpoint,
		Path:   fmt.Sprintf("/api/v0.2/tokens/nft/%v", tokenID),
	}

	httpRequest, err := http.NewRequest(http.MethodGet, requestURL.String(), nil)
	if err != nil {
		return nil, nil, err
	}

	response, _, err := c.DoRequest(ctx, httpRequest)
	if err != nil {
		return nil, nil, err
	}

	if err := json.Unmarshal(response.Result, &tokenInfo); err != nil {
		return nil, nil, err
	}

	// save
	if err := database.Global().Clauses(clause.OnConflict{
		UpdateAll: true,
	}).Create(&tokenInfo).Error; err != nil {
		logrus.Error(err)
		return nil, nil, err
	}

	return &tokenInfo, response, nil
}

func New() *Client {
	return &Client{
		httpClient: http.DefaultClient,
	}
}

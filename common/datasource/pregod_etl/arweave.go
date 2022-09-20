package pregod_etl

import (
	"context"
	"net/url"

	"github.com/go-resty/resty/v2"
	"github.com/google/go-querystring/query"
	"github.com/sirupsen/logrus"
)

func (c *Client) GetArweaveTransactions(ctx context.Context, parameter GetArweaveTransactionsRequest) (*GetArweaveTransactionsResponse, error) {
	values, err := query.Values(parameter)
	if err != nil {
		logrus.Errorf("[pregod_etl client] GetArweaveTransactions: query.Values error, %v", err)

		return nil, err
	}

	url := &url.URL{
		Scheme:   "http",
		Host:     c.endpoint,
		Path:     "/networks/arweave/transactions",
		RawQuery: values.Encode(),
	}

	logrus.Info("[pregod_etl client] GetArweaveTransactions, request = ", url.String())

	client := resty.New()
	result := GetArweaveTransactionsResponse{}

	_, err = client.R().SetResult(&result).Get(url.String())
	if err != nil {
		logrus.Errorf("[token] EnsToNFT url: %v, err: %v", url, err)
		return nil, err
	}

	return &result, nil
}

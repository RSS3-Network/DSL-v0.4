package pregod_etl

import (
	"context"
	"net/http"
	"net/url"

	"github.com/google/go-querystring/query"
	"github.com/naturalselectionlabs/pregod/common/utils/httpx"
	"github.com/sirupsen/logrus"
)

func (c *Client) GetArweaveTransactions(ctx context.Context, parameter GetArweaveTransactionsRequest) (*GetArweaveTransactionsResponse, error) {
	values, err := query.Values(parameter)
	if err != nil {
		logrus.Errorf("[pregod_etl client] GetArweaveTransactions: query.Values error, %v", err)

		return nil, err
	}

	url := &url.URL{
		Scheme:   "https",
		Host:     c.endpoint,
		Path:     "/networks/arweave/transactions",
		RawQuery: values.Encode(),
	}

	logrus.Info("[pregod_etl client] GetArweaveTransactions, request = ", url.String())
	request, err := httpx.NewRequest(http.MethodGet, url.String(), nil)
	if err != nil {
		logrus.Errorf("[pregod_etl client] GetArweaveTransactions: NewRequest error, %v", err)

		return nil, err
	}

	var result *GetArweaveTransactionsResponse

	err = httpx.DoRequest(ctx, c.httpClient, request, &result)
	if err != nil {
		logrus.Errorf("[pregod_etl client] GetArweaveTransactions: DoRequest error, %v", err)

		return nil, err
	}

	return result, nil
}

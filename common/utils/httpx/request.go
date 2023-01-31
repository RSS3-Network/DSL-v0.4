package httpx

import (
	"context"
	"encoding/json"
	"io"
	"net/http"

	"github.com/naturalselectionlabs/pregod/common/utils/loggerx"
	"go.uber.org/zap"

	"github.com/naturalselectionlabs/pregod/common/protocol"
)

func NewRequest(method, rawURL string, body interface{}) (*http.Request, error) {
	var readWriter io.ReadWriter

	if body != nil {
		if err := json.NewEncoder(readWriter).Encode(body); err != nil {
			return nil, err
		}
	}

	request, err := http.NewRequest(method, rawURL, readWriter)
	if err != nil {
		return nil, err
	}

	request.Header.Set("Accept", protocol.ContentTypeJSON)

	return request, nil
}

func DoRequest(_ context.Context, client *http.Client, request *http.Request, response interface{}) error {
	httpResponse, err := client.Do(request)
	if err != nil {
		return err
	}

	reader := httpResponse.Body

	data, err := io.ReadAll(reader)
	if err != nil {
		return err
	}

	_ = reader.Close()

	if err = json.Unmarshal(data, &response); err != nil {
		loggerx.Global().Error("DoRequest error", zap.Error(err), zap.Any("data", data))

		return err
	}

	return nil
}

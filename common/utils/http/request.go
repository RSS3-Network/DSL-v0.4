package http

import (
	"context"
	"encoding/json"
	"io"
	"net/http"

	"github.com/sirupsen/logrus"
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

	request.Header.Set("Accept", "application/json")

	return request, nil
}

func DoRequest(_ context.Context, client *http.Client, request *http.Request, response interface{}) error {
	httpResponse, err := client.Do(request)
	if err != nil {
		logrus.Errorf("DoRequest: client.Do, request: %v,  error: %v", request.URL, err)
		return err
	}

	reader := httpResponse.Body

	data, err := io.ReadAll(reader)
	if err != nil {
		logrus.Errorf("DoRequest: io ReadAll error, request: %v, error: %v", request.URL, err)

		return err
	}

	_ = reader.Close()

	if err = json.Unmarshal(data, &response); err != nil {
		logrus.Errorf("DoRequest: json unmarshal error, request: %v, response: %v, error: %v", request.URL, string(data), err)
		return err
	}

	return nil
}

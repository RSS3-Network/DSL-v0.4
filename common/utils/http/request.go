package http

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
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
		return err
	}

	defer func() {
		_ = httpResponse.Body.Close()
	}()

	if err = json.NewDecoder(httpResponse.Body).Decode(&response); err != nil {
		return err
	}

	return nil
}

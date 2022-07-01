package ipfs

import (
	"context"
	"io"
	"net/http"
	"strings"
)

func GetFileByURL(ctx context.Context, url string) ([]byte, error) {
	response, err := http.Get(strings.Replace(url, "ipfs://", "https://rss3.infura-ipfs.io/ipfs/", 1))
	if err != nil {
		return nil, err
	}

	defer func() {
		_ = response.Body.Close()
	}()

	return io.ReadAll(response.Body)
}

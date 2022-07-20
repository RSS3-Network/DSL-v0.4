package ipfs

import (
	"context"
	"io"
	"net/http"
	"strings"
)

func GetDirectURL(ctx context.Context, url string) string {
	return strings.Replace(url, "ipfs://", "https://ipfs.rss3.page/ipfs/", 1)
}

func GetFileByURL(ctx context.Context, url string) ([]byte, error) {
	response, err := http.Get(GetDirectURL(ctx, url))
	if err != nil {
		return nil, err
	}

	defer func() {
		_ = response.Body.Close()
	}()

	return io.ReadAll(response.Body)
}

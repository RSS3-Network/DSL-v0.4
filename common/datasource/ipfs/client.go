package ipfs

import (
	"context"
	"errors"
	"io"
	"net/http"
	"strings"
	"time"
)

var IPFSTimeoutError = errors.New("ipfs timeout")

func GetDirectURL(url string) string {
	if s := strings.Split(url, "/ipfs/"); len(s) == 2 {
		url = "https://ipfs.rss3.page/ipfs/" + s[1]
	}

	return strings.Replace(url, "ipfs://", "https://ipfs.rss3.page/ipfs/", 1)
}

func GetFileByURL(url string) ([]byte, error) {
	var body []byte
	var err error
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(time.Millisecond*800))
	defer cancel()

	go func(ctx context.Context) {
		response, e := http.Get(GetDirectURL(url))
		if e != nil {
			err = e
			return
		}

		defer func() {
			_ = response.Body.Close()
		}()

		body, err = io.ReadAll(response.Body)
	}(ctx)

	select {
	case <-ctx.Done():
		return body, err
	case <-time.After(time.Duration(time.Second * 10)):
		return nil, IPFSTimeoutError
	}
}

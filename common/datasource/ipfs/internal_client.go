package ipfs

import (
	"context"
	"io"
	"net/http"
	"strings"
	"time"

	"github.com/sirupsen/logrus"
)

type InternalClient struct {
	internalIPFS string
}

func New() *InternalClient {
	return &InternalClient{
		internalIPFS: "https://ipfs.rss3.page/ipfs/",
	}
}

func (c *InternalClient) GetDirectURL(url string) string {
	if s := strings.Split(url, "/ipfs/"); len(s) == 2 {
		url = c.internalIPFS + s[1]
	}

	return strings.Replace(url, "ipfs://", c.internalIPFS, 1)
}

func (c *InternalClient) GetFileByURL(url string) ([]byte, error) {
	var body []byte
	var err error
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(time.Second*10))
	defer cancel()

	go func(ctx context.Context) {
		response, e := http.Get(c.GetDirectURL(url))
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
		logrus.Errorf("[ipfs] GetFileByURL timeout, url = %v", c.GetDirectURL(url))
		return nil, IPFSTimeoutError
	}
}

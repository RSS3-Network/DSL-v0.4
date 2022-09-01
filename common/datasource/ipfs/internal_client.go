package ipfs

import (
	"io"
	"net/http"
	"strings"
)

type InternalClient struct {
	internalIPFS string
}

func New() *InternalClient {
	return &InternalClient{
		// internalIPFS: "http://ipfs-cluster.pregod:8080/ipfs/",
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
	response, err := http.Get(GetDirectURL(url))
	if err != nil {
		return nil, err
	}

	defer func() {
		_ = response.Body.Close()
	}()

	return io.ReadAll(response.Body)
}

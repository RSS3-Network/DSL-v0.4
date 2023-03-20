package metadata_url

import (
	"errors"
	"io"
	"net/http"
	"strings"
)

var IPFSTimeoutError = errors.New("ipfs timeout")

var IPFSClient *InternalClient

type InternalClient struct {
	internalIPFS string
}

func New(ipfs string) {
	IPFSClient = &InternalClient{
		internalIPFS: ipfs,
	}
}

func GetDirectURL(directURL string) string {
	s := strings.Split(directURL, "/ipfs/")

	switch {
	case len(s) == 2:
		directURL = IPFSClient.internalIPFS + s[1]
	case strings.HasPrefix(directURL, "ar://"):
		directURL = strings.Replace(directURL, "ar://", "https://arweave.net/", 1)
	case strings.HasPrefix(directURL, "ipfs://"):
		directURL = strings.Replace(directURL, "ipfs://", IPFSClient.internalIPFS, 1)
	}

	return directURL
}

func GetFileByURL(url string) ([]byte, error) {
	response, err := http.Get(GetDirectURL(url))
	if err != nil {
		return nil, err
	}

	defer func() {
		_ = response.Body.Close()
	}()

	return io.ReadAll(response.Body)
}

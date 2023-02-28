package metadata_url

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
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
	if s := strings.Split(directURL, "/ipfs/"); len(s) == 2 {
		directURL = IPFSClient.internalIPFS + s[1]
	}

	contentURL, err := url.Parse(directURL)
	if err != nil {
		return directURL
	}

	switch contentURL.Scheme {
	case "ar":
		directURL = fmt.Sprintf("https://arweave.net/%s", contentURL.Host)
	case "ipfs":
		directURL = IPFSClient.internalIPFS + contentURL.Host
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

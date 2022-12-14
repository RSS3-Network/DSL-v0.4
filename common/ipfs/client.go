package ipfs

import (
	"errors"
	"fmt"
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

func GetDirectURL(url string) string {
	if s := strings.Split(url, "/ipfs/"); len(s) == 2 {
		url = IPFSClient.internalIPFS + s[1]
	}

	return strings.Replace(url, "ipfs://", IPFSClient.internalIPFS, 1)
}

func BuildURL(cid string) string {
	return fmt.Sprintf("%s/%s", strings.TrimRight(IPFSClient.internalIPFS, "/"), cid)
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

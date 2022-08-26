package ipfs

import (
	"io"
	"net/http"
	"strings"
)

func GetDirectURL(url string) string {
	if s := strings.Split(url, "/ipfs/"); len(s) == 2 {
		url = "https://ipfs.rss3.page/ipfs/" + s[1]
	}

	return strings.Replace(url, "ipfs://", "https://ipfs.rss3.page/ipfs/", 1)
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

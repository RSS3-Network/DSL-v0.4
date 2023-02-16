package nft

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/shopspring/decimal"
)

var (
	CryptoPunksName = "CryptoPunks"
	CryptoPunksDes  = "CryptoPunks launched as a fixed set of 10,000 items in mid-2017 and became one of the inspirations for the ERC-721 standard. They have been featured in places like The New York Times, Christie’s of London, Art|Basel Miami, and The PBS NewsHour."
)

type CkMetadata struct {
	Name  string `json:"name"`
	Bio   string `json:"bio"`
	Image string `json:"image_url_png"`
}

func GetCkMetadata(ctx context.Context, tokenId decimal.Decimal) (*CkMetadata, error) {
	endpoint := fmt.Sprintf("%s%s", "https://api.cryptokitties.co/v3/kitties/", tokenId)

	request, err := http.NewRequestWithContext(ctx, http.MethodGet, endpoint, nil)
	if err != nil {
		return nil, fmt.Errorf("new request: %w", err)
	}

	// request.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/89.0.4389.90 Safari/537.36")
	httpClient := &http.Client{}

	response, err := httpClient.Do(request)
	if err != nil {
		return nil, fmt.Errorf("do request: %w", err)
	}

	defer func() {
		_ = response.Body.Close()
	}()
	ckMetadata := &CkMetadata{}
	err = json.NewDecoder(response.Body).Decode(ckMetadata)
	if err != nil {
		return nil, err
	}

	return ckMetadata, nil
}

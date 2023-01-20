package alchemy

import "time"

type GetAssetTransfersParameter struct {
	FromBlock         string   `json:"fromBlock,omitempty"`
	ToBlock           string   `json:"toBlock,omitempty"`
	FromAddress       string   `json:"fromAddress,omitempty"`
	ToAddress         string   `json:"toAddress,omitempty"`
	ContractAddresses []string `json:"contractAddresses,omitempty"`
	MaxCount          string   `json:"maxCount,omitempty"`
	ExcludeZeroValue  bool     `json:"excludeZeroValue"`
	PageKey           string   `json:"pageKey,omitempty"`
	WithMetadata      bool     `json:"withMetadata,omitempty"`
	Category          []string `json:"category,omitempty"`
	Order             string   `json:"order,omitempty"`
}

type GetAssetTransfersResult struct {
	Transfers []struct {
		BlockNum        string      `json:"blockNum"`
		Hash            string      `json:"hash"`
		From            string      `json:"from"`
		To              string      `json:"to"`
		Value           float64     `json:"value"`
		Erc721TokenId   interface{} `json:"erc721TokenId"`
		Erc1155Metadata interface{} `json:"erc1155Metadata"`
		TokenId         interface{} `json:"tokenId"`
		Asset           string      `json:"asset"`
		Category        string      `json:"category"`
		RawContract     struct {
			Value   string      `json:"value"`
			Address interface{} `json:"address"`
			Decimal string      `json:"decimal"`
		}
	}
	PageKey string `json:"pageKey"`
}

type GetNFTsParameter struct {
	Owner             string   `url:"owner,omitempty"`
	PageKey           string   `url:"pageKey,omitempty"`
	ContractAddresses []string `url:"contractAddresses,omitempty"`
	WithMetadata      bool     `url:"withMetadata"`
}

type GetNFTsResult struct {
	OwnedNFTs  []NFTResult `json:"ownedNfts"`
	PageKey    string      `json:"pageKey"`
	TotalCount int         `json:"totalCount"`
	BlockHash  string      `json:"blockHash"`
}

type NFTResult struct {
	Contract struct {
		Address string `json:"address"`
	} `json:"contract"`
	ID struct {
		TokenID       string `json:"tokenId"`
		TokenMetadata struct {
			TokenType string `json:"tokenType"`
		} `json:"tokenMetadata"`
	} `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	TokenURI    struct {
		Raw     string `json:"raw"`
		Gateway string `json:"gateway"`
	} `json:"tokenUri"`
	Media []struct {
		Raw     string `json:"raw"`
		Gateway string `json:"gateway"`
	} `json:"media"`
	Metadata struct {
		Image      string      `json:"image"`
		Attributes interface{} `json:"attributes"`
	} `json:"metadata"`
	TimeLastUpdated time.Time `json:"timeLastUpdated"`
}

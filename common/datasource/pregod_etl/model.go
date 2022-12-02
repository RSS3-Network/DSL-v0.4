package pregod_etl

import (
	"github.com/shopspring/decimal"
)

type GetLogsRequest struct {
	Address         string `url:"address,omitempty"`
	TopicFirst      string `url:"topic_first,omitempty"`
	TopicSecond     string `url:"topic_second,omitempty"`
	TopicThird      string `url:"topic_third,omitempty"`
	TopicFourth     string `url:"topic_fourth,omitempty"`
	TransactionHash string `url:"transaction_hash,omitempty"`
	BlockNumberFrom int64  `url:"block_number_from,omitempty"`
	BlockNumberTo   int64  `url:"block_number_to,omitempty"`
	Cursor          string `url:"cursor,omitempty"`
	Order           string `url:"order"`
	Limit           int    `url:"limit"`
}

type GetLogsResponse struct {
	Cursor string        `json:"cursor,omitempty"`
	Total  int64         `json:"total"`
	Result []EthereumLog `json:"result"`
}

// EthereumLog is the model entity for the EthereumLog schema.
type EthereumLog struct {
	// LogIndex holds the value of the "log_index" field.
	LogIndex decimal.Decimal `json:"log_index,omitempty"`
	// BlockNumber holds the value of the "block_number" field.
	BlockNumber decimal.Decimal `json:"block_number,omitempty"`
	// BlockHash holds the value of the "block_hash" field.
	BlockHash string `json:"block_hash,omitempty"`
	// TransactionHash holds the value of the "transaction_hash" field.
	TransactionHash string `json:"transaction_hash,omitempty"`
	// TransactionIndex holds the value of the "transaction_index" field.
	TransactionIndex decimal.Decimal `json:"transaction_index,omitempty"`
	// Address holds the value of the "address" field.
	Address string `json:"address,omitempty"`
	// TopicFirst holds the value of the "topic_first" field.
	TopicFirst string `json:"topic_first,omitempty"`
	// TopicSecond holds the value of the "topic_second" field.
	TopicSecond string `json:"topic_second,omitempty"`
	// TopicThird holds the value of the "topic_third" field.
	TopicThird string `json:"topic_third,omitempty"`
	// TopicFourth holds the value of the "topic_fourth" field.
	TopicFourth string `json:"topic_fourth,omitempty"`
	// Data holds the value of the "data" field.
	Data string `json:"data,omitempty"`
}

type GetArweaveTransactionsRequest struct {
	BlockHeightFrom string   `url:"block_height_from,omitempty"`
	BlockHeightTo   string   `url:"block_height_to,omitempty"`
	Owner           string   `url:"owner,omitempty"`
	Target          string   `url:"target,omitempty"`
	Tags            []string `url:"tags,omitempty"`
	Cursor          string   `url:"cursor,omitempty"`
	Order           string   `url:"order,omitempty"`
	Limit           int      `url:"limit,omitempty"`
}

type GetArweaveTransactionsResponse struct {
	Cursor string               `json:"cursor,omitempty"`
	Total  int64                `json:"total"`
	Result []ArweaveTransaction `json:"result"`
}

type ArweaveTransaction struct {
	// ID of the ent.
	ID string `json:"id,omitempty"`
	// Format holds the value of the "format" field.
	Format uint `json:"format,omitempty"`
	// Height holds the value of the "height" field.
	Height decimal.Decimal `json:"height,omitempty"`
	// LastTransaction holds the value of the "last_transaction" field.
	LastTransaction string `json:"last_transaction,omitempty"`
	// Owner holds the value of the "owner" field.
	Owner string `json:"owner,omitempty"`
	// Tags holds the value of the "tags" field.
	Tags []TagPlaintext `json:"tags,omitempty"`
	// Target holds the value of the "target" field.
	Target string `json:"target,omitempty"`
	// Quantity holds the value of the "quantity" field.
	Quantity string `json:"quantity,omitempty"`
	// DataRoot holds the value of the "data_root" field.
	DataRoot string `json:"data_root,omitempty"`
	// DataSize holds the value of the "data_size" field.
	DataSize string `json:"data_size,omitempty"`
	// Data holds the value of the "data" field.
	Data string `json:"data,omitempty"`
	// Reward holds the value of the "reward" field.
	Reward string `json:"reward,omitempty"`
	// Signature holds the value of the "signature" field.
	Signature string `json:"signature,omitempty"`
}

type TagPlaintext tagMarshal

type tagMarshal struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

type ArweaveData struct {
	Wnft struct {
		Fee              decimal.Decimal `json:"fee"`
		Name             string          `json:"name"`
		Nonce            decimal.Decimal `json:"nonce"`
		Owner            string          `json:"owner"`
		Price            decimal.Decimal `json:"price"`
		Supply           decimal.Decimal `json:"supply"`
		Symbol           string          `json:"symbol"`
		ChainId          decimal.Decimal `json:"chainId"`
		ImageURI         string          `json:"imageURI"`
		Renderer         string          `json:"renderer"`
		Description      string          `json:"description"`
		MediaAssetId     decimal.Decimal `json:"mediaAssetId"`
		ProxyAddress     string          `json:"proxyAddress"`
		FundingRecipient string          `json:"fundingRecipient"`
	} `json:"wnft"`
	Digest  string `json:"digest"`
	Content struct {
		Body      string          `json:"body"`
		Title     string          `json:"title"`
		Timestamp decimal.Decimal `json:"timestamp"`
	} `json:"content"`
	Version    string `json:"version"`
	Authorship struct {
		Algorithm struct {
			Hash string `json:"hash"`
			Name string `json:"name"`
		} `json:"algorithm"`
		Signature           string `json:"signature"`
		SigningKey          string `json:"signingKey"`
		Contributor         string `json:"contributor"`
		SigningKeyMessage   string `json:"signingKeyMessage"`
		SigningKeySignature string `json:"signingKeySignature"`
	} `json:"authorship"`
}

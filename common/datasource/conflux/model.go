package conflux

type JsonRPCRequest struct {
	Jsonrpc string `json:"jsonrpc"`
	Method  string `json:"method"`
	Params  []any  `json:"params"`
	Id      int    `json:"id"`
}

type GetAccountParameter struct {
	Address        string `json:"account" url:"account,omitempty"`
	MinEpochNumber int64  `json:"minEpochNumber" url:"minEpochNumber,omitempty"`
	MaxEpochNumber int64  `json:"maxEpochNumber" url:"maxEpochNumber,omitempty"`
	Limit          int64  `json:"limit" url:"limit,omitempty"`
	Sort           string `json:"sort" url:"sort,omitempty"`
	TransferType   string `json:"transferType" url:"transferType,omitempty"`
}

type GetBlockTransactionsParameter struct {
	BlockNumber int64
}

type Transaction struct {
	Hash             string                    `json:"hash"`
	Nonce            string                    `json:"nonce"`
	BlockHash        string                    `json:"blockHash"`
	TransactionIndex string                    `json:"transactionIndex"`
	From             string                    `json:"from"`
	To               string                    `json:"to"`
	Value            string                    `json:"value"`
	GasPrice         string                    `json:"gasPrice"`
	Gas              string                    `json:"gas"`
	ContractCreated  string                    `json:"contractCreated"`
	Data             string                    `json:"data"`
	StorageLimit     string                    `json:"storageLimit"`
	EpochHeight      string                    `json:"epochHeight"`
	ChainId          string                    `json:"chainId"`
	Status           string                    `json:"status"`
	V                string                    `json:"v"`
	R                string                    `json:"r"`
	S                string                    `json:"s"`
	Receipt          ConfluxTransactionReceipt `json:"-"`
}

type BlockEpochResp struct {
	Jsonrpc string       `json:"jsonrpc"`
	Id      int          `json:"id"`
	Result  ConfluxBlock `json:"result"`
}

type ConfluxBlock struct {
	Hash                  string        `json:"hash"`
	ParentHash            string        `json:"parentHash"`
	Height                string        `json:"height"`
	Miner                 string        `json:"miner"`
	DeferredStateRoot     string        `json:"deferredStateRoot"`
	DeferredReceiptsRoot  string        `json:"deferredReceiptsRoot"`
	DeferredLogsBloomHash string        `json:"deferredLogsBloomHash"`
	Blame                 string        `json:"blame"`
	TransactionsRoot      string        `json:"transactionsRoot"`
	EpochNumber           string        `json:"epochNumber"`
	BlockNumber           string        `json:"blockNumber"`
	GasLimit              string        `json:"gasLimit"`
	GasUsed               string        `json:"gasUsed"`
	Timestamp             string        `json:"timestamp"`
	Difficulty            string        `json:"difficulty"`
	PowQuality            string        `json:"powQuality"`
	RefereeHashes         []string      `json:"refereeHashes"`
	Adaptive              bool          `json:"adaptive"`
	Nonce                 string        `json:"nonce"`
	Size                  string        `json:"size"`
	Custom                []string      `json:"custom"`
	PosReference          string        `json:"posReference"`
	Transactions          []Transaction `json:"transactions"`
}

type ConfluxTransactionReceipt struct {
	TransactionHash string `json:"transactionHash"`
	Index           string `json:"index"`
	BlockHash       string `json:"blockHash"`
	EpochNumber     string `json:"epochNumber"`
	From            string `json:"from"`
	To              string `json:"to"`
	GasUsed         string `json:"gasUsed"`
	GasFee          string `json:"gasFee"`
}

type ConfluxTransactionReceiptResp struct {
	Jsonrpc string                    `json:"jsonrpc"`
	Id      int                       `json:"id"`
	Result  ConfluxTransactionReceipt `json:"result"`
}

type ConfluxScanResp[T any] struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    struct {
		Total int `json:"total"`
		List  []T `json:"list"`
	} `json:"data"`
}

type ConfluxScanTransferBrief struct {
	EpochNumber     int64  `json:"epochNumber"`
	BlockIndex      int64  `json:"blockIndex"`
	TxIndex         int64  `json:"txIndex"`
	TxLogIndex      int64  `json:"txLogIndex"`
	From            string `json:"from"`
	To              string `json:"to"`
	Type            string `json:"type"`
	Timestamp       int64  `json:"timestamp"`
	TransactionHash string `json:"transactionHash"`
	ChainId         int64  `json:"chainId"`
	Nonce           int64  `json:"nonce"`
	Method          string `json:"method"`
	Status          int64  `json:"status"`
	GasFee          string `json:"gasFee"`
	StorageFee      string `json:"storageFee"`
	Input           string `json:"input"`
	Amount          string `json:"amount"`
}

type ConfluxScanTxBrief struct {
	EpochNumber      int64  `json:"epochNumber"`
	BlockPosition    int64  `json:"blockPosition"`
	TransactionIndex int64  `json:"transactionIndex"`
	Nonce            string `json:"nonce"`
	Hash             string `json:"hash"`
	From             string `json:"from"`
	To               string `json:"to"`
	Value            string `json:"value"`
	GasPrice         string `json:"gasPrice"`
	GasFee           string `json:"gasFee"`
	Timestamp        int64  `json:"timestamp"`
	Status           int64  `json:"status"`
	ContractCreated  string `json:"contractCreated"`
	Method           string `json:"method"`
}

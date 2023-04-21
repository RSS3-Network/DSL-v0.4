package conflux

import "github.com/shopspring/decimal"

type JsonRPCRequest struct {
	Jsonrpc string `json:"jsonrpc"`
	Method  string `json:"method"`
	Params  []any  `json:"params"`
	Id      int    `json:"id"`
}

type GetAccountTxParameter struct {
	Address        string          `json:"account" url:"account,omitempty"`
	MinEpochNumber decimal.Decimal `json:"minEpochNumber" url:"minEpochNumber,omitempty"`
	Limit          int64           `json:"limit" url:"limit,omitempty"`
	Sort           string          `json:"sort" url:"sort,omitempty"`
	WithInput      bool            `json:"withInput" url:"withInput,omitempty"`
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
	// ContractCreated         interface{}   `json:"contractCreated"`
	// Logs                    []interface{} `json:"logs"`
	// LogsBloom               string        `json:"logsBloom"`
	// StateRoot               string        `json:"stateRoot"`
	// OutcomeStatus           string        `json:"outcomeStatus"`
	// TxExecErrorMsg          interface{}   `json:"txExecErrorMsg"`
	// GasCoveredBySponsor     bool          `json:"gasCoveredBySponsor"`
	// StorageCoveredBySponsor bool          `json:"storageCoveredBySponsor"`
	// StorageCollateralized   string        `json:"storageCollateralized"`
	// StorageReleased         []interface{} `json:"storageReleased"`
}

type ConfluxTransactionReceiptResp struct {
	Jsonrpc string                    `json:"jsonrpc"`
	Id      int                       `json:"id"`
	Result  ConfluxTransactionReceipt `json:"result"`
}

type ConfluxScanAccountTxResp struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    struct {
		Total int                  `json:"total"`
		List  []ConfluxScanTxBrief `json:"list"`
	} `json:"data"`
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

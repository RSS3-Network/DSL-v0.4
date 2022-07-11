package arweave

type Info struct {
	Network          string `json:"network"`
	Version          int    `json:"version"`
	Release          int    `json:"release"`
	Height           int    `json:"height"`
	Current          string `json:"current"`
	Blocks           int    `json:"blocks"`
	Peers            int    `json:"peers"`
	QueueLength      int    `json:"queue_length"`
	NodeStateLatency int    `json:"node_state_latency"`
}

type Block struct {
	Nonce         string        `json:"nonce"`
	PreviousBlock string        `json:"previous_block"`
	Timestamp     int           `json:"timestamp"`
	LastRetarget  int           `json:"last_retarget"`
	Diff          int           `json:"diff"`
	Height        int           `json:"height"`
	Hash          string        `json:"hash"`
	IndepHash     string        `json:"indep_hash"`
	Txs           []interface{} `json:"txs"`
	TxRoot        string        `json:"tx_root"`
	TxTree        []interface{} `json:"tx_tree"`
	WalletList    string        `json:"wallet_list"`
	RewardAddr    string        `json:"reward_addr"`
	Tags          []interface{} `json:"tags"`
	RewardPool    int           `json:"reward_pool"`
	WeaveSize     int           `json:"weave_size"`
	BlockSize     int           `json:"block_size"`
	PoA           BlockPoA      `json:"poa"`
}

type BlockPoA struct {
	Option   string `json:"option"`
	TxPath   string `json:"tx_path"`
	DataPath string `json:"data_path"`
	Chunk    string `json:"chunk"`
}

package protocol

const (
	Version                = "v1.1.0"
	Build                  = ""
	ExchangeJob            = "pregod11.jobs"
	IndexerWorkQueue       = "pregod11.indexer.work"
	IndexerWorkRoutingKey  = "pregod11.indexer.work"
	IndexerAssetQueue      = "pregod11.indexer.asset"
	IndexerAssetRoutingKey = "pregod11.indexer.asset"
	ContentTypeJSON        = "application/json"

	ExchangeRefresh = "pregod11.refresh"

	IndexVirtual int64 = -1
)

package protocol

const (
	Version               = "v1.1.0"
	Build                 = ""
	ExchangeJob           = "pregod11.jobs"
	IndexerWorkQueue      = "pregod11.indexer.work"
	IndexerWorkRoutingKey = "pregod11.indexer.work"
	// https://www.notion.so/rss3/io-APIkey-eccfd84c6afc420c9294e14aff9ddf34
	IndexerWorkQueueIO      = "pregod11.indexer-io.work"
	IndexerWorkRoutingKeyIO = "pregod11.indexer-io.work"
	IndexerAssetQueue       = "pregod11.indexer.asset"
	IndexerAssetRoutingKey  = "pregod11.indexer.asset"
	ContentTypeJSON         = "application/json"

	ExchangeRefresh = "pregod11.refresh"

	IndexVirtual int64 = -1
)

var WorkQ2RoutingKey = map[string]string{
	IndexerWorkQueue:   IndexerWorkRoutingKey,
	IndexerWorkQueueIO: IndexerWorkRoutingKeyIO,
}

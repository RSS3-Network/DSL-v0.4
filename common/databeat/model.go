package databeat

var IndexedCompletion = DataModel{
	Name: "pregod-v1-indexed-completion",
	Fields: []string{
		"address",
		"network",
		"transactions",
		"transfers",
	},
}

var DBMetrics = DataModel{
	Name: "pregod-v1-metrics",
	Fields: []string{
		"notes_count",
		"addresses_count",
		"transfers_count",
		"profiles_count",
		"profiles_per_platform",
		"top_20_addresses",
	},
}

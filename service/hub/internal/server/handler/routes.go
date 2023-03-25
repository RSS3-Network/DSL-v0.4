package handler

const (
	PathGetNotes        = "/notes/:address"
	PathGetAssets       = "/assets/:address"
	PathGetExchanges    = "/exchanges/:exchange_type"
	PathGetPlatformList = "/platforms/:platform_type"
	PathGetProfiles     = "/profiles/:address"
	PathGetNameResolve  = "/ns/:address"
	PathGetTransaction  = "/tx/:hash"

	PathBatchGetSocialNotes = "/notes/social"
	PathBatchGetNotes       = "/notes"
	PathBatchGetProfiles    = "/profiles"

	PathPostAPIKey = "/apikey/apply"
	PathGetAPIKey  = "/apikey/:address"
)

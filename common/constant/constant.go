package constant

// should not use built-in type string as key for value; define your own type to avoid collisions
type HeaderCtxKey string

const (
	// https://www.notion.so/rss3/io-APIkey-eccfd84c6afc420c9294e14aff9ddf34
	API_KEY_HEADER = "X-API-KEY"
	IO_API_Key     = "25092829-33b4-4ed2-a466-ee4873dc58d5"

	HEADER_CTX_KEY = HeaderCtxKey("header")
)

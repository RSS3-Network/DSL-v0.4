package allowlist

import "strings"

var allowList = map[string]string{
	"0xd7bc7107712e2999d3fcf55c92b465801c40974c": "POAP",
}

func Allow(address string) bool {
	_, exists := allowList[strings.ToLower(address)]

	return exists
}

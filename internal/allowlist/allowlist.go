package allowlist

import (
	"strings"
	"sync"
)

var allowList = map[string]string{
	"0xd7bc7107712e2999d3fcf55c92b465801c40974c": "POAP",
}

var spamList = map[string]string{
	"0xf79631521c474984f17d656a05e0d317d8755b20": "Layer Zero Ape",
}

func init() {
	AllowList = New()

	for address, name := range allowList {
		AllowList.Add(address, name)
	}

	SpamList = New()

	for address, name := range spamList {
		SpamList.Add(address, name)
	}
}

var (
	AllowList *List
	SpamList  *List
)

type List struct {
	addressMap sync.Map
}

func (l *List) Add(address string, name string) {
	l.addressMap.Store(strings.ToLower(address), name)
}

func (l *List) Contains(address string) bool {
	_, exists := l.addressMap.Load(strings.ToLower(address))

	return exists
}

func New() *List {
	return &List{
		addressMap: sync.Map{},
	}
}

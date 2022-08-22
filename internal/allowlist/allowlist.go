package allowlist

import (
	"strings"
	"sync"
)

var allowList = map[string]string{
	"0xd7bc7107712e2999d3fcf55c92b465801c40974c": "POAP",
	"0xd1feccf6881970105dfb2b654054174007f0e07e": "Lens",
	"0x57B7bf6f792a6181Ec5aFB88cE7bcE330a9d1b67": "Lens",
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

func (l *List) Keys() []string {
	keys := make([]string, 0)

	l.addressMap.Range(func(key, value interface{}) bool {
		keys = append(keys, key.(string))

		return true
	})

	return keys
}

func New() *List {
	return &List{
		addressMap: sync.Map{},
	}
}

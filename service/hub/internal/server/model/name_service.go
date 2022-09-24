package model

type NameServiceResult struct {
	ENS                string `json:"ens"`
	Crossbell          string `json:"crossbell"`
	Lens               string `json:"lens"`
	SpaceID            string `json:"spaceid"`
	UnstoppableDomains string `json:"unstoppable_domains"`
	Address            string `json:"address"`
}

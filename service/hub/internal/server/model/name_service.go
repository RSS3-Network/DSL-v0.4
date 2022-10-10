package model

type NameServiceResult struct {
	ENS                string `json:"ens"`
	Crossbell          string `json:"crossbell"`
	Lens               string `json:"lens"`
	SpaceID            string `json:"spaceid"`
	UnstoppableDomains string `json:"unstoppable_domains"`
	Bit                string `json:"bit"`
	Address            string `json:"address"`
}

type BitResult struct {
	Result struct {
		Error string `json:"errmsg"`
		Data  struct {
			Account     string `json:"account"`
			AccountInfo struct {
				Address string `json:"owner_key"`
			} `json:"account_info"`
		} `json:"data"`
	} `json:"result"`
}

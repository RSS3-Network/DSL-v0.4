package unstoppable_domains

type UnstoppableDomain struct {
	Name string `json:"name"`
	// Description string `json:"description"`
	// Properties  struct {
	//	Records struct {
	//		CryptoETHAddress               string `json:"crypto.ETH.address"`
	//		CryptoMATICVersionERC20Address string `json:"crypto.MATIC.version.ERC20.address"`
	//		CryptoMATICVersionMATICAddress string `json:"crypto.MATIC.version.MATIC.address"`
	//	} `json:"records"`
	// } `json:"properties"`
	ExternalUrl string `json:"external_url"`
	// Image       string `json:"image"`
	ImageUrl string `json:"image_url"`
	// Attributes  []struct {
	//	TraitType string      `json:"trait_type"`
	//	Value     interface{} `json:"value"`
	// } `json:"attributes"`
	// ImageData       string `json:"image_data"`
	// BackgroundColor string `json:"background_color"`
}

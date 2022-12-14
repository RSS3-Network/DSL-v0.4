package juicebox

type Metadata struct {
	Name          string `json:"name"`
	Description   string `json:"description"`
	LogoURI       string `json:"logoUri"`
	InfoURI       string `json:"infoUri"`
	Twitter       string `json:"twitter"`
	Discord       string `json:"discord"`
	PayButton     string `json:"payButton"`
	PayDisclosure string `json:"payDisclosure"`
	Version       int    `json:"version"`
}

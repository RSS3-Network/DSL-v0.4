package nftscan

type Result struct {
	Code int `json:"code"`
	Msg  any `json:"msg"`
	Data any `json:"data"`
}

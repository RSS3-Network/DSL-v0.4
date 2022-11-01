package iqwiki

type Wiki struct {
	Id      string `json:"id"`
	Content string `json:"content"`
	Images  []struct {
		Id   string `json:"id"`
		Type string `json:"type"`
	} `json:"images"`
}

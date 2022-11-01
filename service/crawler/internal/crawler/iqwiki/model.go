package iqwiki

type Wiki struct {
	Id      string `json:"id"`
	Content string `json:"content"`
	Media   []struct {
		Name string `json:"name"`
		Id   string `json:"id"`
	} `json:"media"`
	Images []struct {
		Id   string `json:"id"`
		Type string `json:"type"`
	} `json:"images"`
	Metadata []struct {
		Id    string `json:"id"`
		Value string `json:"value"`
	} `json:"metadata"`
}

type Reference struct {
	Id  string `json:"id"`
	Url string `json:"url"`
}

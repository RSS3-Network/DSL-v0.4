package snapshot

type Proposal struct {
	Id    string `json:"id"`
	End   int    `json:"end"`
	Body  string `json:"body"`
	Space struct {
		Id   string `json:"id"`
		Name string `json:"name"`
	} `json:"space"`
	Start    int      `json:"start"`
	State    string   `json:"state"`
	Title    string   `json:"title"`
	Author   string   `json:"author"`
	Choices  []string `json:"choices"`
	Created  int      `json:"created"`
	Snapshot string   `json:"snapshot"`
}

type Space struct {
	Id      string        `json:"id"`
	Name    string        `json:"name"`
	About   string        `json:"about"`
	Admins  []interface{} `json:"admins"`
	Symbol  string        `json:"symbol"`
	Filters struct {
		MinScore    int  `json:"minScore"`
		OnlyMembers bool `json:"onlyMembers"`
	} `json:"filters"`
	Members []string `json:"members"`
	Network string   `json:"network"`
	Plugins struct {
		SafeSnap struct {
			Address string `json:"address"`
		} `json:"safeSnap"`
	} `json:"plugins"`
	Strategies []struct {
		Name   string `json:"name"`
		Params struct {
			Symbol string `json:"symbol"`
		} `json:"params"`
	} `json:"strategies"`
}

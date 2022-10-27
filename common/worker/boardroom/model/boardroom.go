package model

type Protocol struct {
	Cname          string  `json:"cname"`
	Name           string  `json:"name"`
	TotalProposals int32   `json:"totalProposals"`
	TotalVotes     int32   `json:"totalVotes"`
	UniqueVoters   int32   `json:"uniqueVoters"`
	Tokens         []Token `json:"tokens"`
	Icons          []Icon  `json:"icons"`
}

type Token struct {
	Network         string        `json:"network"`
	ContractAddress string        `json:"contractAddress"`
	Adapter         string        `json:"adapter"`
	Symbol          string        `json:"symbol"`
	MarketPrices    []MarketPrice `json:"marketPrices"`
}

type MarketPrice struct {
	Currency string  `json:"currency"`
	Price    float32 `json:"price"`
}

type Icon struct {
	Adapter string `json:"adapter"`
	Size    string `json:"size"`
	Url     string `json:"url"`
}

type Proposal struct {
	RefId          string   `json:"refId"`
	Id             string   `json:"id"`
	StartTimestamp string   `json:"startTimestamp"`
	EndTimestamp   string   `json:"endTimestamp"`
	Title          string   `json:"title"`
	Content        string   `json:"content"`
	Protocol       string   `json:"protocol"`
	Adapter        string   `json:"adapter"`
	StartTime      BlockNum `json:"startTime"`
	EndTime        BlockNum `json:"endTime"`
	CurrentState   string   `json:"currentState"`
	Results        []Result `json:"results"`
	IndexedResults []Result `json:"indexedResults"`
	Choices        []string `json:"choices"`
	Events         []Event  `json:"events"`
	Proposer       string   `json:"proposer"`
	TotalVotes     int32    `json:"totalVotes"`
	BlockNumber    int64    `json:"blockNumber"`
	ExternalUrl    string   `json:"externalUrl"`
	Type           string   `json:"type"`
}

type BlockNum struct {
	BlockNumber int64 `json:"blockNumber"`
}

type Result struct {
	Total  float32 `json:"total"`
	Choice float32 `json:"choice"`
}

type Event struct {
	Timestamp int32    `json:"timestamp"`
	Time      BlockNum `json:"time"`
	Event     string   `json:"event"`
}

type Vote struct {
	Protocol     string `json:"protocol"`
	ProposalInfo struct {
		Title     string `json:"title"`
		StartTime struct {
			Timestamp int64 `json:"timestamp"`
		} `json:"startTime"`
		EndTime struct {
			Timestamp int64 `json:"timestamp"`
		} `json:"endTime"`
		StartTimestamp int64    `json:"startTimestamp"`
		EndTimestamp   int64    `json:"endTimestamp"`
		Choices        []string `json:"choices"`
		Events         []Event  `json:"events"`
		Type           string   `json:"type"`
		CurrentState   string   `json:"currentState"`
	} `json:"proposalInfo"`
	Adapter       string  `json:"adapter"`
	Address       string  `json:"address"`
	RefId         string  `json:"refId"`
	ProposalRefId string  `json:"proposalRefId"`
	Power         float32 `json:"power"`
	Time          struct {
		Timestamp int64 `json:"timestamp"`
	} `json:"time"`
	Choice     int    `json:"choice"`
	ProposalId string `json:"proposalId"`
	Timestamp  string `json:"timestamp"`
}

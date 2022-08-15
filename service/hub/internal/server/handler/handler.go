package handler

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/naturalselectionlabs/pregod/common/database/model"
	utils "github.com/naturalselectionlabs/pregod/common/utils/interface"
	rabbitmq "github.com/rabbitmq/amqp091-go"
	"gorm.io/gorm"
)

const (
	DefaultLimit = 500

	// path
	GetNotes     = "/notes/"
	PostNotes    = "/notes"
	GetProfiles  = "/profiles/"
	PostProfiles = "/profiles"
	GetNS        = "/ns/"

	EsIndex = "pregod-v1-visit-path"
)

type Handler struct {
	DatabaseClient     *gorm.DB
	RabbitmqConnection *rabbitmq.Connection
	RabbitmqChannel    *rabbitmq.Channel
}

type Response struct {
	Total         int64           `json:"total,omitempty"`
	Cursor        string          `json:"cursor,omitempty"`
	Result        any             `json:"result,omitempty"`
	AddressStatus []model.Address `json:"address_status,omitempty"`
	Message       string          `json:"message,omitempty"`
}

type GetRequest struct {
	Address     string    `param:"address" json:"address" validate:"required"`
	Limit       int       `query:"limit" json:"limit"`
	Cursor      string    `query:"cursor" json:"cursor"`
	Type        []string  `query:"type" json:"type"`
	Tag         []string  `query:"tag" json:"tag" validate:"required_with=Type"`
	Network     []string  `query:"network" json:"network"`
	Platform    []string  `query:"platform" json:"platform"`
	Timestamp   time.Time `query:"timestamp" json:"timestamp"`
	Hash        string    `query:"hash" json:"hash"`
	IncludePoap bool      `query:"include_poap" json:"include_poap"`
	Refresh     bool      `query:"refresh" json:"refresh"`
	Reindex     bool      `query:"reindex" json:"reindex"`
	Page        int       `query:"page" json:"page"`
	QueryStatus bool      `query:"query_status" json:"query_status"`
}

type GetExchangeRequest struct {
	ExchangeType string   `param:"exchange_type"`
	Cursor       int      `query:"cursor"`
	Name         []string `query:"name"`
	Network      []string `query:"network"`
}

type GetPlatformRequest struct {
	PlatformType string   `param:"platform_type"`
	Network      []string `query:"network"`
}

type GetAssetRequest struct {
	Address      string   `param:"address" validate:"required"`
	Network      []string `query:"network"`
	TokenAddress string   `query:"token_address" validate:"required_with=TokenId"`
	TokenId      string   `query:"token_id"`
	Cursor       string   `query:"cursor"`
	Limit        int      `query:"limit"`
	Refresh      bool     `query:"refresh"`
	BlockSpam    *bool    `query:"block_spam"` // Default true
}

type BatchGetNotesRequest struct {
	Address     []string  `json:"address" validate:"required"`
	Type        []string  `json:"type"`
	Tag         []string  `json:"tag" validate:"required_with=Type"`
	Network     []string  `json:"network"`
	Platform    []string  `json:"platform"`
	Timestamp   time.Time `json:"timestamp"`
	Limit       int       `json:"limit"`
	Cursor      string    `json:"cursor"`
	Refresh     bool      `json:"refresh"`
	IncludePoap bool      `json:"include_poap"`
	Page        int       `json:"page"`
	QueryStatus bool      `json:"query_status"`
}

type BatchGetProfilesRequest struct {
	Address  []string `json:"address" validate:"required"`
	Network  []string `json:"network"`
	Platform []string `json:"platform"`
	Refresh  bool     `query:"refresh"`
}

type APIKeyRequest struct {
	Address string `json:"address" validate:"required"`
}

type Transactions []model.Transaction

// Len()
func (t Transactions) Len() int {
	return len(t)
}

// Less()
func (t Transactions) Less(i, j int) bool {
	return t[i].Timestamp.Unix() > t[j].Timestamp.Unix() ||
		(t[i].Timestamp.Unix() == t[j].Timestamp.Unix() && t[i].Index > t[j].Index)
}

// Swap()
func (t Transactions) Swap(i, j int) {
	t[i], t[j] = t[j], t[i]
}

func APIReport(path string, apiKey interface{}) {
	report := map[string]interface{}{
		"index":   EsIndex,
		"path":    path,
		"ts":      time.Now().Format("2006-01-02 15:04:05"),
		"count":   true,
		"api_key": apiKey,
	}

	output, _ := json.Marshal(report)
	fmt.Printf("[DATABEAT]%v\n", string(output))
}

func FilterReport(path string, request interface{}) {
	b, _ := json.Marshal(request)
	report := make(map[string]interface{})

	err := json.Unmarshal(b, &report)
	if err != nil {
		return
	}

	if path == PostNotes {
		for _, address := range report["address"].([]interface{}) {
			batchReport := map[string]interface{}{
				"index":   EsIndex,
				"path":    path,
				"ts":      time.Now().Format("2006-01-02 15:04:05"),
				"address": address,
			}
			output, _ := json.Marshal(batchReport)
			fmt.Printf("[DATABEAT]%v\n", string(output))
		}
		delete(report, "address")
	}

	for k, v := range report {
		if utils.IfInterfaceValueIsNil(v) {
			delete(report, k)
		}
	}

	report["index"] = EsIndex
	report["path"] = path
	report["ts"] = time.Now().Format("2006-01-02 15:04:05")

	output, _ := json.Marshal(report)
	fmt.Printf("[DATABEAT]%v\n", string(output))
}

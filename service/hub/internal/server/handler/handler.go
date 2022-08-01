package handler

import (
	"encoding/json"
	"fmt"
	"reflect"
	"time"

	"github.com/naturalselectionlabs/pregod/common/database/model"
	rabbitmq "github.com/rabbitmq/amqp091-go"
	"gorm.io/gorm"
)

const (
	DefaultLimit = 500

	// path
	GetNotes    = "/notes/"
	PostNotes   = "/notes"
	GetProfiles = "/profiles/"
	GetNS       = "/ns/"
)

type Handler struct {
	DatabaseClient     *gorm.DB
	RabbitmqConnection *rabbitmq.Connection
	RabbitmqChannel    *rabbitmq.Channel
}

type Response struct {
	Total   int64  `json:"total,omitempty"`
	Cursor  string `json:"cursor,omitempty"`
	Result  any    `json:"result,omitempty"`
	Message string `json:"message,omitempty"`
}

type GetRequest struct {
	Address     string    `param:"address" validate:"required"`
	Limit       int       `query:"limit"`
	Cursor      string    `query:"cursor"`
	Type        []string  `query:"type"`
	Tag         []string  `query:"tag" validate:"required_with=Type"`
	Network     []string  `query:"network"`
	Platform    []string  `query:"platform"`
	Timestamp   time.Time `query:"timestamp"`
	Hash        string    `query:"hash"`
	IncludePoap bool      `query:"include_poap"`
	Refresh     bool      `query:"refresh"`
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
}

type BatchGetNotesRequest struct {
	Address     []string  `json:"address" validate:"required"`
	Type        []string  `query:"type"`
	Tag         []string  `query:"tag" validate:"required_with=Type"`
	Network     []string  `json:"network"`
	Platform    []string  `json:"platform"`
	Timestamp   time.Time `json:"timestamp"`
	Limit       int       `json:"limit"`
	Cursor      string    `json:"cursor"`
	Refresh     bool      `json:"refresh"`
	IncludePoap bool      `json:"include_poap"`
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

func EsReport(path string, request interface{}) {
	types := reflect.TypeOf(request)
	values := reflect.ValueOf(request)
	report := map[string]interface{}{
		"index": "pregod-hub-1",
		"path":  path,
		"ts":    time.Now().Format("2006-01-02 15:04:05"),
	}
	addresses := []string{}

	if kind := values.Kind(); kind != reflect.Struct {
		return
	}

	for i := 0; i < values.NumField(); i++ {
		key := types.Field(i).Tag.Get("query")
		value := values.Field(i)

		if len(key) == 0 {
			key = types.Field(i).Tag.Get("json")
		}

		if value.IsZero() {
			continue
		}

		// address
		if types.Field(i).Tag.Get("param") == "address" {
			addresses = append(addresses, value.String())
			continue
		}

		switch value.Kind() {
		case reflect.String:
			report[key] = value.String()
		case reflect.Int:
			report[key] = value.Int()
		case reflect.Bool:
			report[key] = value.Bool()
		default:
			if value.Len() == 0 {
				continue
			}
			list := []string{}
			for index := 0; index < value.Len(); index++ {
				list = append(list, value.Index(index).String())
			}
			if key == "address" {
				addresses = append(addresses, list...)
				continue
			}
			report[key] = list
		}
	}

	switch path {
	case PostNotes:
		for _, address := range addresses {
			batchReport := map[string]interface{}{
				"index":   "pregod-hub-1",
				"path":    path,
				"ts":      time.Now().Format("2006-01-02 15:04:05"),
				"address": address,
			}
			output, _ := json.Marshal(batchReport)
			fmt.Printf("[DATABEAT]%v\n", string(output))
		}
	default:
		if len(addresses) > 0 {
			report["address"] = addresses[0]
		}
	}

	output, _ := json.Marshal(report)
	fmt.Printf("[DATABEAT]%v\n", string(output))
}

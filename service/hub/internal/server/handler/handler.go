package handler

import (
	"time"

	"github.com/naturalselectionlabs/pregod/common/database/model"
	rabbitmq "github.com/rabbitmq/amqp091-go"
	"gorm.io/gorm"
)

const (
	DefaultLimit         = 500
	DefaultBatchGetLimit = 20
)

type Handler struct {
	DatabaseClient     *gorm.DB
	RabbitmqConnection *rabbitmq.Connection
	RabbitmqChannel    *rabbitmq.Channel
}

type Response struct {
	Total  int64  `json:"total"`
	Cursor string `json:"cursor,omitempty"`
	Result any    `json:"result"`
}

type GetRequest struct {
	Address   string    `param:"address"`
	Limit     int       `query:"limit"`
	Cursor    string    `query:"cursor"`
	Type      []string  `query:"type"`
	Tag       string    `query:"tag"`
	Network   []string  `query:"network"`
	Platform  []string  `query:"platform"`
	Timestamp time.Time `query:"timestamp"`
	Hash      string    `query:"hash"`
	Cache     bool      `query:"cache"`
}

type GetExchangeRequest struct {
	ExchangeType string   `param:"exchange_type"`
	Cursor       int      `query:"cursor"`
	Name         []string `query:"name"`
	Network      []string `query:"network"`
}

type BatchGetNoteListRequest struct {
	Timestamp time.Time `json:"timestamp"`
	Limit     int       `json:"limit"`
	Cursor    string    `json:"cursor"`
	Cache     bool      `json:"cache"`
	Filter    []Filter  `json:"filter"`
}

type Filter struct {
	Address  string   `json:"address"`
	Type     []string `json:"type"`
	Tag      string   `json:"tag"`
	Network  []string `json:"network"`
	Platform []string `json:"platform"`
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

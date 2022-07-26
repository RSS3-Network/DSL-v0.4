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
	Tag         string    `query:"tag" validate:"required_with=Type"`
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

type BatchGetNotesRequest struct {
	Address     []string  `json:"address" validate:"required"`
	Type        []string  `query:"type"`
	Tag         string    `query:"tag" validate:"required_with=Type"`
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

package model

import "time"

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
}

type GetExchangeRequest struct {
	ExchangeType string   `param:"exchange_type"`
	Cursor       int      `query:"cursor"`
	Name         []string `query:"name"`
	Network      []string `query:"network"`
}

package jschema

import (
	"encoding/json"
	"math/big"
	"reflect"
	"time"
)

// Handler is the custom handler for a special type
type Handler func() *Schema

func (s Schemas) AddHandler(v interface{}, h Handler) {
	r := s.RefT(reflect.TypeOf(v))
	s.handlers[r] = h
}

func (s Schemas) getHandler(r Ref) Handler {
	if h, has := s.handlers[r]; has {
		return h
	}
	return nil
}

func (s Schemas) AddTimeHandler() {
	s.AddHandler(time.Time{}, func() *Schema {
		return &Schema{
			Description: "time.Time",
			Title:       "Time",
			Type:        TypeString,
		}
	})
}

func (s Schemas) AddBigIntHandler() {
	s.AddHandler(big.Int{}, func() *Schema {
		return &Schema{
			Description: "math/big.Int",
			Title:       "Int",
			Type:        TypeNumber,
		}
	})
}

func (s Schemas) AddJSONRawMessageHandler() {
	s.AddHandler(json.RawMessage{}, func() *Schema {
		return &Schema{
			Description: "encoding/json.RawMessage",
			Title:       "RawMessage",
		}
	})
}

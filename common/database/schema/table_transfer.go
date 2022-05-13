package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

type Transfer struct {
	ent.Schema
}

func (Transfer) Fields() []ent.Field {
	return []ent.Field{
		field.String("transaction_hash"),
		field.Int("transaction_log_index"),
		field.String("address_from"),
		field.String("address_to"),
		field.String("token_address"),
		field.String("token_id"),
	}
}

func (Transfer) Edges() []ent.Edge {
	return nil
}

func (Transfer) Mixin() []ent.Mixin {
	return []ent.Mixin{
		TimeMixin{},
	}
}

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

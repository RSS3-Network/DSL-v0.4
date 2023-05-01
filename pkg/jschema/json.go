package jschema

import (
	"reflect"
)

func (s Schemas) JSON() map[string]*Schema {
	m := map[string]*Schema{}

	for id, v := range s.types {
		m[id] = v
	}

	return m
}

type Tag struct {
	name      string
	ignore    bool
	omitempty bool
	str       bool
}

func parseJSONTag(st reflect.StructTag) *Tag {
	v := st.Get("json")

	if v == "" {
		return nil
	}

	if v == "-" {
		return &Tag{
			ignore: true,
		}
	}

	name, t := parseTag(v)

	return &Tag{
		name:      name,
		omitempty: t.Contains("omitempty"),
		str:       t.Contains("string"),
	}
}

func (s *Schema) ChangeRoot(p string) {
}

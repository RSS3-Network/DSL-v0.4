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
	Name      string
	Ignore    bool
	Omitempty bool
	String    bool
}

func ParseJSONTag(st reflect.StructTag) *Tag {
	v := st.Get("json")

	if v == "" {
		return nil
	}

	if v == "-" {
		return &Tag{
			Ignore: true,
		}
	}

	name, t := parseTag(v)

	return &Tag{
		Name:      name,
		Omitempty: t.Contains("omitempty"),
		String:    t.Contains("string"),
	}
}

func (s *Schema) ChangeRoot(p string) {
}

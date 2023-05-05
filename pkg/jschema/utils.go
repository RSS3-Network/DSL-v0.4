package jschema

import (
	"bytes"
	"encoding/gob"
	"reflect"
)

func (s Schemas) Define(v interface{}) *Schema {
	return s.DefineT(reflect.TypeOf(v))
}

func (s *Schemas) GetSchema(v interface{}) *Schema {
	r := s.Ref(v)
	return s.types[r.ID]
}

func (s *Schemas) SetSchema(target interface{}, v *Schema) {
	s.Define(target)
	ss := s.GetSchema(target)
	title := ss.Title
	desc := ss.Description
	*ss = *v
	ss.Title = title
	ss.Description = desc
}

func (s *Schema) Clone() *Schema {
	buf := &bytes.Buffer{}
	err := gob.NewEncoder(buf).Encode(s)
	if err != nil {
		panic(err)
	}

	n := &Schema{}
	err = gob.NewDecoder(buf).Decode(n)
	if err != nil {
		panic(err)
	}

	return n
}

func (s *Schemas) AnyOf(list ...interface{}) *Schema {
	ss := []*Schema{}

	for _, v := range list {
		ss = append(ss, s.Define(v))
	}

	return &Schema{
		AnyOf: ss,
	}
}

func (s *Schemas) Const(v JVal) *Schema {
	ss := s.Define(v)
	ss.Enum = []JVal{v}
	return ss
}

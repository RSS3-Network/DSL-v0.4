// Package jschema ...
package jschema

import (
	"fmt"
	"reflect"
)

type Schemas struct {
	defs     string
	types    Types
	handlers map[Ref]Handler
	names    map[string]map[string]int
}

type Types map[string]*Schema

func New(defs string) Schemas {
	if defs == "" {
		defs = "#/$defs"
	}

	return Schemas{
		defs:     defs,
		types:    Types{},
		handlers: map[Ref]Handler{},
		names:    map[string]map[string]int{},
	}
}

// Schema is designed for typescript conversion.
type Schema struct {
	// string, number, boolean, null, array, object
	Type SchemaType `json:"type,omitempty"`

	Title string `json:"title,omitempty"`

	Description string `json:"description,omitempty"`

	Ref *Ref `json:"$ref,omitempty"`

	Nullable *bool `json:"nullable,omitempty"`

	AnyOf []*Schema `json:"anyOf,omitempty"`
	Enum  []JVal    `json:"enum,omitempty"`

	Properties           Properties `json:"properties,omitempty"`
	PatternProperties    Properties `json:"patternProperties,omitempty"`
	Required             []string   `json:"required,omitempty"`
	AdditionalProperties *bool      `json:"additionalProperties,omitempty"`

	Items    *Schema `json:"items,omitempty"`
	MinItems *int    `json:"minItems,omitempty"`
	MaxItems *int    `json:"maxItems,omitempty"`
}

type SchemaType string

const (
	TypeString  SchemaType = "string"
	TypeNumber  SchemaType = "number"
	TypeObject  SchemaType = "object"
	TypeArray   SchemaType = "array"
	TypeBool    SchemaType = "boolean"
	TypeNull    SchemaType = "null"
	TypeUnknown SchemaType = "unknown"
)

type Properties map[string]*Schema

type JVal interface{}

func (s Schemas) has(r Ref) bool {
	_, has := s.types[r.ID]
	return has
}

func (s Schemas) add(r Ref, scm *Schema) {
	if r.Unique() {
		s.types[r.ID] = scm
	}
}

func (s Schemas) DefineT(t reflect.Type) *Schema {
	r := s.RefT(t)
	if s.has(r) {
		return &Schema{Ref: &r}
	}

	scm := &Schema{}
	s.add(r, scm)

	if r.Package != "" {
		scm.Title = r.Name
		scm.Description = fmt.Sprintf("%s.%s", r.Package, r.Name)
	}

	if h := s.getHandler(r); h != nil {
		*scm = *h()
		return scm
	}

	switch t.Kind() {
	case reflect.Interface:
		scm.Type = TypeObject

	case reflect.Bool:
		scm.Type = TypeBool

	case reflect.String:
		scm.Type = TypeString

	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64,
		reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64,
		reflect.Float32, reflect.Float64,
		reflect.Uintptr, reflect.Complex64, reflect.Complex128:
		scm.Type = TypeNumber

	case reflect.Slice:
		el := s.DefineT(t.Elem())
		scm.Type = TypeArray
		scm.Items = el

	case reflect.Array:
		el := s.DefineT(t.Elem())
		l := t.Len()
		*scm = Schema{
			Type:     TypeArray,
			Items:    el,
			MinItems: &l,
			MaxItems: &l,
		}

	case reflect.Map:
		*scm = Schema{
			Type: TypeObject,
			PatternProperties: map[string]*Schema{
				"": s.DefineT(t.Elem()),
			},
		}

	case reflect.Struct:
		scm.Type = TypeObject
		scm.Properties = Properties{}
		scm.AdditionalProperties = new(bool)
		for i := 0; i < t.NumField(); i++ {
			f := t.Field(i)

			if !f.IsExported() {
				continue
			}

			tag := parseJSONTag(f.Tag)

			if tag != nil && tag.ignore {
				continue
			}

			n := f.Name

			p := s.DefineT(f.Type)

			desc := f.Tag.Get("description")
			if desc != "" {
				p.Description = desc
			}

			if tag != nil {
				if tag.name != "" {
					n = tag.name
				}
				if tag.str {
					p.Type = TypeString
				}
			}
			if tag == nil || !tag.omitempty {
				scm.Required = append(scm.Required, n)
			}

			scm.Properties[n] = p
		}

	case reflect.Ptr:
		*scm = *s.DefineT(t.Elem())

		if scm.Ref != nil {
			scm.AnyOf = []*Schema{{Ref: scm.Ref}}
			scm.Ref = nil
		}

		boolTrue := true
		scm.Nullable = &boolTrue

	default:
		scm.Type = TypeUnknown
	}

	if r.Unique() {
		scm = &Schema{
			Ref: &r,
		}
	}

	return scm
}

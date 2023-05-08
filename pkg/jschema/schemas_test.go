package jschema_test

import (
	"encoding/json"
	"math/big"
	"reflect"
	"testing"
	"time"

	"github.com/naturalselectionlabs/pregod/pkg/jschema"
	"github.com/ysmood/got"
)

func TestTypeName(t *testing.T) {
	g := got.T(t)

	g.Eq(reflect.TypeOf(1).PkgPath(), "")
}

func TestNil(t *testing.T) {
	g := got.T(t)

	type A struct {
		A *A
	}

	c := jschema.New("")

	out := c.Define(A{})

	g.Eq(g.JSON(g.ToJSONString(out)), map[string]interface{}{
		"$ref": `#/$defs/A`, /* len=42 */
	})
}

func TestCommonSchema(t *testing.T) {
	g := got.T(t)

	c := jschema.New("")

	type Node2 struct {
		Map map[string]int
		Any interface{}
	}

	type Node1 struct {
		Str     string
		Num     int    `json:"num,omitempty"`
		Bool    bool   `json:"bool"`
		Ignore  string `json:"-"`
		Slice   []Node1
		Arr     [2]int
		Obj     *Node2
		private int
	}

	c.Define(Node1{})
	c.Define(Node2{})

	g.Eq(g.JSON(g.ToJSONString(c.Define(Node1{}))), map[string]interface{}{
		"$ref": "#/$defs/Node1",
	})

	g.Eq(g.JSON(g.ToJSONString(c.JSON())), map[string]interface{} /* len=2 */ {
		"Node1": map[string]interface{} /* len=6 */ {
			`additionalProperties` /* len=20 */ : false,
			"description":                        `github.com/naturalselectionlabs/pregod/pkg/jschema_test.Node1`, /* len=61 */
			"properties": map[string]interface{} /* len=6 */ {
				"Arr": map[string]interface{} /* len=4 */ {
					"items": map[string]interface{}{
						"type": "number",
					},
					"maxItems": 2.0,
					"minItems": 2.0,
					"type":     "array",
				},
				"Obj": map[string]interface{} /* len=2 */ {
					"anyOf": []interface{} /* len=1 cap=1 */ {
						map[string]interface{}{
							"$ref": "#/$defs/Node2",
						},
					},
					"nullable": true,
				},
				"Slice": map[string]interface{} /* len=2 */ {
					"items": map[string]interface{}{
						"$ref": "#/$defs/Node1",
					},
					"type": "array",
				},
				"Str": map[string]interface{}{
					"type": "string",
				},
				"bool": map[string]interface{}{
					"type": "boolean",
				},
				"num": map[string]interface{}{
					"type": "number",
				},
			},
			"required": []interface{} /* len=5 cap=8 */ {
				"Str",
				"bool",
				"Slice",
				"Arr",
				"Obj",
			},
			"title": "Node1",
			"type":  "object",
		},
		"Node2": map[string]interface{} /* len=6 */ {
			`additionalProperties` /* len=20 */ : false,
			"description":                        `github.com/naturalselectionlabs/pregod/pkg/jschema_test.Node2`, /* len=61 */
			"properties": map[string]interface{} /* len=2 */ {
				"Any": map[string]interface{}{
					"type": "object",
				},
				"Map": map[string]interface{} /* len=2 */ {
					`patternProperties` /* len=17 */ : map[string]interface{}{
						"": map[string]interface{}{
							"type": "number",
						},
					},
					"type": "object",
				},
			},
			"required": []interface{} /* len=2 cap=2 */ {
				"Map",
				"Any",
			},
			"title": "Node2",
			"type":  "object",
		},
	})
}

func TestHandler(t *testing.T) {
	g := got.T(t)

	c := jschema.New("")

	type A struct {
		Str string
	}
	type B struct {
		A A
	}

	c.AddHandler(A{}, func() *jschema.Schema {
		return &jschema.Schema{
			Type: "number",
		}
	})

	c.Define(B{})

	g.Eq(g.JSON(g.ToJSONString(c.JSON())), map[string]interface{} /* len=2 */ {
		"A": map[string]interface{}{
			"type": "number",
		},
		"B": map[string]interface{} /* len=6 */ {
			`additionalProperties` /* len=20 */ : false,
			"description":                        `github.com/naturalselectionlabs/pregod/pkg/jschema_test.B`, /* len=57 */
			"properties": map[string]interface{}{
				"A": map[string]interface{}{
					"type": "number",
				},
			},
			"required": []interface{} /* len=1 cap=1 */ {
				"A",
			},
			"title": "B",
			"type":  "object",
		},
	})
}

func TestTime(t *testing.T) {
	g := got.T(t)

	c := jschema.New("")
	c.AddTimeHandler()
	c.Define(time.Now())

	g.Eq(g.JSON(g.ToJSONString(c.JSON())), map[string]interface{}{
		`Time` /* len=37 */ : map[string]interface{} /* len=3 */ {
			"description": "time.Time",
			"title":       "Time",
			"type":        "string",
		},
	})
}

func TestBigInt(t *testing.T) {
	g := got.T(t)

	c := jschema.New("")
	c.AddBigIntHandler()
	c.Define(big.Int{})

	g.Eq(g.JSON(g.ToJSONString(c.JSON())), map[string]interface{}{
		`Int` /* len=36 */ : map[string]interface{} /* len=3 */ {
			"description": "math/big.Int",
			"title":       "Int",
			"type":        "string",
		},
	})
}

func TestNameConflict(t *testing.T) {
	g := got.T(t)

	c := jschema.New("")

	type Time struct {
		Name string
	}

	c.Define(time.Time{})
	c.Define(Time{})

	g.Eq(g.JSON(g.ToJSONString(c.JSON())), map[string]interface{} /* len=2 */ {
		"Time": map[string]interface{} /* len=4 */ {
			`additionalProperties` /* len=20 */ : false,
			"description":                        "time.Time",
			"title":                              "Time",
			"type":                               "object",
		},
		"Time1": map[string]interface{} /* len=6 */ {
			`additionalProperties` /* len=20 */ : false,
			"description":                        `github.com/naturalselectionlabs/pregod/pkg/jschema_test.Time`, /* len=60 */
			"properties": map[string]interface{}{
				"Name": map[string]interface{}{
					"type": "string",
				},
			},
			"required": []interface{} /* len=1 cap=1 */ {
				"Name",
			},
			"title": "Time",
			"type":  "object",
		},
	})
}

func TestRawMessage(t *testing.T) {
	g := got.T(t)

	c := jschema.New("")
	c.AddJSONRawMessageHandler()

	type A struct {
		A json.RawMessage
	}

	c.Define(A{})

	g.Eq(g.JSON(g.ToJSONString(c.JSON())), map[string]interface{} /* len=2 */ {
		"A": map[string]interface{} /* len=6 */ {
			`additionalProperties` /* len=20 */ : false,
			"description":                        `github.com/naturalselectionlabs/pregod/pkg/jschema_test.A`, /* len=57 */
			"properties": map[string]interface{}{
				"A": map[string]interface{} /* len=2 */ {
					"description": `encoding/json.RawMessage`, /* len=24 */
					"title":       "RawMessage",
				},
			},
			"required": []interface{} /* len=1 cap=1 */ {
				"A",
			},
			"title": "A",
			"type":  "object",
		},
		"RawMessage": map[string]interface{} /* len=2 */ {
			"description": `encoding/json.RawMessage`, /* len=24 */
			"title":       "RawMessage",
		},
	})
}

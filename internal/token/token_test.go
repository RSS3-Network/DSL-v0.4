package token_test

import (
	"encoding/json"
	"github.com/naturalselectionlabs/pregod/internal/token"
	"reflect"
	"strings"
	"testing"
)

var str_0 = `{"name": "State of Mind 6377", "description": "Tools for collective imagination.", "image": "ipfs://bafybeibe57itbnezurdo4xml5722cw3eay5peb5jz43vmlvit3rkhtkj74?id=6377", "animation_url": "ipfs://bafybeibe57itbnezurdo4xml5722cw3eay5peb5jz43vmlvit3rkhtkj74?id=6377", "properties": {"number": 6377, "name": "State of Mind"}}`
var str_1 = `{"name":"RSS3 Whitepaper v1.0","description":"RSS3 Whitepaper v1.0 - Commemorative & Limited Edition","external_link":"https://rss3.io/RSS3-Whitepaper.pdf","attributes":[{"trait_type":"Author(s)","value":"Natural Selection Labs"},{"trait_type":"Edition","value":"First Edition"},{"trait_type":"Edition Language","value":"English"},{"trait_type":"File Format","value":"PDF"},{"trait_type":"No.","value":1383},{"trait_type":"Released Data","value":1610323200,"display_type":"date"}],"image":"ipfs://bafybeieqhs3fvot4es2rmza767dmgagwchlkgo5htv77ncashfmp6df7mq/rss3-whitepaper-no-1383.png","animation_url":"ipfs://bafybeiezsmiihbs67xoshf2ace3333jerk6twg6q3dbca5eqtmgcnbpboa/rss3-whitepaper-no-1383.glb"}`
var str_2 = `{"name":"RSS3 Whitepaper v1.0","description":"RSS3 Whitepaper v1.0 - Commemorative & Limited Edition","external_link":"https://rss3.io/RSS3-Whitepaper.pdf","attributes":[{"trait_type":"Author(s)","value":"Natural Selection Labs"},{"trait_type":"Edition","value":"First Edition"},{"trait_type":"Edition Language","value":"English"},{"trait_type":"File Format","value":"PDF"},{"trait_type":"No.","value":1383},{"trait_type":"Released Data","value":1610323200,"display_type":"date"}],"image":"ipfs://bafybeieqhs3fvot4es2rmza767dmgagwchlkgo5htv77ncashfmp6df7mq/rss3-whitepaper-no-1383.png","animation_url":"ipfs://bafybeiezsmiihbs67xoshf2ace3333jerk6twg6q3dbca5eqtmgcnbpboa/rss3-whitepaper-no-1383.glb","properties": {
        "name": {
            "type": "string",
            "description": "nihao"
        },
        "description": {
            "type": "string",
            "description": "handsome"
        },
        "image": {
            "type": "string",
            "description": "this is image"
        },
		"number":"1234"
    }}`

func TestUnmarshall(t *testing.T) {
	var metadata token.Metadata

	if err := json.Unmarshal([]byte(str_1), &metadata); err != nil {
		t.Fatal(err)
	}

	t.Log(metadata)

	x := metadataToAttributes(metadata)
	y, _ := json.Marshal(x)
	t.Log(string(y))
}

func metadataToAttributes(metadata token.Metadata) []token.MetadataAttribute {
	attributeMap := make(map[string]any)

	for _, attribute := range metadata.Attributes {
		attributeMap[attribute.TraitType] = attribute.Value
		attributeMap[attribute.DisplayType] = attribute.DisplayType
	}

	types := reflect.TypeOf(&token.MetadataProperty{})
	field1 := types.Elem().Field(1)
	description := field1.Tag.Get("json")
	for key, value := range metadata.Properties {
		t := reflect.TypeOf(value)
		if strings.EqualFold(t.Kind().String(), reflect.Map.String()) {
			if temp, ok := value.(map[string]any); ok {
				value = temp[description]
			}
		}
		attributeMap[key] = value
	}

	for _, trait := range metadata.Traits {
		attributeMap[trait.TraitType] = trait.Value
		attributeMap[trait.DisplayType] = trait.DisplayType
	}

	var attributes []token.MetadataAttribute

	for traitType, value := range attributeMap {
		attributes = append(attributes, token.MetadataAttribute{
			TraitType: traitType,
			Value:     value,
		})
	}

	return attributes
}

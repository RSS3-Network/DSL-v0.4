package token

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
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
	var (
		client   Client
		metadata Metadata
	)

	assert.NoError(t, json.Unmarshal([]byte(str_1), &metadata))

	t.Log(metadata)

	attributes := client.metadataToAttributes(metadata)

	result, err := json.Marshal(attributes)
	assert.NoError(t, err)

	t.Log(string(result))
}

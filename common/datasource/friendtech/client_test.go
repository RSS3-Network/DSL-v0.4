package friendtech

import (
	"context"
	"gotest.tools/assert"
	"testing"
)

func TestClient_GetUserMeta(t *testing.T) {
	c := NewClient()
	ctx := context.Background()
	t.Run("test friend.tech public api", func(t *testing.T) {
		resp, err := c.GetUserMeta(ctx, "0x8888888198FbdC8c017870cC5d3c96D0cf15C4F0")
		assert.NilError(t, err)
		t.Log(resp)
	})

	t.Run("test friend.tech public api", func(t *testing.T) {
		resp, err := c.GetUserMeta(ctx, "0xfd7232e66a69e1ae01e1e0ea8fab4776e2d325a9")
		assert.NilError(t, err)
		t.Log(resp)
	})
}

func TestClient_GetUserMetaByID(t *testing.T) {
	c := NewClient()
	ctx := context.Background()
	t.Run("test friend.tech public api", func(t *testing.T) {
		resp, err := c.GetUserMetaByID(ctx, 12123)
		assert.NilError(t, err)
		t.Log(resp)
	})

	t.Run("test friend.tech public api", func(t *testing.T) {
		resp, err := c.GetUserMetaByID(ctx, 10)
		assert.NilError(t, err)
		t.Log(resp)
	})
}

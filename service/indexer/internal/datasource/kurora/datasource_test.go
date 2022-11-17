package kurora_test

import (
	"context"
	"sync"
	"testing"
	"time"

	"github.com/naturalselectionlabs/pregod/common/database/model"
	"github.com/naturalselectionlabs/pregod/common/ethclientx"
	"github.com/naturalselectionlabs/pregod/common/protocol"
	"github.com/naturalselectionlabs/pregod/service/indexer/internal/config"
	"github.com/naturalselectionlabs/pregod/service/indexer/internal/datasource"
	"github.com/naturalselectionlabs/pregod/service/indexer/internal/datasource/kurora"
	"github.com/stretchr/testify/assert"
)

var once sync.Once

func initialize(t *testing.T) {
	once.Do(func() {
		config.Initialize()

		ethereumClientMap, err := ethclientx.Dial(config.ConfigIndexer.RPC, protocol.EthclientNetworks)
		assert.NoError(t, err)

		for network, ethereumClient := range ethereumClientMap {
			ethclientx.ReplaceGlobal(network, ethereumClient)
		}
	})
}

func TestKurora_Handle(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
	defer cancel()

	initialize(t)

	kuroraDatasource, err := kurora.New(ctx, config.ConfigIndexer.Kurora.Endpoint)
	assert.NoError(t, err)

	type fields struct {
		datasource datasource.Datasource
	}

	type arguments struct {
		ctx     context.Context
		message protocol.Message
	}

	testcases := []struct {
		name      string
		fields    fields
		arguments arguments
		want      assert.ValueAssertionFunc
		wantError assert.ErrorAssertionFunc
	}{
		{
			name: "kallydev.eth on Crossbell",
			fields: fields{
				datasource: kuroraDatasource,
			},
			arguments: arguments{
				ctx: ctx,
				message: protocol.Message{
					Address: "0x000000a52a03835517e9d193b3c27626e1bc96b1",
					Network: protocol.NetworkCrossbell,
				},
			},
			want: func(t assert.TestingT, i interface{}, i2 ...interface{}) bool {
				return assert.Greater(t, len(i.([]model.Transaction)), 0, i2...)
			},
			wantError: assert.NoError,
		},
		{
			name: "diygod.eth on Crossbell",
			fields: fields{
				datasource: kuroraDatasource,
			},
			arguments: arguments{
				ctx: ctx,
				message: protocol.Message{
					Address: "0xc8b960d09c0078c18dcbe7eb9ab9d816bcca8944",
					Network: protocol.NetworkCrossbell,
				},
			},
			want: func(t assert.TestingT, i interface{}, i2 ...interface{}) bool {
				return assert.Greater(t, len(i.([]model.Transaction)), 0, i2...)
			},
			wantError: assert.NoError,
		},
		{
			name: "black hole on Crossbell",
			fields: fields{
				datasource: kuroraDatasource,
			},
			arguments: arguments{
				ctx: ctx,
				message: protocol.Message{
					Address: "0x0000000000000000000000000000000000000000",
					Network: protocol.NetworkCrossbell,
				},
			},
			want: func(t assert.TestingT, i interface{}, i2 ...interface{}) bool {
				return assert.Equal(t, len(i.([]model.Transaction)), 0, i2...)
			},
			wantError: assert.NoError,
		},
	}

	for _, testcase := range testcases {
		t.Run(testcase.name, func(t *testing.T) {
			transactions, err := testcase.fields.datasource.Handle(testcase.arguments.ctx, &testcase.arguments.message)
			if testcase.wantError(t, err, "Handle(%v, %v)", testcase.arguments.ctx, testcase.arguments.message) {
				return
			}

			testcase.want(t, transactions, "Handle(%v, %v)", testcase.arguments.ctx, testcase.arguments.message)
		})
	}
}

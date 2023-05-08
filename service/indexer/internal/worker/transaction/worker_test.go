package transaction

import (
	"context"
	"encoding/json"
	"reflect"
	"testing"
	"time"

	"github.com/lib/pq"

	"github.com/alecthomas/repr"
	"github.com/naturalselectionlabs/pregod/common/cache"
	configx "github.com/naturalselectionlabs/pregod/common/config"
	"github.com/naturalselectionlabs/pregod/common/database"
	"github.com/naturalselectionlabs/pregod/common/database/model"
	"github.com/naturalselectionlabs/pregod/common/database/model/metadata"
	"github.com/naturalselectionlabs/pregod/common/ethclientx"
	"github.com/naturalselectionlabs/pregod/common/metadata_url"
	"github.com/naturalselectionlabs/pregod/common/protocol"
	"github.com/naturalselectionlabs/pregod/common/protocol/filter"
	"github.com/naturalselectionlabs/pregod/service/indexer/internal/config"
	"github.com/naturalselectionlabs/pregod/service/indexer/internal/worker"
	"github.com/stretchr/testify/assert"
)

var tokenWorker worker.Worker

func init() {
	config.Initialize()

	db, err := database.Dial(config.ConfigIndexer.Postgres.String(), true)
	if err != nil {
		panic(err)
	}

	database.ReplaceGlobal(db)

	cache.Dial(config.ConfigIndexer.Redis)

	metadata_url.New("https://ipfs.rss3.page/ipfs/")

	rpcConfig := &configx.RPC{
		General: configx.RPCNetwork{
			Ethereum: &configx.RPCEndpoint{
				WebSocket: "https://rpc.ankr.com/eth",
			},
		},
	}

	ethereumClientMap, err := ethclientx.Dial(rpcConfig, []string{
		protocol.NetworkEthereum,
	})
	if err != nil {
		panic(err)
	}

	ethclientx.ReplaceGlobal(protocol.NetworkEthereum, ethereumClientMap[protocol.NetworkEthereum])

	tokenWorker = New()
}

func Test_service_Name(t *testing.T) {
	tests := []struct {
		name string
		want string
	}{
		{
			name: "test transaction worker Name",
			want: "transaction",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := tokenWorker

			assert.NotNil(t, s.Jobs())

			if got := s.Name(); got != tt.want {
				t.Errorf("service.Name() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_service_Networks(t *testing.T) {
	tests := []struct {
		name string
		want []string
	}{
		{
			name: "test transaction worker Networks",
			want: []string{
				protocol.NetworkEthereum,
				protocol.NetworkPolygon,
				protocol.NetworkBinanceSmartChain,
				protocol.NetworkCrossbell,
				protocol.NetworkXDAI,
				protocol.NetworkZkSync,
				protocol.NetworkArbitrum,
				protocol.NetworkOptimism,
				protocol.NetworkAvalanche,
				protocol.NetworkCelo,
				protocol.NetworkFantom,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := tokenWorker
			if got := s.Networks(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("service.Networks() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_service_Initialize(t *testing.T) {
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name:    "test transaction worker Initialize",
			args:    args{context.Background()},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := tokenWorker
			if err := s.Initialize(tt.args.ctx); (err != nil) != tt.wantErr {
				t.Errorf("service.Initialize() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_service_Handle(t *testing.T) {
	location, err := time.LoadLocation("UTC")
	if err != nil {
		panic(err)
	}

	fakeTime := time.Date(2022, 6, 14, 0, 0, 0, 0, location)

	sourceData := `{"data":"0x","topics":["0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef","0x0000000000000000000000000000000000000000000000000000000000000000","0x000000000000000000000000d8da6bf26964af9d7eed9e03e53415d37aa96045","0x0000000000000000000000000000000000000000000000000000000000000596"],"address":"0x87acae6df21385a74ed4fb55a1a29354e9bdc6c1","removed":false,"logIndex":"0x32","blockHash":"0x6cbcd453c560f75d778caf28984f51ec8d38cb3b3f999004db2041ec6e85c979","blockNumber":"0xf09c54","transactionHash":"0xb7eb6c5820471dcf8b3827091e58ac636c5a2ebe58aceb3f48888e0e3a02fc5b","transactionIndex":"0x29"}`

	type args struct {
		ctx          context.Context
		message      *protocol.Message
		transactions []model.Transaction
	}
	tests := []struct {
		name    string
		args    args
		want    []model.Transaction
		wantErr bool
	}{
		{
			name: "Handle: empty",
			args: args{
				ctx:          context.Background(),
				message:      &protocol.Message{},
				transactions: []model.Transaction{},
			},
			want:    []model.Transaction{},
			wantErr: true,
		},
		{
			name: "Handle: vitalik.eth",
			args: args{
				ctx: context.Background(),
				message: &protocol.Message{
					Address: "0xd8da6bf26964af9d7eed9e03e53415d37aa96045",
					Network: protocol.NetworkEthereum,
				},
				transactions: []model.Transaction{
					{
						BlockNumber: 0,
						Timestamp:   fakeTime,
						Hash:        "0xb7eb6c5820471dcf8b3827091e58ac636c5a2ebe58aceb3f48888e0e3a02fc5b",
						AddressFrom: "0xd8da6bf26964af9d7eed9e03e53415d37aa96045",
						AddressTo:   "0x87acae6df21385a74ed4fb55a1a29354e9bdc6c1",
						Network:     protocol.NetworkEthereum,
						Source:      "ethereum",
						SourceData:  []byte(sourceData),
						Transfers: []model.Transfer{
							{
								TransactionHash: "0xb7eb6c5820471dcf8b3827091e58ac636c5a2ebe58aceb3f48888e0e3a02fc5b",
								Timestamp:       fakeTime,
								Index:           50,
								AddressFrom:     "0x0000000000000000000000000000000000000000",
								AddressTo:       "0xd8da6bf26964af9d7eed9e03e53415d37aa96045",
								Network:         protocol.NetworkEthereum,
								Source:          "ethereum",
								SourceData:      []byte(sourceData),
								Metadata:        metadata.Default,
							},
						},
					},
				},
			},
			want: []model.Transaction{
				{
					BlockNumber: 0,
					Timestamp:   fakeTime,
					Hash:        "0xb7eb6c5820471dcf8b3827091e58ac636c5a2ebe58aceb3f48888e0e3a02fc5b",
					Owner:       "0xd8da6bf26964af9d7eed9e03e53415d37aa96045",
					AddressFrom: "0xd8da6bf26964af9d7eed9e03e53415d37aa96045",
					AddressTo:   "0x87acae6df21385a74ed4fb55a1a29354e9bdc6c1",
					Network:     protocol.NetworkEthereum,
					Source:      "ethereum",
					Tag:         filter.TagCollectible,
					Type:        filter.CollectibleMint,
					SourceData:  []byte(sourceData),
					Transfers: []model.Transfer{
						{
							TransactionHash: "0xb7eb6c5820471dcf8b3827091e58ac636c5a2ebe58aceb3f48888e0e3a02fc5b",
							Timestamp:       fakeTime,
							Index:           50,
							AddressFrom:     "0x0000000000000000000000000000000000000000",
							AddressTo:       "0xd8da6bf26964af9d7eed9e03e53415d37aa96045",
							Network:         protocol.NetworkEthereum,
							Source:          "ethereum",
							SourceData:      []byte(sourceData),
							Tag:             filter.TagCollectible,
							Type:            filter.CollectibleMint,
							Metadata:        []byte(`{"id":"1430","cost":{"name":"Ethereum","image":"https://assets.coingecko.com/coins/images/279/large/ethereum.png","value":"150000000000000000","symbol":"ETH","decimals":18,"standard":"Native","value_display":"0.15"},"name":"#1430","image":"https://ipfs.rss3.page/ipfs/QmPJ1TLZYq4XRFqFMuFzFiHxqVgVeMrxEesHMDFjFXnqky","value":"1","symbol":"VOP","standard":"ERC-721","collection":"Voyager Pass","value_display":"1","contract_address":"0x87acae6df21385a74ed4fb55a1a29354e9bdc6c1"}`),
							RelatedUrls: pq.StringArray{
								"https://opensea.io/assets/0x87acae6df21385a74ed4fb55a1a29354e9bdc6c1/1430",
								"https://looksrare.org/collections/0x87acae6df21385a74ed4fb55a1a29354e9bdc6c1/1430",
							},
						},
					},
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := tokenWorker
			got, err := s.Handle(tt.args.ctx, tt.args.message, tt.args.transactions)
			if err != nil && tt.wantErr {
				return
			}

			if err != nil {
				t.Errorf("service.Handle() error = %v, wantErr %v", err, tt.wantErr)

				return
			}

			// check metadata
			for index := range got {
				for index_transfers := range got[index].Transfers {
					g := map[string]interface{}{}
					w := map[string]interface{}{}
					if err := json.Unmarshal(got[index].Transfers[index_transfers].Metadata, &g); err != nil {
						t.Error(err)
					}
					if err := json.Unmarshal(tt.want[index].Transfers[index_transfers].Metadata, &w); err != nil {
						t.Error(err)
					}
					for k, v := range g {
						assert.EqualValues(t, w[k], v, "key is %s", k)
					}
				}
			}

			// check other field
			for index := range got {
				for index_transfers := range got[index].Transfers {
					got[index].Transfers[index_transfers].Metadata = []byte("")
					tt.want[index].Transfers[index_transfers].Metadata = []byte("")
				}
			}
			if !reflect.DeepEqual(got, tt.want) {
				got[0].SourceData = []byte{}
				got[0].Transfers[0].SourceData = []byte{}
				tt.want[0].SourceData = []byte{}
				tt.want[0].Transfers[0].SourceData = []byte{}

				repr.Println(got)
				repr.Println(tt.want)

				t.Errorf("service.Handle() = %v, want %v", got, tt.want)
			}
		})
	}
}

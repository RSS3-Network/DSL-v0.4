package transaction

import (
	"context"
	"encoding/json"
	"github.com/lib/pq"
	"github.com/naturalselectionlabs/pregod/common/ethclientx"
	"github.com/naturalselectionlabs/pregod/common/ipfs"
	"reflect"
	"testing"
	"time"

	"github.com/alecthomas/repr"
	"github.com/naturalselectionlabs/pregod/common/cache"
	configx "github.com/naturalselectionlabs/pregod/common/config"
	"github.com/naturalselectionlabs/pregod/common/database"
	"github.com/naturalselectionlabs/pregod/common/database/model"
	"github.com/naturalselectionlabs/pregod/common/database/model/metadata"
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

	ipfs.New("https://ipfs.rss3.page/ipfs/")

	rpcConfig := &configx.RPC{
		General: configx.RPCNetwork{
			Ethereum: &configx.RPCEndpoint{
				HTTP: "https://rpc.rss3.dev/networks/ethereum",
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
			assert.Nil(t, s.Jobs())
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
	sourceDataZkSync := `{"tx":{"op":{"to":"0x000000a52a03835517e9d193b3c27626e1bc96b1","fee":"435000000000000000","from":"0x000000a52a03835517e9d193b3c27626e1bc96b1","type":"Transfer","nonce":8,"token":1,"amount":"0","accountId":1087980,"signature":{"pubKey":"6bcf3cf9121bd9db9b8cc0112bbd430d0a774637634b33cbb64a99f3fdcb3d2a","signature":"1a5ec0461ae66bb819f8a7997b7f935a031759be94bbc64fb0c575c395f2e9956f65ea4a8566149a90e418e0b0518e82c9e2679d0671b082089269d8b2860403"},"validFrom":0,"validUntil":4294967295},"status":"finalized","txHash":"0x3ad53192eb24de7c476d6e6d6b1edfa8068313831615503878c8f49efa52d4e7","batchId":1481025,"createdAt":"2022-06-08T13:28:23.927989Z","failReason":"","blockNumber":81578},"ethSignature":""}`
	sourceDataZkSyncNFT := `{"tx": {"op": {"to": "0xbaffff8509fc36ca4c6bccea3ae4c5fe53286892", "fee": "0", "from": "0x37719d7662a616e466b4d0f139a38e032946d503", "type": "Transfer", "nonce": 11, "token": 424218, "amount": "1", "accountId": 1101474, "signature": {"pubKey": "39010e78e70c93d0795324b039e38bb2649a7604485e32c4d40d651c33a028a6", "signature": "9493d8835e42bc9728078a78043be5eab3ff5d06ef16d82941136f6e46074b8832e7b01d8a469518e393c02f766b69dbc6e152d9dcbe861838dbe1dda2a4fc02"}, "validFrom": 0, "validUntil": 4294967295}, "status": "finalized", "txHash": "0x060914bad793e53410914c974a180267fc28fc0a83141a32a33f5316a7ab2852", "batchId": 1527108, "createdAt": "2022-06-14T02:59:37.921502Z", "failReason": "", "blockNumber": 84575}, "ethSignature": ""}`

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
			wantErr: false,
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
		{
			name: "Handle: zksync",
			args: args{
				ctx: context.Background(),
				message: &protocol.Message{
					Address: "0x000000a52a03835517e9d193b3c27626e1bc96b1",
					Network: protocol.NetworkZkSync,
				},
				transactions: []model.Transaction{
					{
						BlockNumber: 0,
						Timestamp:   fakeTime,
						Hash:        "0x3ad53192eb24de7c476d6e6d6b1edfa8068313831615503878c8f49efa52d4e7",
						AddressFrom: "0x000000a52a03835517e9d193b3c27626e1bc96b1",
						AddressTo:   "0x000000a52a03835517e9d193b3c27626e1bc96b1",
						Network:     protocol.NetworkZkSync,
						Source:      "zksync",
						SourceData:  []byte(sourceDataZkSync),
						Transfers: []model.Transfer{
							{
								TransactionHash: "0x3ad53192eb24de7c476d6e6d6b1edfa8068313831615503878c8f49efa52d4e7",
								Timestamp:       fakeTime,
								Index:           0,
								AddressFrom:     "0x000000a52a03835517e9d193b3c27626e1bc96b1",
								AddressTo:       "0x000000a52a03835517e9d193b3c27626e1bc96b1",
								Network:         protocol.NetworkZkSync,
								Source:          "zksync",
								SourceData:      []byte(sourceDataZkSync),
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
					Hash:        "0x3ad53192eb24de7c476d6e6d6b1edfa8068313831615503878c8f49efa52d4e7",
					Owner:       "0x000000a52a03835517e9d193b3c27626e1bc96b1",
					AddressFrom: "0x000000a52a03835517e9d193b3c27626e1bc96b1",
					AddressTo:   "0x000000a52a03835517e9d193b3c27626e1bc96b1",
					Network:     protocol.NetworkZkSync,
					Source:      "zksync",
					Tag:         filter.TagTransaction,
					Type:        filter.TransactionTransfer,
					SourceData:  []byte(sourceDataZkSync),
					Transfers: []model.Transfer{
						{
							TransactionHash: "0x3ad53192eb24de7c476d6e6d6b1edfa8068313831615503878c8f49efa52d4e7",
							Timestamp:       fakeTime,
							Tag:             filter.TagTransaction,
							Type:            filter.TransactionTransfer,
							Index:           0,
							AddressFrom:     "0x000000a52a03835517e9d193b3c27626e1bc96b1",
							AddressTo:       "0x000000a52a03835517e9d193b3c27626e1bc96b1",
							Network:         protocol.NetworkZkSync,
							Source:          "zksync",
							SourceData:      []byte(sourceDataZkSync),
							Metadata:        json.RawMessage(`{"name":"Dai Stablecoin","image":"https://assets.coingecko.com/coins/images/9956/large/4943.png","value":"0","symbol":"DAI","decimals":18,"standard":"ERC-20","value_display":"0","contract_address":"0x6b175474e89094c44da98b954eedeac495271d0f"}`),
						},
					},
				},
			},
			wantErr: false,
		},
		{
			name: "Handle: zksync NFT",
			args: args{
				ctx: context.Background(),
				message: &protocol.Message{
					Address: "0x37719d7662a616e466b4d0f139a38e032946d503",
					Network: protocol.NetworkZkSync,
				},
				transactions: []model.Transaction{
					{
						BlockNumber: 0,
						Timestamp:   fakeTime,
						Hash:        "0x060914bad793e53410914c974a180267fc28fc0a83141a32a33f5316a7ab2852",
						AddressFrom: "0x37719d7662a616e466b4d0f139a38e032946d503",
						AddressTo:   "0xbaffff8509fc36ca4c6bccea3ae4c5fe53286892",
						Network:     protocol.NetworkZkSync,
						Source:      "zksync",
						SourceData:  []byte(sourceDataZkSyncNFT),
						Transfers: []model.Transfer{
							{
								TransactionHash: "0x060914bad793e53410914c974a180267fc28fc0a83141a32a33f5316a7ab2852",
								Timestamp:       fakeTime,
								Index:           0,
								AddressFrom:     "0x37719d7662a616e466b4d0f139a38e032946d503",
								AddressTo:       "0xbaffff8509fc36ca4c6bccea3ae4c5fe53286892",
								Network:         protocol.NetworkZkSync,
								Source:          "zksync",
								SourceData:      []byte(sourceDataZkSyncNFT),
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
					Hash:        "0x060914bad793e53410914c974a180267fc28fc0a83141a32a33f5316a7ab2852",
					Owner:       "0x37719d7662a616e466b4d0f139a38e032946d503",
					AddressFrom: "0x37719d7662a616e466b4d0f139a38e032946d503",
					AddressTo:   "0xbaffff8509fc36ca4c6bccea3ae4c5fe53286892",
					Network:     protocol.NetworkZkSync,
					Tag:         filter.TagCollectible,
					Type:        filter.CollectibleTransfer,
					Source:      "zksync",
					SourceData:  []byte(sourceDataZkSyncNFT),
					Transfers: []model.Transfer{
						{
							TransactionHash: "0x060914bad793e53410914c974a180267fc28fc0a83141a32a33f5316a7ab2852",
							Timestamp:       fakeTime,
							Index:           0,
							AddressFrom:     "0x37719d7662a616e466b4d0f139a38e032946d503",
							AddressTo:       "0xbaffff8509fc36ca4c6bccea3ae4c5fe53286892",
							Network:         protocol.NetworkZkSync,
							Source:          "zksync",
							Tag:             filter.TagCollectible,
							Type:            filter.CollectibleTransfer,
							SourceData:      []byte(sourceDataZkSyncNFT),
							Metadata:        []byte(`{"id":"424218","name":"","symbol":"NFT-424218","standard":"ERC-721","value_display":null,"contract_address":"0x4a64471047696f0ee2dfbff8a92fd91c3a060cf2"}`),
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
			if (err != nil) != tt.wantErr {
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

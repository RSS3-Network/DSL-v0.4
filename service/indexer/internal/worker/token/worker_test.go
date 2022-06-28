package token

import (
	"context"
	"encoding/json"
	"reflect"
	"testing"
	"time"

	"github.com/naturalselectionlabs/pregod/common/protocol/filter"

	"github.com/naturalselectionlabs/pregod/common/database"
	"github.com/naturalselectionlabs/pregod/common/database/model"
	"github.com/naturalselectionlabs/pregod/common/protocol"
	"github.com/naturalselectionlabs/pregod/service/indexer/internal/worker"
	"github.com/stretchr/testify/assert"
)

var tokenWorker worker.Worker

func init() {
	db, err := database.Dial("postgres://postgres:password@127.0.0.1:5432/pregod2", true)
	if err != nil {
		panic(err)
	}
	tokenWorker = New(db)
}

func Test_service_Name(t *testing.T) {
	tests := []struct {
		name string
		want string
	}{
		{
			name: "test token worker Name",
			want: "token",
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
			name: "test token worker Networks",
			want: []string{
				protocol.NetworkEthereum,
				protocol.NetworkPolygon,
				protocol.NetworkBinanceSmartChain,
				protocol.NetworkZkSync,
				protocol.NetworkCrossbell,
				protocol.NetworkXDAI,
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
			name:    "test token worker Initialize",
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

	sourceData := `{"amount": "1", "token_id": "1", "log_index": "360", "block_hash": "0xbc1bad9a35b90beb6133090dead9ecbbdf7a66685da86176ce5fa1c6f7288b90", "to_address": "0xe02a52a553acf14cd5552e53d48dc0fc072978d8", "block_number": "28360240", "from_address": "0x0000000000000000000000000000000000000000", "contract_type": "ERC1155", "token_address": "0xb3a5104f3d934fffab52cfa5edd4968d0bbaa470", "block_timestamp": "2022-05-15T19:54:11.000Z", "transaction_hash": "0x7d3a23577c600b76fa7528f1a7cae8631cc7fa786a4b5923d676aed2f619121d", "transaction_type": "Single", "transaction_index": "55"}`
	sourceDataZkSync := `{"tx": {"op": {"to": "0x000000a52a03835517e9d193b3c27626e1bc96b1", "fee": "435000000000000000", "from": "0x000000a52a03835517e9d193b3c27626e1bc96b1", "type": "Transfer", "nonce": 8, "token": 1, "amount": "0", "accountId": 1087980, "signature": {"pubKey": "6bcf3cf9121bd9db9b8cc0112bbd430d0a774637634b33cbb64a99f3fdcb3d2a", "signature": "1a5ec0461ae66bb819f8a7997b7f935a031759be94bbc64fb0c575c395f2e9956f65ea4a8566149a90e418e0b0518e82c9e2679d0671b082089269d8b2860403"}, "validFrom": 0, "validUntil": 4294967295}, "status": "finalized", "txHash": "0x3ad53192eb24de7c476d6e6d6b1edfa8068313831615503878c8f49efa52d4e7", "batchId": 1481025, "createdAt": "2022-06-08T13:28:23.927989Z", "failReason": "", "blockNumber": 81578}, "ethSignature": ""}`
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
			name: "Handle: jeff",
			args: args{
				ctx: context.Background(),
				message: &protocol.Message{
					Address: "0xe02a52A553ACf14cD5552E53D48Dc0fC072978D8",
					Network: protocol.NetworkPolygon,
				},
				transactions: []model.Transaction{
					{
						BlockNumber: 0,
						Timestamp:   fakeTime,
						Hash:        "0x7d3a23577c600b76fa7528f1a7cae8631cc7fa786a4b5923d676aed2f619121d",
						AddressFrom: "0x0000000000000000000000000000000000000000",
						AddressTo:   "0xe02a52a553acf14cd5552e53d48dc0fc072978d8",
						Network:     protocol.NetworkPolygon,
						Source:      "moralis",
						SourceData:  []byte(sourceData),
						Transfers: []model.Transfer{
							{
								TransactionHash: "0x7d3a23577c600b76fa7528f1a7cae8631cc7fa786a4b5923d676aed2f619121d",
								Timestamp:       fakeTime,
								Index:           360,
								AddressFrom:     "0x0000000000000000000000000000000000000000",
								AddressTo:       "0xe02a52a553acf14cd5552e53d48dc0fc072978d8",
								Network:         protocol.NetworkPolygon,
								Source:          "moralis",
								SourceData:      []byte(sourceData),
								Metadata:        []byte("{}"),
							},
						},
					},
				},
			},
			want: []model.Transaction{
				{
					BlockNumber: 0,
					Timestamp:   fakeTime,
					Hash:        "0x7d3a23577c600b76fa7528f1a7cae8631cc7fa786a4b5923d676aed2f619121d",
					AddressFrom: "0x0000000000000000000000000000000000000000",
					AddressTo:   "0xe02a52a553acf14cd5552e53d48dc0fc072978d8",
					Network:     protocol.NetworkPolygon,
					Source:      "moralis",
					Tag:         filter.TagCollectible,
					SourceData:  []byte(sourceData),
					Transfers: []model.Transfer{
						{
							TransactionHash: "0x7d3a23577c600b76fa7528f1a7cae8631cc7fa786a4b5923d676aed2f619121d",
							Timestamp:       time.Date(2022, 6, 14, 0, 0, 0, 0, location),
							Index:           360,
							AddressFrom:     "0x0000000000000000000000000000000000000000",
							AddressTo:       "0xe02a52a553acf14cd5552e53d48dc0fc072978d8",
							Network:         protocol.NetworkPolygon,
							Source:          "moralis",
							SourceData:      []byte(sourceData),
							Tag:             filter.TagCollectible,
							Type:            filter.TransactionMint,
							Metadata:        []byte(`{"token":{"token_id":"1","token_value":"1","token_address":"0xb3a5104f3d934fffab52cfa5edd4968d0bbaa470","token_standard":"ERC1155"}}`),
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
					Address: "0x000000A52a03835517E9d193B3c27626e1Bc96b1",
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
								Timestamp:       time.Date(2022, 6, 14, 0, 0, 0, 0, location),
								Index:           0,
								AddressFrom:     "0x000000a52a03835517e9d193b3c27626e1bc96b1",
								AddressTo:       "0x000000a52a03835517e9d193b3c27626e1bc96b1",
								Network:         protocol.NetworkZkSync,
								Source:          "zksync",
								SourceData:      []byte(sourceDataZkSync),
								Metadata:        []byte("{}"),
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
					AddressFrom: "0x000000a52a03835517e9d193b3c27626e1bc96b1",
					AddressTo:   "0x000000a52a03835517e9d193b3c27626e1bc96b1",
					Network:     protocol.NetworkZkSync,
					Source:      "zksync",
					Tag:         filter.TagTransaction,
					SourceData:  []byte(sourceDataZkSync),
					Transfers: []model.Transfer{
						{
							TransactionHash: "0x3ad53192eb24de7c476d6e6d6b1edfa8068313831615503878c8f49efa52d4e7",
							Timestamp:       time.Date(2022, 6, 14, 0, 0, 0, 0, location),
							Tag:             filter.TagTransaction,
							Index:           0,
							AddressFrom:     "0x000000a52a03835517e9d193b3c27626e1bc96b1",
							AddressTo:       "0x000000a52a03835517e9d193b3c27626e1bc96b1",
							Network:         protocol.NetworkZkSync,
							Source:          "zksync",
							SourceData:      []byte(sourceDataZkSync),
							Type:            filter.TransactionTransfer,
							Metadata:        []byte(`{"token":{"symbol":"DAI","decimals":18,"token_id":"1","token_value":"0","token_address":"0x6b175474e89094c44da98b954eedeac495271d0f","token_standard":"ERC-20"}}`),
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
								Timestamp:       time.Date(2022, 6, 14, 0, 0, 0, 0, location),
								Index:           0,
								AddressFrom:     "0x37719d7662a616e466b4d0f139a38e032946d503",
								AddressTo:       "0xbaffff8509fc36ca4c6bccea3ae4c5fe53286892",
								Network:         protocol.NetworkZkSync,
								Source:          "zksync",
								SourceData:      []byte(sourceDataZkSyncNFT),
								Metadata:        []byte("{}"),
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
					AddressFrom: "0x37719d7662a616e466b4d0f139a38e032946d503",
					AddressTo:   "0xbaffff8509fc36ca4c6bccea3ae4c5fe53286892",
					Network:     protocol.NetworkZkSync,
					Tag:         filter.TagCollectible,
					Source:      "zksync",
					SourceData:  []byte(sourceDataZkSyncNFT),
					Transfers: []model.Transfer{
						{
							TransactionHash: "0x060914bad793e53410914c974a180267fc28fc0a83141a32a33f5316a7ab2852",
							Timestamp:       time.Date(2022, 6, 14, 0, 0, 0, 0, location),
							Index:           0,
							AddressFrom:     "0x37719d7662a616e466b4d0f139a38e032946d503",
							AddressTo:       "0xbaffff8509fc36ca4c6bccea3ae4c5fe53286892",
							Network:         protocol.NetworkZkSync,
							Source:          "zksync",
							Tag:             filter.TagCollectible,
							SourceData:      []byte(sourceDataZkSyncNFT),
							Type:            filter.CollectibleTrade,
							Metadata:        []byte(`{"token":{"symbol":"NFT-424218","token_id":"424218","token_value":"1","nft_metadata":{"id":424218,"symbol":"NFT-424218","address":"0x4a64471047696f0ee2dfbff8a92fd91c3a060cf2","serialId":0,"creatorId":1101474,"contentHash":"0xd252497a6db751c63bb23eb1493e8461b280bdc398bcc494e6bb5bab2d04935a","creatorAddress":"0x37719d7662a616e466b4d0f139a38e032946d503","currentFactory":"0x7c770595a2be9a87cf49b35ea9bc534f1a59552d","withdrawnFactory":""},"token_address":"0x4a64471047696f0ee2dfbff8a92fd91c3a060cf2","token_standard":"ERC-721"}}`),
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
					g := map[string]map[string]interface{}{}
					w := map[string]map[string]interface{}{}
					if err := json.Unmarshal(got[index].Transfers[index_transfers].Metadata, &g); err != nil {
						t.Error(err)
					}
					if err := json.Unmarshal(tt.want[index].Transfers[index_transfers].Metadata, &w); err != nil {
						t.Error(err)
					}
					for k, v := range g["token"] {
						assert.EqualValues(t, v, w["token"][k], "key is %s", k)
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
				t.Errorf("service.Handle() = %v, want %v", got, tt.want)
			}
		})
	}
}

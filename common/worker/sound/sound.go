package sound

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/Khan/genqlient/graphql"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/naturalselectionlabs/pregod/common/cache"
	"github.com/naturalselectionlabs/pregod/common/database/model"
	"github.com/naturalselectionlabs/pregod/common/database/model/metadata"
	"github.com/naturalselectionlabs/pregod/common/datasource/ethereum"
	"github.com/naturalselectionlabs/pregod/common/datasource/ethereum/contract/sound"
	"github.com/naturalselectionlabs/pregod/common/datasource/ethereum/contract/sound/contract"
	"github.com/naturalselectionlabs/pregod/common/ethclientx"
	"github.com/naturalselectionlabs/pregod/common/protocol"
	"github.com/naturalselectionlabs/pregod/common/protocol/filter"
	"github.com/naturalselectionlabs/pregod/common/utils"
	"github.com/naturalselectionlabs/pregod/common/utils/loggerx"
	"github.com/naturalselectionlabs/pregod/common/utils/opentelemetry"
	"github.com/naturalselectionlabs/pregod/common/worker/sound/schema"
	"github.com/naturalselectionlabs/pregod/internal/token"
	"github.com/samber/lo"
	"github.com/shopspring/decimal"
	"go.opentelemetry.io/otel"
	"go.uber.org/zap"
	"math/big"
	"net/http"
	"time"
)

//go:generate go run -mod=mod github.com/Khan/genqlient ./schema/genqlient.yaml

var APIKey string

type httpClient struct{}

var _ graphql.Doer = (*httpClient)(nil)

func (h *httpClient) Do(request *http.Request) (*http.Response, error) {
	request.Header.Set("X-Sound-Client-Key", APIKey)

	return http.DefaultClient.Do(request)
}

func GetEditionInfo(ctx context.Context, contractAddress string, editionId string) (*schema.ReleaseContractResponse, error) {
	// read cache
	cacheKey := fmt.Sprintf("sound:edition:%s:%s", contractAddress, editionId)
	if cache, err := cache.Global().Get(ctx, cacheKey).Result(); err == nil {
		var result schema.ReleaseContractResponse

		if err = json.Unmarshal([]byte(cache), &result); err == nil {
			return &result, nil
		}
	}

	graphqlClient := graphql.NewClient("https://api.sound.xyz/graphql", &httpClient{})

	result, err := schema.ReleaseContract(context.Background(), graphqlClient, contractAddress, editionId)
	if err != nil {
		loggerx.Global().Error("failed to query sound api", zap.Error(err))

		return nil, err
	}

	if result == nil {
		return nil, fmt.Errorf("failed to query sound api")
	}

	// write cache
	go cache.Global().Set(ctx, cacheKey, result, 24*time.Hour)

	return result, nil
}

func GetEditionNfts(ctx context.Context, contractAddress string, editionId string) (*schema.ReleaseNftsResponse, error) {
	graphqlClient := graphql.NewClient("https://api.sound.xyz/graphql", &httpClient{})

	var nfts *schema.ReleaseNftsResponse

	var cursor string
	for {
		response, err := schema.ReleaseNfts(context.Background(), graphqlClient, contractAddress, editionId, cursor)
		if err != nil {
			loggerx.Global().Error("failed to query sound api", zap.Error(err))

			return nil, err
		}

		if response == nil {
			return nil, fmt.Errorf("failed to query sound api")
		}

		if nfts == nil {
			nfts = response
		} else {
			nfts.ReleaseContract.NftsPaginated.Edges = append(nfts.ReleaseContract.NftsPaginated.Edges, response.GetReleaseContract().NftsPaginated.Edges...)
		}

		if !response.GetReleaseContract().NftsPaginated.PageInfo.HasNextPage {
			return nfts, nil
		}

		cursor = response.GetReleaseContract().NftsPaginated.PageInfo.EndCursor
	}

	return nfts, nil
}

func HandleReceipt(ctx context.Context, transaction *model.Transaction) (results []model.Transfer, err error) {
	tracer := otel.Tracer("sound")
	_, trace := tracer.Start(ctx, "sound:handleReceipt")

	defer func() { opentelemetry.Log(trace, transaction, results, err) }()

	// source
	var sourceData ethereum.SourceData
	if err := json.Unmarshal(transaction.SourceData, &sourceData); err != nil {
		return nil, fmt.Errorf("failed to unmarshal source data: %w", err)
	}

	// ethereum client
	client, err := ethclientx.Global(transaction.Network)
	if err != nil {
		return nil, fmt.Errorf("failed to create ethereum client: %w", err)
	}

	// handle logs
	for _, log := range sourceData.Receipt.Logs {
		if len(log.Topics) == 0 {
			continue
		}

		sourceData, err := json.Marshal(log)
		if err != nil {
			loggerx.Global().Error("sound: marshal log error", zap.Error(err))

			continue
		}

		transfer := model.Transfer{
			AddressFrom:     transaction.AddressFrom,
			AddressTo:       transaction.AddressTo,
			TransactionHash: transaction.Hash,
			Timestamp:       transaction.Timestamp,
			BlockNumber:     big.NewInt(transaction.BlockNumber),
			Index:           int64(log.Index),
			Network:         transaction.Network,
			SourceData:      sourceData,
			Platform:        transaction.Platform,
			Source:          transaction.Source,
			Tag:             filter.TagCollectible,
			Type:            filter.CollectibleMusic,
			RelatedUrls:     []string{utils.GetTxHashURL(protocol.NetworkEthereum, transaction.Hash)},
		}

		var transfers []model.Transfer
		switch log.Topics[0] {
		case sound.EventHashEditionCreatedV1:
			artistContract, err := contract.NewArtist(log.Address, client)
			if err != nil {
				loggerx.Global().Error("sound: new artist contract error", zap.Error(err))

				continue
			}

			transfers, err = HandleEditionCreatedV1(ctx, artistContract, log, &transfer)
		case sound.EventHashEditionCreatedV3:
			artistContract, err := contract.NewArtistV3(log.Address, client)
			if err != nil {
				loggerx.Global().Error("sound: new artist v3 contract error", zap.Error(err))

				continue
			}

			transfers, err = HandleEditionCreatedV3(ctx, artistContract, log, &transfer)
		case sound.EventHashEditionPurchasedV3:
			artistContract, err := contract.NewArtistV3(log.Address, client)
			if err != nil {
				loggerx.Global().Error("sound: new artist v3 contract error", zap.Error(err))

				continue
			}

			err = HandleEditionPurchasedV3(ctx, artistContract, log, &transfer)
		case sound.EventHashEditionPurchasedV5:
			artistContract, err := contract.NewArtistV5(log.Address, client)
			if err != nil {
				loggerx.Global().Error("sound: new artist v5 contract error", zap.Error(err))

				continue
			}

			err = HandleEditionPurchasedV5(ctx, artistContract, log, &transfer)
		default:
			continue
		}

		if err != nil {
			loggerx.Global().Error("sound: handle receipt error", zap.Error(err))

			continue
		}

		if len(transfers) == 0 {
			results = append(results, transfer)
		} else {
			results = append(results, transfers...)
		}
	}

	return results, nil
}

func HandleEditionCreatedV1(ctx context.Context, artistContract *contract.Artist, log *types.Log, transfer *model.Transfer) (transfers []model.Transfer, err error) {
	tracer := otel.Tracer("sound")
	_, trace := tracer.Start(ctx, "sound:HandleEditionCreatedV1")

	defer func() { opentelemetry.Log(trace, log, transfers, err) }()

	event, err := artistContract.ParseEditionCreated(*log)
	if err != nil {
		loggerx.Global().Error("failed to parse edition created event", zap.Error(err))

		return nil, err
	}

	// get edition
	edition, err := GetEditionNfts(ctx, log.Address.String(), event.EditionId.String())
	if err != nil {
		loggerx.Global().Error("failed to get edition info", zap.Error(err), zap.String("contract", log.Address.String()), zap.String("edition_id", event.EditionId.String()))

		return nil, err
	}

	for index, nft := range edition.GetReleaseContract().NftsPaginated.Edges {
		tokenID, _ := decimal.NewFromString(nft.Node.TokenId)
		metadata, err := GetMetadata(ctx, &log.Address, tokenID.BigInt(), edition.GetReleaseContract().CoverImage.Url)
		if err != nil {
			loggerx.Global().Error("failed to get metadata", zap.Error(err), zap.String("contract", log.Address.String()), zap.String("token_id", tokenID.String()))

			continue
		}

		metadata.Action = filter.CollectibleMusicRelease
		value, _ := decimal.NewFromString(edition.GetReleaseContract().Price)
		metadata.Cost.Value = lo.ToPtr(value)
		metadata.Cost.ValueDisplay = lo.ToPtr(value.Shift(-int32(metadata.Cost.Decimals)))

		releaseTransfer := *transfer
		releaseTransfer.Index = int64(index)
		releaseTransfer.Metadata, err = json.Marshal(metadata)
		releaseTransfer.RelatedUrls = append(transfer.RelatedUrls, edition.GetReleaseContract().OpenseaUrl)
		transfers = append(transfers, releaseTransfer)
	}

	return transfers, nil
}

func HandleEditionCreatedV3(ctx context.Context, artistContract *contract.ArtistV3, log *types.Log, transfer *model.Transfer) (transfers []model.Transfer, err error) {
	tracer := otel.Tracer("sound")
	_, trace := tracer.Start(ctx, "sound:HandleEditionCreatedV3")

	defer func() { opentelemetry.Log(trace, log, transfers, err) }()

	event, err := artistContract.ParseEditionCreated(*log)
	if err != nil {
		loggerx.Global().Error("failed to parse edition created event", zap.Error(err))

		return nil, err
	}

	// get edition
	edition, err := GetEditionNfts(ctx, log.Address.String(), event.EditionId.String())
	if err != nil {
		loggerx.Global().Error("failed to get edition info", zap.Error(err), zap.String("contract", log.Address.String()), zap.String("edition_id", event.EditionId.String()))

		return nil, err
	}

	for index, nft := range edition.GetReleaseContract().NftsPaginated.Edges {
		tokenID, _ := decimal.NewFromString(nft.Node.TokenId)
		metadata, err := GetMetadata(ctx, &log.Address, tokenID.BigInt(), edition.GetReleaseContract().CoverImage.Url)
		if err != nil {
			loggerx.Global().Error("failed to get metadata", zap.Error(err), zap.String("contract", log.Address.String()), zap.String("token_id", tokenID.String()))

			continue
		}

		metadata.Action = filter.CollectibleMusicRelease

		releaseTransfer := *transfer
		releaseTransfer.Index = int64(index)
		releaseTransfer.Metadata, err = json.Marshal(metadata)
		releaseTransfer.RelatedUrls = append(transfer.RelatedUrls, edition.GetReleaseContract().OpenseaUrl)
		transfers = append(transfers, releaseTransfer)
	}

	return transfers, nil
}

func HandleEditionPurchasedV3(ctx context.Context, artistContract *contract.ArtistV3, log *types.Log, transfer *model.Transfer) (err error) {
	tracer := otel.Tracer("sound")
	_, trace := tracer.Start(ctx, "sound:HandleEditionPurchasedV3")

	defer func() { opentelemetry.Log(trace, log, transfer, err) }()

	// parse event
	event, err := artistContract.ParseEditionPurchased(*log)
	if err != nil {
		loggerx.Global().Error("failed to parse edition purchased event", zap.Error(err))

		return fmt.Errorf("failed to parse edition purchased event: %w", err)
	}

	// get edition information
	edition, err := GetEditionInfo(ctx, log.Address.String(), event.TokenId.String())
	if err != nil {
		loggerx.Global().Error("failed to get edition info", zap.Error(err), zap.String("contract", log.Address.String()), zap.String("token_id", event.TokenId.String()))

		return err
	}

	// get nft metadata
	metadata, err := GetMetadata(ctx, &log.Address, event.TokenId, edition.GetReleaseContract().CoverImage.Url)
	if err != nil {
		loggerx.Global().Error("failed to get metadata", zap.Error(err), zap.String("contract", log.Address.String()), zap.String("token_id", event.TokenId.String()))

		return err
	}

	metadata.Action = filter.CollectibleMusicBuyEdition
	metadata.SetValue(decimal.NewFromInt32(int32(event.NumSold)))
	value, _ := decimal.NewFromString(edition.GetReleaseContract().Price)
	metadata.Cost.Value = lo.ToPtr(value)
	metadata.Cost.ValueDisplay = lo.ToPtr(value.Shift(-int32(metadata.Cost.Decimals)))
	transfer.Metadata, err = json.Marshal(metadata)
	transfer.RelatedUrls = append(transfer.RelatedUrls, edition.GetReleaseContract().OpenseaUrl)
	return nil
}

func HandleEditionPurchasedV5(ctx context.Context, artistContract *contract.ArtistV5, log *types.Log, transfer *model.Transfer) (err error) {
	tracer := otel.Tracer("sound")
	_, trace := tracer.Start(ctx, "sound:HandleEditionPurchasedV5")

	defer func() { opentelemetry.Log(trace, log, transfer, err) }()

	// parse event
	event, err := artistContract.ParseEditionPurchased(*log)
	if err != nil {
		loggerx.Global().Error("failed to parse edition purchased event", zap.Error(err))

		return fmt.Errorf("failed to parse edition purchased event: %w", err)
	}

	// get edition information
	edition, err := GetEditionInfo(ctx, log.Address.String(), event.TokenId.String())
	if err != nil {
		loggerx.Global().Error("failed to get edition info", zap.Error(err), zap.String("contract", log.Address.String()), zap.String("token_id", event.TokenId.String()))

		return err
	}

	// get nft metadata
	metadata, err := GetMetadata(ctx, &log.Address, event.TokenId, edition.GetReleaseContract().CoverImage.Url)
	if err != nil {
		loggerx.Global().Error("failed to get metadata", zap.Error(err), zap.String("contract", log.Address.String()), zap.String("token_id", event.TokenId.String()))

		return err
	}

	metadata.Action = filter.CollectibleMusicBuyEdition
	metadata.SetValue(decimal.NewFromInt32(int32(event.NumSold)))
	value, _ := decimal.NewFromString(edition.GetReleaseContract().Price)
	metadata.Cost.Value = lo.ToPtr(value)
	metadata.Cost.ValueDisplay = lo.ToPtr(value.Shift(-int32(metadata.Cost.Decimals)))
	transfer.Metadata, err = json.Marshal(metadata)
	transfer.RelatedUrls = append(transfer.RelatedUrls, edition.GetReleaseContract().OpenseaUrl)
	return nil
}

func GetMetadata(ctx context.Context, contractAddress *common.Address, tokenId *big.Int, image string) (soundMetadata *metadata.Token, err error) {
	tracer := otel.Tracer("sound")
	_, trace := tracer.Start(ctx, "sound:GetMetadata")
	soundMetadata = &metadata.Token{}
	defer func() { opentelemetry.Log(trace, nil, soundMetadata, err) }()

	// get token metadata
	tokenClient := token.New()

	erc721, err := tokenClient.ERC721(ctx, protocol.NetworkEthereum, contractAddress.String(), tokenId)
	if err != nil {
		return nil, fmt.Errorf("failed to get erc721: %w", err)
	}

	nft, err := erc721.ToNFT(tokenId)
	if err != nil {
		return nil, fmt.Errorf("erc721 to nft %s/%d: %w", contractAddress, tokenId, err)
	}

	soundMetadata, err = nft.ToMetadata()
	if err != nil {
		return nil, fmt.Errorf("nft to metadata %s/%d: %w", contractAddress, tokenId, err)
	}

	if soundMetadata.Image == "" {
		soundMetadata.Image = image
	}

	nativeToken, err := tokenClient.Native(ctx, protocol.NetworkEthereum)
	if err != nil {
		return nil, err
	}

	costToken := metadata.Token{
		Name:     nativeToken.Name,
		Symbol:   nativeToken.Symbol,
		Decimals: nativeToken.Decimals,
		Standard: protocol.TokenStandardNative,
		Image:    nativeToken.Logo,
	}

	soundMetadata.Cost = &costToken

	return soundMetadata, nil
}

package lens

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"math/big"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/naturalselectionlabs/pregod/common/datasource/ethereum/contract/erc721"

	"github.com/naturalselectionlabs/pregod/common/metadata_url"
	"github.com/samber/lo"
	"github.com/shopspring/decimal"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/core/types"
	kurora_client "github.com/naturalselectionlabs/kurora/client"
	kurora "github.com/naturalselectionlabs/kurora/common/contract/lens"
	"github.com/naturalselectionlabs/pregod/common/database/model"
	"github.com/naturalselectionlabs/pregod/common/database/model/metadata"
	"github.com/naturalselectionlabs/pregod/common/database/model/social"
	"github.com/naturalselectionlabs/pregod/common/datasource/ethereum"
	"github.com/naturalselectionlabs/pregod/common/datasource/ethereum/contract/lens"
	"github.com/naturalselectionlabs/pregod/common/datasource/ethereum/contract/lens/contract"
	lenscontract "github.com/naturalselectionlabs/pregod/common/datasource/ethereum/contract/lens/contract"
	"github.com/naturalselectionlabs/pregod/common/ethclientx"
	"github.com/naturalselectionlabs/pregod/common/protocol"
	"github.com/naturalselectionlabs/pregod/common/protocol/filter"
	"github.com/naturalselectionlabs/pregod/common/utils"
	"github.com/naturalselectionlabs/pregod/common/utils/loggerx"
	"github.com/naturalselectionlabs/pregod/common/utils/opentelemetry"
	"github.com/sirupsen/logrus"
	"go.opentelemetry.io/otel"
	"go.uber.org/zap"
)

type LensContent struct {
	Description string             `json:"description"`
	Content     string             `json:"content"`
	ExternalURL string             `json:"external_url"`
	AppId       string             `json:"appId"`
	Media       []LensContentMedia `json:"media"`
	CreatedOn   time.Time          `json:"createdOn"`
	Attributes  []struct {
		TraitType string `json:"traitType"`
		Key       string `json:"key"`
		Value     string `json:"value"`
	} `json:"attributes"`
}

type LensContentMedia struct {
	Item string `json:"item"`
	Type string `json:"type"`
}

type FormatOption struct {
	ContentURI       string
	ContentType      string
	Content          json.RawMessage
	Transfer         *model.Transfer
	ProfileIdPointed *big.Int
	PubIdPointed     *big.Int
	ProfileId        *big.Int
	PubId            *big.Int
	Handle           string
}

const (
	Post    = "post"
	Comment = "comment"
	Share   = "mirror"
)

var ErrorNotFoundInKurora = errors.New("not found")

var SupportLensPlatform = map[string]bool{
	protocol.PlatformLens:                 true,
	protocol.PlatformLensOrb:              true,
	protocol.PlatformLensLenster:          true,
	protocol.PlatformLensLenstube:         true,
	protocol.PlatformLensLenstubeBytes:    true,
	protocol.PlatformLensLensterCrowdfund: true,
}

func (c *Client) GetProfile(profileID *big.Int) (*social.Profile, error) {
	if profileID == nil || profileID.Int64() == 0 {
		return nil, fmt.Errorf("empty profile")
	}

	lensHubContract, err := lenscontract.NewHub(lens.HubProxyContractAddress, c.ethClient)
	if err != nil {
		loggerx.Global().Error("[common] lens: NewHub err", zap.Error(err))

		return nil, err
	}

	owner, err := lensHubContract.OwnerOf(&bind.CallOpts{}, profileID)
	if err != nil {
		loggerx.Global().Error("[common] lens: OwnerOf error", zap.Error(err))

		return nil, err
	}

	result, err := lensHubContract.GetProfile(&bind.CallOpts{}, profileID)
	if err != nil {
		loggerx.Global().Error("[common] lens: GetProfile err", zap.Error(err))

		return nil, err
	}

	profile := &social.Profile{
		Address:     strings.ToLower(owner.String()),
		Network:     protocol.NetworkPolygon,
		Platform:    protocol.PlatformLens,
		Source:      protocol.SourceKurora,
		Name:        result.Handle,
		Handle:      result.Handle,
		URL:         fmt.Sprintf("https://lenster.xyz/u/%v", result.Handle),
		ProfileUris: []string{metadata_url.GetDirectURL(result.ImageURI)},
	}

	return profile, nil
}

func (c *Client) HandleReceipt(ctx context.Context, transaction *model.Transaction) (transfers []model.Transfer, err error) {
	tracer := otel.Tracer("worker_lens")
	_, trace := tracer.Start(ctx, "worker_len:handleReceipt")

	defer func() { opentelemetry.Log(trace, transaction, transfers, err) }()

	var sourceData ethereum.SourceData
	if err := json.Unmarshal(transaction.SourceData, &sourceData); err != nil {
		return nil, fmt.Errorf("failed to unmarshal source data: %w", err)
	}

	// Parse logs
	for _, log := range sourceData.Receipt.Logs {
		// Filter anonymous log
		if len(log.Topics) == 0 {
			continue
		}

		lensContract, err := contract.NewEvents(log.Address, c.ethClient)
		if err != nil {
			logrus.Errorf("[lens worker] handleReceipt: new events error, %v", err)

			continue
		}

		// common attributes
		transfer := model.Transfer{
			TransactionHash: transaction.Hash,
			Timestamp:       transaction.Timestamp,
			BlockNumber:     big.NewInt(transaction.BlockNumber),
			Index:           int64(log.Index),
			Network:         transaction.Network,
			Platform:        protocol.PlatformLens,
			Source:          transaction.Source,
			RelatedUrls:     []string{utils.GetTxHashURL(protocol.NetworkPolygon, transaction.Hash)},
		}

		var handleErr error
		var batchTransfers []model.Transfer

		switch log.Topics[0] {
		case lens.EventHashPostCreated:
			handleErr = c.HandlePostCreated(ctx, *lensContract, transaction, &transfer, *log)
		case lens.EventHashCommentCreated:
			handleErr = c.HandleCommentCreated(ctx, *lensContract, transaction, &transfer, *log)
		case lens.EventHashProfileCreated:
			handleErr = c.HandleProfileCreated(ctx, *lensContract, transaction, &transfer, *log)
		case lens.EventHashMirrorCreated:
			handleErr = c.HandleMirrorCreated(ctx, *lensContract, transaction, &transfer, *log)
		case lens.EventHashFollowed:
			batchTransfers, handleErr = c.HandleFollowed(ctx, *lensContract, transaction, &transfer, *log)
		case lens.EventHashFollowNFTTransferred:
			handleErr = c.HandleFollowNFTTransferred(ctx, *lensContract, transaction, &transfer, *log)
		case lens.EventHashCollected:
			handleErr = c.HandleCollected(ctx, *lensContract, transaction, &transfer, *log)
		default:
			continue
		}
		if handleErr != nil {
			continue
		}

		if len(batchTransfers) > 0 {
			transfers = append(transfers, batchTransfers...)
			continue
		}

		transfers = append(transfers, transfer)
	}

	if len(transfers) > 0 && transfers[0].Tag == filter.TagSocial && (transfers[0].Type == filter.SocialFollow || transfers[0].Type == filter.SocialMint) {
		return c.HandleTransfer(ctx, transaction, transfers)
	}

	return transfers, nil
}

func (c *Client) HandlePostCreated(ctx context.Context, lensContract contract.Events, transaction *model.Transaction, transfer *model.Transfer, log types.Log) error {
	event, err := lensContract.EventsFilterer.ParsePostCreated(log)
	if err != nil {
		loggerx.Global().Error("[lens worker] HandlePostCreated: ParsePostCreated error", zap.Error(err))

		return err
	}

	profile, err := c.GetProfile(event.ProfileId)
	if err != nil {
		return err
	}

	err = c.FormatContent(ctx, &FormatOption{
		ContentURI:  event.ContentURI,
		ContentType: Post,
		Transfer:    transfer,
		ProfileId:   event.ProfileId,
		PubId:       event.PubId,
		Handle:      profile.Handle,
	})
	if err != nil {
		loggerx.Global().Error("[lens worker] HandlePostCreated: FormatContent error", zap.Error(err))

		return err
	}

	transaction.Owner = strings.ToLower(profile.Address)
	transfer.AddressFrom = transaction.Owner

	transfer.Timestamp = time.Unix(event.Timestamp.Int64(), 0)
	transfer.Tag, transfer.Type = filter.UpdateTagAndType(filter.TagSocial, transfer.Tag, filter.SocialPost, transfer.Type)
	if SupportLensPlatform[transfer.Platform] {
		transfer.RelatedUrls = append(transfer.RelatedUrls, c.GetLensRelatedURL(ctx, event.ProfileId, event.PubId))
	}

	return nil
}

func (c *Client) HandleCommentCreated(ctx context.Context, lensContract contract.Events, transaction *model.Transaction, transfer *model.Transfer, log types.Log) error {
	event, err := lensContract.EventsFilterer.ParseCommentCreated(log)
	if err != nil {
		loggerx.Global().Error("[lens worker] handleCommentCreated: ParseCommentCreated error", zap.Error(err))

		return err
	}

	profile, err := c.GetProfile(event.ProfileId)
	if err != nil {
		return err
	}

	err = c.FormatContent(ctx, &FormatOption{
		ContentURI:       event.ContentURI,
		ContentType:      Comment,
		Transfer:         transfer,
		ProfileId:        event.ProfileId,
		PubId:            event.PubId,
		ProfileIdPointed: event.ProfileIdPointed,
		PubIdPointed:     event.PubIdPointed,
		Handle:           profile.Handle,
	})

	if err != nil {
		loggerx.Global().Error("[lens worker] handleCommentCreated: FormatContent error", zap.Error(err))

		return err
	}

	transaction.Owner = strings.ToLower(profile.Address)
	transfer.AddressFrom = transaction.Owner

	transfer.Timestamp = time.Unix(event.Timestamp.Int64(), 0)
	transfer.Tag, transfer.Type = filter.UpdateTagAndType(filter.TagSocial, transfer.Tag, filter.SocialComment, transfer.Type)
	if SupportLensPlatform[transfer.Platform] {
		transfer.RelatedUrls = append(transfer.RelatedUrls, c.GetLensRelatedURL(ctx, event.ProfileId, event.PubId))
	}

	return nil
}

func (c *Client) HandleProfileCreated(ctx context.Context, lensContract contract.Events, transaction *model.Transaction, transfer *model.Transfer, log types.Log) error {
	event, err := lensContract.EventsFilterer.ParseProfileCreated(log)
	if err != nil {
		loggerx.Global().Error("[lens worker] handleProfileCreated: ParseProfileCreated error", zap.Error(err))

		return err
	}

	profile := social.Profile{
		Address:     strings.ToLower(event.To.String()),
		Platform:    protocol.PlatformLens,
		Network:     transaction.Network,
		Source:      transaction.Platform,
		Action:      filter.SocialCreate,
		URL:         fmt.Sprintf("https://lenster.xyz/u/%v", event.Handle),
		Handle:      event.Handle,
		ProfileUris: []string{metadata_url.GetDirectURL(event.ImageURI)},
	}

	rawMetadata, err := json.Marshal(&profile)
	if err != nil {
		loggerx.Global().Error("[lens worker] handleProfileCreated: json marshal error", zap.Error(err))

		return err
	}

	transfer.Timestamp = time.Unix(event.Timestamp.Int64(), 0)
	transfer.AddressFrom = strings.ToLower(event.Creator.String())
	transfer.AddressTo = strings.ToLower(event.To.String())
	transfer.Metadata = rawMetadata
	transfer.Tag, transfer.Type = filter.UpdateTagAndType(filter.TagSocial, transfer.Tag, filter.SocialProfile, transfer.Type)
	transfer.RelatedUrls = append(transfer.RelatedUrls, fmt.Sprintf("https://lenster.xyz/u/%v", event.Handle))

	transaction.Owner = transfer.AddressFrom
	return nil
}

func (c *Client) HandleMirrorCreated(ctx context.Context, lensContract contract.Events, transaction *model.Transaction, transfer *model.Transfer, log types.Log) error {
	event, err := lensContract.EventsFilterer.ParseMirrorCreated(log)
	if err != nil {
		loggerx.Global().Error("[lens worker] handleMirrorCreated: ParseMirrorCreated error", zap.Error(err))
		return err
	}

	contentURI, err := c.GetContentURI(ctx, event.ProfileIdPointed, event.PubIdPointed)
	if err != nil {
		loggerx.Global().Error("[lens worker] HandleMirrorCreated: GetContentURI error", zap.Error(err))
		return err
	}

	profile, err := c.GetProfile(event.ProfileId)
	if err != nil {
		loggerx.Global().Error("[lens worker] HandleFollowNFTTransferred: GetProfile error", zap.Error(err))

		return err
	}

	err = c.FormatContent(ctx, &FormatOption{
		ContentURI:       contentURI,
		ContentType:      Share,
		Transfer:         transfer,
		ProfileId:        event.ProfileId,
		PubId:            event.PubId,
		ProfileIdPointed: event.ProfileIdPointed,
		PubIdPointed:     event.PubIdPointed,
		Handle:           profile.Handle,
	})
	if err != nil {
		loggerx.Global().Error("[lens worker] HandleMirrorCreated: FormatContent error", zap.Error(err))
		return err
	}

	transaction.Owner = strings.ToLower(profile.Address)
	transfer.AddressFrom = transaction.Owner

	transfer.Timestamp = time.Unix(event.Timestamp.Int64(), 0)
	transfer.Tag, transfer.Type = filter.UpdateTagAndType(filter.TagSocial, transfer.Tag, filter.SocialShare, transfer.Type)
	if SupportLensPlatform[transfer.Platform] {
		transfer.RelatedUrls = append(transfer.RelatedUrls, c.GetLensRelatedURL(ctx, event.ProfileId, event.PubId))
	}

	return nil
}

func (c *Client) HandleFollowed(ctx context.Context, lensContract contract.Events, transaction *model.Transaction, transfer *model.Transfer, log types.Log) ([]model.Transfer, error) {
	event, err := lensContract.EventsFilterer.ParseFollowed(log)
	if err != nil {
		loggerx.Global().Error("[lens worker] HandleFollowed: ParseFollowed error", zap.Error(err))
		return nil, err
	}

	transaction.Owner = strings.ToLower(event.Follower.String())
	transfer.AddressFrom = transaction.Owner

	transfers := make([]model.Transfer, 0)
	for index, profileID := range event.ProfileIds {
		profile, err := c.GetProfile(profileID)
		if err != nil {
			loggerx.Global().Error("[lens worker] HandleFollowed: GetProfile error", zap.Error(err))
			return nil, err
		}

		transfer.Index = int64(index)
		transfer.RelatedUrls = []string{
			utils.GetTxHashURL(protocol.NetworkPolygon, transaction.Hash),
			fmt.Sprintf("https://lenster.xyz/u/%v", profile.Handle),
		}

		if transfer.Metadata, err = json.Marshal(profile); err != nil {
			loggerx.Global().Error("[lens worker] HandleFollowed: json marshal error", zap.Error(err))

			return nil, err
		}

		transfer.Timestamp = time.Unix(event.Timestamp.Int64(), 0)
		transfer.Tag, transfer.Type = filter.UpdateTagAndType(filter.TagSocial, transfer.Tag, filter.SocialFollow, transfer.Type)
		transfers = append(transfers, *transfer)
	}

	return transfers, nil
}

func (c *Client) HandleFollowNFTTransferred(ctx context.Context, lensContract contract.Events, transaction *model.Transaction, transfer *model.Transfer, log types.Log) error {
	event, err := lensContract.EventsFilterer.ParseFollowNFTTransferred(log)
	if err != nil {
		loggerx.Global().Error("[lens worker] HandleFollowNFTTransferred: ParseFollowNFTTransferred error", zap.Error(err))

		return err
	}

	// unfollow: burn
	if event.To != ethereum.AddressGenesis {
		return fmt.Errorf("not supported")
	}

	profile, err := c.GetProfile(event.ProfileId)
	if err != nil {
		loggerx.Global().Error("[lens worker] HandleFollowNFTTransferred: GetProfile error", zap.Error(err))

		return err
	}

	transaction.Owner = strings.ToLower(event.From.String())
	transfer.AddressFrom = transaction.Owner

	if transfer.Metadata, err = json.Marshal(profile); err != nil {
		loggerx.Global().Error("[lens worker] HandleFollowNFTTransferred: json marshal error", zap.Error(err))

		return err
	}

	transfer.Timestamp = time.Unix(event.Timestamp.Int64(), 0)
	transfer.Tag, transfer.Type = filter.UpdateTagAndType(filter.TagSocial, transfer.Tag, filter.SocialUnfollow, transfer.Type)
	transfer.RelatedUrls = append(transfer.RelatedUrls, fmt.Sprintf("https://lenster.xyz/u/%v", profile.Handle))

	return nil
}

func (c *Client) HandleCollected(ctx context.Context, lensContract contract.Events, transaction *model.Transaction, transfer *model.Transfer, log types.Log) error {
	event, err := lensContract.EventsFilterer.ParseCollected(log)
	if err != nil {
		loggerx.Global().Error("[lens worker] HandleCollected: ParseCollected error", zap.Error(err))

		return err
	}

	var contentURI string
	var content json.RawMessage

	pubs, err := c.kuroraClient.FetchDatasetLensPublications(ctx, kurora_client.DatasetLensPublicationQuery{
		ProfileID:     lo.ToPtr[decimal.Decimal](decimal.NewFromBigInt(event.ProfileId, 0)),
		PublicationID: lo.ToPtr[decimal.Decimal](decimal.NewFromBigInt(event.PubId, 0)),
		Limit:         lo.ToPtr[int](1),
	})

	if err != nil || len(pubs) == 0 {
		contentURI, err = c.GetContentURI(ctx, event.ProfileId, event.PubId)
		if err != nil {
			loggerx.Global().Error("[lens worker] HandleCollected: GetContentURI error", zap.Error(err))

			return err
		}
	} else {
		contentURI = pubs[0].ContentURI
		content = pubs[0].Content
	}

	profile, err := c.GetProfile(event.ProfileId)
	if err != nil {
		loggerx.Global().Error("[lens worker] HandleCollected: GetProfile error", zap.Error(err))
		return err
	}

	err = c.FormatContent(ctx, &FormatOption{
		ContentURI: contentURI,
		Content:    content,
		Transfer:   transfer,
		ProfileId:  event.ProfileId,
		PubId:      event.PubId,
		Handle:     profile.Handle,
	})

	if err != nil {
		loggerx.Global().Error("[lens worker] HandleCollected: FormatContent error", zap.Error(err))

		return err
	}

	transaction.Owner = strings.ToLower(event.Collector.String())
	transfer.AddressFrom = transaction.Owner

	transfer.Timestamp = time.Unix(event.Timestamp.Int64(), 0)
	transfer.Tag, transfer.Type = filter.UpdateTagAndType(filter.TagSocial, transfer.Tag, filter.SocialMint, transfer.Type)
	transfer.RelatedUrls = append(transfer.RelatedUrls, c.GetLensRelatedURL(ctx, event.ProfileId, event.PubId))

	return nil
}

func (c *Client) HandleTransfer(ctx context.Context, transaction *model.Transaction, transfers []model.Transfer) ([]model.Transfer, error) {
	var sourceData ethereum.SourceData
	if err := json.Unmarshal(transaction.SourceData, &sourceData); err != nil {
		return nil, fmt.Errorf("failed to unmarshal source data: %w", err)
	}

	for _, log := range sourceData.Receipt.Logs {
		if len(log.Topics) != 4 || log.Topics[0] != erc721.EventHashTransfer {
			continue
		}

		transferFiltered := lo.Must(erc721.NewERC721Filterer(ethereum.AddressGenesis, nil))

		event, err := transferFiltered.ParseTransfer(*log)
		if err != nil {
			zap.L().Error("[lens worker] HandleTransfer: ParseTransfer error", zap.Error(err), zap.String("transaction", transaction.Hash))

			return nil, err
		}

		if event.From != common.HexToAddress(transaction.Owner) {
			continue
		}

		transaction.Owner = strings.ToLower(event.To.String())
		for index := range transfers {
			transfers[index].AddressFrom = transaction.Owner
		}
	}

	return transfers, nil
}

func (c *Client) GetContentURI(ctx context.Context, profileId *big.Int, pubId *big.Int) (string, error) {
	// rpc
	ethclient, err := ethclientx.Global(protocol.NetworkPolygon)
	if err != nil {
		return "", err
	}
	iLensHub, err := contract.NewILensHub(lens.HubProxyContractAddress, ethclient)
	if err != nil {
		logrus.Errorf("[lens worker] NewILensHub error, %v", err)

		return "", err
	}

	return iLensHub.GetContentURI(&bind.CallOpts{}, profileId, pubId)
}

func (c *Client) GetLensRelatedURL(ctx context.Context, profileId *big.Int, pubId *big.Int) string {
	return fmt.Sprintf("https://lenster.xyz/posts/%v-%v", kurora.EncodeID(profileId), kurora.EncodeID(pubId))
}

func (c *Client) GetLensAuthorURL(ctx context.Context, externalURL string, handle string) []string {
	if len(externalURL) == 0 && len(handle) == 0 {
		return nil
	}

	var author []string
	if len(externalURL) == 0 {
		author = append(author, fmt.Sprintf("https://lenster.xyz/u/%v", handle))
	} else {
		author = append(author, externalURL)
	}

	if len(handle) == 0 {
		list := strings.FieldsFunc(externalURL, func(r rune) bool {
			return r == '/' || r == '@' || r == ':'
		})
		author = append(author, list[len(list)-1])

		return author
	}

	author = append(author, handle)

	return author
}

// GetContent get content from arweave
func (c *Client) GetContent(ctx context.Context, uri string, lensContent *LensContent) error {
	// get content
	content, err := metadata_url.GetFileByURL(uri)
	if err != nil {
		loggerx.Global().Error("[lens worker] GetContent: get file failed", zap.Error(err), zap.String("ipfs", uri))

		return err
	}

	if err = json.Unmarshal(content, &lensContent); err != nil {
		loggerx.Global().Error("[lens worker] GetContent: json unmarshal error", zap.Error(err), zap.String("json", string(content)), zap.String("ipfs", uri))

		return err
	}

	if len(lensContent.Content) == 0 && len(lensContent.Media) == 0 {
		return fmt.Errorf("content is nil, ipfs:%v", uri)
	}

	return nil
}

// CreatePost create the basic Post struct
func (c *Client) CreatePost(ctx context.Context, lensContent *LensContent, handle string) *metadata.Post {
	post := &metadata.Post{
		Body: lensContent.Content,
	}

	// handle the target post's author for lenster
	post.Author = c.GetLensAuthorURL(ctx, lensContent.ExternalURL, handle)

	for _, media := range lensContent.Media {
		post.Media = append(post.Media, metadata.Media{
			Address:  metadata_url.GetDirectURL(media.Item),
			MimeType: media.Type,
		})
	}
	return post
}

func (c *Client) FormatContent(ctx context.Context, opt *FormatOption) error {
	lensContent := LensContent{}
	if opt.Content == nil || json.Unmarshal(opt.Content, &lensContent) != nil {
		err := c.GetContent(ctx, opt.ContentURI, &lensContent)
		if err != nil {
			loggerx.Global().Error("[lens worker] FormatContent: GetContent error", zap.Error(err))

			return err
		}
	}

	// handle transfer fields
	opt.Transfer.Platform = lensContent.AppId

	// special handling for Share and Comment
	postFinal := &metadata.Post{}

	switch opt.ContentType {
	case Share:
		profile, err := c.GetProfile(opt.ProfileIdPointed)
		if err != nil {
			loggerx.Global().Error("[lens worker] FormatContent: GetProfile error", zap.Error(err))
			return err
		}

		postFinal.Target = c.CreatePost(ctx, &lensContent, profile.Handle)

		// get the correct type on Lens
		if len(lensContent.Attributes) > 0 && len(c.FormatTypeOnPlatform(lensContent.Attributes[0].Value)) > 0 {
			postFinal.Target.TypeOnPlatform = []string{c.FormatTypeOnPlatform(lensContent.Attributes[0].Value)}
		}

		// get the pub time of the target
		if !lensContent.CreatedOn.IsZero() {
			postFinal.Target.CreatedAt = lensContent.CreatedOn.Format(time.RFC3339)
		}

		// target
		postFinal.Target.ProfileID = opt.ProfileIdPointed
		postFinal.Target.PublicationID = opt.PubIdPointed
		postFinal.Target.TargetURL = c.GetLensRelatedURL(ctx, opt.ProfileIdPointed, opt.PubIdPointed)

		// post
		postFinal.TypeOnPlatform = []string{opt.ContentType}
		postFinal.Author = c.GetLensAuthorURL(ctx, "", opt.Handle)

	case Comment:
		postFinal = c.CreatePost(ctx, &lensContent, opt.Handle)
		postFinal.TypeOnPlatform = []string{opt.ContentType}

		contentURI, err := c.GetContentURI(ctx, opt.ProfileIdPointed, opt.PubIdPointed)
		if err != nil {
			loggerx.Global().Error("[lens worker] FormatContent-Comment: GetContentURI error", zap.Error(err))
			return err
		}
		// get the target content
		targetContent := LensContent{}
		err = c.GetContent(ctx, contentURI, &targetContent)
		if err != nil {
			loggerx.Global().Error("[lens worker] FormatContent-Comment: GetContentURI error", zap.Error(err))
			return err
		}

		// target
		postFinal.Target = c.CreatePost(ctx, &targetContent, "")
		postFinal.Target.ProfileID = opt.ProfileIdPointed
		postFinal.Target.PublicationID = opt.PubIdPointed
		postFinal.Target.TargetURL = c.GetLensRelatedURL(ctx, opt.ProfileIdPointed, opt.PubIdPointed)

		// get the correct type on Lens
		if len(targetContent.Attributes) > 0 && len(c.FormatTypeOnPlatform(targetContent.Attributes[0].Value)) > 0 {
			postFinal.Target.TypeOnPlatform = []string{c.FormatTypeOnPlatform(targetContent.Attributes[0].Value)}
		}

		// get the pub time of the target
		if !targetContent.CreatedOn.IsZero() {
			postFinal.Target.CreatedAt = targetContent.CreatedOn.Format(time.RFC3339)
		}

	default:
		postFinal = c.CreatePost(ctx, &lensContent, opt.Handle)
		postFinal.TypeOnPlatform = []string{opt.ContentType}
	}

	postFinal.ProfileID = opt.ProfileId
	postFinal.PublicationID = opt.PubId

	rawMetadata, err := json.Marshal(postFinal)
	if err != nil {
		return err
	}

	opt.Transfer.Metadata = rawMetadata

	return nil
}

func (c *Client) FormatTypeOnPlatform(input string) string {
	switch input {
	// treat all these types as post
	case "post", "image", "text_only", "video":
		return Post
	}

	return ""
}

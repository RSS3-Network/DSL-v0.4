package lens

import (
	"context"
	"encoding/json"
	"fmt"
	"math/big"
	"net/url"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/naturalselectionlabs/pregod/common/database/model"
	"github.com/naturalselectionlabs/pregod/common/database/model/metadata"
	"github.com/naturalselectionlabs/pregod/common/database/model/social"
	"github.com/naturalselectionlabs/pregod/common/datasource/ethereum"
	"github.com/naturalselectionlabs/pregod/common/datasource/ethereum/contract/lens"
	"github.com/naturalselectionlabs/pregod/common/datasource/ethereum/contract/lens/contract"
	lenscontract "github.com/naturalselectionlabs/pregod/common/datasource/ethereum/contract/lens/contract"
	"github.com/naturalselectionlabs/pregod/common/ethclientx"
	"github.com/naturalselectionlabs/pregod/common/ipfs"
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
	Transfer         *model.Transfer
	ProfileIdPointed *big.Int
	PubIdPointed     *big.Int
	ProfileId        *big.Int
	PubId            *big.Int
}

const (
	Post    = "post"
	Comment = "comment"
	Share   = "mirror"

	NotSupportedError = "not supported"
)

var SupportLensPlatform = map[string]bool{
	protocol.PlatformLens:                 true,
	protocol.PlatformLensOrb:              true,
	protocol.PlatformLensLenster:          true,
	protocol.PlatformLensLenstube:         true,
	protocol.PlatformLensLenstubeBytes:    true,
	protocol.PlatformLensLensterCrowdfund: true,
}

func (c *Client) GetProfile(blockNumber *big.Int, address string, profileID *big.Int) (*social.Profile, error) {
	if len(address) == 0 && profileID == nil {
		return nil, fmt.Errorf("empty profile")
	}

	ethereumClient, err := ethclientx.Global(protocol.NetworkPolygon)
	if err != nil {
		loggerx.Global().Error("[common] lens: ethclientx.Global err", zap.Error(err))

		return nil, err
	}

	lensHubContract, err := lenscontract.NewHub(lens.HubProxyContractAddress, ethereumClient)
	if err != nil {
		loggerx.Global().Error("[common] lens: NewHub err", zap.Error(err))

		return nil, err
	}

	if profileID == nil {
		profileID, err = lensHubContract.DefaultProfile(&bind.CallOpts{}, common.HexToAddress(address))
		if err != nil {
			loggerx.Global().Error("[common] lens: Handle DefaultProfile err", zap.Error(err))

			return nil, err
		}
	}

	if profileID.Int64() == 0 {
		loggerx.Global().Error("[common] lens: Handle getDefaultProfile is nil", zap.String("address", address))

		return nil, fmt.Errorf("DefaultProfile is nil")
	}

	if len(address) == 0 {
		owner, err := lensHubContract.OwnerOf(&bind.CallOpts{}, profileID)
		if err != nil {
			loggerx.Global().Error("[common] lens: OwnerOf error", zap.Error(err))

			return nil, err
		}

		address = owner.String()
	}

	result, err := lensHubContract.GetProfile(&bind.CallOpts{}, profileID)
	if err != nil {
		loggerx.Global().Error("[common] lens: GetProfile err", zap.Error(err))

		return nil, err
	}

	profile := &social.Profile{
		Address:     strings.ToLower(address),
		Network:     protocol.NetworkPolygon,
		Platform:    protocol.PlatformLens,
		Source:      protocol.SourceKurora,
		Name:        result.Handle,
		Handle:      result.Handle,
		URL:         fmt.Sprintf("https://lenster.xyz/u/%v", result.Handle),
		ProfileUris: []string{ipfs.GetDirectURL(result.ImageURI)},
	}

	return profile, nil
}

func (c *Client) HandleReceipt(ctx context.Context, transaction *model.Transaction) (transfers []model.Transfer, err error) {
	tracer := otel.Tracer("worker_lens")
	_, trace := tracer.Start(ctx, "worker_len:handleReceipt")

	defer func() { opentelemetry.Log(trace, transaction, transfers, err) }()

	// rpc
	ethclient, err := ethclientx.Global(protocol.NetworkPolygon)
	if err != nil {
		return nil, err
	}

	var sourceData ethereum.SourceData
	if err := json.Unmarshal(transaction.SourceData, &sourceData); err != nil {
		return nil, fmt.Errorf("failed to unmarshal source data: %w", err)
	}

	// parse log
	for _, log := range sourceData.Receipt.Logs {
		lensContract, err := contract.NewEvents(log.Address, ethclient)
		if err != nil {
			logrus.Errorf("[lens worker] handleReceipt: new events error, %v", err)

			continue
		}

		sourceData, err := json.Marshal(log)
		if err != nil {
			logrus.Errorf("[lens worker] marshal source data error, %v", err)

			continue
		}

		// common attributes
		transfer := model.Transfer{
			TransactionHash: transaction.Hash,
			Timestamp:       transaction.Timestamp,
			BlockNumber:     big.NewInt(transaction.BlockNumber),
			Index:           int64(log.Index),
			Network:         transaction.Network,
			SourceData:      sourceData,
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

	return transfers, nil
}

func (c *Client) HandlePostCreated(ctx context.Context, lensContract contract.Events, transaction *model.Transaction, transfer *model.Transfer, log types.Log) (err error) {
	event, err := lensContract.EventsFilterer.ParsePostCreated(log)
	if err != nil {
		loggerx.Global().Error("[lens worker] HandlePostCreated: ParsePostCreated error", zap.Error(err))

		return err
	}

	err = c.FormatContent(ctx, &FormatOption{
		ContentURI:  event.ContentURI,
		ContentType: Post,
		Transfer:    transfer,
		ProfileId:   event.ProfileId,
		PubId:       event.PubId,
	})
	if err != nil {
		loggerx.Global().Error("[lens worker] HandlePostCreated: FormatContent error", zap.Error(err))

		return err
	}

	profile, err := c.GetProfile(transfer.BlockNumber, "", event.ProfileId)
	if err != nil {
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

func (c *Client) HandleCommentCreated(ctx context.Context, lensContract contract.Events, transaction *model.Transaction, transfer *model.Transfer, log types.Log) (err error) {
	event, err := lensContract.EventsFilterer.ParseCommentCreated(log)
	if err != nil {
		loggerx.Global().Error("[lens worker] handleCommentCreated: ParseCommentCreated error", zap.Error(err))

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
	})

	if err != nil {
		loggerx.Global().Error("[lens worker] handleCommentCreated: FormatContent error", zap.Error(err))

		return err
	}

	profile, err := c.GetProfile(transfer.BlockNumber, "", event.ProfileId)
	if err != nil {
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

func (c *Client) HandleProfileCreated(ctx context.Context, lensContract contract.Events, transaction *model.Transaction, transfer *model.Transfer, log types.Log) (err error) {
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
		ProfileUris: []string{ipfs.GetDirectURL(event.ImageURI)},
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

func (c *Client) HandleMirrorCreated(ctx context.Context, lensContract contract.Events, transaction *model.Transaction, transfer *model.Transfer, log types.Log) (err error) {
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

	err = c.FormatContent(ctx, &FormatOption{
		ContentURI:       contentURI,
		ContentType:      Share,
		Transfer:         transfer,
		ProfileId:        event.ProfileId,
		PubId:            event.PubId,
		ProfileIdPointed: event.ProfileIdPointed,
		PubIdPointed:     event.PubIdPointed,
	})
	if err != nil {
		loggerx.Global().Error("[lens worker] HandleMirrorCreated: FormatContent error", zap.Error(err))
		return err
	}

	profile, err := c.GetProfile(transfer.BlockNumber, "", event.ProfileId)
	if err != nil {
		loggerx.Global().Error("[lens worker] HandleFollowNFTTransferred: GetProfile error", zap.Error(err))

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

func (c *Client) HandleFollowed(ctx context.Context, lensContract contract.Events, transaction *model.Transaction, transfer *model.Transfer, log types.Log) (transfers []model.Transfer, err error) {
	event, err := lensContract.EventsFilterer.ParseFollowed(log)
	if err != nil {
		loggerx.Global().Error("[lens worker] HandleFollowed: ParseFollowed error", zap.Error(err))
		return nil, err
	}

	transaction.Owner = strings.ToLower(event.Follower.String())
	transfer.AddressFrom = transaction.Owner

	for index, profileID := range event.ProfileIds {
		profile, err := c.GetProfile(transfer.BlockNumber, "", profileID)
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

func (c *Client) HandleFollowNFTTransferred(ctx context.Context, lensContract contract.Events, transaction *model.Transaction, transfer *model.Transfer, log types.Log) (err error) {
	event, err := lensContract.EventsFilterer.ParseFollowNFTTransferred(log)
	if err != nil {
		loggerx.Global().Error("[lens worker] HandleFollowNFTTransferred: ParseFollowNFTTransferred error", zap.Error(err))

		return err
	}

	// unfollow: burn
	if event.To != ethereum.AddressGenesis {
		return fmt.Errorf("not supported")
	}

	profile, err := c.GetProfile(transfer.BlockNumber, "", event.ProfileId)
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
	profileIdHex := []byte(hexutil.EncodeBig(profileId))
	pubIdHex := []byte(hexutil.EncodeBig(pubId))

	if len(profileIdHex)%2 == 1 {
		profileIdHex = append(profileIdHex[:2], append([]byte("0"), profileIdHex[2:]...)...)
	}
	if len(pubIdHex)%2 == 1 {
		pubIdHex = append(pubIdHex[:2], append([]byte("0"), pubIdHex[2:]...)...)
	}

	return fmt.Sprintf("https://lenster.xyz/posts/%v-%v", string(profileIdHex), string(pubIdHex))
}

// GetContent get content from arweave
func (c *Client) GetContent(ctx context.Context, uri string, lensContent *LensContent) error {
	// get content
	content, err := ipfs.GetFileByURL(uri)
	if err != nil {
		logrus.Errorf("[lens worker] handleReceipt: getContent error, %v, ipfs: %v", err, uri)
		return err
	}

	if err = json.Unmarshal(content, &lensContent); err != nil {
		logrus.Errorf("[lens worker] handleReceipt: json unmarshal error, %v, json: %v, ipfs: %v", err, string(content), uri)
		return err
	}

	if len(lensContent.Content) == 0 && len(lensContent.Media) == 0 {
		return fmt.Errorf("content is nil, ipfs:%v", uri)
	}

	return nil
}

// CreatePost create the basic Post struct
func (c *Client) CreatePost(ctx context.Context, lensContent *LensContent) *metadata.Post {
	post := &metadata.Post{
		Body: lensContent.Content,
	}

	// handle the target post's author for lenster
	if lensContent.ExternalURL != "" {
		externalUrl, err := url.Parse(lensContent.ExternalURL)
		if err == nil {
			post.Author = []string{
				lensContent.ExternalURL,
			}

			switch externalUrl.Host {
			case "lenster.xyz":
				// lenster url example: https://lenster.xyz/u/username.lens
				path := strings.Split(externalUrl.Path, "/")
				post.Author = []string{
					lensContent.ExternalURL,
					path[2],
				}
			case "orb.ac":
				// orb.ac url example: https://orb.ac/@username
				path := strings.Split(externalUrl.Path, "@")
				post.Author = []string{
					fmt.Sprintf("https://lenster.xyz/u/%v", path[1]),
					path[1],
				}
			}

		}
	}

	for _, media := range lensContent.Media {
		post.Media = append(post.Media, metadata.Media{
			Address:  ipfs.GetDirectURL(media.Item),
			MimeType: media.Type,
		})
	}
	return post
}

func (c *Client) FormatContent(ctx context.Context, opt *FormatOption) error {
	lensContent := LensContent{}
	err := c.GetContent(ctx, opt.ContentURI, &lensContent)
	if err != nil {
		loggerx.Global().Error("[lens worker] FormatContent: GetContent error", zap.Error(err))
		return err
	}
	// handle transfer fields
	opt.Transfer.Platform = lensContent.AppId

	post := c.CreatePost(ctx, &lensContent)
	post.TypeOnPlatform = []string{opt.ContentType}

	// special handling for Share and Comment
	postFinal := &metadata.Post{}

	switch opt.ContentType {
	case Share:
		// get the correct type on Lens
		if len(lensContent.Attributes) > 0 {
			post.TypeOnPlatform = []string{c.FormatTypeOnPlatform(lensContent.Attributes[0].Value)}
		}
		// get the pub time of the target

		if !lensContent.CreatedOn.IsZero() {
			post.CreatedAt = lensContent.CreatedOn.Format(time.RFC3339)
		}

		// target
		postFinal.Target = post
		postFinal.Target.ProfileID = opt.ProfileIdPointed
		postFinal.Target.PublicationID = opt.PubIdPointed
		post.TargetURL = c.GetLensRelatedURL(ctx, opt.ProfileIdPointed, opt.PubIdPointed)

		// post
		postFinal.TypeOnPlatform = []string{opt.ContentType}

	case Comment:
		postFinal = post
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
		postFinal.Target = c.CreatePost(ctx, &targetContent)
		postFinal.Target.ProfileID = opt.ProfileIdPointed
		postFinal.Target.PublicationID = opt.PubIdPointed
		postFinal.Target.TargetURL = c.GetLensRelatedURL(ctx, opt.ProfileIdPointed, opt.PubIdPointed)

		// get the correct type on Lens
		if len(targetContent.Attributes) > 0 {
			postFinal.Target.TypeOnPlatform = []string{c.FormatTypeOnPlatform(targetContent.Attributes[0].Value)}
		}

		// get the pub time of the target
		if !targetContent.CreatedOn.IsZero() {
			postFinal.Target.CreatedAt = targetContent.CreatedOn.Format(time.RFC3339)
		}

	default:
		postFinal = post
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
	return input
}

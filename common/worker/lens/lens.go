package lens

import (
	"context"
	"encoding/json"
	"fmt"
	"math/big"
	"net/http"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/naturalselectionlabs/pregod/common/database/model"
	"github.com/naturalselectionlabs/pregod/common/database/model/metadata"
	"github.com/naturalselectionlabs/pregod/common/database/model/social"
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

type Client struct {
	httpClient *http.Client
}

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
}

const (
	Post    = "post"
	Comment = "comment"
	Share   = "mirror"
)

func New() *Client {
	return &Client{
		httpClient: http.DefaultClient,
	}
}

func (c *Client) GetProfile(address string) (*social.Profile, error) {
	ethereumClient, err := ethclientx.Global(protocol.NetworkPolygon)
	if err != nil {
		logrus.Errorf("[common] lens: ethclientx.Global err, %v", err)

		return nil, err
	}

	lensHubContract, err := lenscontract.NewHub(lens.HubProxyContractAddress, ethereumClient)
	if err != nil {
		logrus.Errorf("[common] lens: NewHub err, %v", err)

		return nil, err
	}

	profileID, err := lensHubContract.DefaultProfile(&bind.CallOpts{}, common.HexToAddress(address))
	if err != nil {
		logrus.Errorf("[common] lens: Handle DefaultProfile err, %v", err)

		return nil, err
	}

	if profileID.Int64() == 0 {
		logrus.Infof("[common] lens: Handle getDefaultProfile is nil, address: %v", address)

		return nil, nil
	}

	result, err := lensHubContract.GetProfile(&bind.CallOpts{}, profileID)
	if err != nil {
		logrus.Errorf("[common] lens: GetProfile err, %v", err)

		return nil, err
	}

	profile := &social.Profile{
		Address:     address,
		Network:     protocol.NetworkPolygon,
		Platform:    protocol.PlatformLens,
		Source:      protocol.PlatformLens,
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

	receipt, err := ethclient.TransactionReceipt(ctx, common.HexToHash(transaction.Hash))
	if err != nil {
		logrus.Errorf("[lens worker] ethereumClient TransactionReceipt error, %v", err)

		return nil, err
	}

	// parse log
	for _, log := range receipt.Logs {
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
			BlockNumber:     big.NewInt(transaction.BlockNumber),
			Index:           int64(log.Index),
			Network:         transaction.Network,
			AddressFrom:     transaction.Owner,
			SourceData:      sourceData,
			Source:          protocol.SourcePregodETL,
			Platform:        protocol.PlatformLens,
		}

		var eventType string
		var handleErr error
		switch log.Topics[0] {
		case lens.EventHashPostCreated:
			handleErr = c.HandlePostCreated(ctx, *lensContract, &transfer, *log)
			eventType = "handlePostCreated"
		case lens.EventHashCommentCreated:
			handleErr = c.HandleCommentCreated(ctx, *lensContract, &transfer, *log)
			eventType = "handleCommentCreated"
		case lens.EventHashProfileCreated:
			handleErr = c.HandleProfileCreated(ctx, *lensContract, transaction, &transfer, *log)
			eventType = "handleProfileCreated"
		case lens.EventHashMirrorCreated:
			handleErr = c.HandleMirrorCreated(ctx, *lensContract, &transfer, *log)
			eventType = "handleMirrorCreated"
		default:
			continue
		}
		if handleErr != nil {
			logrus.Errorf("[lens worker] handleReceipt: %s error, %v", eventType, handleErr)

			continue
		}

		if transfer.Platform == protocol.PlatformLens || transfer.Platform == protocol.PlatformLenster {
			transfers = append(transfers, transfer)
		}
	}

	return transfers, nil
}

func (c *Client) HandlePostCreated(ctx context.Context, lensContract contract.Events, transfer *model.Transfer, log types.Log) (err error) {
	event, err := lensContract.EventsFilterer.ParsePostCreated(log)
	if err != nil {
		loggerx.Global().Error("[lens worker] HandlePostCreated: ParsePostCreated error", zap.Error(err))

		return err
	}

	err = c.FormatContent(ctx, &FormatOption{
		ContentURI:  event.ContentURI,
		ContentType: Post,
		Transfer:    transfer,
	})

	transfer.Tag, transfer.Type = filter.UpdateTagAndType(filter.TagSocial, transfer.Tag, filter.SocialPost, transfer.Type)
	transfer.RelatedUrls = append(transfer.RelatedUrls, c.GetLensRelatedURL(ctx, event.ProfileId, event.PubId))

	return nil
}

func (c *Client) HandleCommentCreated(ctx context.Context, lensContract contract.Events, transfer *model.Transfer, log types.Log) (err error) {
	event, err := lensContract.EventsFilterer.ParseCommentCreated(log)
	if err != nil {
		loggerx.Global().Error("[lens worker] handleCommentCreated: ParseCommentCreated error", zap.Error(err))

		return err
	}

	err = c.FormatContent(ctx, &FormatOption{
		ContentURI:       event.ContentURI,
		ContentType:      Comment,
		Transfer:         transfer,
		ProfileIdPointed: event.ProfileIdPointed,
		PubIdPointed:     event.PubIdPointed,
	})

	if err != nil {
		loggerx.Global().Error("[lens worker] handleCommentCreated: FormatContent error", zap.Error(err))

		return err
	}

	transfer.Timestamp = time.Unix(event.Timestamp.Int64(), 0)
	transfer.Tag, transfer.Type = filter.UpdateTagAndType(filter.TagSocial, transfer.Tag, filter.SocialComment, transfer.Type)
	transfer.RelatedUrls = append(transfer.RelatedUrls, c.GetLensRelatedURL(ctx, event.ProfileId, event.PubId))

	return nil
}

func (c *Client) HandleProfileCreated(ctx context.Context, lensContract contract.Events, transaction *model.Transaction, transfer *model.Transfer, log types.Log) (err error) {
	event, err := lensContract.EventsFilterer.ParseProfileCreated(log)
	if err != nil {
		loggerx.Global().Error("[lens worker] handleProfileCreated: ParseProfileCreated error", zap.Error(err))

		return err
	}

	transfer.Timestamp = time.Unix(event.Timestamp.Int64(), 0)
	transfer.AddressFrom = strings.ToLower(event.Creator.String())
	transfer.AddressTo = strings.ToLower(event.To.String())

	profile := social.Profile{
		Address:     strings.ToLower(event.To.String()),
		Platform:    protocol.PlatformLens,
		Network:     transaction.Network,
		Source:      transaction.Platform,
		Type:        filter.SocialProfileCreate,
		URL:         fmt.Sprintf("https://lenster.xyz/u/%v", event.Handle),
		Handle:      event.Handle,
		ProfileUris: []string{ipfs.GetDirectURL(event.ImageURI)},
	}

	rawMetadata, err := json.Marshal(&profile)
	if err != nil {
		return err
	}

	transfer.Metadata = rawMetadata
	transfer.Tag, transfer.Type = filter.UpdateTagAndType(filter.TagSocial, transfer.Tag, filter.SocialProfileCreate, transfer.Type)
	transfer.RelatedUrls = []string{
		fmt.Sprintf("https://lenster.xyz/u/%v", event.Handle),
		utils.GetTxHashURL(protocol.NetworkPolygon, transfer.TransactionHash),
	}

	transaction.Tag, transaction.Type = filter.UpdateTagAndType(filter.TagSocial, transaction.Tag, filter.SocialProfile, transaction.Type)

	return nil
}

func (c *Client) HandleMirrorCreated(ctx context.Context, lensContract contract.Events, transfer *model.Transfer, log types.Log) (err error) {
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
		ProfileIdPointed: event.ProfileIdPointed,
		PubIdPointed:     event.PubIdPointed,
	})
	if err != nil {
		loggerx.Global().Error("[lens worker] HandleMirrorCreated: FormatContent error", zap.Error(err))
		return err
	}

	transfer.Timestamp = time.Unix(event.Timestamp.Int64(), 0)
	transfer.Tag, transfer.Type = filter.UpdateTagAndType(filter.TagSocial, transfer.Tag, filter.SocialShare, transfer.Type)
	transfer.RelatedUrls = append(transfer.RelatedUrls, c.GetLensRelatedURL(ctx, event.ProfileId, event.PubId))

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

	if lensContent.ExternalURL != "" {
		post.Author = []string{lensContent.ExternalURL}
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
	c.GetContent(ctx, opt.ContentURI, &lensContent)

	// handle transfer fields
	opt.Transfer.Platform = lensContent.AppId

	opt.Transfer.RelatedUrls = []string{
		utils.GetTxHashURL(protocol.NetworkPolygon, opt.Transfer.TransactionHash),
	}

	post := c.CreatePost(ctx, &lensContent)

	// special handling for Share and Comment
	var postFinal *metadata.Post
	switch opt.ContentType {
	case Share:
		// get the correct type on Lens
		post.TypeOnPlatform = []string{lensContent.Attributes[0].Value}
		// get the pub time of the target
		post.CreatedAt = lensContent.CreatedOn.Format(time.RFC3339)
		post.TargetURL = c.GetLensRelatedURL(ctx, opt.ProfileIdPointed, opt.PubIdPointed)

		postFinal = &metadata.Post{
			Target:         post,
			TypeOnPlatform: []string{opt.ContentType},
		}
	case Comment:
		postFinal = post
		contentURI, err := c.GetContentURI(ctx, opt.ProfileIdPointed, opt.PubIdPointed)
		if err != nil {
			loggerx.Global().Error("[lens worker] FormatContent-Share: GetContentURI error", zap.Error(err))
			return err
		}
		// get the target content
		targetContent := LensContent{}
		c.GetContent(ctx, contentURI, &targetContent)

		postFinal.Target = c.CreatePost(ctx, &targetContent)
		// get the correct type on Lens
		postFinal.Target.TypeOnPlatform = []string{targetContent.Attributes[0].Value}
		// get the pub time of the target
		postFinal.Target.CreatedAt = targetContent.CreatedOn.Format(time.RFC3339)
		postFinal.Target.TargetURL = c.GetLensRelatedURL(ctx, opt.ProfileIdPointed, opt.PubIdPointed)

	default:
		post.TypeOnPlatform = []string{opt.ContentType}
		postFinal = post
	}

	rawMetadata, err := json.Marshal(postFinal)
	if err != nil {
		return err
	}

	opt.Transfer.Metadata = rawMetadata

	return nil
}

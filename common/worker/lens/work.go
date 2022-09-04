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
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/naturalselectionlabs/pregod/common/database/model"
	"github.com/naturalselectionlabs/pregod/common/database/model/metadata"
	"github.com/naturalselectionlabs/pregod/common/database/model/social"
	"github.com/naturalselectionlabs/pregod/common/datasource/ethereum/contract/lens"
	"github.com/naturalselectionlabs/pregod/common/datasource/ethereum/contract/lens/contract"
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
	ethereumClient *ethclient.Client
	httpClient     *http.Client
}

type LensContent struct {
	Description string             `json:"description"`
	Content     string             `json:"content"`
	ExternalURL string             `json:"external_url"`
	AppId       string             `json:"appId"`
	Media       []LensContentMedia `json:"media"`
}

type LensContentMedia struct {
	Item string `json:"item"`
	Type string `json:"type"`
}

func New(ethereumClientMap map[string]*ethclient.Client) *Client {
	return &Client{
		ethereumClient: ethereumClientMap[protocol.NetworkPolygon],
		httpClient:     http.DefaultClient,
	}
}

func (c *Client) HandleReceipt(ctx context.Context, transaction *model.Transaction) (transfers []model.Transfer, err error) {
	tracer := otel.Tracer("worker_lens")
	_, trace := tracer.Start(ctx, "worker_len:handleReceipt")

	defer func() { opentelemetry.Log(trace, transaction, transfers, err) }()

	// rpc
	receipt, err := c.ethereumClient.TransactionReceipt(ctx, common.HexToHash(transaction.Hash))
	if err != nil {
		logrus.Errorf("[lens worker] ethereumClient TransactionReceipt error, %v", err)

		return nil, err
	}

	// parse log
	for _, log := range receipt.Logs {
		lensContract, err := contract.NewEvents(log.Address, c.ethereumClient)
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

		transfers = append(transfers, transfer)
	}

	return transfers, nil
}

func (c *Client) HandlePostCreated(ctx context.Context, lensContract contract.Events, transfer *model.Transfer, log types.Log) (err error) {
	event, err := lensContract.EventsFilterer.ParsePostCreated(log)
	if err != nil {
		loggerx.Global().Error("[lens worker] handleReceipt: ParsePostCreated error", zap.Error(err))

		return err
	}

	transfer.Timestamp = time.Unix(event.Timestamp.Int64(), 0)

	// get content
	content, err := ipfs.GetFileByURL(event.ContentURI)
	if err != nil {
		logrus.Errorf("[lens worker] handleReceipt: getContent error, %v, ipfs: %v", err, event.ContentURI)
		return err
	}

	lensContent := LensContent{}
	if err = json.Unmarshal(content, &lensContent); err != nil {
		logrus.Errorf("[lens worker] handleReceipt: json unmarshal error, %v, json: %v, ipfs: %v", err, string(content), event.ContentURI)
		return err
	}

	post := &metadata.Post{
		Body: lensContent.Content,
	}
	for _, media := range lensContent.Media {
		post.Media = append(post.Media, metadata.Media{
			Address:  ipfs.GetDirectURL(media.Item),
			MimeType: media.Type,
		})
	}

	if len(post.Body) == 0 && len(post.Media) == 0 {
		return fmt.Errorf("content is nil, ipfs:%v", event.ContentURI)
	}

	rawMetadata, err := json.Marshal(post)
	if err != nil {
		return err
	}

	transfer.Platform = lensContent.AppId
	transfer.Metadata = rawMetadata
	transfer.Tag, transfer.Type = filter.UpdateTagAndType(filter.TagSocial, transfer.Tag, filter.SocialPost, transfer.Type)
	transfer.RelatedUrls = []string{
		c.GetLensRelatedURL(ctx, event.ProfileId, event.PubId),
		utils.GetTxHashURL(protocol.NetworkPolygon, transfer.TransactionHash),
	}

	return nil
}

func (c *Client) HandleCommentCreated(ctx context.Context, lensContract contract.Events, transfer *model.Transfer, log types.Log) (err error) {
	event, err := lensContract.EventsFilterer.ParseCommentCreated(log)
	if err != nil {
		loggerx.Global().Error("[lens worker] handleCommentCreated: ParsePostCreated error", zap.Error(err))

		return err
	}

	transfer.Timestamp = time.Unix(event.Timestamp.Int64(), 0)

	// get content
	content, err := ipfs.GetFileByURL(event.ContentURI)
	if err != nil {
		logrus.Errorf("[lens worker] handleCommentCreated: getContent error, %v, ipfs: %v", err, event.ContentURI)
		return err
	}

	lensContent := LensContent{}
	if err = json.Unmarshal(content, &lensContent); err != nil {
		logrus.Errorf("[lens worker] handleCommentCreated: json unmarshal error, %v, json: %v, ipfs: %v", err, string(content), event.ContentURI)
		return err
	}

	post := &metadata.Post{
		Body: lensContent.Content,
	}
	for _, media := range lensContent.Media {
		post.Media = append(post.Media, metadata.Media{
			Address:  ipfs.GetDirectURL(media.Item),
			MimeType: media.Type,
		})
	}

	if len(post.Body) == 0 && len(post.Media) == 0 {
		return fmt.Errorf("content is nil, ipfs:%v", event.ContentURI)
	}

	// get content pointed
	post.Target, err = c.GetContenPointed(ctx, event.ProfileIdPointed, event.PubIdPointed, c.ethereumClient)
	if err != nil {
		logrus.Errorf("[lens worker] handleCommentCreated: getContenPointed error, %v", err)
		return err
	}

	rawMetadata, err := json.Marshal(post)
	if err != nil {
		return err
	}

	transfer.Platform = lensContent.AppId
	transfer.Metadata = rawMetadata
	transfer.Tag, transfer.Type = filter.UpdateTagAndType(filter.TagSocial, transfer.Tag, filter.SocialComment, transfer.Type)
	transfer.RelatedUrls = []string{
		c.GetLensRelatedURL(ctx, event.ProfileId, event.PubId),
		utils.GetTxHashURL(protocol.NetworkPolygon, transfer.TransactionHash),
	}

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
		Address:  strings.ToLower(event.To.String()),
		Platform: protocol.PlatformLens,
		Network:  transaction.Network,
		Source:   transaction.Platform,
		Type:     filter.SocialProfileCreate,
		URL:      ipfs.GetDirectURL(event.ImageURI),
		Handle:   event.Handle,
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
		loggerx.Global().Error("[lens worker] handleMirrorCreated: ParsePostCreated error", zap.Error(err))

		return err
	}

	transfer.Timestamp = time.Unix(event.Timestamp.Int64(), 0)

	// get content
	share, err := c.GetContenPointed(ctx, event.ProfileId, event.PubId, c.ethereumClient)
	if err != nil {
		logrus.Errorf("[lens worker] handleMirrorCreated: getContenPointed error, %v", err)
		return err
	}

	// get content pointed
	share.Target, err = c.GetContenPointed(ctx, event.ProfileIdPointed, event.PubIdPointed, c.ethereumClient)
	if err != nil {
		logrus.Errorf("[lens worker] handleMirrorCreated: getContenPointed error, %v", err)
		return err
	}

	rawMetadata, err := json.Marshal(share)
	if err != nil {
		return err
	}

	transfer.Metadata = rawMetadata
	transfer.Tag, transfer.Type = filter.UpdateTagAndType(filter.TagSocial, transfer.Tag, filter.SocialShare, transfer.Type)
	transfer.RelatedUrls = []string{
		c.GetLensRelatedURL(ctx, event.ProfileId, event.PubId),
		utils.GetTxHashURL(protocol.NetworkPolygon, transfer.TransactionHash),
	}

	return nil
}

func (c *Client) GetContenPointed(ctx context.Context, profileId *big.Int, pubId *big.Int, ethereumClient *ethclient.Client) (*metadata.Post, error) {
	// rpc
	iLensHub, err := contract.NewILensHub(lens.HubProxyContractAddress, ethereumClient)
	if err != nil {
		logrus.Errorf("[lens worker] NewILensHub error, %v", err)

		return nil, err
	}

	contentURI, err := iLensHub.GetContentURI(&bind.CallOpts{}, profileId, pubId)
	if err != nil {
		logrus.Errorf("[lens worker] iLensHub GetContentURI error, %v", err)

		return nil, err
	}

	// get content
	content, err := ipfs.GetFileByURL(contentURI)
	if err != nil {
		logrus.Errorf("[lens worker] getContenPointed: getContent error, %v, ipfs: %v", err, contentURI)
		return nil, err
	}

	lensContent := LensContent{}
	if err = json.Unmarshal(content, &lensContent); err != nil {
		logrus.Errorf("[lens worker] getContenPointed: json unmarshal error, %v, json: %v, ipfs: %v", err, string(content), contentURI)
		return nil, err
	}

	post := &metadata.Post{
		Body: lensContent.Content,
	}
	for _, media := range lensContent.Media {
		post.Media = append(post.Media, metadata.Media{
			Address:  ipfs.GetDirectURL(media.Item),
			MimeType: media.Type,
		})
	}

	if len(post.Body) == 0 && len(post.Media) == 0 {
		return nil, fmt.Errorf("content is nil, ipfs:%v", contentURI)
	}

	return post, nil
}

func (c *Client) GetLensRelatedURL(ctx context.Context, profileId *big.Int, pubId *big.Int) string {
	profileIdHex := []byte(hexutil.EncodeBig(profileId))
	pubIdHex := []byte(hexutil.EncodeBig(pubId))

	if len(profileIdHex) == 3 {
		profileIdHex = append(profileIdHex[:2], 48, profileIdHex[2])
	}
	if len(pubIdHex) == 3 {
		pubIdHex = append(pubIdHex[:2], 48, pubIdHex[2])
	}

	return fmt.Sprintf("https://lenster.xyz/posts/%v-%v", string(profileIdHex), string(pubIdHex))
}

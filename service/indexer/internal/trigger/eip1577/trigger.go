package eip1577

import (
	"archive/tar"
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"math/big"
	"net/http"
	"path/filepath"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	shell "github.com/ipfs/go-ipfs-api"
	"github.com/mmcdole/gofeed"
	"github.com/naturalselectionlabs/pregod/common/database"
	"github.com/naturalselectionlabs/pregod/common/database/model"
	"github.com/naturalselectionlabs/pregod/common/database/model/metadata"
	"github.com/naturalselectionlabs/pregod/common/datasource/ethereum"
	"github.com/naturalselectionlabs/pregod/common/ethclientx"
	"github.com/naturalselectionlabs/pregod/common/protocol"
	"github.com/naturalselectionlabs/pregod/common/protocol/filter"
	"github.com/naturalselectionlabs/pregod/common/utils/shedlock"
	"github.com/naturalselectionlabs/pregod/service/indexer/internal/trigger"
	"github.com/samber/lo"
	goens "github.com/wealdtech/go-ens/v3"
	"gorm.io/gorm/clause"
)

const (
	ProtocolIPNS = "ipns"
	ProtocolIPFS = "ipfs"
)

var (
	ErrorUnknownFeedType     = errors.New("unknown feed type")
	ErrorNoDefaultDomain     = errors.New("no default domain")
	ErrorNoContentHash       = errors.New("no content hash")
	ErrorInvalidContentURI   = errors.New("invalid content uri")
	ErrorUnsupportedProtocol = errors.New("unsupported protocol")
	ErrorUnsupportedNetwork  = errors.New("unsupported network")
)

var _ trigger.Trigger = (*internal)(nil)

type internal struct {
	employer     *shedlock.Employer
	ipfsEndpoint string
}

func (i *internal) Name() string {
	return "eip1577"
}

func (i *internal) Networks() []string {
	return []string{
		protocol.NetworkEthereum, // ENS
	}
}

func (i *internal) Handle(ctx context.Context, message *protocol.Message) error {
	lockerName := fmt.Sprintf("eip1577:%s@trigger", message.Address)

	if !i.employer.DoLock(lockerName, time.Minute*2) {
		return fmt.Errorf("%v locked", lockerName)
	}

	ctx, cancel := context.WithCancel(ctx)

	go func(ctx context.Context) {
		for {
			time.Sleep(time.Second)

			if err := i.employer.Renewal(ctx, lockerName, time.Minute); err != nil {
				return
			}
		}
	}(ctx)

	defer cancel()

	name, err := i.reverseResolveName(ctx, message.Network, common.HexToAddress(message.Address))
	if err != nil {
		return err
	}

	message.Name = name

	contentHash, err := i.resolveContentHash(ctx, message.Network, name)
	if err != nil {
		if errors.Is(err, ErrorNoDefaultDomain) || errors.Is(err, ErrorNoContentHash) {
			return nil
		}

		return err
	}

	contentURI, err := goens.ContenthashToString(contentHash)
	if err != nil {
		return err
	}

	response, err := i.fetchContentURI(ctx, contentURI)
	if err != nil {
		return err
	}

	defer func() {
		_ = response.Close()
	}()

	transactions, err := i.handleFile(ctx, message, tar.NewReader(response.Output))
	if err != nil {
		return err
	}

	return i.storeTransactions(ctx, message, transactions)
}

func (i *internal) fetchContentURI(ctx context.Context, contentURI string) (*shell.Response, error) {
	splits := strings.Split(contentURI, "/")

	if len(splits) != 3 {
		return nil, ErrorInvalidContentURI
	}

	var (
		err             error
		contentProtocol = splits[1]
		contentHash     = splits[2]
	)

	if contentProtocol != ProtocolIPNS && contentProtocol != ProtocolIPFS {
		return nil, ErrorUnsupportedProtocol
	}

	ipfsShell := shell.NewShell(i.ipfsEndpoint)

	if contentProtocol == ProtocolIPNS {
		if contentHash, err = ipfsShell.Resolve(contentHash); err != nil {
			return nil, err
		}
	}

	return ipfsShell.Request("get", contentHash).Option("create", true).Send(context.Background())
}

func (i *internal) handleFile(ctx context.Context, message *protocol.Message, reader *tar.Reader) ([]model.Transaction, error) {
	var (
		transactions = make([]model.Transaction, 0)
		buffer       = bytes.NewBuffer([]byte{})
	)

	for {
		header, err := reader.Next()
		if err != nil {
			if errors.Is(err, io.EOF) {
				break
			}

			return nil, err
		}

		if header.Typeflag != tar.TypeReg {
			continue
		}

		n, err := io.Copy(buffer, reader)
		if err != nil {
			return nil, err
		}

		internalTransactions, err := i.handleData(ctx, message, filepath.Base(header.Name), buffer.Next(int(n)))
		if err != nil {
			if errors.Is(err, ErrorUnknownFeedType) {
				continue
			}

			return nil, err
		}

		transactions = append(transactions, internalTransactions...)
	}

	return transactions, nil
}

func (i *internal) handleData(ctx context.Context, message *protocol.Message, filename string, data []byte) ([]model.Transaction, error) {
	switch contentType := http.DetectContentType(data); contentType {
	case "text/xml; charset=utf-8", "text/plain; charset=utf-8":
		switch filename {
		case "article.json":
			return nil, ErrorUnknownFeedType
		case "planet.json":
			return i.buildTransactionsByPlanet(ctx, message, data)
		default:
			return i.buildTransactions(ctx, message, data)
		}
	default:
		return nil, ErrorUnknownFeedType
	}
}

func (i *internal) buildTransactions(ctx context.Context, message *protocol.Message, data []byte) ([]model.Transaction, error) {
	feedType := gofeed.DetectFeedType(bytes.NewReader(data))

	if feedType == gofeed.FeedTypeUnknown {
		return nil, ErrorUnknownFeedType
	}

	parser := gofeed.NewParser()

	feed, err := parser.Parse(bytes.NewReader(data))
	if err != nil {
		return nil, err
	}

	transactions := make([]model.Transaction, 0, len(feed.Items))

	for _, item := range feed.Items {
		if item.PublishedParsed == nil {
			return nil, ErrorUnknownFeedType
		}

		transaction, err := i.buildTransaction(ctx, message, item.Title, item.Content, item.Link, *item.PublishedParsed)
		if err != nil {
			return nil, err
		}

		transactions = append(transactions, *transaction)
	}

	return transactions, nil
}

func (i *internal) buildTransactionsByPlanet(ctx context.Context, message *protocol.Message, data []byte) ([]model.Transaction, error) {
	var planet Planet

	if err := json.Unmarshal(data, &planet); err != nil {
		return nil, err
	}

	transactions := make([]model.Transaction, 0, len(planet.Articles))

	for _, article := range planet.Articles {
		transaction, err := i.buildTransaction(ctx, message, article.Title, article.Content, article.Link, time.UnixMilli(int64(article.Created*1000)+time.Date(2001, 0, 0, 0, 0, 0, 0, time.UTC).UnixMilli()))
		if err != nil {
			return nil, err
		}

		transactions = append(transactions, *transaction)
	}

	return transactions, nil
}

func (i *internal) buildTransaction(ctx context.Context, message *protocol.Message, title, content, url string, createdAt time.Time) (*model.Transaction, error) {
	success := true

	metadataRaw, err := json.Marshal(metadata.Post{
		Title: title,
		Body:  content,
	})
	if err != nil {
		return nil, err
	}

	if strings.HasPrefix(url, "/") {
		url = fmt.Sprintf("https://%s.limo%s", message.Name, url)
	}

	hash := crypto.Keccak256Hash([]byte(url)).String()

	return &model.Transaction{
		Timestamp:   createdAt,
		Hash:        hash,
		Owner:       message.Address,
		AddressFrom: message.Address,
		Network:     protocol.NetworkEIP1577,
		Platform:    protocol.NetworkEIP1577,
		Source:      "IPFS",
		Tag:         filter.TagSocial,
		Type:        filter.SocialPost,
		Success:     &success,
		Transfers: []model.Transfer{
			{
				TransactionHash: hash,
				Timestamp:       createdAt,
				BlockNumber:     big.NewInt(0),
				Tag:             filter.TagSocial,
				Type:            filter.SocialPost,
				AddressFrom:     message.Address,
				Metadata:        metadataRaw,
				Network:         protocol.NetworkEIP1577,
				Platform:        protocol.NetworkEIP1577,
				Source:          "IPFS",
				RelatedUrls: []string{
					url,
				},
			},
		},
	}, nil
}

func (i *internal) storeTransactions(ctx context.Context, message *protocol.Message, transactions []model.Transaction) error {
	databaseClient := database.Global()

	databaseTransaction := databaseClient.Begin()

	clauseOnConflict := clause.OnConflict{
		UpdateAll: true,
	}

	if err := databaseTransaction.Clauses(clauseOnConflict).Create(&transactions).Error; err != nil {
		_ = databaseTransaction.Rollback()

		return err
	}

	transfers := lo.Map(transactions, func(transaction model.Transaction, _ int) model.Transfer {
		return transaction.Transfers[0]
	})

	if err := databaseTransaction.Clauses(clauseOnConflict).Create(&transfers).Error; err != nil {
		_ = databaseTransaction.Rollback()

		return err
	}

	return databaseTransaction.Commit().Error
}

func (i *internal) buildRegistry(ctx context.Context, network string, ethereumClient *ethclient.Client) (*goens.Registry, error) {
	// Reverse resolver
	var contractAddress common.Address

	switch network {
	case protocol.NetworkEthereum:
		contractAddress = common.HexToAddress("0x00000000000C2E074eC69A0dFb2997BA6C7d2e1e")
	default:
		return nil, ErrorUnsupportedNetwork
	}

	return goens.NewRegistryAt(ethereumClient, contractAddress)
}

func (i *internal) reverseResolveName(ctx context.Context, network string, address common.Address) (string, error) {
	ethereumClient, err := ethclientx.Global(network)
	if err != nil {
		return "", err
	}

	registry, err := i.buildRegistry(ctx, network, ethereumClient)
	if err != nil {
		return "", err
	}

	reverseResolverAddress, err := registry.ResolverAddress(fmt.Sprintf("%x.addr.reverse", address.Bytes()))
	if err != nil {
		return "", err
	}

	if reverseResolverAddress == ethereum.AddressGenesis {
		return "", ErrorNoDefaultDomain
	}

	reverseResolver, err := goens.NewReverseResolverAt(ethereumClient, reverseResolverAddress)
	if err != nil {
		return "", err
	}

	return reverseResolver.Name(address)
}

func (i *internal) resolveContentHash(ctx context.Context, network, name string) ([]byte, error) {
	ethereumClient, err := ethclientx.Global(network)
	if err != nil {
		return nil, err
	}

	registry, err := i.buildRegistry(ctx, network, ethereumClient)
	if err != nil {
		return nil, err
	}

	resolver, err := registry.Resolver(name)
	if err != nil {
		return nil, err
	}

	contentHash, err := resolver.Contenthash()
	if err != nil {
		return nil, err
	}

	if len(contentHash) == 0 {
		return nil, ErrorNoContentHash
	}

	return contentHash, nil
}

func New(employer *shedlock.Employer, ipfsEndpoint string) trigger.Trigger {
	return &internal{
		employer:     employer,
		ipfsEndpoint: ipfsEndpoint,
	}
}

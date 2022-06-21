package ens

import (
	"context"
	"embed"
	"encoding/json"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/naturalselectionlabs/pregod/common/database/model"
	"github.com/naturalselectionlabs/pregod/common/protocol"
	"github.com/naturalselectionlabs/pregod/common/shedlock"
	"github.com/naturalselectionlabs/pregod/service/crawler/internal/config"
	"github.com/naturalselectionlabs/pregod/service/crawler/internal/crawler"
	"github.com/naturalselectionlabs/pregod/service/crawler/internal/crawler/ens/contract"
	rabbitmq "github.com/rabbitmq/amqp091-go"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

var (
	_ crawler.Crawler = (*service)(nil)

	//go:embed contract/event.abi
	abiFileSystem embed.FS
)

type service struct {
	ethClient       *ethclient.Client
	abiClient       abi.ABI
	databaseClient  *gorm.DB
	rabbitmqChannel *rabbitmq.Channel
	employer        *shedlock.Employer
}

func New(databaseClient *gorm.DB,
	rabbitmqChannel *rabbitmq.Channel,
	employer *shedlock.Employer,
	config *config.Config,
) crawler.Crawler {
	crawler := &service{
		databaseClient:  databaseClient,
		rabbitmqChannel: rabbitmqChannel,
		employer:        employer,
	}

	var err error

	// get ethclient
	crawler.ethClient, err = ethclient.Dial(config.Gateway.EthEndpoint)
	if err != nil {
		logrus.Errorf("[crawler] ens: ethclient Dial error, %v", err)

		return nil
	}

	// get abi
	abiFile, err := abiFileSystem.Open("contract/event.abi")
	if err != nil {
		logrus.Errorf("[crawler] ens: open abi file error, %v", err)

		return nil
	}

	crawler.abiClient, err = abi.JSON(abiFile)
	if err != nil {
		logrus.Errorf("[crawler] ens: abi file parse error, %v", err)

		return nil
	}

	return crawler
}

func (s *service) Name() string {
	return "crawler:ens_contract"
}

func (s *service) Run() error {
	ctx := context.Background()

	go func() {
		time.Sleep(time.Second)

		s.employer.Renewal(ctx, s.Name(), 10*time.Second)
	}()

	if err := s.subscribeEns(); err != nil {
		return err
	}

	return nil
}

func (s *service) subscribeEns() error {
	query := ethereum.FilterQuery{
		Addresses: []common.Address{
			common.HexToAddress("0x283Af0B28c62C092C9727F1Ee09c02CA627EB7F5"),
		},
	}

	logs := make(chan types.Log)
	sub, err := s.ethClient.SubscribeFilterLogs(context.Background(), query, logs)
	if err != nil {
		logrus.Errorf("[crawler] ens: ethclient SubscribeFilterLogs error, %v", err)

		return err
	}

	for {
		select {
		case err := <-sub.Err():
			logrus.Errorf("[crawler] ens: ethclient subscribe error, %v", err)
		case vLog := <-logs:
			if vLog.Topics[0] == contract.TopicHashNameRegistered {

				data := contract.NameRegisteredData{}

				// parse contract log
				if err := s.abiClient.UnpackIntoInterface(&data, "NameRegistered", vLog.Data); err != nil {
					logrus.Errorf("[crawler] ens: parse data into NameRegistered error, %v", err)

					continue
				}

				// get owner by topics
				owner := common.HexToAddress(vLog.Topics[2].Hex())

				// get block details
				block, err := s.ethClient.BlockByNumber(context.Background(), big.NewInt(int64(vLog.BlockNumber)))
				if err != nil {
					logrus.Errorf("[crawler] ens: get block error, %v", err)

					continue
				}

				// save ens data into db
				ens := &model.Domains{
					TransactionHash: vLog.TxHash.Bytes(),
					Type:            "ens",
					Name:            data.Name,
					AddressOwner:    owner.Bytes(),
					ExpiredAt:       time.Unix(data.Expires.Int64(), 0),
					Source:          "crawler",
					BlockTimestamp:  time.Unix(int64(block.Time()), 0),
				}

				if err := s.databaseClient.Create(ens).Error; err != nil {
					logrus.Errorf("[crawler] ens: db insert error, %v", err)

					continue
				}

				// create a rabbitmq job to index the latest user data
				go s.createRabbitmqJob(owner.String())
			}
		}
	}
}

// createRabbitmqJob create a rabbitmq job to index the latest user data
func (s *service) createRabbitmqJob(address string) error {
	networks := []string{
		protocol.NetworkEthereum, protocol.NetworkPolygon, protocol.NetworkBinanceSmartChain,
		protocol.NetworkArweave, protocol.NetworkXDAI, protocol.NetworkZkSync, protocol.NetworkCrossbell,
	}

	for _, network := range networks {
		message := protocol.Message{
			Address: address,
			Network: network,
		}

		messageData, err := json.Marshal(&message)
		if err != nil {
			return err
		}

		if err := s.rabbitmqChannel.Publish(protocol.ExchangeJob, "", false, false, rabbitmq.Publishing{
			ContentType: protocol.ContentTypeJSON,
			Body:        messageData,
		}); err != nil {
			return err
		}
	}

	return nil
}

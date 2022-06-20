package ens

import (
	"embed"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/naturalselectionlabs/pregod/service/crawler/internal/crawler"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type service struct {
	ethClient      *ethclient.Client
	abiClient      abi.ABI
	databaseClient *gorm.DB
}

var (
	_ crawler.Crawler = (*service)(nil)

	//go:embed contract/event.abi
	abiFileSystem embed.FS
)

func New(databaseClient *gorm.DB, gateway string) (c crawler.Crawler, err error) {
	s := &service{
		databaseClient: databaseClient,
	}

	// get ethclient
	s.ethClient, err = ethclient.Dial(gateway)
	if err != nil {
		logrus.Errorf("[crawler] ens: ethclient Dial error, %v", err)

		return nil, err
	}

	// get abi
	abiFile, err := abiFileSystem.Open("contract/event.abi")
	if err != nil {
		logrus.Errorf("[crawler] ens: open abi file error, %v", err)

		return
	}

	s.abiClient, err = abi.JSON(abiFile)
	if err != nil {
		logrus.Errorf("[crawler] ens: abi file parse error, %v", err)

		return
	}
}

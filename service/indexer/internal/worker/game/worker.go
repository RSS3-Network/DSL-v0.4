package game

import (
	"context"
	"embed"
	"encoding/csv"

	"github.com/naturalselectionlabs/pregod/common/database"
	"github.com/naturalselectionlabs/pregod/common/database/model/game"
	"github.com/naturalselectionlabs/pregod/internal/token"
	"github.com/naturalselectionlabs/pregod/service/indexer/internal/worker"
	"gorm.io/gorm/clause"
)

var _ worker.Worker = (*service)(nil)

//go:embed contract/contract.csv
var gameContract embed.FS
var gameNetwork []string

type service struct {
	tokenClient *token.Client
}

func (s *service) Name() string {
	return "game"
}

func (s *service) Networks() []string {
	return gameNetwork
}

func New() worker.Worker {
	return &service{
		tokenClient: token.New(),
	}
}

func (s *service) Initialize(ctx context.Context) error {
	file, err := gameContract.Open("contract/contract.csv")
	if err != nil {
		return err
	}
	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		return err
	}

	contract := make([]game.GameContract, 0)
	for i, record := range records {
		if i == 0 {
			continue
		}

		contract = append(contract, game.GameContract{
			ContractAddress: record[0],
			Network:         record[1],
			Platform:        record[2],
		})

	}

	if len(contract) == 0 {
		return nil
	}

	if err := database.Global().
		Model((*game.GameContract)(nil)).
		Clauses(clause.OnConflict{
			DoNothing: true,
		}).
		Create(contract).
		Error; err != nil {
		return err
	}

	return nil
}

package dao_test

import (
	"math/big"
	"testing"
	"time"

	"github.com/naturalselectionlabs/pregod/common/database"
	"github.com/naturalselectionlabs/pregod/common/database/model"
	"github.com/naturalselectionlabs/pregod/service/hub/internal/config"
	"github.com/naturalselectionlabs/pregod/service/hub/internal/server/dao"
	"github.com/stretchr/testify/assert"
)

func init() {
	config.Initialize()

	db, err := database.Dial(config.ConfigHub.Postgres.String(), false)
	if err != nil {
		panic(err)
	}
	database.ReplaceGlobal(db)
}

func setup() {
	if err := database.Global().Migrator().CreateTable(&model.Transfer{}); err != nil {
		panic(err)
	}

	ts := time.Date(2022, 12, 1, 1, 1, 1, 1, time.UTC)

	database.Global().CreateInBatches([]model.Transfer{
		{
			TransactionHash: "0x00",
			Timestamp:       ts,
			BlockNumber:     big.NewInt(0),
			Tag:             "social",
			Type:            "post",
			Index:           0,
			AddressFrom:     "0x00",
			AddressTo:       "0x00",
			Metadata:        []byte(`{"body":"你好"}`), // 2
			Network:         "eth",
			Platform:        "",
		},
		{
			TransactionHash: "0x01",
			Timestamp:       ts,
			BlockNumber:     big.NewInt(0),
			Tag:             "social",
			Type:            "post",
			Index:           0,
			AddressFrom:     "0x00",
			AddressTo:       "0x00",
			Metadata:        []byte(`{"body":"hello, world"}`), // 2
			Network:         "eth",
			Platform:        "",
		},
		{
			TransactionHash: "0x02",
			Timestamp:       ts,
			BlockNumber:     big.NewInt(0),
			Tag:             "social",
			Type:            "post",
			Index:           0,
			AddressFrom:     "0x00",
			AddressTo:       "0x00",
			Metadata:        []byte(`{"body":"hello! 你好 hello, world"}`), // 5
			Network:         "eth",
			Platform:        "",
		},
		{
			TransactionHash: "0x03",
			Timestamp:       ts,
			BlockNumber:     big.NewInt(0),
			Tag:             "social",
			Type:            "post",
			Index:           0,
			AddressFrom:     "0x01", // other user
			AddressTo:       "0x00",
			Metadata:        []byte(`{"body":"hello! 你好 hello, world"}`),
			Network:         "eth",
			Platform:        "",
		},
	}, 100)
}

func teardown() {
	database.Global().Migrator().DropTable(&model.Transfer{})
}

func TestGetWordsCountPercentileByAddress(t *testing.T) {
	setup()
	defer teardown()

	count, percentage, err := dao.GetWordsCountPercentileByAddress("0x00")
	assert.NoError(t, err)
	assert.EqualValues(t, 2+2+5, count)
	assert.EqualValues(t, 30, percentage)
}

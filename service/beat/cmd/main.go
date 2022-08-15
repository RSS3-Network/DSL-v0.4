package main

import (
	"encoding/json"
	"fmt"

	"github.com/naturalselectionlabs/pregod/common/protocol"
	"github.com/naturalselectionlabs/pregod/common/utils/logger"
	"github.com/naturalselectionlabs/pregod/service/beat/internal/config"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"go.uber.org/zap"
)

var rootCommand = cobra.Command{
	Use:           "beat",
	SilenceUsage:  true,
	SilenceErrors: true,
	Version:       protocol.Version,
}

func main() {
	// init db
	config.Initialize()

	if err := logger.Initialize(string(config.ConfigBeat.Mode)); err != nil {
		logrus.Fatalln(err)
	}

	rootCommand.RunE = beat

	if err := rootCommand.Execute(); err != nil {
		logger.Global().Fatal("beat execution failed", zap.Error(err))
	}
}

const (
	EsIndex = "pregod-v1-metrics"

	notesCountSQL        = `select reltuples::bigint as total from pg_class where relname = 'transfers'`
	transactionsCountSQL = `select reltuples::bigint as total from pg_class where relname = 'transactions'`
	sampleTotal          = 400000
	sampleSQL            = `SELECT COUNT(DISTINCT owner) FROM (SELECT owner FROM transactions limit 400000) AS temp`
)

func beat(cmd *cobra.Command, args []string) error {
	var notesCount int64
	var addressesCount float64

	var transactionsCount int64
	var sampleCount int64

	// notes count
	if err := config.ConfigBeat.DatabaseClient.Raw(notesCountSQL).Scan(&notesCount).Error; err != nil {
		return err
	}

	// unique addresses
	if err := config.ConfigBeat.DatabaseClient.Raw(transactionsCountSQL).Scan(&transactionsCount).Error; err != nil {
		return err
	}
	if transactionsCount > sampleTotal {
		if err := config.ConfigBeat.DatabaseClient.Raw(sampleSQL).Scan(&sampleCount).Error; err != nil {
			return err
		}
		addressesCount = (float64(sampleCount) / sampleTotal) * float64(transactionsCount)
	} else {
		if err := config.ConfigBeat.DatabaseClient.Raw(sampleSQL).Scan(&addressesCount).Error; err != nil {
			return err
		}
	}

	report := map[string]interface{}{
		"index":           EsIndex,
		"notes_count":     notesCount,
		"addresses_count": addressesCount,
	}

	output, err := json.Marshal(report)
	if err != nil {
		return err
	}

	fmt.Printf("[DATABEAT]%v\n", string(output))

	return nil
}

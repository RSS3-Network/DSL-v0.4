package main

import (
	"github.com/naturalselectionlabs/pregod/common/databeat"
	"github.com/naturalselectionlabs/pregod/common/protocol"
	"github.com/naturalselectionlabs/pregod/common/utils/loggerx"
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

	if err := loggerx.Initialize(string(config.ConfigBeat.Mode)); err != nil {
		logrus.Fatalln(err)
	}

	rootCommand.RunE = beat

	if err := rootCommand.Execute(); err != nil {
		loggerx.Global().Fatal("beat execution failed", zap.Error(err))
	}
}

const (
	notesCountSQL     = `select reltuples::bigint as total from pg_class where relname = 'transfers'`
	addressesCountSQL = `select reltuples::bigint as total from pg_class where relname = 'address'`
	sampleTotal       = 400000
	sampleSQL         = `SELECT COUNT(DISTINCT owner) FROM (SELECT owner FROM transactions limit 400000) AS temp`

	transfersPerTagSQL = `SELECT tag, COUNT(*) FROM (SELECT tag FROM transfers order by updated_at desc limit 100000) as temp GROUP BY tag;`

	profilesCountSQL       = `select reltuples::bigint as total from pg_class where relname = 'profiles'`
	profilesPerPlatformSQL = `SELECT platform, COUNT(*) FROM profiles GROUP BY platform;`

	top20AddressesSQL = `SELECT address, count FROM address ORDER BY count DESC LIMIT 20`
)

func beat(cmd *cobra.Command, args []string) error {
	var notesCount int64
	var addressesCount float64

	// notes count
	if err := config.ConfigBeat.DatabaseClient.Raw(notesCountSQL).Scan(&notesCount).Error; err != nil {
		return err
	}

	// addresses count
	if err := config.ConfigBeat.DatabaseClient.Raw(addressesCountSQL).Scan(&addressesCount).Error; err != nil {
		return err
	}

	// transfers per tag in recent 100000 transfers
	var transfersPerTag []struct {
		Tag   string `json:"tag"`
		Count int64  `json:"count"`
	}
	if err := config.ConfigBeat.DatabaseClient.Raw(transfersPerTagSQL).Scan(&transfersPerTag).Error; err != nil {
		return err
	}

	// profiles Count
	var profilesCount int64
	if err := config.ConfigBeat.DatabaseClient.Raw(profilesCountSQL).Scan(&profilesCount).Error; err != nil {
		return err
	}

	// profiles per platform
	var profilesPerPlatform []struct {
		Platform string `json:"platform"`
		Count    int64  `json:"count"`
	}
	if err := config.ConfigBeat.DatabaseClient.Raw(profilesPerPlatformSQL).Scan(&profilesPerPlatform).Error; err != nil {
		return err
	}

	// top 20 addresses
	var top20Addresses []struct {
		Address string `json:"address"`
		Count   int64  `json:"count"`
	}
	if err := config.ConfigBeat.DatabaseClient.Raw(top20AddressesSQL).Scan(&top20Addresses).Error; err != nil {
		return err
	}

	_, _ = databeat.DBMetrics.Beat(map[string]interface{}{
		"notes_count":           notesCount,
		"addresses_count":       addressesCount,
		"transfers_count":       transfersPerTag,
		"profiles_count":        profilesCount,
		"profiles_per_platform": profilesPerPlatform,
		"top_20_addresses":      top20Addresses,
	})

	return nil
}

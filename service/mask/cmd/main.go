package main

import (
	"bufio"
	"encoding/json"
	"errors"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/naturalselectionlabs/pregod/common/database/model"
	"github.com/naturalselectionlabs/pregod/common/databeat"
	"github.com/naturalselectionlabs/pregod/common/protocol"
	"github.com/naturalselectionlabs/pregod/common/utils/loggerx"
	"github.com/naturalselectionlabs/pregod/service/mask/internal/config"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/willf/bloom"
	"go.uber.org/multierr"
	"go.uber.org/zap"
)

const (
	MASK_ADDRESSES_FILE = "https://s3.amazonaws.com/public.firefly/wallet-addr.txt"
)

var rootCommand = cobra.Command{
	Use:           "mask",
	SilenceUsage:  true,
	SilenceErrors: true,
	Version:       protocol.Version,
}

func main() {
	// init db
	config.Initialize()

	if err := loggerx.Initialize(string(config.ConfigMask.Mode)); err != nil {
		logrus.Fatalln(err)
	}

	rootCommand.RunE = calculateSocial

	if err := rootCommand.Execute(); err != nil {
		loggerx.Global().Fatal("beat execution failed", zap.Error(err))
	}
}

type SocialRes struct {
	Owner    string `gorm:"column:owner"`
	Network  string `gorm:"column:network"`
	Platform string `gorm:"column:platform"`
	Type     string `gorm:"column:type"`
}

// https://mask.notion.site/RSS3-Social-Statistic-4fa9dcdeb2594b94b3f4a7cb979bc536
// Execute at 1:00 UTC every day, fetch 500K addresses, count their social activities in the last 24 hours.
// There are basically two ideas:
//    1. Iterate through these 500K addresses in batches, query from the database, just like what Mask is doing now.
//    2. Iterate directly through the social data of the last 24 hours in the database, and then check if they are in these 500K addresses.
// As of 20230413, the daily social data volume seems to be less than 100K, and many of them should not be among these 500K addresses, so the cost of the second method is much lower.

// Steps:
// 1. Download the file, put 500K lines into the Bloom filter, and delete the file.
// 2. Prepare an empty map to store the results.
// 3. Iterate through the social data of the last 24 hours in the database, and then check if they are in these 500K addresses.
// 4. Submit the results to Mask in batches.
func calculateSocial(*cobra.Command, []string) error {
	utcNow := time.Now().UTC()
	dbClient := config.ConfigMask.DatabaseClient
	loggerx.Global().Info("start to calculate social", zap.Time("utc now", utcNow))

	filter, err := downloadAndInitBloomFilter()
	if err != nil {
		return err
	}

	result := map[string]map[string]uint{}

	// Iterate through yesterday's social data, fetch 30 minutes of data each time, a total of 48 times.
	yesterdayStart := utcNow.Add(-24 * time.Hour).Truncate(24 * time.Hour)
	yesterdayEnd := yesterdayStart.Add(24 * time.Hour)
	loggerx.Global().Info("yesterday start", zap.Time("yesterday start", yesterdayStart))
	loggerx.Global().Info("yesterday end", zap.Time("yesterday end", yesterdayEnd))
	for start := yesterdayStart; start.Before(yesterdayEnd); start = start.Add(30 * time.Minute) {
		end := start.Add(30 * time.Minute)
		if end.After(yesterdayEnd) {
			end = yesterdayEnd
		}

		findRes := []SocialRes{}

		if err := dbClient.Model(&model.Transaction{}).Select("owner", "network", "platform", "type").Where(
			"timestamp >= ? AND timestamp < ? AND tag = ?",
			start.Format("2006-01-02 15:04:05 -0700"), // You need to specify the time zone, otherwise, this value will be treated as Eastern Time by default
			end.Format("2006-01-02 15:04:05 -0700"),
			"social",
		).Find(&findRes).Error; err != nil {
			return err
		}
		loggerx.Global().Info("find social data", zap.Time("start", start), zap.Time("end", end), zap.Int("count", len(findRes)))

		for _, item := range findRes {
			socialKey := getSocialKey(item.Network, item.Platform, item.Type)
			if socialKey != "" && filter.TestString(item.Owner) {
				if _, ok := result[item.Owner]; !ok {
					result[item.Owner] = map[string]uint{}
				}
				result[item.Owner][socialKey]++
			}
		}
		loggerx.Global().Info("result size", zap.Int("result size", len(result)))
	}

	return send2Mask(result)
}

func downloadAndInitBloomFilter() (*bloom.BloomFilter, error) {
	filter := bloom.NewWithEstimates(500000, 0.01)

	// download
	tmpFile, err := os.CreateTemp("", "mask")
	if err != nil {
		return nil, err
	}
	defer os.Remove(tmpFile.Name())
	resp, err := http.Get(MASK_ADDRESSES_FILE)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// read by line
	scanner := bufio.NewScanner(resp.Body)
	for scanner.Scan() {
		filter.AddString(strings.ToLower(strings.TrimSpace(scanner.Text())))
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}
	loggerx.Global().Info("filter init done", zap.Uint("filter size", filter.Cap()))

	// test
	if !filter.TestString("0xe02a52a553acf14cd5552e53d48dc0fc072978d8") {
		loggerx.Global().Error("test failed")
		return nil, errors.New("test failed")
	}

	return filter, nil
}

func getSocialKey(network, platform, _type string) string {
	platforms := []string{"Lens", "Mirror", "Farcaster"}
	networks := []string{"EIP-1577"}

	matched := ""
	for _, p := range platforms {
		if p == platform {
			matched = p
			break
		}
	}
	for _, n := range networks {
		if n == network {
			matched = n
			break
		}
	}
	if matched != "" {
		matched = strings.ReplaceAll(matched, "-", "")
		return strings.ToLower(matched) + "_" + strings.ToLower(_type)
	}

	return ""
}

var beatForMask = databeat.DataModel{
	Name: "pregod-mask-social-udpate",
	Fields: []string{
		"address",
		"value",
	},
}

func send2Mask(res map[string]map[string]uint) error {
	var reserr error

	chunkSize := 50
	chunkMap := map[string]map[string]uint{}
	for address, value := range res {
		chunkMap[address] = value
		if len(chunkMap) >= chunkSize {
			if err := sendChunk2Mask(chunkMap); err != nil {
				reserr = multierr.Append(reserr, err)
			}

			chunkMap = map[string]map[string]uint{}
		}
	}

	return reserr
}

// TODO
func sendChunk2Mask(chunk map[string]map[string]uint) error {
	for address, value := range chunk {
		// TODO: here should be a real Mask API call
		// value json, max 2000 bytes
		valueJson, err := json.Marshal(value)
		if err != nil {
			return err
		}
		if len(valueJson) > 2000 {
			valueJson = valueJson[:2000]
		}

		_, _ = beatForMask.Beat(map[string]interface{}{
			"address": address,
			"value":   string(valueJson),
		})
	}

	return nil
}

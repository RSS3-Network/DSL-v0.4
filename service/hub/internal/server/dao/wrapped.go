package dao

import (
	"context"
	"encoding/json"
	"fmt"
	"regexp"
	"strings"
	"unicode/utf8"

	"github.com/go-resty/resty/v2"
	"github.com/naturalselectionlabs/pregod/common/database"
	dbModel "github.com/naturalselectionlabs/pregod/common/database/model"
	"github.com/naturalselectionlabs/pregod/common/database/model/metadata"
	bridge "github.com/naturalselectionlabs/pregod/common/database/model/transaction"
	"github.com/naturalselectionlabs/pregod/common/protocol/filter"
	"github.com/naturalselectionlabs/pregod/service/hub/internal/server/handler/wrapped/lens"
	"github.com/naturalselectionlabs/pregod/service/hub/internal/server/model"
	lop "github.com/samber/lo/parallel"
	"go.opentelemetry.io/otel"
)

var (
	// https://www.notion.so/rss3/social-x-4ec64be6f03146809c3d3d46abac0264
	wordsCountPercentiles = []uint{0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 2, 2, 2, 2, 2, 2, 3, 3, 3, 3, 4, 4, 4, 5, 6, 6, 7, 7, 8, 9, 10, 11, 12, 13, 14, 15, 17, 19, 20, 22, 24, 27, 29, 33, 36, 40, 43, 47, 52, 56, 62, 68, 73, 80, 87, 99, 109, 120, 133, 149, 166, 184, 205, 231, 255, 282, 314, 352, 395, 436, 475, 511, 571, 631, 696, 767, 854, 939, 1034, 1136, 1241, 1369, 1489, 1630, 1777, 1963, 2183, 2420, 2687, 3078, 3468, 3917, 4493, 5103, 5816, 6654, 7901, 10567, 16244, 33439}
	nonAcsii              = regexp.MustCompile(`([^\x00-\x7F])`)
	lensClient            = lens.NewClient()
)

func CountSocial(c context.Context, request model.GetRequest) (model.SocialResult, error) {
	tracer := otel.Tracer("dao.CountSocial")
	_, postgresSnap := tracer.Start(c, "postgres")

	defer postgresSnap.End()

	request.Address = strings.ToLower(request.Address)

	var result model.SocialResult

	// the where condition for all the queries
	condition := fmt.Sprintf("WHERE owner = '%s' AND DATE_PART('year', timestamp) = '%d'", request.Address, 2022)

	// count post, comment, follow
	database.Global().
		Raw(fmt.Sprintf(`SELECT
					COUNT(*) FILTER ( WHERE type = 'post') AS post,
					COUNT(*) FILTER ( WHERE type = 'comment') AS comment
				 FROM transactions 
				 %s`, condition)).Scan(&result)

	// followers and followings from farcaster
	countStruct := struct{ Follower, Following int64 }{}

	database.EthDb().
		Raw(fmt.Sprintf(`SELECT follower_count as follower, following_count as following
				FROM dataset_farcaster.profiles
				WHERE '%s' = ANY (signer_address);`, request.Address)).Scan(&countStruct)

	result.Following += countStruct.Following
	result.Follower += countStruct.Follower

	// followers and followings from Lens
	_ = lensClient.GetFollowStat(c, &result, request.Address)

	// followers and followings from crossbell
	_ = getCSBFollowStats(c, &result, request.Address)

	// get hashes of the longest and the shortest posts
	hashStruct := struct{ Longest, Shortest string }{}

	database.Global().
		Raw(fmt.Sprintf(`SELECT (SELECT transaction_hash
        FROM (SELECT transaction_hash, metadata ::jsonb ->> 'body' AS body
              FROM transfers
              WHERE type = 'post'
                AND transaction_hash IN
                    (SELECT hash
                     FROM transactions
                     %s)) sub
        WHERE sub.body IS NOT NULL
        ORDER BY LENGTH(body::TEXT) DESC
        LIMIT 1) AS longest,
       (SELECT transaction_hash
        FROM (SELECT transaction_hash, metadata ::jsonb ->> 'body' AS body
              FROM transfers
              WHERE type = 'post'
                AND transaction_hash IN
                    (SELECT hash
                     FROM transactions
                     %s)) sub
        WHERE sub.body IS NOT NULL
        ORDER BY LENGTH(body::TEXT) ASC
        LIMIT 1) AS shortest;`, condition, condition)).Scan(&hashStruct)

	result.LongestPost = getPost(hashStruct.Longest)
	result.ShortestPost = getPost(hashStruct.Shortest)

	database.Global().
		Model(&dbModel.Transaction{}).
		Select("platform AS name, COUNT(*) AS count").
		Where("owner = ? ", request.Address).
		Where("tag = ? ", filter.TagSocial).
		Where("platform != ? ", "").
		Where("DATE_PART('year', timestamp) = ?", "2022").
		Group("platform").
		Order("count DESC").
		Scan(&result.List)

	// calculate web3 social score

	scoreStruct := struct{ Platform, Score int64 }{}

	database.Global().Raw(`SELECT COUNT(DISTINCT source) as Platform, COUNT(*) as Score
			FROM transfers
			WHERE transaction_hash IN (SELECT hash
									   FROM transactions
									   WHERE owner = ?
										 AND tag = 'social'
										 AND type IN ('post', 'comment')
										 AND network = 'crossbell'
										 AND DATE_PART('year', timestamp) = '2022')
			  AND platform IN ('xSync', 'OperatorSync');`, request.Address).Scan(&scoreStruct)

	result.SocialScore = scoreStruct.Platform*50 + scoreStruct.Score*5

	// words count percentile
	totalWord, wordPercentile, err := GetWordsCountPercentileByAddress(request.Address)
	if err != nil {
		return result, err
	}
	result.TotalWord = totalWord
	result.WordPercentile = wordPercentile - 1

	return result, nil
}

func CountSearch(c context.Context, request model.GetRequest) (model.SearchResult, error) {
	tracer := otel.Tracer("dao.CountSearch")
	_, postgresSnap := tracer.Start(c, "postgres")

	defer postgresSnap.End()

	request.Address = strings.ToLower(request.Address)

	var result model.SearchResult

	// count post, comment, follow
	database.Global().
		Model(&dbModel.Address{}).
		Select("count").
		Where("address = ?", request.Address).
		Scan(&result)

	return result, nil
}

func CountGas(c context.Context, request model.GetRequest) (model.GasResult, error) {
	tracer := otel.Tracer("dao.CountGas")
	_, postgresSnap := tracer.Start(c, "postgres")

	defer postgresSnap.End()

	request.Address = strings.ToLower(request.Address)

	var result model.GasResult

	// the where condition for all the queries
	condition := fmt.Sprintf("address_from = '%s' AND network = 'ethereum' AND DATE_PART('year', timestamp) = '%d' AND fee IS NOT NULL", request.Address, 2022)

	// calculate gas: total and highest
	database.Global().
		Model(&dbModel.Transaction{}).
		Select("SUM(fee::numeric) as total, MAX(fee::NUMERIC) as highest").
		Where(condition).
		Scan(&result)

	database.Global().
		Model(&dbModel.Transaction{}).
		Select("hash as highest_hash, timestamp as highest_date").
		Where(condition).
		Order("fee::NUMERIC DESC").
		Limit(1).
		Scan(&result)

	return result, nil
}

func CountTransaction(c context.Context, request model.GetRequest) (model.TxResult, error) {
	tracer := otel.Tracer("dao.CountTransaction")
	_, postgresSnap := tracer.Start(c, "postgres")

	defer postgresSnap.End()

	request.Address = strings.ToLower(request.Address)

	var result model.TxResult

	// transactions initiated by the address
	database.Global().
		Model(&dbModel.Transaction{}).
		Select("network, COUNT(*) AS total").
		Where("address_from = ?", request.Address).
		Group("network").
		Scan(&result.Initiate)

	// transactions received by the address
	database.Global().
		Model(&dbModel.Transaction{}).
		Select("network, COUNT(*) AS total").
		Where("address_to = ?", request.Address).
		Group("network").
		Scan(&result.Receive)

	database.Global().Raw(`
		SELECT COUNT(timestamp), TO_CHAR(timestamp,'MMDD') AS date
		FROM transfers
		WHERE transaction_hash IN (SELECT hash
								   FROM transactions
								   WHERE owner = ?
									 AND DATE_PART('year', timestamp) = '2022')
		GROUP BY date`, request.Address).
		Scan(&result.Heatmap)

	return result, nil
}

func GetNFT(c context.Context, request model.GetRequest) (model.NFTResult, error) {
	tracer := otel.Tracer("dao.GetNFT")
	_, postgresSnap := tracer.Start(c, "postgres")

	defer postgresSnap.End()

	request.Address = strings.ToLower(request.Address)
	// the where condition for all the queries
	condition := fmt.Sprintf("WHERE owner = '%s' AND tag = 'collectible' AND (type = 'mint'  OR type = 'trade') AND DATE_PART('year', timestamp) = '%d'", request.Address, 2022)

	var result model.NFTResult

	var list []model.NFT

	database.Global().
		Raw(fmt.Sprintf(`SELECT address_from AS from, address_to AS to, metadata, timestamp, type, platform
			FROM transfers
			WHERE transaction_hash IN (SELECT hash
									   FROM transactions
									   %s)
			  AND (type = 'mint'  OR type = 'trade')
			ORDER BY timestamp DESC;`, condition)).Scan(&list)

	result.Total = len(list)

	for i, current := range list {
		var nft metadata.Token
		err := json.Unmarshal(current.Metadata, &nft)
		if err != nil {
			continue
		}

		// if current.Type == filter.CollectibleMint {
		//	result.Mint = append(result.Mint, nft)
		//	continue
		// }
		//
		// if current.From == request.Address {
		//	result.Sold = append(result.Sold, nft)
		// } else if current.To == request.Address {
		//	result.Bought = append(result.Bought, nft)
		// }
		//
		// if i == 0 {
		//	result.Last = &model.NFTSingle{
		//		Metadata:  nft,
		//		Timestamp: current.Timestamp,
		//	}
		// }

		if i == len(list)-1 {
			result.First = &model.NFTSingle{
				Metadata:  nft,
				Timestamp: current.Timestamp,
				Platform:  current.Platform,
			}
		}
	}

	return result, nil
}

func GetDApp(c context.Context, request model.GetRequest) (model.DAppResult, error) {
	tracer := otel.Tracer("dao.GetDApp")
	_, postgresSnap := tracer.Start(c, "postgres")

	defer postgresSnap.End()

	request.Address = strings.ToLower(request.Address)

	var result model.DAppResult

	database.Global().
		Model(&dbModel.Transaction{}).
		Select("platform AS name, COUNT(*) AS count").
		Where("owner = ? ", request.Address).
		Where("platform != ? ", "").
		Where("DATE_PART('year', timestamp) = ?", "2022").
		Group("platform").
		Order("count DESC").
		Scan(&result.List)

	return result, nil
}

func GetDeFi(c context.Context, request model.GetRequest) (model.DeFiResult, error) {
	tracer := otel.Tracer("dao.GetDeFi")
	_, postgresSnap := tracer.Start(c, "postgres")

	defer postgresSnap.End()

	request.Address = strings.ToLower(request.Address)

	var result model.DeFiResult

	// DeFi platform list
	database.Global().
		Model(&dbModel.Transaction{}).
		Select("platform AS name, COUNT(*) AS count").
		Where("owner = ? ", request.Address).
		Where("tag = ? ", filter.TagExchange).
		Where("platform != ? ", "").
		Where("DATE_PART('year', timestamp) = ?", "2022").
		Group("platform").
		Order("count DESC").
		Scan(&result.PlatformList)

	// swap pairs
	database.Global().
		Raw(`SELECT metadata::jsonb -> 'from' ->> 'name' AS from,metadata::jsonb -> 'to' ->> 'name' AS to
			FROM transfers
			WHERE transaction_hash IN (SELECT hash
									   FROM "transactions"
									   WHERE OWNER = ?
										 AND tag = 'exchange'
										 AND type = 'swap'
										 AND DATE_PART('year', TIMESTAMP) = '2022')
			GROUP BY "from","to"`, request.Address).
		Scan(&result.SwapPair)

	// liquidity actions
	var liquidityJson []json.RawMessage
	database.Global().
		Raw(`select metadata
			FROM transfers
			WHERE transaction_hash IN (SELECT hash
									   FROM "transactions"
									   WHERE OWNER = ?
										 AND tag = 'exchange'
										 AND type = 'liquidity'
										 AND DATE_PART('year', TIMESTAMP) = '2022')`, request.Address).
		Scan(&liquidityJson)

	for _, m := range liquidityJson {
		var current metadata.Liquidity
		err := json.Unmarshal(m, &current)
		if err != nil {
			continue
		}

		switch current.Action {
		case "add":
			result.Liquidity.Add = append(result.Liquidity.Add, current.Tokens...)
		case "remove":
			result.Liquidity.Remove = append(result.Liquidity.Remove, current.Tokens...)
		case "supply":
			result.Liquidity.Supply = append(result.Liquidity.Supply, current.Tokens...)
		case "withdraw":
			result.Liquidity.Withdraw = append(result.Liquidity.Withdraw, current.Tokens...)
		case "borrow":
			result.Liquidity.Borrow = append(result.Liquidity.Borrow, current.Tokens...)
		case "repay":
			result.Liquidity.Repay = append(result.Liquidity.Repay, current.Tokens...)
		case "collect":
			result.Liquidity.Collect = append(result.Liquidity.Collect, current.Tokens...)
		}
	}

	// bridge actions
	var bridgeJson []json.RawMessage
	database.Global().
		Raw(`select metadata
			FROM transfers
			WHERE transaction_hash IN (SELECT hash
									   FROM "transactions"
									   WHERE OWNER = ?
										 AND tag = 'transaction'
										 AND type = 'bridge'
										 AND DATE_PART('year', TIMESTAMP) = '2022')`, request.Address).
		Scan(&bridgeJson)

	for _, m := range bridgeJson {
		var current bridge.Bridge
		err := json.Unmarshal(m, &current)
		if err != nil {
			continue
		}
		result.Bridge = append(result.Bridge, current)
	}

	return result, nil
}

// Get words count and percentile
// example: 你提交的内容包含 m 个词，超越了 n% 的用户
func GetWordsCountPercentileByAddress(address string) (uint, uint, error) {
	db := database.Global()

	var metadata []string

	if err := db.Model(&dbModel.Transfer{}).
		Select("metadata").
		Where("address_from = ? AND tag = 'social' AND type in ('post', 'comment') AND DATE_PART('year', timestamp) = '2022'", address).
		Scan(&metadata).Error; err != nil {
		return 0, 0, err
	}

	var wordsCount uint

	for _, v := range metadata {
		md := make(map[string]interface{})
		if err := json.Unmarshal([]byte(v), &md); err != nil {
			continue
		}
		if c, ok := md["body"].(string); ok {
			wordsCount += count(c)
		}
	}

	return wordsCount, getWordsCountPercentile(wordsCount), nil
}

func getWordsCountPercentile(wordsCount uint) uint {
	for i, v := range wordsCountPercentiles {
		if wordsCount < v {
			return uint(i)
		}
	}
	return 100
}

func getPost(hash string) *model.Post {
	// Get post by hash
	var post model.Post

	database.Global().Model(&dbModel.Transfer{}).
		Select("transaction_hash as hash, metadata, timestamp, platform").
		Where("transaction_hash = ? ", hash).
		First(&post)

	return &post
}

func count(s string) uint {
	runeCount := 0
	res := nonAcsii.ReplaceAllStringFunc(s, func(ss string) string {
		trimedSs := strings.TrimSpace(ss)
		if trimedSs != "" {
			runeCount += utf8.RuneCountInString(trimedSs)
		}
		return " "
	})
	count := runeCount + len(strings.Fields(res))
	return uint(count)
}

func getCSBFollowStats(ctx context.Context, result *model.SocialResult, address string) error {
	client := resty.New()
	client.SetBaseURL("https://indexer.crossbell.io/")

	type Character struct {
		CharacterId int    `json:"characterId"`
		Handle      string `json:"handle"`
	}

	type CharacterList struct {
		List []Character `json:"list"`
	}

	// get CSB character list
	var characterList CharacterList

	_, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetContext(ctx).
		SetResult(&characterList).
		Get(fmt.Sprintf("/v1/addresses/%s/characters", address))
	if err != nil {
		return err
	}

	response := struct {
		Count int64 `json:"count"`
	}{}

	lop.ForEach(characterList.List, func(character Character, i int) {
		// get following
		_, err := client.R().
			SetHeader("Content-Type", "application/json").
			SetQueryParam("limit", "0").
			SetQueryParam("linkType", "follow").
			SetContext(ctx).
			SetResult(&response).
			Get(fmt.Sprintf("/v1/characters/%d/links", character.CharacterId))

		if err == nil {
			result.Following += response.Count
		}

		// get follower
		_, err = client.R().
			SetHeader("Content-Type", "application/json").
			SetQueryParam("limit", "0").
			SetQueryParam("linkType", "follow").
			SetContext(ctx).
			SetResult(&response).
			Get(fmt.Sprintf("/v1/characters/%d/backlinks", character.CharacterId))
		if err == nil {
			result.Follower += response.Count
		}
	})

	return nil
}

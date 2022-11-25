package dao

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"

	"github.com/naturalselectionlabs/pregod/common/database"
	dbModel "github.com/naturalselectionlabs/pregod/common/database/model"
	"github.com/naturalselectionlabs/pregod/common/database/model/metadata"
	bridge "github.com/naturalselectionlabs/pregod/common/database/model/transaction"
	"github.com/naturalselectionlabs/pregod/common/protocol/filter"
	"github.com/naturalselectionlabs/pregod/service/hub/internal/server/model"
	"go.opentelemetry.io/otel"
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
					COUNT(*) FILTER ( WHERE type = 'comment') AS comment,
					COUNT(*) FILTER ( WHERE type = 'follow') AS follow
				 FROM transactions 
				 %s`, condition)).Scan(&result)

	// get hashes of the longest and the shortest posts
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
        LIMIT 1) AS longest_hash,
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
        LIMIT 1) AS shortest_hash;`, condition, condition)).Scan(&result)

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
	condition := fmt.Sprintf("address_from = '%s' AND DATE_PART('year', timestamp) = '%d' AND fee IS NOT NULL", request.Address, 2022)

	// calculate gas: total and highest
	database.Global().
		Model(&dbModel.Transaction{}).
		Select("SUM(fee::numeric) as total, MAX(fee::NUMERIC) as highest").
		Where(condition).
		Scan(&result)

	database.Global().
		Model(&dbModel.Transaction{}).
		Select("hash as highest_hash").
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

	return result, nil
}

func GetNFT(c context.Context, request model.GetRequest) (model.NFTResult, error) {
	tracer := otel.Tracer("dao.GetNFT")
	_, postgresSnap := tracer.Start(c, "postgres")

	defer postgresSnap.End()

	request.Address = strings.ToLower(request.Address)
	// the where condition for all the queries
	condition := fmt.Sprintf("WHERE owner = '%s' AND tag = 'collectible' AND type = 'trade' AND DATE_PART('year', timestamp) = '%d'", request.Address, 2022)

	var result model.NFTResult

	var list []model.NFT

	database.Global().
		Raw(fmt.Sprintf(`SELECT address_from AS from, address_to AS to, metadata, timestamp
			FROM transfers
			WHERE transaction_hash IN (SELECT hash
									   FROM transactions
									   %s)
			  AND type = 'trade'
			  AND address_from = '%s' 
			ORDER BY timestamp DESC;`, condition, request.Address)).Scan(&list)

	result.Total = len(list)

	for i, current := range list {
		var nft metadata.Token
		err := json.Unmarshal(current.Metadata, &nft)
		if err != nil {
			continue
		}

		if current.From == request.Address {
			result.Sold = append(result.Sold, nft)
		} else if current.To == request.Address {
			result.Bought = append(result.Bought, nft)
		}

		if i == 0 {
			result.Last = &model.NFTSingle{
				Metadata:  nft,
				Timestamp: current.Timestamp,
			}
		}

		if i == len(list)-1 {
			result.First = &model.NFTSingle{
				Metadata:  nft,
				Timestamp: current.Timestamp,
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

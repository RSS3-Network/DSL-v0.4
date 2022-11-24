package dao

import (
	"context"
	"fmt"
	"strings"

	"github.com/naturalselectionlabs/pregod/common/database"
	dbModel "github.com/naturalselectionlabs/pregod/common/database/model"
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
        FROM (SELECT transaction_hash, metadata ::jsonb -> 'body' AS body
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
        FROM (SELECT transaction_hash, metadata ::jsonb -> 'body' AS body
              FROM transfers
              WHERE type = 'post'
                AND transaction_hash IN
                    (SELECT hash
                     FROM transactions
                     %s)) sub
        WHERE sub.body IS NOT NULL
        ORDER BY LENGTH(body::TEXT) ASC
        LIMIT 1) AS shortest_hash;`, condition, condition)).Scan(&result)

	return result, nil
}

func CountSearch(c context.Context, request model.GetRequest) (model.SearchResult, error) {
	tracer := otel.Tracer("dao.CountSocial")
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

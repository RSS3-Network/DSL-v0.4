package dao

import (
	"context"
	"strings"

	"github.com/naturalselectionlabs/pregod/common/database"
	"github.com/naturalselectionlabs/pregod/common/database/model/social"
	"github.com/naturalselectionlabs/pregod/service/hub/internal/server/model"
	"github.com/sirupsen/logrus"
	"go.opentelemetry.io/otel"
	"gorm.io/gorm/clause"
)

// getProfiles get profile data from database
func GetProfiles(c context.Context, request model.GetRequest) ([]social.Profile, error) {
	tracer := otel.Tracer("getProfiles")
	_, postgresSnap := tracer.Start(c, "postgres")

	defer postgresSnap.End()

	dbResult := make([]social.Profile, 0)

	sql := database.Global().Model(&social.Profile{}).Where("LOWER(address) = ? ",
		strings.ToLower(request.Address),
	)

	if len(request.Network) > 0 {
		for i, v := range request.Network {
			request.Network[i] = strings.ToLower(v)
		}
		sql = sql.Where("network IN ?", request.Network)
	}

	if len(request.Platform) > 0 {
		for i, v := range request.Platform {
			request.Platform[i] = strings.ToLower(v)
		}
		sql = sql.Where("platform IN ?", request.Platform)
	}

	sql = sql.Limit(model.DefaultLimit)

	if err := sql.Find(&dbResult).Error; err != nil {
		logrus.Errorf("[profile] getProfilesDatabase error, %v", err)

		return nil, err
	}

	return dbResult, nil
}

// batchGetProfiles get profile data from database
func BatchGetProfiles(c context.Context, request model.BatchGetProfilesRequest) ([]social.Profile, error) {
	tracer := otel.Tracer("batchGetProfiles")
	_, postgresSnap := tracer.Start(c, "postgres")

	defer postgresSnap.End()

	dbResult := make([]social.Profile, 0)

	if len(request.Address) > model.DefaultLimit {
		request.Address = request.Address[:model.DefaultLimit]
	}

	for i, v := range request.Address {
		request.Address[i] = strings.ToLower(v)
	}

	sql := database.Global().Model(&social.Profile{}).Where("LOWER(address) IN ? ", request.Address)

	if len(request.Network) > 0 {
		for i, v := range request.Network {
			request.Network[i] = strings.ToLower(v)
		}
		sql = sql.Where("network IN ?", request.Network)
	}

	if len(request.Platform) > 0 {
		for i, v := range request.Platform {
			request.Platform[i] = strings.ToLower(v)
		}
		sql = sql.Where("platform IN ?", request.Platform)
	}

	if err := sql.Find(&dbResult).Error; err != nil {
		return nil, err
	}

	return dbResult, nil
}

func UpsertProfiles(c context.Context, profiles []social.Profile) {
	err := database.Global().Model(&social.Profile{}).
		Clauses(clause.OnConflict{
			UpdateAll: true,
			DoUpdates: clause.AssignmentColumns([]string{"updated_at"}),
		}).
		Create(profiles).Error
	if err != nil {
		logrus.Errorf("[profile] upsertProfiles error, %v", err)
	}
}

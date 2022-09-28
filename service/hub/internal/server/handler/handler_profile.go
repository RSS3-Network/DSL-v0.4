package handler

import (
	"context"
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
	"github.com/naturalselectionlabs/pregod/common/database"
	"github.com/naturalselectionlabs/pregod/common/database/model/social"
	"github.com/naturalselectionlabs/pregod/common/protocol"
	"github.com/naturalselectionlabs/pregod/common/worker/crossbell"
	"github.com/naturalselectionlabs/pregod/common/worker/ens"
	"github.com/naturalselectionlabs/pregod/common/worker/lens"
	lop "github.com/samber/lo/parallel"
	"github.com/sirupsen/logrus"
	"go.opentelemetry.io/otel"
	"gorm.io/gorm/clause"
)

var ProfilePlatformList = []string{
	protocol.PlatformEns,
	protocol.PlatformLens,
	protocol.PlatformCrossbell,
}

// GetProfilesFunc supported filter:
// - address
// - network
// - platform
func (h *Handler) GetProfilesFunc(c echo.Context) error {
	go h.apiReport(GetProfiles, c.Get("API-KEY"))
	tracer := otel.Tracer("GetProfilesFunc")
	ctx, httpSnap := tracer.Start(c.Request().Context(), "http")

	defer httpSnap.End()

	request := GetRequest{}

	if err := c.Bind(&request); err != nil {
		return BadRequest(c)
	}

	if err := c.Validate(&request); err != nil {
		return err
	}

	go h.filterReport(GetProfiles, request)

	profileList, err := h.getProfiles(ctx, request)
	if err != nil {
		return InternalError(c)
	}

	return c.JSON(http.StatusOK, &Response{
		Total:  int64(len(profileList)),
		Result: profileList,
	})
}

func (h *Handler) BatchGetProfilesFunc(c echo.Context) error {
	go h.apiReport(PostProfiles, c.Get("API-KEY"))
	tracer := otel.Tracer("BatchGetProfilesFunc")
	ctx, httpSnap := tracer.Start(c.Request().Context(), "http")

	defer httpSnap.End()

	request := BatchGetProfilesRequest{}

	if err := c.Bind(&request); err != nil {
		return BadRequest(c)
	}

	if err := c.Validate(&request); err != nil {
		return err
	}

	go h.filterReport(PostProfiles, request)

	if len(request.Address) > DefaultLimit {
		request.Address = request.Address[:DefaultLimit]
	}

	profileList, err := h.batchGetProfiles(ctx, request)
	if err != nil {
		return InternalError(c)
	}

	return c.JSON(http.StatusOK, &Response{
		Total:  int64(len(profileList)),
		Result: profileList,
	})
}

func (h *Handler) getProfiles(c context.Context, request GetRequest) ([]social.Profile, error) {
	m := make(map[string]bool)
	result := make([]social.Profile, 0)

	profiles, _ := h.getProfilesDatabase(c, request)

	for _, profile := range profiles {
		m[profile.Platform] = true
		result = append(result, profile)
	}

	lop.ForEach(ProfilePlatformList, func(platform string, i int) {
		if m[platform] {
			go h.getProfilesFromPlatform(c, platform, request.Address, true)

			return
		}

		list, err := h.getProfilesFromPlatform(c, platform, request.Address, false)
		if err == nil && len(list) > 0 {
			result = append(result, list...)
		}
	})

	err := database.Global().Model(&social.Profile{}).
		Clauses(clause.OnConflict{
			UpdateAll: true,
		}).
		Create(result).Error
	if err != nil {
		logrus.Errorf("[profile] getProfiles error, %v", err)

		return nil, err
	}

	return result, nil
}

func (h *Handler) batchGetProfiles(c context.Context, request BatchGetProfilesRequest) ([]social.Profile, error) {
	return nil, nil
}

func (h *Handler) getProfilesFromPlatform(c context.Context, platform, address string, update bool) ([]social.Profile, error) {
	var profile *social.Profile
	var profiles []social.Profile
	var err error

	switch platform {
	case protocol.PlatformEns:
		ensClient := ens.New()
		profile, err = ensClient.GetProfile(address)
	case protocol.PlatformLens:
		lensClient := lens.New()
		profile, err = lensClient.GetProfile(address)
	case protocol.PlatformCrossbell:
		csbClient := crossbell.New()
		profiles, err = csbClient.GetProfile(address)
	}

	if profile != nil {
		profiles = append(profiles, *profile)
	}

	if update && err == nil && len(profiles) > 0 {
		_ = database.Global().Model(&social.Profile{}).
			Clauses(clause.OnConflict{
				UpdateAll: true,
			}).
			Create(profiles).Error
	}

	return profiles, err
}

// getProfiles get profile data from database
func (h *Handler) getProfilesDatabase(c context.Context, request GetRequest) ([]social.Profile, error) {
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

	sql = sql.Limit(DefaultLimit)

	if err := sql.Find(&dbResult).Error; err != nil {
		logrus.Errorf("[profile] getProfilesDatabase error, %v", err)

		return nil, err
	}

	return dbResult, nil
}

// batchGetProfiles get profile data from database
func (h *Handler) batchGetProfilesDatabase(c context.Context, request BatchGetProfilesRequest) ([]social.Profile, error) {
	tracer := otel.Tracer("batchGetProfiles")
	_, postgresSnap := tracer.Start(c, "postgres")

	defer postgresSnap.End()

	dbResult := make([]social.Profile, 0)

	if len(request.Address) > DefaultLimit {
		request.Address = request.Address[:DefaultLimit]
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

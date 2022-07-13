package handler

import (
	"context"
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

	"github.com/naturalselectionlabs/pregod/common/protocol"

	"github.com/labstack/echo/v4"
	dbModel "github.com/naturalselectionlabs/pregod/common/database/model"
	"github.com/naturalselectionlabs/pregod/common/protocol/filter"
	"go.opentelemetry.io/otel"
)

type ProfileResult struct {
	Hash       string          `json:"hash"`
	Network    string          `json:"network"`
	Platform   string          `json:"platform"`
	ProfileURL string          `json:"profile_url"`
	Metadata   json.RawMessage `json:"metadata"`
}

// GetProfileListFunc supported filter:
// - address
// - network
// - platform
func (h *Handler) GetProfileListFunc(c echo.Context) error {
	tracer := otel.Tracer("GetProfileListFunc")
	ctx, httpSnap := tracer.Start(c.Request().Context(), "http")

	defer httpSnap.End()

	request := GetRequest{}

	if err := c.Bind(&request); err != nil {
		return err
	}

	profileList, total, err := h.getProfileListDatabase(ctx, request)
	if err != nil {
		return err
	}

	if len(profileList) == 0 {
		return c.JSON(http.StatusOK, &Response{
			Result: make([]dbModel.Transfer, 0),
		})
	}

	var cursor string
	if total > int64(DefaultLimit) {
		cursor = profileList[len(profileList)-1].TransactionHash
	}

	result, err := formatProfileResult(profileList)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, &Response{
		Total:  total,
		Cursor: cursor,
		Result: result,
	})
}

// getTransferListDatabase get transfer data from database
func (h *Handler) getProfileListDatabase(c context.Context, request GetRequest) ([]dbModel.Transfer, int64, error) {
	tracer := otel.Tracer("getProfileListDatabase")
	_, postgresSnap := tracer.Start(c, "postgres")

	defer postgresSnap.End()

	dbResult := make([]dbModel.Transfer, 0)
	total := int64(0)
	sql := h.DatabaseClient.Model(&dbModel.Transfer{}).Where("LOWER(address_from) = ? OR LOWER(address_to) = ?",
		request.Address,
		request.Address,
	)

	if len(request.Cursor) > 0 {
		cursorInt, err := strconv.Atoi(request.Cursor)
		if err != nil {
			return nil, total, err
		}

		if cursorInt > 0 {
			sql = sql.Offset(cursorInt * DefaultLimit)
		}
	}

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

	sql = sql.Where("tag = ?", filter.TagSocial)
	sql = sql.Where("\"type\" IN ? ", []string{filter.SocialProfile})
	sql = sql.Limit(DefaultLimit)

	if err := sql.Count(&total).Find(&dbResult).Error; err != nil {
		return nil, total, err
	}

	return dbResult, total, nil
}

type CrossbellProfilestruct struct {
	Crossbell struct {
		Event     string `json:"event"`
		Character struct {
			Id       int    `json:"id"`
			Uri      string `json:"uri"`
			Metadata struct {
				Bio              string   `json:"bio"`
				Name             string   `json:"name"`
				Type             string   `json:"type"`
				Avatars          []string `json:"avatars"`
				ConnectedAvatars []string `json:"connected_avatars"`
			} `json:"metadata"`
		} `json:"character"`
	} `json:"crossbell"`
}

// formatProfileResult format profile result by extracting ProfileURL from metadata
func formatProfileResult(profileList []dbModel.Transfer) ([]ProfileResult, error) {
	result := make([]ProfileResult, 0)
	for _, profile := range profileList {

		temp := ProfileResult{
			Hash:     profile.TransactionHash,
			Network:  profile.Network,
			Platform: profile.Platform,
			Metadata: profile.Metadata,
		}

		switch profile.Network {
		case protocol.NetworkCrossbell:
			tempStructure := &CrossbellProfilestruct{}
			if err := json.Unmarshal(profile.Metadata, &tempStructure); err != nil {
				return nil, err
			}
			temp.ProfileURL = tempStructure.Crossbell.Character.Metadata.Avatars[0]
		default:
		}

		result = append(result, temp)
	}
	return result, nil
}

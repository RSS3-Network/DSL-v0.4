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

	profileList, err := h.getProfileListDatabase(ctx, request)
	if err != nil {
		return err
	}

	if len(profileList) == 0 {
		return c.JSON(http.StatusOK, &Response{
			Result: make([]dbModel.Transfer, 0),
		})
	}

	result, err := formatProfileResult(profileList)
	if err != nil {
		return BadRequest(c)
	}

	return c.JSON(http.StatusOK, &Response{
		Total:  int64(len(result)),
		Result: result,
	})
}

// getTransferListDatabase get transfer data from database
func (h *Handler) getProfileListDatabase(c context.Context, request GetRequest) ([]dbModel.Transfer, error) {
	tracer := otel.Tracer("getProfileListDatabase")
	_, postgresSnap := tracer.Start(c, "postgres")

	defer postgresSnap.End()

	dbResult := make([]dbModel.Transfer, 0)

	sql := h.DatabaseClient.Model(&dbModel.Transfer{}).Where("LOWER(address_from) = ? OR LOWER(address_to) = ?",
		request.Address,
		request.Address,
	)

	if len(request.Cursor) > 0 {
		cursorInt, err := strconv.Atoi(request.Cursor)
		if err != nil {
			return nil, err
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

	if err := sql.Find(&dbResult).Error; err != nil {
		return nil, err
	}

	return dbResult, nil
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
func formatProfileResult(profileList []dbModel.Transfer) ([]dbModel.Profile, error) {
	result := make([]dbModel.Profile, 0)
	for _, profile := range profileList {
		temp := dbModel.Profile{
			Network:    profile.Network,
			Platform:   profile.Platform,
			SourceData: profile.Metadata,
		}

		switch profile.Network {
		case protocol.NetworkCrossbell:
			tempStructure := &CrossbellProfilestruct{}
			if err := json.Unmarshal(profile.Metadata, &tempStructure); err != nil {
				return nil, err
			}

			switch tempStructure.Crossbell.Event {
			case "CharacterCreated":
				temp.Address = tempStructure.Crossbell.Character.Uri

				temp.Name = tempStructure.Crossbell.Character.Metadata.Name
				temp.Handle = tempStructure.Crossbell.Character.Metadata.Name
				temp.Bio = tempStructure.Crossbell.Character.Metadata.Bio

				if len(tempStructure.Crossbell.Character.Metadata.Avatars) > 0 {
					for _, avatar := range tempStructure.Crossbell.Character.Metadata.Avatars {
						temp.ProfileUris = append(temp.ProfileUris, avatar)
					}
				}

				if len(tempStructure.Crossbell.Character.Metadata.ConnectedAvatars) > 0 {
					for _, avatar := range tempStructure.Crossbell.Character.Metadata.ConnectedAvatars {
						temp.SocialUris = append(temp.SocialUris, avatar)
					}
				}
			case "SetCharacterUri":
				// TODO: update URI
				continue
			case "SetHandle":
				// TODO: update handle
				continue
			}

		default:
		}
		result = append(result, temp)
	}
	return result, nil
}

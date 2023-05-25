package handler

import (
	"context"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/mattn/go-mastodon"
	"github.com/naturalselectionlabs/pregod/service/hub/internal/config"
	"github.com/naturalselectionlabs/pregod/service/hub/internal/server/model"

	"go.opentelemetry.io/otel"
)

func (h *Handler) GetMastodonFunc(c echo.Context) error {
	go h.apiReport(model.GetMastodon, c)
	tracer := otel.Tracer("GetMastodonFunc")
	ctx, httpSnap := tracer.Start(c.Request().Context(), "http")

	defer httpSnap.End()

	request := model.GetRequest{}

	if err := c.Bind(&request); err != nil {
		return BadRequest(c)
	}

	if err := c.Validate(&request); err != nil {
		return ValidateFailed(c)
	}

	go h.filterReport(model.GetMastodon, request, c)

	client, err := dial(config.ConfigHub.Mastodon.Server, config.ConfigHub.Mastodon.Username, config.ConfigHub.Mastodon.Password)
	if err != nil {
		return ErrorResp(c, err, http.StatusInternalServerError, ErrorCodeBadRequest)
	}

	contentList, err := h.service.GetMastodonContent(ctx, request, client)
	if err != nil {
		return ErrorResp(c, err, http.StatusInternalServerError, ErrorCodeInternalError)
	}

	total := int64(len(contentList))
	return c.JSON(http.StatusOK, &model.Response{
		Total:  &total,
		Result: contentList,
	})
}

func dial(server, username, password string) (*mastodon.Client, error) {
	clientName := time.Now().String()
	app, err := mastodon.RegisterApp(context.Background(), &mastodon.AppConfig{
		Server:     server,
		ClientName: clientName,
		Scopes:     "read write follow",
	})
	if err != nil {
		return nil, err
	}

	c := mastodon.NewClient(&mastodon.Config{
		Server:       server,
		ClientID:     app.ClientID,
		ClientSecret: app.ClientSecret,
	})

	err = c.Authenticate(context.Background(), username, password)
	if err != nil {
		return nil, err
	}

	return c, nil
}

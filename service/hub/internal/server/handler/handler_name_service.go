package handler

import (
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
	"github.com/naturalselectionlabs/pregod/common/worker/ens"
	"github.com/naturalselectionlabs/pregod/service/hub/internal/config"
	"go.opentelemetry.io/otel"
)

func (h *Handler) GetENSResolve(c echo.Context) error {
	tracer := otel.Tracer("GetENSResolve")
	_, httpSnap := tracer.Start(c.Request().Context(), "http")

	defer httpSnap.End()

	request := GetRequest{}

	if err := c.Bind(&request); err != nil {
		return BadRequest(c)
	}

	if len(request.Address) == 0 {
		return c.JSON(http.StatusBadRequest, &Response{})
	}
	request.Address = strings.ToLower(request.Address)

	result, err := ens.Resolve(config.ConfigHub.RPC.General.Ethereum.HTTP, request.Address)
	if err != nil {
		return BadRequest(c)
	}

	return c.JSON(http.StatusOK, result)
}

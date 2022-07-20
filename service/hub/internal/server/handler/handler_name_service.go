package handler

import (
	"net/http"

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
		return AddressIsEmpty(c)
	}

	// error here means the address doesn't have a primary ENS, and can be ignored
	result, _ := ens.Resolve(config.ConfigHub.RPC.General.Ethereum.HTTP, request.Address)

	return c.JSON(http.StatusOK, result)
}

package handler

import (
	"fmt"
	"net/http"

	"github.com/ethereum/go-ethereum/crypto"
	"github.com/google/uuid"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
	"github.com/naturalselectionlabs/pregod/common/database/model"
	"github.com/sirupsen/logrus"
	"go.opentelemetry.io/otel"
)

func (h *Handler) PostAPIKeyFunc(c echo.Context) error {
	// tracer
	tracer := otel.Tracer("PostAPIKeyFunc")
	_, httpSnap := tracer.Start(c.Request().Context(), "http")

	defer httpSnap.End()

	request := APIKeyRequest{}

	if err := c.Bind(&request); err != nil {
		return BadRequest(c)
	}

	if err := c.Validate(&request); err != nil {
		return err
	}

	// check if key already exists
	item := model.APIKey{}
	if err := h.DatabaseClient.
		Where("address = ?").
		First(&item).Error; err == nil && len(item.UUID) > 0 {
		return fmt.Errorf("This address has already applied for API-KEY.")
	}

	// generate api key
	item = model.APIKey{
		Address: request.Address,
		UUID:    uuid.New().String(),
	}
	if err := h.DatabaseClient.Create(&item).Error; err != nil {
		logrus.Errorf("PostAPIKeyFunc: create sql error, %v", err)
		return InternalError(c)
	}

	return c.JSON(http.StatusOK, &item)
}

func (h *Handler) GetAPIKeyFunc(c echo.Context) error {
	// tracer
	tracer := otel.Tracer("GetAPIKeyFunc")
	_, httpSnap := tracer.Start(c.Request().Context(), "http")

	defer httpSnap.End()

	request := APIKeyRequest{}

	if err := c.Bind(&request); err != nil {
		return BadRequest(c)
	}

	if err := c.Validate(&request); err != nil {
		return err
	}

	wallet, err := session.Get("user_wallet", c)
	if err != nil {
		return BadRequest(c)
	}

	hash := wallet.Values["hash"].(string)
	signature := wallet.Values["signature"].(string)

	sigPublicKeyECDSA, err := crypto.SigToPub([]byte(hash), []byte(signature))
	if err != nil {
		logrus.Errorf("GetAPIKeyFunc: crypto.SigToPub error, %v", err)
		return InternalError(c)
	}

	address := crypto.PubkeyToAddress(*sigPublicKeyECDSA)

	if address.String() != request.Address {
		return InternalError(c)
	}

	item := model.APIKey{}
	h.DatabaseClient.Where("address = ?").First(&item)

	return c.JSON(http.StatusOK, &item)
}

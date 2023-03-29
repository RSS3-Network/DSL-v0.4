package handler

import (
	"errors"
	"net/http"

	"github.com/ethereum/go-ethereum/crypto"
	"github.com/google/uuid"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
	"github.com/naturalselectionlabs/pregod/common/database"
	dbModel "github.com/naturalselectionlabs/pregod/common/database/model"
	"github.com/naturalselectionlabs/pregod/service/hub/internal/server/model"
	"github.com/sirupsen/logrus"
	"go.opentelemetry.io/otel"
)

func (h *Handler) PostAPIKeyFunc(c echo.Context) error {
	// tracer
	tracer := otel.Tracer("PostAPIKeyFunc")
	_, httpSnap := tracer.Start(c.Request().Context(), "http")

	defer httpSnap.End()

	request := model.APIKeyRequest{}

	if err := c.Bind(&request); err != nil {
		return BadRequest(c)
	}

	if err := c.Validate(&request); err != nil {
		return ValidateFailed(c)
	}

	// check if key already exists
	item := dbModel.APIKey{}
	if err := database.Global().
		Where("address = ?").
		First(&item).Error; err == nil && len(item.UUID) > 0 {
		return ErrorResp(c, errors.New("this address has already applied for API-KEY"), http.StatusBadRequest, ErrorCodeKeyAlreadyExists)
	}

	// generate api key
	item = dbModel.APIKey{
		Address: request.Address,
		UUID:    uuid.New().String(),
	}
	if err := database.Global().Create(&item).Error; err != nil {
		logrus.Errorf("PostAPIKeyFunc: create sql error, %v", err)
		return ErrorResp(c, errors.New("PostAPIKeyFunc: create sql error"), http.StatusInternalServerError, ErrorCodeInternalError)
	}

	return c.JSON(http.StatusOK, &item)
}

func (h *Handler) GetAPIKeyFunc(c echo.Context) error {
	// tracer
	tracer := otel.Tracer("GetAPIKeyFunc")
	_, httpSnap := tracer.Start(c.Request().Context(), "http")

	defer httpSnap.End()

	request := model.APIKeyRequest{}

	if err := c.Bind(&request); err != nil {
		return BadRequest(c)
	}

	if err := c.Validate(&request); err != nil {
		return ValidateFailed(c)
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
		return ErrorResp(c, errors.New("GetAPIKeyFunc: crypto.SigToPub error"), http.StatusBadRequest, ErrorCodeSigToPubError)
	}

	address := crypto.PubkeyToAddress(*sigPublicKeyECDSA)

	if address.String() != request.Address {
		return ErrorResp(c, errors.New("GetAPIKeyFunc: address is not match"), http.StatusBadRequest, ErrorCodeAddressIsNotMatch)
	}

	item := dbModel.APIKey{}
	database.Global().Where("address = ?").First(&item)

	return c.JSON(http.StatusOK, &item)
}

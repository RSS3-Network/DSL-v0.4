package handler

import (
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
)

type ErrorResponse struct {
	Error string `json:"error"`
}

func BadRequest(c echo.Context) error {
	return c.JSON(http.StatusBadRequest, &ErrorResponse{
		Error: "Please check your request and try again.",
	})
}

func AddressIsEmpty(c echo.Context) error {
	return c.JSON(http.StatusBadRequest, &ErrorResponse{
		Error: "At least one address is required",
	})
}

func AddressIsInvalid(c echo.Context) error {
	return c.JSON(http.StatusBadRequest, &ErrorResponse{
		Error: "The address provided is invalid. You can use 0x, ENS, Crossbell, or Lens addresses.",
	})
}

func InternalError(c echo.Context) error {
	return c.JSON(http.StatusInternalServerError, &ErrorResponse{
		Error: "An internal error has occurred, please try again later.",
	})
}

func (h *Handler) ErrorFunc(err error, c echo.Context) {
	var (
		httpCode     = http.StatusInternalServerError
		errorMessage = err.Error()
	)

	if httpError, ok := err.(*echo.HTTPError); ok {
		httpCode = httpError.Code
		errorMessage = httpError.Message.(string)
	}

	_ = c.JSON(httpCode, &ErrorResponse{
		Error: strings.ToLower(errorMessage),
	})
}

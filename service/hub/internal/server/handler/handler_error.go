package handler

import (
	"errors"
	"fmt"
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
)

var (
	ErrorInvalidParameter = errors.New("invalid parameter")
	ErrorInternalDatabase = errors.New("internal database error")
)

type ErrorResponse struct {
	Error string `json:"error"`
}

func (h *Handler) ErrorFunc(err error, c echo.Context) {
	var (
		httpCode     = http.StatusInternalServerError
		errorMessage = err.Error()
	)

	if httpError, ok := err.(*echo.HTTPError); ok {
		httpCode = httpError.Code
		errorMessage = fmt.Sprintf("%s", httpError.Message)
	}

	_ = c.JSON(httpCode, &ErrorResponse{
		Error: strings.ToLower(errorMessage),
	})
}

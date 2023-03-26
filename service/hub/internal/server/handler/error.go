package handler

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
)

type ErrorResponse struct {
	Error     string `json:"error"`
	ErrorCode int    `json:"error_code"`
}

const (
	ErrorCodeBadRequest                = 1001
	ErrorCodeBadParams                 = 1002
	ErrorCodeAddressIsEmpty            = 1003
	ErrorCodeAddressIsInvalid          = 1004
	ErrorCodeInternalError             = 1005
	ErrorCodeNotSupportContract        = 1006
	ErrorCodeGetTransactionByHashWrong = 1007
)

func ErrorResp(c echo.Context, err error, errorCode int) error {
	return c.JSON(http.StatusBadRequest, &ErrorResponse{
		Error:     err.Error(),
		ErrorCode: errorCode,
	})
}

func BadRequest(c echo.Context) error {
	return c.JSON(http.StatusBadRequest, &ErrorResponse{
		Error:     "Please check your request and try again.",
		ErrorCode: ErrorCodeBadRequest,
	})
}

func BadParams(c echo.Context, tag []string, typeX []string) error {
	return c.JSON(http.StatusBadRequest, &ErrorResponse{
		Error:     fmt.Sprintf("Please check your param combination and try again. Tag: %s. Type: %s", tag, typeX),
		ErrorCode: ErrorCodeBadParams,
	})
}

func AddressIsEmpty(c echo.Context) error {
	return c.JSON(http.StatusBadRequest, &ErrorResponse{
		Error:     "At least one address is required",
		ErrorCode: ErrorCodeAddressIsEmpty,
	})
}

func AddressIsInvalid(c echo.Context) error {
	return c.JSON(http.StatusOK, &ErrorResponse{
		Error:     "The address provided is invalid. You can use a 0x, ENS, Crossbell, or Lens address.",
		ErrorCode: ErrorCodeAddressIsInvalid,
	})
}

func InternalError(c echo.Context) error {
	return c.JSON(http.StatusInternalServerError, &ErrorResponse{
		Error:     "An internal error has occurred, please try again later.",
		ErrorCode: ErrorCodeInternalError,
	})
}

func ErrorFunc(err error, c echo.Context) {
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

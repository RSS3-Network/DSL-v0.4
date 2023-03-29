package handler

import (
	"context"
	"net/http"
	"strings"

	"github.com/google/uuid"
	"github.com/gorilla/websocket"
	"github.com/labstack/echo/v4"
	"github.com/naturalselectionlabs/pregod/common/constant"
	dbModel "github.com/naturalselectionlabs/pregod/common/database/model"
	"github.com/naturalselectionlabs/pregod/service/hub/internal/server/dao"
	"github.com/naturalselectionlabs/pregod/service/hub/internal/server/middlewarex"
	"github.com/naturalselectionlabs/pregod/service/hub/internal/server/model"
	ws "github.com/naturalselectionlabs/pregod/service/hub/internal/server/websocket"

	"go.opentelemetry.io/otel"
)

// GetNotesFunc HTTP handler for action API
// parse query parameters, query and assemble data
func (h *Handler) GetNotesFunc(c echo.Context) error {
	go h.apiReport(model.GetNotes, c)
	tracer := otel.Tracer("GetNotesFunc")
	ctx, httpSnap := tracer.Start(c.Request().Context(), "http")

	defer httpSnap.End()

	request := model.GetRequest{}

	if err := c.Bind(&request); err != nil {
		return BadRequest(c)
	}

	if err := c.Validate(&request); err != nil {
		return ValidateFailed(c)
	}

	// api report
	if len(request.Cursor) == 0 {
		go h.filterReport(model.GetNotes, request, c)
	}

	if request.Limit <= 0 || request.Limit > model.DefaultLimit {
		request.Limit = model.DefaultLimit
	}

	if request.ActionLimit <= 0 {
		request.ActionLimit = model.DefaultActionLimit
	}

	// header into ctx
	ctx = context.WithValue(ctx, constant.HEADER_CTX_KEY, c.Request().Header)

	var (
		transactions []dbModel.Transaction
		total        int64
		err          error
	)
	response := &model.Response{}

	// nft feed for rara
	if strings.HasPrefix(request.Address, "nft:") {
		request.Address = strings.Split(request.Address, "nft:")[1]
		transactions, total, err = h.service.GetNftFeeds(ctx, request)
	} else {
		transactions, total, err = h.service.GetNotes(ctx, request)

		var addressStatus []dbModel.Address
		if request.QueryStatus {
			addressStatus, _ = dao.GetAddress(ctx, []string{request.Address})
		}

		response.AddressStatus = addressStatus
	}

	if err != nil {
		return ErrorResp(c, err, http.StatusInternalServerError, ErrorCodeInternalError)
	}

	if request.CountOnly {
		return c.JSON(http.StatusOK, &model.Response{
			Total: &total,
		})
	}

	var cursor string
	if total > int64(request.Limit) {
		cursor = transactions[len(transactions)-1].Hash
	}

	response.Total = &total
	response.Cursor = cursor
	response.Result = transactions

	return c.JSON(http.StatusOK, response)
}

// BatchGetNotesFunc query multiple addresses and filters
func (h *Handler) BatchGetNotesFunc(c echo.Context) error {
	go h.apiReport(model.PostNotes, c)
	tracer := otel.Tracer("BatchGetNotesFunc")
	ctx, httpSnap := tracer.Start(c.Request().Context(), "BatchGetNotesFunc:http")

	defer httpSnap.End()

	request := model.BatchGetNotesRequest{}

	if err := c.Bind(&request); err != nil {
		return BadRequest(c)
	}

	if len(request.Cursor) == 0 {
		go h.filterReport(model.PostNotes, request, c)
	}

	if len(request.Address) == 0 {
		return AddressIsEmpty(c)
	}

	// check address
	if len(request.Address) > model.DefaultLimit {
		request.Address = request.Address[:model.DefaultLimit]
	}

	if request.Limit <= 0 || request.Limit > model.DefaultLimit {
		request.Limit = model.DefaultLimit
	}

	if request.ActionLimit <= 0 {
		request.ActionLimit = model.DefaultActionLimit
	}

	for i, v := range request.Address {
		address, err := middlewarex.ResolveAddress(c, v, request.IgnoreContract)
		if err != nil {
			return ErrorResp(c, err, http.StatusBadRequest, ErrorCodeNotSupportContract)
		}
		request.Address[i] = address
	}

	// header into ctx
	ctx = context.WithValue(ctx, constant.HEADER_CTX_KEY, c.Request().Header)

	transactions, total, err := h.service.BatchGetNotes(ctx, request)
	if err != nil {
		return ErrorResp(c, err, http.StatusInternalServerError, ErrorCodeInternalError)
	}

	var addressStatus []dbModel.Address
	if request.QueryStatus {
		addressStatus, _ = dao.GetAddress(ctx, request.Address)
	}

	if request.CountOnly {
		return c.JSON(http.StatusOK, &model.Response{
			Total: &total,
		})
	}

	var cursor string
	if total > int64(request.Limit) {
		cursor = transactions[len(transactions)-1].Hash
	}

	if total == 0 {
		return c.JSON(http.StatusOK, &model.Response{
			Result:        make([]dbModel.Transaction, 0),
			AddressStatus: addressStatus,
		})
	}

	return c.JSON(http.StatusOK, &model.Response{
		Total:         &total,
		Cursor:        cursor,
		Result:        transactions,
		AddressStatus: addressStatus,
	})
}

// BatchGetSocialNotesFunc query multiple addresses
func (h *Handler) BatchGetSocialNotesFunc(c echo.Context) error {
	go h.apiReport(model.PostSocialNotes, c)
	tracer := otel.Tracer("BatchGetSocialNotesFunc")
	ctx, httpSnap := tracer.Start(c.Request().Context(), "BatchGetSocialNotesFunc:http")

	defer httpSnap.End()

	request := model.BatchGetSocialNotesRequest{}

	if err := c.Bind(&request); err != nil {
		return BadRequest(c)
	}

	if len(request.Cursor) == 0 {
		go h.filterReport(model.PostSocialNotes, request, c)
	}

	if len(request.Address) == 0 {
		return AddressIsEmpty(c)
	}

	// check address
	if len(request.Address) > model.DefaultLimit {
		request.Address = request.Address[:model.DefaultLimit]
	}
	for i, v := range request.Address {
		address, err := middlewarex.ResolveAddress(c, v, true)
		if err != nil {
			return ErrorResp(c, err, http.StatusBadRequest, ErrorCodeNotSupportContract)
		}
		request.Address[i] = address
	}

	transactions, total, err := h.service.BatchGetSocialNotes(ctx, request)
	if err != nil {
		return ErrorResp(c, err, http.StatusInternalServerError, ErrorCodeInternalError)
	}

	if request.CountOnly {
		return c.JSON(http.StatusOK, &model.Response{
			Total: &total,
		})
	}

	var cursor string
	if total > int64(request.Limit) {
		cursor = transactions[len(transactions)-1].Hash
	}

	if total == 0 {
		return c.JSON(http.StatusOK, &model.Response{
			Result: make([]dbModel.Transaction, 0),
		})
	}

	return c.JSON(http.StatusOK, &model.Response{
		Total:  &total,
		Cursor: cursor,
		Result: transactions,
	})
}

func (h *Handler) GetNotesWsFunc(c echo.Context) error {
	ws.GetClientMaps()

	clientId := uuid.New().String()

	conn, err := (&websocket.Upgrader{ReadBufferSize: 1024, WriteBufferSize: 1024, CheckOrigin: func(r *http.Request) bool { return true }}).Upgrade(c.Response(), c.Request(), nil)
	if err != nil {
		return err
	}

	client := &ws.WSClient{Hub: h.service.WsHub, Conn: conn, Send: make(chan []byte, 256), ClientId: []byte(clientId), AddressMap: make(map[string]struct{})}

	client.Hub.Register <- client

	ws.ReplaceGlobal(clientId, client)

	go h.service.SubscribeIndexerRefreshMessage(client)
	go client.WriteMsg()
	go client.ReadMsg()

	return nil
}

func (h *Handler) GetTransactionByHashFunc(c echo.Context) error {
	go h.apiReport(model.GetTransactionByHash, c)
	tracer := otel.Tracer("GetTransactionByHashFunc")
	ctx, httpSnap := tracer.Start(c.Request().Context(), "GetTransactionByHashFunc:http")

	defer httpSnap.End()

	request := model.GetTransactionRequest{}

	if err := c.Bind(&request); err != nil {
		return BadRequest(c)
	}

	if err := c.Validate(&request); err != nil {
		return ValidateFailed(c)
	}

	transaction, err := h.service.GetTransactionByHash(ctx, request)
	if err != nil {
		return ErrorResp(c, err, http.StatusBadRequest, ErrorCodeGetTransactionByHashError)
	}

	return c.JSON(http.StatusOK, &model.Response{
		Result: transaction,
	})
}

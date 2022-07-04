package handler

import (
	"context"
	"encoding/json"
	"net/http"
	"strings"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/naturalselectionlabs/pregod/common/database/model"
	"github.com/naturalselectionlabs/pregod/common/protocol"
	"github.com/naturalselectionlabs/pregod/common/protocol/filter"
	rabbitmq "github.com/rabbitmq/amqp091-go"
	"go.opentelemetry.io/otel"
)

type GetActionListRequest struct {
	Address   string    `param:"address"`
	Limit     int       `query:"limit"`
	Cursor    string    `query:"cursor"`
	Type      []string  `query:"type"`
	Tag       string    `query:"tag"`
	Network   []string  `query:"network"`
	Platform  []string  `query:"platform"`
	Timestamp time.Time `query:"timestamp"`
	Hash      string    `query:"hash"`
}

// GetActionListFunc HTTP handler for action API
// parse query parameters, query and assemble data
func (h *Handler) GetActionListFunc(c echo.Context) error {
	tracer := otel.Tracer("GetActionListFunc")
	ctx, httpSnap := tracer.Start(c.Request().Context(), "http")

	defer httpSnap.End()

	request := GetActionListRequest{}
	if err := c.Bind(&request); err != nil {
		return err
	}

	if request.Limit <= 0 || request.Limit > DefaultLimit {
		request.Limit = DefaultLimit
	}

	var types []string
	for _, t := range request.Type {
		if filter.CheckTypeValid(request.Tag, t) {
			types = append(types, t)
		}
	}

	if len(request.Type) > 0 && len(types) == 0 {
		return c.JSON(http.StatusOK, &Response{
			Result: make([]model.Transaction, 0),
		})
	}

	request.Type = types

	go func() {
		// create a rabbitmq job to index the latest user data
		_, rabbitmqSnap := tracer.Start(ctx, "rabbitmq")

		defer rabbitmqSnap.End()

		networks := []string{
			protocol.NetworkEthereum, protocol.NetworkPolygon, protocol.NetworkBinanceSmartChain,
			protocol.NetworkArweave, protocol.NetworkXDAI, protocol.NetworkZkSync, protocol.NetworkCrossbell,
		}

		for _, network := range networks {
			message := protocol.Message{
				Address: request.Address,
				Network: network,
			}

			messageData, err := json.Marshal(&message)
			if err != nil {
				return
			}

			if err := h.RabbitmqChannel.Publish(protocol.ExchangeJob, protocol.IndexerWorkRoutingKey, false, false, rabbitmq.Publishing{
				ContentType: protocol.ContentTypeJSON,
				Body:        messageData,
			}); err != nil {
				return
			}
		}
	}()

	transactionList, total, err := h.getTransactionListDatabase(ctx, request)
	if err != nil {
		return err
	}

	if len(transactionList) == 0 {
		return c.JSON(http.StatusOK, &Response{
			Result: make([]model.Transaction, 0),
		})
	}

	transactionHashList := make([]string, 0)
	for _, transactionHash := range transactionList {
		transactionHashList = append(transactionHashList, transactionHash.Hash)
	}

	transferList, err := h.getActionListDatabase(ctx, transactionHashList, request)
	if err != nil {
		return err
	}

	transferMap := make(map[string][]model.Transfer)
	for _, transfer := range transferList {
		transferMap[transfer.TransactionHash] = append(transferMap[transfer.TransactionHash], transfer)
	}

	for index := range transactionList {
		transactionList[index].Transfers = transferMap[transactionList[index].Hash]
	}

	var cursor string
	if total > int64(request.Limit) {
		cursor = transactionList[len(transactionList)-1].Hash
	}

	return c.JSON(http.StatusOK, &Response{
		Total:  total,
		Cursor: cursor,
		Result: transactionList,
	})
}

// getTransactionListDatabase get transaction data from database
func (h *Handler) getTransactionListDatabase(c context.Context, request GetActionListRequest) ([]model.Transaction, int64, error) {
	tracer := otel.Tracer("getActionListDatabase")
	_, postgresSnap := tracer.Start(c, "postgres")

	defer postgresSnap.End()

	transactionList := make([]model.Transaction, 0)
	total := int64(0)
	sql := h.DatabaseClient.
		Model(&model.Transaction{}).
		Where("LOWER(address_from) = ? OR LOWER(address_to) = ?",
			strings.ToLower(request.Address),
			strings.ToLower(request.Address),
		)

	if len(request.Cursor) > 0 {
		var lastItem model.Transaction

		if err := h.DatabaseClient.Where("hash = ?", strings.ToLower(request.Cursor)).First(&lastItem).Error; err != nil {
			return nil, 0, err
		}

		sql = sql.Where("timestamp < ? OR (timestamp = ? AND index < ?)", lastItem.Timestamp, lastItem.Timestamp, lastItem.Index)
	}

	if len(request.Hash) > 0 {
		sql = sql.Where("hash = ?", request.Hash)
	}

	if len(request.Tag) > 0 {
		sql = sql.Where("tag = ?", request.Tag)
	}

	if len(request.Network) > 0 {
		sql = sql.Where("network IN ?", request.Network)
	}

	if len(request.Platform) > 0 {
		sql = sql.Where("platform IN ?", request.Platform)
	}

	if len(request.Timestamp.String()) > 0 {
		sql = sql.Where("timestamp > ?", request.Timestamp)
	}

	if err := sql.Count(&total).Limit(request.Limit).Order("timestamp DESC, index DESC").Find(&transactionList).Error; err != nil {
		return nil, 0, err
	}

	return transactionList, total, nil
}

// getActionListDatabase get transfer data from database
func (h *Handler) getActionListDatabase(c context.Context, transactionHashList []string, request GetActionListRequest) ([]model.Transfer, error) {
	tracer := otel.Tracer("getActionListDatabase")
	_, postgresSnap := tracer.Start(c, "postgres")

	defer postgresSnap.End()

	transfers := make([]model.Transfer, 0)

	sql := h.DatabaseClient.Model(&model.Transfer{}).Where("LOWER(address_from) = ? OR LOWER(address_to) = ?",
		strings.ToLower(request.Address),
		strings.ToLower(request.Address),
	)

	if len(request.Tag) > 0 && len(request.Type) > 0 {
		sql = sql.Where("\"type\" IN ? ", request.Type)
	}

	if err := sql.
		Where("transaction_hash IN ?", transactionHashList).
		Find(&transfers).Error; err != nil {
		return nil, err
	}

	return transfers, nil
}

package handler

import (
	"context"
	"encoding/json"
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
	"github.com/naturalselectionlabs/pregod/common/database/model"
	"github.com/naturalselectionlabs/pregod/common/protocol"
	rabbitmq "github.com/rabbitmq/amqp091-go"
	"go.opentelemetry.io/otel"
)

type GetActionListRequest struct {
	Address     string   `param:"address"`
	Limit       int      `query:"limit"`
	Cursor      string   `query:"cursor"`
	Types       []string `query:"types"`
	Tags        []string `query:"tags"`
	ItemSources []string `query:"item_sources"`
	Networks    []string `query:"networks"`
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

			if err := h.RabbitmqChannel.Publish(protocol.ExchangeJob, "", false, false, rabbitmq.Publishing{
				ContentType: protocol.ContentTypeJSON,
				Body:        messageData,
			}); err != nil {
				return
			}
		}
	}()

	transactionList, total, err := h.getTransactionListDatabse(ctx, request)
	if err != nil {
		return err
	}

	if len(transactionList) == 0 {
		return c.JSON(http.StatusOK, &Response{})
	}

	transactionHashList := []string{}
	for _, transactionHash := range transactionList {
		transactionHashList = append(transactionHashList, transactionHash.Hash)
	}

	transferList, err := h.getActionListDatabase(ctx, transactionHashList, request)
	if err != nil {
		return err
	}

	transferMap := make(map[string][]*model.Transfer)
	for _, transfer := range transferList {
		transferMap[transfer.TransactionHash] = append(transferMap[transfer.TransactionHash], transfer)
	}

	for _, transaction := range transactionList {
		transaction.Transfers = transferMap[transaction.Hash]
	}

	return c.JSON(http.StatusOK, &Response{
		Total:  total,
		Cursor: transactionList[len(transactionList)-1].Hash,
		Result: transactionList,
	})
}

// getTransactionListDatabse get transaction data from database
func (h *Handler) getTransactionListDatabse(c context.Context, request GetActionListRequest) ([]*model.Transaction, int64, error) {
	tracer := otel.Tracer("getActionListDatabase")
	_, postgresSnap := tracer.Start(c, "postgres")

	defer postgresSnap.End()

	transactionList := []*model.Transaction{}
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

		sql = sql.Where("timestamp <= ?", lastItem.Timestamp).
			Where(
				"hash != ? AND index < ?",
				lastItem.Hash, lastItem.Index,
			)
	}

	if len(request.Tags) > 0 {
		sql = sql.Where("tag IN ?", request.Tags)
	}

	if len(request.Networks) > 0 {
		sql = sql.Where("network IN ?", request.Networks)
	}

	if err := sql.Limit(request.Limit).Order("timestamp DESC").Find(&transactionList).
		Limit(-1).Count(&total).Error; err != nil {
		return nil, 0, err
	}

	return transactionList, total, nil
}

// getActionListDatabase get transfer data from database
func (h *Handler) getActionListDatabase(c context.Context, transactionHashList []string, request GetActionListRequest) ([]*model.Transfer, error) {
	tracer := otel.Tracer("getActionListDatabase")
	_, postgresSnap := tracer.Start(c, "postgres")

	defer postgresSnap.End()

	transfers := make([]*model.Transfer, 0)

	sql := h.DatabaseClient.Model(&model.Transfer{})

	if len(request.Types) > 0 {
		sql = sql.Where("\"type\" in ? ", request.Types)
	}

	if err := sql.
		Where("transaction_hash IN ?", transactionHashList).
		Find(&transfers).Error; err != nil {
		return nil, err
	}

	return transfers, nil
}

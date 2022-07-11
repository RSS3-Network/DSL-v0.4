package handler

import (
	"context"
	"encoding/json"
	"net/http"
	"sort"
	"strings"

	"github.com/labstack/echo/v4"
	dbModel "github.com/naturalselectionlabs/pregod/common/database/model"
	"github.com/naturalselectionlabs/pregod/common/protocol"
	"github.com/naturalselectionlabs/pregod/common/protocol/filter"
	rabbitmq "github.com/rabbitmq/amqp091-go"
	lop "github.com/samber/lo/parallel"
	"github.com/sirupsen/logrus"
	"go.opentelemetry.io/otel"
)

// GetNoteListFunc HTTP handler for action API
// parse query parameters, query and assemble data
func (h *Handler) GetNoteListFunc(c echo.Context) error {
	tracer := otel.Tracer("GetActionListFunc")
	ctx, httpSnap := tracer.Start(c.Request().Context(), "http")

	defer httpSnap.End()

	request := GetRequest{}

	if err := c.Bind(&request); err != nil {
		return c.JSON(http.StatusBadRequest, http.StatusText(http.StatusBadRequest))
	}

	request.Address = strings.ToLower(request.Address)

	if request.Limit <= 0 || request.Limit > DefaultLimit {
		request.Limit = DefaultLimit
	}

	var types []string
	for _, t := range request.Type {
		t = strings.ToLower(t)
		if filter.CheckTypeValid(request.Tag, t) {
			types = append(types, t)
		}
	}

	request.Type = types

	transactionList, total, err := h.getTransactionListDatabase(ctx, request)
	if err != nil {
		return c.JSON(http.StatusBadRequest, http.StatusText(http.StatusBadRequest))
	}

	// publish mq message
	if request.Refresh || len(transactionList) == 0 {
		go h.publishIndexerMessage(ctx, request.Address)
	}

	if len(transactionList) == 0 {
		return c.JSON(http.StatusOK, &Response{
			Result: make([]dbModel.Transaction, 0),
		})
	}

	transactionHashList := make([]string, 0)
	for _, transactionHash := range transactionList {
		transactionHashList = append(transactionHashList, transactionHash.Hash)
	}

	transferList, err := h.getTransferListDatabase(ctx, transactionHashList, request.Type)
	if err != nil {
		return c.JSON(http.StatusBadRequest, http.StatusText(http.StatusBadRequest))
	}

	transferMap := make(map[string][]dbModel.Transfer)
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
func (h *Handler) getTransactionListDatabase(c context.Context, request GetRequest) ([]dbModel.Transaction, int64, error) {
	tracer := otel.Tracer("getNoteListDatabase")
	_, postgresSnap := tracer.Start(c, "postgres")

	defer postgresSnap.End()

	transactionList := make([]dbModel.Transaction, 0)
	total := int64(0)
	sql := h.DatabaseClient.
		Model(&dbModel.Transaction{}).
		Where("LOWER(address_from) = ? OR LOWER(address_to) = ?",
			// address was already converted to lowercase
			request.Address,
			request.Address,
		)

	if len(request.Cursor) > 0 {
		var lastItem dbModel.Transaction

		if err := h.DatabaseClient.Where("hash = ?", strings.ToLower(request.Cursor)).First(&lastItem).Error; err != nil {
			return nil, 0, err
		}

		sql = sql.Where("timestamp < ? OR (timestamp = ? AND index < ?)", lastItem.Timestamp, lastItem.Timestamp, lastItem.Index)
	}

	if len(request.Hash) > 0 {
		sql = sql.Where("hash = ?", strings.ToLower(request.Hash))
	}

	if len(request.Tag) > 0 {
		sql = sql.Where("tag = ?", strings.ToLower(request.Tag))
	}

	if len(request.Type) > 0 {
		// type was already converted to lowercase
		sql = sql.Where("\"type\" IN ?", request.Type)
	}

	if len(request.Network) > 0 {
		for i, v := range request.Network {
			request.Network[i] = strings.ToLower(v)
		}

		sql = sql.Where("network IN ?", request.Network)
	}

	if len(request.Platform) > 0 {
		for i, v := range request.Platform {
			request.Platform[i] = strings.ToLower(v)
		}
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

// getTransferListDatabase get transfer data from database
func (h *Handler) getTransferListDatabase(c context.Context, transactionHashList []string, requestType []string) ([]dbModel.Transfer, error) {
	tracer := otel.Tracer("getNoteListDatabase")
	_, postgresSnap := tracer.Start(c, "postgres")

	defer postgresSnap.End()

	transfers := make([]dbModel.Transfer, 0)

	sql := h.DatabaseClient.Model(&dbModel.Transfer{})

	if len(requestType) > 0 {
		sql = sql.Where("\"type\" IN ? ", requestType)
	}

	if err := sql.
		Where("transaction_hash IN ?", transactionHashList).
		Find(&transfers).Error; err != nil {
		return nil, err
	}

	return transfers, nil
}

// BatchGetNoteListFunc query multiple addresses and filters
func (h *Handler) BatchGetNoteListFunc(c echo.Context) error {
	tracer := otel.Tracer("BatchGetNoteListFunc")
	ctx, httpSnap := tracer.Start(c.Request().Context(), "BatchGetNoteListFunc:http")

	defer httpSnap.End()

	request := BatchGetNoteListRequest{}

	if err := c.Bind(&request); err != nil {
		return c.JSON(http.StatusBadRequest, http.StatusText(http.StatusBadRequest))
	}

	if request.Limit <= 0 || request.Limit > DefaultLimit {
		request.Limit = DefaultLimit
	}

	if len(request.List) > DefaultBatchGetLimit {
		request.List = request.List[:DefaultBatchGetLimit]
	}

	transactionList, total, err := h.batchGetTransactionListDatabase(ctx, request)
	if err != nil {
		return c.JSON(http.StatusBadRequest, http.StatusText(http.StatusBadRequest))
	}

	if total == 0 {
		return c.JSON(http.StatusOK, &Response{
			Result: make([]dbModel.Transaction, 0),
		})
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

func (h *Handler) batchGetTransactionListDatabase(ctx context.Context, request BatchGetNoteListRequest) ([]dbModel.Transaction, int64, error) {
	tracer := otel.Tracer("batchGetTransactionListDatabase")
	_, postgresSnap := tracer.Start(ctx, "postgres")

	defer postgresSnap.End()

	if len(request.List) == 0 {
		return nil, 0, nil
	}

	results := make([][]dbModel.Transaction, len(request.List))
	total := int64(0)
	lastItem := dbModel.Transaction{}

	if len(request.Cursor) > 0 {
		if err := h.DatabaseClient.Where("hash = ?", strings.ToLower(request.Cursor)).First(&lastItem).Error; err != nil {
			return nil, 0, err
		}
	}

	lop.ForEach(request.List, func(reqFilter Filter, i int) {
		var types []string
		count := int64(0)
		transactionList := make([]dbModel.Transaction, 0)

		// check type
		for _, t := range reqFilter.Type {
			t = strings.ToLower(t)
			if filter.CheckTypeValid(reqFilter.Tag, t) {
				types = append(types, t)
			}
		}
		reqFilter.Type = types

		// query transaction
		sql := h.DatabaseClient.Model(&dbModel.Transaction{}).
			Where("LOWER(address_from) = ? OR LOWER(address_to) = ?",
				strings.ToLower(reqFilter.Address),
				strings.ToLower(reqFilter.Address),
			)

		if len(request.Cursor) > 0 {
			sql = sql.Where("timestamp < ? OR (timestamp = ? AND index < ?)", lastItem.Timestamp, lastItem.Timestamp, lastItem.Index)
		}

		if len(reqFilter.Tag) > 0 {
			sql = sql.Where("tag = ?", reqFilter.Tag)
		}

		if len(reqFilter.Type) > 0 {
			sql = sql.Where("\"type\" IN ?", reqFilter.Type)
		}

		if len(reqFilter.Network) > 0 {
			sql = sql.Where("network IN ?", reqFilter.Network)
		}

		if len(reqFilter.Platform) > 0 {
			sql = sql.Where("platform IN ?", reqFilter.Platform)
		}

		if len(request.Timestamp.String()) > 0 {
			sql = sql.Where("timestamp > ?", request.Timestamp)
		}

		if err := sql.Count(&count).Limit(request.Limit).Order("timestamp DESC, index DESC").Find(&transactionList).Error; err != nil {
			logrus.Errorf("batchGetTransactionListDatabase: query transaction error, %v", err)
			return
		}

		// query transfer
		transactionHashList := make([]string, 0)
		for _, transactionHash := range transactionList {
			transactionHashList = append(transactionHashList, transactionHash.Hash)
		}

		transferList, err := h.getTransferListDatabase(ctx, transactionHashList, reqFilter.Type)
		if err != nil {
			logrus.Errorf("batchGetTransactionListDatabase: query transfer error, %v", err)
			return
		}

		transferMap := make(map[string][]dbModel.Transfer)
		for _, transfer := range transferList {
			transferMap[transfer.TransactionHash] = append(transferMap[transfer.TransactionHash], transfer)
		}

		for index := range transactionList {
			transactionList[index].Transfers = transferMap[transactionList[index].Hash]
		}

		// publish mq message
		if request.Refresh || len(transactionHashList) == 0 {
			go h.publishIndexerMessage(ctx, reqFilter.Address)
		}

		results[i] = transactionList
		total += count
	})

	// sort
	transactions := make([]dbModel.Transaction, 0)

	for _, result := range results {
		transactions = append(transactions, result...)
	}

	sort.Sort(Transactions(transactions))

	if len(transactions) > request.Limit {
		transactions = transactions[:request.Limit]
	}

	return transactions, total, nil
}

// publishIndexerMessage create a rabbitmq job to index the latest user data
func (h *Handler) publishIndexerMessage(ctx context.Context, address string) {
	tracer := otel.Tracer("publishIndexerMessage")
	_, rabbitmqSnap := tracer.Start(ctx, "rabbitmq")

	defer rabbitmqSnap.End()

	networks := []string{
		protocol.NetworkEthereum, protocol.NetworkPolygon, protocol.NetworkBinanceSmartChain,
		protocol.NetworkArweave, protocol.NetworkXDAI, protocol.NetworkZkSync, protocol.NetworkCrossbell,
	}

	for _, network := range networks {
		message := protocol.Message{
			Address: address,
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
}

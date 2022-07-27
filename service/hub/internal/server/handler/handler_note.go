package handler

import (
	"context"
	"encoding/json"
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
	dbModel "github.com/naturalselectionlabs/pregod/common/database/model"
	"github.com/naturalselectionlabs/pregod/common/protocol"
	"github.com/naturalselectionlabs/pregod/common/protocol/filter"
	rabbitmq "github.com/rabbitmq/amqp091-go"
	"go.opentelemetry.io/otel"
)

// GetNotesFunc HTTP handler for action API
// parse query parameters, query and assemble data
func (h *Handler) GetNotesFunc(c echo.Context) error {
	tracer := otel.Tracer("GetNotesFunc")
	ctx, httpSnap := tracer.Start(c.Request().Context(), "http")

	defer httpSnap.End()

	request := GetRequest{}

	if err := c.Bind(&request); err != nil {
		return BadRequest(c)
	}

	if err := c.Validate(&request); err != nil {
		return err
	}

	if request.Limit <= 0 || request.Limit > DefaultLimit {
		request.Limit = DefaultLimit
	}

	var types []string
	for _, t := range request.Type {
		t = strings.ToLower(t)

		if filter.CheckTypeValid(request.Tag, t) {
			types = append(types, t)
		}
		// by default POAPs are not returned
		if t == filter.CollectiblePoap {
			request.IncludePoap = true
		}
	}

	request.Type = types

	transactions, total, err := h.getTransactions(ctx, request)
	if err != nil {
		return InternalError(c)
	}

	// publish mq message
	if request.Refresh || len(transactions) == 0 {
		go h.publishIndexerMessage(ctx, request.Address)
	}

	if len(transactions) == 0 {
		return c.JSON(http.StatusOK, &Response{
			Result: make([]dbModel.Transaction, 0),
		})
	}

	transactionHashes := make([]string, 0)
	for _, transactionHash := range transactions {
		transactionHashes = append(transactionHashes, transactionHash.Hash)
	}

	transfers, err := h.getTransfers(ctx, transactionHashes, request.Type)
	if err != nil {
		return InternalError(c)
	}

	transferMap := make(map[string][]dbModel.Transfer)
	for _, transfer := range transfers {
		transferMap[transfer.TransactionHash] = append(transferMap[transfer.TransactionHash], transfer)
	}

	for index := range transactions {
		transactions[index].Transfers = transferMap[transactions[index].Hash]
	}

	var cursor string
	if total > int64(request.Limit) {
		cursor = transactions[len(transactions)-1].Hash
	}

	return c.JSON(http.StatusOK, &Response{
		Total:  total,
		Cursor: cursor,
		Result: transactions,
	})
}

// getTransactions get transaction data from database
func (h *Handler) getTransactions(c context.Context, request GetRequest) ([]dbModel.Transaction, int64, error) {
	tracer := otel.Tracer("getTransactions")
	_, postgresSnap := tracer.Start(c, "postgres")

	defer postgresSnap.End()

	transactions := make([]dbModel.Transaction, 0)
	total := int64(0)
	sql := h.DatabaseClient.
		Model(&dbModel.Transaction{}).
		Where("owner = ?", request.Address) // address was already converted to lowercase

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

	if !request.IncludePoap {
		sql = sql.Where("type != ?", filter.CollectiblePoap)
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

	if request.Timestamp.Unix() > 0 {
		sql = sql.Where("timestamp > ?", request.Timestamp)
	}

	if err := sql.Count(&total).Limit(request.Limit).Order("timestamp DESC, index DESC").Find(&transactions).Error; err != nil {
		return nil, 0, err
	}

	return transactions, total, nil
}

// getTransfers get transfer data from database
func (h *Handler) getTransfers(c context.Context, transactionHashes []string, requestType []string) ([]dbModel.Transfer, error) {
	tracer := otel.Tracer("getTransfers")
	_, postgresSnap := tracer.Start(c, "postgres")

	defer postgresSnap.End()

	transfers := make([]dbModel.Transfer, 0)

	sql := h.DatabaseClient.Model(&dbModel.Transfer{})

	if len(requestType) > 0 {
		sql = sql.Where("\"type\" IN ? ", requestType)
	}

	if err := sql.
		Where("transaction_hash IN ?", transactionHashes).
		Find(&transfers).Error; err != nil {
		return nil, err
	}

	return transfers, nil
}

// BatchGetNotesFunc query multiple addresses and filters
func (h *Handler) BatchGetNotesFunc(c echo.Context) error {
	tracer := otel.Tracer("BatchGetNotesFunc")
	ctx, httpSnap := tracer.Start(c.Request().Context(), "BatchGetNotesFunc:http")

	defer httpSnap.End()

	request := BatchGetNotesRequest{}

	if err := c.Bind(&request); err != nil {
		return BadRequest(c)
	}

	if request.Limit <= 0 || request.Limit > DefaultLimit {
		request.Limit = DefaultLimit
	}

	var err error
	var transactions []dbModel.Transaction
	var total int64

	if len(request.Address) == 0 {
		return AddressIsEmpty(c)
	}

	if len(request.Address) > DefaultLimit {
		request.Address = request.Address[:DefaultLimit]
	}
	transactions, total, err = h.batchGetTransactions(ctx, request)

	if err != nil {
		return InternalError(c)
	}

	if total == 0 {
		return c.JSON(http.StatusOK, &Response{
			Result: make([]dbModel.Transaction, 0),
		})
	}

	var cursor string
	if total > int64(request.Limit) {
		cursor = transactions[len(transactions)-1].Hash
	}

	return c.JSON(http.StatusOK, &Response{
		Total:  total,
		Cursor: cursor,
		Result: transactions,
	})
}

func (h *Handler) batchGetTransactions(ctx context.Context, request BatchGetNotesRequest) ([]dbModel.Transaction, int64, error) {
	tracer := otel.Tracer("batchGetTransactions")
	_, postgresSnap := tracer.Start(ctx, "postgres")

	defer postgresSnap.End()

	transactions := make([]dbModel.Transaction, 0)
	total := int64(0)

	for i, v := range request.Address {
		request.Address[i] = strings.ToLower(v)
	}

	sql := h.DatabaseClient.
		Model(&dbModel.Transaction{}).
		Where("owner IN ?", request.Address)

	if len(request.Cursor) > 0 {
		var lastItem dbModel.Transaction

		if err := h.DatabaseClient.Where("hash = ?", strings.ToLower(request.Cursor)).First(&lastItem).Error; err != nil {
			return nil, 0, err
		}

		sql = sql.Where("timestamp < ? OR (timestamp = ? AND index < ?)", lastItem.Timestamp, lastItem.Timestamp, lastItem.Index)
	}

	if len(request.Tag) > 0 {
		sql = sql.Where("tag = ?", strings.ToLower(request.Tag))
	}

	var types []string
	for _, t := range request.Type {
		t = strings.ToLower(t)
		if filter.CheckTypeValid(request.Tag, t) {
			types = append(types, t)
		}

		// by default POAPs are not returned
		if t == filter.CollectiblePoap {
			request.IncludePoap = true
		}
	}

	request.Type = types

	if len(request.Type) > 0 {
		// type was already converted to lowercase
		sql = sql.Where("\"type\" IN ?", request.Type)
	}

	if !request.IncludePoap {
		sql = sql.Where("type != ?", filter.CollectiblePoap)
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

	if request.Timestamp.Unix() > 0 {
		sql = sql.Where("timestamp > ?", request.Timestamp)
	}

	if err := sql.Count(&total).Limit(request.Limit).Order("timestamp DESC, index DESC").Find(&transactions).Error; err != nil {
		return nil, 0, err
	}

	transactionHashes := make([]string, 0)
	for _, transactionHash := range transactions {
		transactionHashes = append(transactionHashes, transactionHash.Hash)
	}

	transfers, err := h.getTransfers(ctx, transactionHashes, request.Type)
	if err != nil {
		return nil, 0, err
	}

	transferMap := make(map[string][]dbModel.Transfer)
	for _, transfer := range transfers {
		transferMap[transfer.TransactionHash] = append(transferMap[transfer.TransactionHash], transfer)
	}

	for index := range transactions {
		transactions[index].Transfers = transferMap[transactions[index].Hash]
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

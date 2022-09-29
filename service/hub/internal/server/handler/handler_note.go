package handler

import (
	"context"
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
	"github.com/lib/pq"
	"github.com/naturalselectionlabs/pregod/common/database"
	dbModel "github.com/naturalselectionlabs/pregod/common/database/model"
	"github.com/naturalselectionlabs/pregod/common/protocol"
	"github.com/naturalselectionlabs/pregod/common/protocol/filter"
	"go.opentelemetry.io/otel"
)

// GetNotesFunc HTTP handler for action API
// parse query parameters, query and assemble data
func (h *Handler) GetNotesFunc(c echo.Context) error {
	go h.apiReport(GetNotes, c)
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

	request.Address = strings.ToLower(request.Address)

	if len(request.Cursor) == 0 {
		go h.filterReport(GetNotes, request)
	}

	if request.Limit <= 0 || request.Limit > DefaultLimit {
		request.Limit = DefaultLimit
	}

	// support many-many relationship between tag and type
	var tags []string
	var types []string
	if len(request.Tag) > 0 {
		for _, tag := range request.Tag {
			tag = strings.ToLower(tag)
			if len(request.Type) > 0 {
				for _, typeX := range request.Type {
					typeX = strings.ToLower(typeX)

					if filter.CheckTypeValid(tag, typeX) {
						tags = append(tags, tag)
						types = append(types, typeX)
						// by default POAPs are not returned
						if typeX == filter.CollectiblePoap {
							request.IncludePoap = true
						}
					}
				}
			} else {
				tags = append(tags, tag)
			}
		}
		if len(tags) == 0 {
			return BadParams(c, request.Tag, request.Type)
		}
		request.Tag = tags
		request.Type = types
	}

	transactions, total, err := h.getTransactions(ctx, request)
	if err != nil {
		return InternalError(c)
	}

	// publish mq message
	if len(request.Cursor) == 0 && (request.Refresh || len(transactions) == 0) {
		h.publishIndexerMessage(ctx, protocol.Message{Address: request.Address, Reindex: request.Reindex})
	}

	if request.CountOnly {
		return c.JSON(http.StatusOK, &Response{
			Total: &total,
		})
	}

	var cursor string
	if total > int64(request.Limit) {
		cursor = transactions[len(transactions)-1].Hash
	}

	var addressStatus []dbModel.Address
	if request.QueryStatus {
		addressStatus, _ = h.getAddress(ctx, []string{request.Address})
	}

	if len(transactions) == 0 {
		return c.JSON(http.StatusOK, &Response{
			Result:        make([]dbModel.Transaction, 0),
			AddressStatus: addressStatus,
		})
	}

	transactionHashes := make([]string, 0)
	for _, transactionHash := range transactions {
		transactionHashes = append(transactionHashes, transactionHash.Hash)
	}

	transfers, err := h.getTransfers(ctx, transactionHashes)
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

	return c.JSON(http.StatusOK, &Response{
		Total:         &total,
		Cursor:        cursor,
		Result:        transactions,
		AddressStatus: addressStatus,
	})
}

// getTransactions get transaction data from database
func (h *Handler) getTransactions(c context.Context, request GetRequest) ([]dbModel.Transaction, int64, error) {
	tracer := otel.Tracer("getTransactions")
	_, postgresSnap := tracer.Start(c, "postgres")

	defer postgresSnap.End()

	transactions := make([]dbModel.Transaction, 0)
	total := int64(0)
	sql := database.Global().
		Model(&dbModel.Transaction{}).
		Where("owner = ?", request.Address). // address was already converted to lowercase
		Where("success IS TRUE")             // Hide failed transactions

	if len(request.Cursor) > 0 {
		var lastItem dbModel.Transaction

		// no need to lowercase
		if err := database.Global().Where("hash = ?", request.Cursor).First(&lastItem).Error; err != nil {
			return nil, 0, err
		}

		sql = sql.Where("timestamp < ? OR (timestamp = ? AND index < ?)", lastItem.Timestamp, lastItem.Timestamp, lastItem.Index)
	}

	if len(request.Hash) > 0 {
		sql = sql.Where("hash = ?", request.Hash) // no need to lowercase
	}

	if len(request.Tag) > 0 {
		sql = sql.Where("tag IN ?", request.Tag)
		if len(request.Type) > 0 {
			// type was already converted to lowercase
			sql = sql.Where("\"type\" IN ?", request.Type)
		}
	}

	if !request.IncludePoap {
		sql = sql.Where("\"type\" != ?", filter.CollectiblePoap)
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
		sql = sql.Where("LOWER(platform) IN ?", request.Platform)
	}

	if request.Timestamp.Unix() > 0 {
		sql = sql.Where("timestamp > ?", request.Timestamp)
	}

	if err := sql.Count(&total).Limit(request.Limit).Offset(request.Page * request.Limit).Order("timestamp DESC, index DESC").Find(&transactions).Error; err != nil {
		return nil, 0, err
	}

	return transactions, total, nil
}

// getTransfers get transfer data from database
func (h *Handler) getTransfers(c context.Context, transactionHashes []string) ([]dbModel.Transfer, error) {
	tracer := otel.Tracer("getTransfers")
	_, postgresSnap := tracer.Start(c, "postgres")

	defer postgresSnap.End()

	transfers := make([]dbModel.Transfer, 0)

	sql := database.Global().Model(&dbModel.Transfer{})

	if err := sql.
		Where("transaction_hash IN (SELECT * FROM UNNEST(?::TEXT[]))", pq.Array(transactionHashes)).
		Find(&transfers).Error; err != nil {
		return nil, err
	}

	return transfers, nil
}

// BatchGetNotesFunc query multiple addresses and filters
func (h *Handler) BatchGetNotesFunc(c echo.Context) error {
	go h.apiReport(PostNotes, c)
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

	if len(request.Cursor) == 0 {
		go h.filterReport(PostNotes, request)
	}

	// support many-many relationship between tag and type
	var tags []string
	var types []string
	if len(request.Tag) > 0 {
		for _, tag := range request.Tag {
			tag = strings.ToLower(tag)
			if len(request.Type) > 0 {
				for _, typeX := range request.Type {
					typeX = strings.ToLower(typeX)

					if filter.CheckTypeValid(tag, typeX) {
						tags = append(tags, tag)
						types = append(types, typeX)
						// by default POAPs are not returned
						if typeX == filter.CollectiblePoap {
							request.IncludePoap = true
						}
					}
				}
			} else {
				tags = append(tags, tag)
			}
		}
		if len(tags) == 0 {
			return BadParams(c, request.Tag, request.Type)
		}
		request.Tag = tags
		request.Type = types
	}

	transactions, total, err = h.batchGetTransactions(ctx, request)
	if err != nil {
		return InternalError(c)
	}

	// publish mq message
	if len(request.Cursor) == 0 && (request.Refresh || len(transactions) == 0) {
		for _, address := range request.Address {
			h.publishIndexerMessage(ctx, protocol.Message{Address: address})
		}
	}

	var addressStatus []dbModel.Address
	if request.QueryStatus {
		addressStatus, _ = h.getAddress(ctx, request.Address)
	}

	if request.CountOnly {
		return c.JSON(http.StatusOK, &Response{
			Total: &total,
		})
	}

	var cursor string
	if total > int64(request.Limit) {
		cursor = transactions[len(transactions)-1].Hash
	}

	if total == 0 {
		return c.JSON(http.StatusOK, &Response{
			Result:        make([]dbModel.Transaction, 0),
			AddressStatus: addressStatus,
		})
	}

	return c.JSON(http.StatusOK, &Response{
		Total:         &total,
		Cursor:        cursor,
		Result:        transactions,
		AddressStatus: addressStatus,
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

	sql := database.Global().
		Model(&dbModel.Transaction{}).
		Where("owner IN ?", request.Address).
		Where("success IS TRUE") // Hide failed transactions

	if len(request.Cursor) > 0 {
		var lastItem dbModel.Transaction

		// no need to lowercase
		if err := database.Global().Where("hash = ?", request.Cursor).First(&lastItem).Error; err != nil {
			return nil, 0, err
		}

		sql = sql.Where("timestamp < ? OR (timestamp = ? AND index < ?)", lastItem.Timestamp, lastItem.Timestamp, lastItem.Index)
	}

	if len(request.Tag) > 0 {
		sql = sql.Where("tag IN ?", request.Tag)
		if len(request.Type) > 0 {
			// type was already converted to lowercase
			sql = sql.Where("\"type\" IN ?", request.Type)
		}
	}

	if !request.IncludePoap {
		sql = sql.Where("\"type\" != ?", filter.CollectiblePoap)
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
		sql = sql.Where("LOWER(platform) IN ?", request.Platform)
	}

	if request.Timestamp.Unix() > 0 {
		sql = sql.Where("timestamp > ?", request.Timestamp)
	}

	if err := sql.Count(&total).Limit(request.Limit).Offset(request.Page * request.Limit).Order("timestamp DESC, index DESC").Find(&transactions).Error; err != nil {
		return nil, 0, err
	}

	transactionHashes := make([]string, 0)
	for _, transactionHash := range transactions {
		transactionHashes = append(transactionHashes, transactionHash.Hash)
	}

	transfers, err := h.getTransfers(ctx, transactionHashes)
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

// getAddress
func (h *Handler) getAddress(ctx context.Context, address []string) ([]dbModel.Address, error) {
	tracer := otel.Tracer("getAddress")
	_, postgresSnap := tracer.Start(ctx, "postgres")

	defer postgresSnap.End()

	addressStatus := make([]dbModel.Address, 0)

	sql := database.Global().Model(&dbModel.Address{})

	if err := sql.
		Where("address IN ?", address).
		Find(&addressStatus).Error; err != nil {
		return nil, err
	}

	return addressStatus, nil
}

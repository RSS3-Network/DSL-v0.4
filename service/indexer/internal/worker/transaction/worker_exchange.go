package transaction

import (
	"context"
	"encoding/csv"
	"errors"
	"fmt"
	"path"
	"time"

	"github.com/naturalselectionlabs/pregod/common/cache"
	"github.com/naturalselectionlabs/pregod/common/database"
	"github.com/naturalselectionlabs/pregod/common/database/model"
	"github.com/naturalselectionlabs/pregod/common/database/model/exchange"
	"github.com/naturalselectionlabs/pregod/common/protocol/filter"
	"github.com/naturalselectionlabs/pregod/internal/allowlist"

	"go.uber.org/zap"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func keyOfCheckCexWallet(address string) string {
	return fmt.Sprintf("check_exchange_wallet.%s", address)
}

func (s *service) loadCentralizedExchangeWallets(ctx context.Context) error {
	rootPath := "asset/cex_wallet"

	files, err := asseFS.ReadDir(rootPath)
	if err != nil {
		return fmt.Errorf("read directory %s: %w", rootPath, err)
	}

	for _, file := range files {
		records, err := s.loadCentralizedExchangeWallet(path.Join(rootPath, file.Name()))
		if err != nil {
			return fmt.Errorf("read file %s: %w", file.Name(), err)
		}

		wallets := make([]exchange.CexWallet, 0, len(records))

		for i, record := range records {
			// Skip table header
			if i == 0 {
				continue
			}

			wallets = append(wallets, exchange.CexWallet{
				WalletAddress: record[0],
				Name:          record[1],
				Source:        record[2],
				Network:       record[3],
			})

			allowlist.AllowList.Add(record[0], record[1])
		}

		if len(wallets) == 0 {
			return nil
		}

		clauses := clause.OnConflict{
			DoNothing: true,
		}

		if err := database.Global().Clauses(clauses).Create(&wallets).Error; err != nil {
			return fmt.Errorf("create centralized exchange wallets: %w", err)
		}
	}

	return nil
}

func (s *service) loadCentralizedExchangeWallet(path string) ([][]string, error) {
	file, err := asseFS.Open(path)
	if err != nil {
		return nil, fmt.Errorf("open file %s: %w", path, err)
	}

	defer func() {
		_ = file.Close()
	}()

	records, err := csv.NewReader(file).ReadAll()
	if err != nil {
		return nil, fmt.Errorf("read file %s: %w", path, err)
	}

	return records, nil
}

func (s *service) isCentralizedExchangeWallet(ctx context.Context, addresses ...string) (*exchange.CexWallet, error) {
	var (
		wallet *exchange.CexWallet
		missed bool
	)

	for _, address := range addresses {
		// Get centralized exchange wallet from cache
		exists, err := cache.GetMsgPack(ctx, keyOfCheckCexWallet(address), &wallet)
		if err != nil {
			return nil, fmt.Errorf("get centralized exchange wallet from cache: %w", err)
		}

		switch {
		case !exists: // No data about whether this address is a centralized exchange wallet in the cache
			missed = true
		case exists && wallet.Name != "": // Matched a centralized exchange wallet
			return wallet, nil
		}
	}

	if !missed { // All wallets have been matched, but no exchange wallets were found
		return nil, nil
	}

	queries := map[string]any{
		"wallet_address": addresses,
	}

	if err := database.Global().WithContext(ctx).Where(queries).First(&wallet).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			for _, address := range addresses {
				// Set the user wallet address to cache
				if err := cache.SetMsgPack(ctx, keyOfCheckCexWallet(address), wallet, 7*24*time.Hour); err != nil {
					return nil, fmt.Errorf("set user wallet address to cache: %w", err)
				}

				zap.L().Debug("set user wallet to cache", zap.String("address", address))
			}

			return nil, nil
		}

		return nil, fmt.Errorf("find centralized exchange wallet from database: %w", err)
	}

	// Set the centralized exchange wallet address to cache
	if err := cache.SetMsgPack(ctx, keyOfCheckCexWallet(wallet.WalletAddress), wallet, 7*24*time.Hour); err != nil {
		return nil, fmt.Errorf("set centralized exchange wallet to cache: %w", err)
	}

	zap.L().Debug("set centralized exchange wallet to cache", zap.String("address", wallet.WalletAddress), zap.String("name", wallet.Name))

	return wallet, nil
}

func (s *service) buildCentralizedExchangeTransfer(ctx context.Context, transfer *model.Transfer, wallet exchange.CexWallet, owner string) error {
	if transfer.Tag = filter.UpdateTag(filter.TagExchange, transfer.Tag); transfer.Tag == filter.TagExchange {
		transfer.Platform = wallet.Name

		switch {
		case transfer.AddressTo == owner:
			transfer.Type = filter.ExchangeWithdraw
		case transfer.AddressFrom == owner:
			transfer.Type = filter.ExchangeDeposit
		}
	}

	return nil
}

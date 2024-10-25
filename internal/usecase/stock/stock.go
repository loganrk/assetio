package stock

import (
	"assetio/internal/domain"
	"assetio/internal/port"
	"context"
	"errors"
	"time"
)

type stockUsecase struct {
	logger port.Logger
	mysql  port.RepositoryMySQL
}

func New(loggerIns port.Logger, mysqlIns port.RepositoryMySQL) domain.StockSvr {
	return &stockUsecase{
		mysql:  mysqlIns,
		logger: loggerIns,
	}
}

func (s *stockUsecase) BuyStock(ctx context.Context, userId int, accountId int, inventoryId int, stockId int, quantity int, amountPerQuantity float64, taxAmount float64) error {
	secuirity, err := s.mysql.GetSecuriryById(ctx, stockId)
	if err != nil {
		return err
	}
	var inventory domain.Inventory

	if inventoryId != 0 {
		inventory, err = s.mysql.GetInventoryByInventoryIdAndAccountId(ctx, inventoryId)
		if err != nil {
			return err
		}
		if inventory.Id == 0 {
			return errors.New("incorrect inventory")
		}

		if accountId != inventory.AccountId {
			return errors.New("incorrect inventory")
		}

		if secuirity.Id != inventory.SecurityId {
			return errors.New("incorrect stock")
		}

	} else {
		inventory, err = s.mysql.InsertInventory(ctx, domain.Inventory{
			AccountId:  accountId,
			SecurityId: secuirity.Id,
		})
		if err != nil {
			return err
		}
	}

	_, err = s.mysql.InsertTransaction(ctx, domain.Transaction{
		AccountId:   accountId,
		InventoryId: inventory.Id,
		Type:        domain.Buy,
		Quantity:    quantity,
		Price:       float64(quantity) * amountPerQuantity,
		Fee:         taxAmount,
		Date:        time.Now(),
	})

	if err != nil {
		return err
	}

	return nil
}

func (s *stockUsecase) SellStock(ctx context.Context, userId int, accountId int, inventoryId int, stockId int, quantity int, amountPerQuantity float64, taxAmount float64) error {
	secuirity, err := s.mysql.GetSecuriryById(ctx, stockId)
	if err != nil {
		return err
	}

	inventory, err := s.mysql.GetInventoryByInventoryIdAndAccountId(ctx, inventoryId)
	if err != nil {
		return err
	}

	if accountId != inventory.AccountId {
		return errors.New("incorrect inventory")
	}

	if secuirity.Id != inventory.SecurityId {
		return errors.New("incorrect stock")
	}

	_, err = s.mysql.InsertTransaction(ctx, domain.Transaction{
		AccountId:   accountId,
		InventoryId: inventory.Id,
		Type:        domain.Sell,
		Quantity:    quantity,
		Price:       float64(quantity) * amountPerQuantity,
		Fee:         taxAmount,
		Date:        time.Now(),
	})

	if err != nil {
		return err
	}
	return nil
}

func (s *stockUsecase) StockDividendAdd(ctx context.Context, userId int, accountId int, inventoryId int, stockId int, quantity int, amountPerQuantity float64) error {
	secuirity, err := s.mysql.GetSecuriryById(ctx, stockId)
	if err != nil {
		return err
	}

	inventory, err := s.mysql.GetInventoryByInventoryIdAndAccountId(ctx, inventoryId)
	if err != nil {
		return err
	}

	if accountId != inventory.AccountId {
		return errors.New("incorrect inventory")
	}

	if secuirity.Id != inventory.SecurityId {
		return errors.New("incorrect stock")
	}

	_, err = s.mysql.InsertTransaction(ctx, domain.Transaction{
		AccountId:   accountId,
		InventoryId: inventory.Id,
		Type:        domain.Dividend,
		Quantity:    quantity,
		Price:       float64(quantity) * amountPerQuantity,
		Date:        time.Now(),
	})

	if err != nil {
		return err
	}
	return nil
}

package mutualFund

import (
	"assetio/internal/constant"
	"assetio/internal/domain"
	"assetio/internal/port"
	"context"
	"errors"
	"time"
)

type mutualFundUsecase struct {
	logger port.Logger
	mysql  port.RepositoryMySQL
}

func New(loggerIns port.Logger, mysqlIns port.RepositoryMySQL) domain.MutualFundSvr {
	return &mutualFundUsecase{
		mysql:  mysqlIns,
		logger: loggerIns,
	}
}

func (m *mutualFundUsecase) BuyMutualFund(ctx context.Context, userId int, accountId int, inventoryId int, stockId int, quantity int, amountPerQuantity float64, taxAmount float64) error {
	secuirity, err := m.mysql.GetSecuriryById(ctx, stockId)
	if err != nil {
		return err
	}

	if secuirity.Type != constant.SECURITY_TYPE_MUTUAL_FUND {
		return errors.New("incorrect stock")
	}

	var inventory domain.Inventory

	if inventoryId != 0 {
		inventory, err = m.mysql.GetInventoryByInventoryIdAndAccountId(ctx, inventoryId)
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
		inventory, err = m.mysql.InsertInventory(ctx, domain.Inventory{
			AccountId:  accountId,
			SecurityId: secuirity.Id,
		})
		if err != nil {
			return err
		}
	}

	_, err = m.mysql.InsertTransaction(ctx, domain.Transaction{
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

func (m *mutualFundUsecase) SellMutualFund(ctx context.Context, userId int, accountId int, inventoryId int, stockId int, quantity int, amountPerQuantity float64, taxAmount float64) error {
	secuirity, err := m.mysql.GetSecuriryById(ctx, stockId)
	if err != nil {
		return err
	}
	if secuirity.Type != constant.SECURITY_TYPE_MUTUAL_FUND {
		return errors.New("incorrect stock")
	}

	inventory, err := m.mysql.GetInventoryByInventoryIdAndAccountId(ctx, inventoryId)
	if err != nil {
		return err
	}

	if accountId != inventory.AccountId {
		return errors.New("incorrect inventory")
	}

	if secuirity.Id != inventory.SecurityId {
		return errors.New("incorrect stock")
	}

	_, err = m.mysql.InsertTransaction(ctx, domain.Transaction{
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

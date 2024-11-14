package validator

import (
	"assetio/internal/domain"
	"errors"
)

func (v validation) StockBuy(request domain.ClientStockBuyRequest) error {
	if request.AccountId == 0 {
		return errors.New("invalid account id")
	}
	if request.UserId == 0 {
		return errors.New("invalid user id")
	}

	if request.StockId == 0 {
		return errors.New("invalid stock id")
	}

	if request.Quantity == 0 {
		return errors.New("invalid quantity")
	}

	if request.AveragePrice == 0 {
		return errors.New("invalid amount per quantity")
	}
	if request.FeeAmount == 0 {
		return errors.New("invalid fee amount")
	}

	return nil
}
func (v validation) StockSell(request domain.ClientStockSellRequest) error {
	if request.AccountId == 0 {
		return errors.New("invalid account id")
	}
	if request.UserId == 0 {
		return errors.New("invalid user id")
	}

	if request.StockId == 0 {
		return errors.New("invalid stock id")
	}

	if request.Quantity == 0 {
		return errors.New("invalid quantity")
	}

	if request.AveragePrice == 0 {
		return errors.New("invalid amount per quantity")
	}
	if request.FeeAmount == 0 {
		return errors.New("invalid fee amount")
	}

	return nil
}

func (v validation) StockSplit(request domain.ClientStockSplitRequest) error {
	if request.AccountId == 0 {
		return errors.New("invalid account id")
	}
	if request.UserId == 0 {
		return errors.New("invalid user id")
	}

	if request.StockId == 0 {
		return errors.New("invalid stock id")
	}

	if request.Quantity == 0 {
		return errors.New("invalid quantity")
	}

	if request.FeeAmount == 0 {
		return errors.New("invalid fee amount")
	}

	return nil
}

func (v validation) StockDividendAdd(request domain.ClientStockDividendAddRequest) error {
	if request.AccountId == 0 {
		return errors.New("invalid account id")
	}
	if request.UserId == 0 {
		return errors.New("invalid user id")
	}

	if request.StockId == 0 {
		return errors.New("invalid stock id")
	}

	if request.Quantity == 0 {
		return errors.New("invalid quantity")
	}

	if request.AmountPerQuantity == 0 {
		return errors.New("invalid amount per quantity")
	}

	return nil
}
func (v validation) StockSummary(request domain.ClientStockSummaryRequest) error {
	if request.AccountId == 0 {
		return errors.New("invalid account id")
	}

	if request.UserId == 0 {
		return errors.New("invalid user id")
	}

	return nil
}
func (v validation) StockInventories(request domain.ClientStockInventoriesRequest) error {

	if request.AccountId == 0 {
		return errors.New("invalid account id")
	}

	if request.UserId == 0 {
		return errors.New("invalid user id")
	}
	if request.StockId == 0 {
		return errors.New("invalid stock id")
	}

	return nil
}
func (v validation) StockInventoryLedgers(request domain.ClientStockInventoryLedgersRequest) error {
	if request.AccountId == 0 {
		return errors.New("invalid account id")
	}

	if request.UserId == 0 {
		return errors.New("invalid user id")
	}
	if request.InventoryId == 0 {
		return errors.New("invalid inventory id")
	}

	return nil
}

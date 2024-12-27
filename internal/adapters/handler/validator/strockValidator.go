package validator

import (
	"assetio/internal/domain"
	"errors"
)

// StockBuy validates the fields in the ClientStockBuyRequest object before proceeding with a stock purchase.
// It checks if the required fields (AccountId, UserId, StockId, Quantity, AveragePrice) are valid (non-zero).
func (v validation) StockBuy(request domain.ClientStockBuyRequest) error {
	if request.AccountId == 0 {
		return errors.New("invalid account id") // AccountId must be non-zero
	}
	if request.UserId == 0 {
		return errors.New("invalid user id") // UserId must be non-zero
	}

	if request.StockId == 0 {
		return errors.New("invalid stock id") // StockId must be non-zero
	}

	if request.Quantity == 0 {
		return errors.New("invalid quantity") // Quantity must be greater than 0
	}

	if request.AveragePrice == 0 {
		return errors.New("invalid amount per quantity") // AveragePrice must be greater than 0
	}

	return nil // Return nil if all validations pass
}

// StockSell validates the fields in the ClientStockSellRequest object before proceeding with a stock sale.
// It checks if the required fields (AccountId, UserId, StockId, Quantity, AveragePrice) are valid (non-zero).
func (v validation) StockSell(request domain.ClientStockSellRequest) error {
	if request.AccountId == 0 {
		return errors.New("invalid account id") // AccountId must be non-zero
	}
	if request.UserId == 0 {
		return errors.New("invalid user id") // UserId must be non-zero
	}

	if request.StockId == 0 {
		return errors.New("invalid stock id") // StockId must be non-zero
	}

	if request.Quantity == 0 {
		return errors.New("invalid quantity") // Quantity must be greater than 0
	}

	if request.AveragePrice == 0 {
		return errors.New("invalid amount per quantity") // AveragePrice must be greater than 0
	}

	return nil // Return nil if all validations pass
}

// StockSplit validates the fields in the ClientStockSplitRequest object before proceeding with a stock split.
// It checks if the required fields (AccountId, UserId, StockId, Quantity) are valid (non-zero).
func (v validation) StockSplit(request domain.ClientStockSplitRequest) error {
	if request.AccountId == 0 {
		return errors.New("invalid account id") // AccountId must be non-zero
	}
	if request.UserId == 0 {
		return errors.New("invalid user id") // UserId must be non-zero
	}

	if request.StockId == 0 {
		return errors.New("invalid stock id") // StockId must be non-zero
	}

	if request.Quantity == 0 {
		return errors.New("invalid quantity") // Quantity must be greater than 0
	}

	return nil // Return nil if all validations pass
}

// StockDividendAdd validates the fields in the ClientStockDividendAddRequest object before adding a stock dividend.
// It checks if the required fields (AccountId, UserId, StockId, Quantity, AmountPerQuantity) are valid (non-zero).
func (v validation) StockDividendAdd(request domain.ClientStockDividendAddRequest) error {
	if request.AccountId == 0 {
		return errors.New("invalid account id") // AccountId must be non-zero
	}
	if request.UserId == 0 {
		return errors.New("invalid user id") // UserId must be non-zero
	}

	if request.StockId == 0 {
		return errors.New("invalid stock id") // StockId must be non-zero
	}

	if request.AmountPerQuantity == 0 {
		return errors.New("invalid amount per quantity") // AmountPerQuantity must be greater than 0
	}

	return nil // Return nil if all validations pass
}

// StockSummary validates the fields in the ClientStockSummaryRequest object before fetching the stock summary.
// It checks if the required fields (AccountId, UserId) are valid (non-zero).
func (v validation) StockSummary(request domain.ClientStockSummaryRequest) error {
	if request.AccountId == 0 {
		return errors.New("invalid account id") // AccountId must be non-zero
	}

	if request.UserId == 0 {
		return errors.New("invalid user id") // UserId must be non-zero
	}

	return nil // Return nil if all validations pass
}

// StockInventories validates the fields in the ClientStockInventoriesRequest object before fetching the stock inventories.
// It checks if the required fields (AccountId, UserId, StockId) are valid (non-zero).
func (v validation) StockInventories(request domain.ClientStockInventoriesRequest) error {

	if request.AccountId == 0 {
		return errors.New("invalid account id") // AccountId must be non-zero
	}

	if request.UserId == 0 {
		return errors.New("invalid user id") // UserId must be non-zero
	}
	if request.StockId == 0 {
		return errors.New("invalid stock id") // StockId must be non-zero
	}

	return nil // Return nil if all validations pass
}

// StockInventoryLedgers validates the fields in the ClientStockInventoryLedgersRequest object before fetching stock inventory ledgers.
// It checks if the required fields (AccountId, UserId, InventoryId) are valid (non-zero).
func (v validation) StockInventoryLedgers(request domain.ClientStockInventoryLedgersRequest) error {
	if request.AccountId == 0 {
		return errors.New("invalid account id") // AccountId must be non-zero
	}

	if request.UserId == 0 {
		return errors.New("invalid user id") // UserId must be non-zero
	}
	if request.InventoryId == 0 {
		return errors.New("invalid inventory id") // InventoryId must be non-zero
	}

	return nil // Return nil if all validations pass
}

// StockDividends validates the fields in the ClientStockDividendsRequest object before fetching stock dividend data.
// It checks if the required fields (AccountId, UserId, StockId) are valid (non-zero).
func (v validation) StockDividends(request domain.ClientStockDividendsRequest) error {
	if request.AccountId == 0 {
		return errors.New("invalid account id") // AccountId must be non-zero
	}

	if request.UserId == 0 {
		return errors.New("invalid user id") // UserId must be non-zero
	}
	if request.StockId == 0 {
		return errors.New("invalid stock id") // InventoryId must be non-zero
	}

	return nil // Return nil if all validations pass
}

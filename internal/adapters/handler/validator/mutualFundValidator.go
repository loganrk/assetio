package validator

import (
	"assetio/internal/domain"
	"errors"
)

// MutualFundBuy validates the fields in the ClientMutualFundBuyRequest object before processing a mutual fund buy.
// It checks if the required fields (AccountId, UserId, MutualFundId, Quantity, AveragePrice) are valid (non-zero or non-empty).
func (v validation) MutualFundBuy(request domain.ClientMutualFundBuyRequest) error {
	if request.AccountId == 0 {
		return errors.New("invalid account id") // AccountId must be non-zero
	}
	if request.UserId == 0 {
		return errors.New("invalid user id") // UserId must be non-zero
	}

	if request.MutualFundId == 0 {
		return errors.New("invalid mutual fund id") // MutualFundId must be non-zero
	}

	if request.Quantity == 0 {
		return errors.New("invalid quantity") // Quantity must be non-zero
	}

	if request.AveragePrice == 0 {
		return errors.New("invalid amount per quantity") // AveragePrice must be non-zero
	}

	return nil // Return nil if all validations pass
}

// MutualFundSell validates the fields in the ClientMutualFundSellRequest object before processing a mutual fund sell.
// It checks if the required fields (AccountId, UserId, MutualFundId, Quantity, AveragePrice) are valid (non-zero or non-empty).
func (v validation) MutualFundSell(request domain.ClientMutualFundSellRequest) error {

	if request.AccountId == 0 {
		return errors.New("invalid account id") // AccountId must be non-zero
	}
	if request.UserId == 0 {
		return errors.New("invalid user id") // UserId must be non-zero
	}

	if request.MutualFundId == 0 {
		return errors.New("invalid mutual fund id") // MutualFundId must be non-zero
	}

	if request.Quantity == 0 {
		return errors.New("invalid quantity") // Quantity must be non-zero
	}

	if request.AveragePrice == 0 {
		return errors.New("invalid amount per quantity") // AveragePrice must be non-zero
	}

	return nil // Return nil if all validations pass
}

// MutualFundAdd validates the fields in the ClientMutualFundAddRequest object before adding mutual fund data.
// It checks if the required fields (AccountId, UserId, MutualFundId, InventoryId, Quantity, AveragePrice) are valid (non-zero or non-empty).
func (v validation) MutualFundAdd(request domain.ClientMutualFundAddRequest) error {

	if request.AccountId == 0 {
		return errors.New("invalid account id") // AccountId must be non-zero
	}
	if request.UserId == 0 {
		return errors.New("invalid user id") // UserId must be non-zero
	}

	if request.MutualFundId == 0 {
		return errors.New("invalid mutual fund id") // MutualFundId must be non-zero
	}

	if request.InventoryId == 0 {
		return errors.New("invalid inventory id") // InventoryId must be non-zero
	}

	if request.Quantity == 0 {
		return errors.New("invalid quantity") // Quantity must be non-zero
	}

	if request.AveragePrice == 0 {
		return errors.New("invalid amount per quantity") // AveragePrice must be non-zero
	}

	return nil // Return nil if all validations pass
}

// MutualFundSummary validates the fields in the ClientMutualFundSummaryRequest object before summarizing mutual fund data.
// It checks if the required fields (AccountId, UserId) are valid (non-zero).
func (v validation) MutualFundSummary(request domain.ClientMutualFundSummaryRequest) error {
	if request.AccountId == 0 {
		return errors.New("invalid account id") // AccountId must be non-zero
	}

	if request.UserId == 0 {
		return errors.New("invalid user id") // UserId must be non-zero
	}

	return nil // Return nil if all validations pass
}

// MutualFundInventory validates the fields in the ClientMutualFundInventoryRequest object before retrieving mutual fund inventory data.
// It checks if the required fields (AccountId, UserId, MutualFundId) are valid (non-zero).
func (v validation) MutualFundInventory(request domain.ClientMutualFundInventoryRequest) error {
	if request.AccountId == 0 {
		return errors.New("invalid account id") // AccountId must be non-zero
	}

	if request.UserId == 0 {
		return errors.New("invalid user id") // UserId must be non-zero
	}
	if request.MutualFundId == 0 {
		return errors.New("invalid mutual fund id") // MutualFundId must be non-zero
	}

	return nil // Return nil if all validations pass
}

// MutualFundInventoryLedgers validates the fields in the ClientMutualFundInventoryLedgersRequest object before retrieving mutual fund inventory ledger data.
// It checks if the required fields (AccountId, UserId, InventoryId) are valid (non-zero).
func (v validation) MutualFundInventoryLedgers(request domain.ClientMutualFundInventoryLedgersRequest) error {
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

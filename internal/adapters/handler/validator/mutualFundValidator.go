package validator

import (
	"assetio/internal/domain"
	"errors"
)

func (v validation) MutualFundBuy(request domain.ClientMutualFundBuyRequest) error {
	if request.AccountId == 0 {
		return errors.New("invalid account id")
	}
	if request.UserId == 0 {
		return errors.New("invalid user id")
	}

	if request.MutualFundId == 0 {
		return errors.New("invalid mutual fund id")
	}

	if request.Quantity == 0 {
		return errors.New("invalid quantity")
	}

	if request.AveragePrice == 0 {
		return errors.New("invalid amount per quantity")
	}

	return nil

}

func (v validation) MutualFundSell(request domain.ClientMutualFundSellRequest) error {

	if request.AccountId == 0 {
		return errors.New("invalid account id")
	}
	if request.UserId == 0 {
		return errors.New("invalid user id")
	}

	if request.MutualFundId == 0 {
		return errors.New("invalid mutual fund id")
	}

	if request.Quantity == 0 {
		return errors.New("invalid quantity")
	}

	if request.AveragePrice == 0 {
		return errors.New("invalid amount per quantity")
	}

	return nil

}

func (v validation) MutualFundAdd(request domain.ClientMutualFundAddRequest) error {

	if request.AccountId == 0 {
		return errors.New("invalid account id")
	}
	if request.UserId == 0 {
		return errors.New("invalid user id")
	}

	if request.MutualFundId == 0 {
		return errors.New("invalid mutual fund id")
	}

	if request.InventoryId == 0 {
		return errors.New("invalid inventory id")
	}

	if request.Quantity == 0 {
		return errors.New("invalid quantity")
	}

	if request.AveragePrice == 0 {
		return errors.New("invalid amount per quantity")
	}

	return nil
}

func (v validation) MutualFundSummary(request domain.ClientMutualFundSummaryRequest) error {
	if request.AccountId == 0 {
		return errors.New("invalid account id")
	}

	if request.UserId == 0 {
		return errors.New("invalid user id")
	}

	return nil
}
func (v validation) MutualFundInventory(request domain.ClientMutualFundInventoryRequest) error {
	if request.AccountId == 0 {
		return errors.New("invalid account id")
	}

	if request.UserId == 0 {
		return errors.New("invalid user id")
	}
	if request.MutualFundId == 0 {
		return errors.New("invalid mutual fund id")
	}

	return nil
}
func (v validation) MutualFundInventoryLedgers(request domain.ClientMutualFundInventoryLedgersRequest) error {
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

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

	if request.AmountPerQuantity == 0 {
		return errors.New("invalid amount per quantity")
	}
	if request.FeeAmount == 0 {
		return errors.New("invalid fee amount")
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

	if request.AmountPerQuantity == 0 {
		return errors.New("invalid amount per quantity")
	}

	if request.FeeAmount == 0 {
		return errors.New("invalid fee amount")
	}

	return nil

}

package request

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"
	"strings"
)

func NewMutualFundBuy(r *http.Request) (*mutualFundBuy, error) {
	var s mutualFundBuy
	if r.Method == http.MethodPost {
		decoder := json.NewDecoder(r.Body)
		err := decoder.Decode(&s)
		if err != nil {
			return &s, err
		}
	} else {
		accountId, _ := strconv.Atoi(r.URL.Query().Get("account_id"))
		s.AccountId = accountId

		mutualFundId, _ := strconv.Atoi(r.URL.Query().Get("mf_id"))
		s.MutualFundId = mutualFundId

		inventoryId, _ := strconv.Atoi(r.URL.Query().Get("inventory_id"))
		s.InventoryId = inventoryId

		quantity, _ := strconv.Atoi(r.URL.Query().Get("quantity"))
		s.Quantity = quantity

		amountPerQuantity, _ := strconv.ParseFloat(strings.TrimSpace(r.URL.Query().Get("amount_per_quantity")), 64)
		s.AmountPerQuantity = amountPerQuantity

		taxAmount, _ := strconv.ParseFloat(strings.TrimSpace(r.URL.Query().Get("tax_amount")), 64)
		s.TaxAmount = taxAmount
	}

	userid, _ := strconv.Atoi(r.URL.Query().Get("uid"))
	s.UserId = userid
	return &s, nil
}

func (s *mutualFundBuy) Validate() error {

	if s.AccountId == 0 {
		return errors.New("invalid account id")
	}
	if s.UserId == 0 {
		return errors.New("invalid user id")
	}

	if s.MutualFundId == 0 {
		return errors.New("invalid mutual fund id")
	}

	if s.Quantity == 0 {
		return errors.New("invalid quantity")
	}

	if s.AmountPerQuantity == 0 {
		return errors.New("invalid amount per quantity")
	}
	if s.TaxAmount == 0 {
		return errors.New("invalid tax amount")
	}

	return nil
}

func NewMutualFundSell(r *http.Request) (*mutualFundSell, error) {
	var s mutualFundSell
	if r.Method == http.MethodPost {
		decoder := json.NewDecoder(r.Body)
		err := decoder.Decode(&s)
		if err != nil {
			return &s, err
		}
	} else {
		accountId, _ := strconv.Atoi(r.URL.Query().Get("account_id"))
		s.AccountId = accountId

		mutualFundId, _ := strconv.Atoi(r.URL.Query().Get("mf_id"))
		s.MutualFundId = mutualFundId

		inventoryId, _ := strconv.Atoi(r.URL.Query().Get("inventory_id"))
		s.InventoryId = inventoryId

		quantity, _ := strconv.Atoi(r.URL.Query().Get("quantity"))
		s.Quantity = quantity

		amountPerQuantity, _ := strconv.ParseFloat(strings.TrimSpace(r.URL.Query().Get("amount_per_quantity")), 64)
		s.AmountPerQuantity = amountPerQuantity

		taxAmount, _ := strconv.ParseFloat(strings.TrimSpace(r.URL.Query().Get("tax_amount")), 64)
		s.TaxAmount = taxAmount
	}

	userid, _ := strconv.Atoi(r.URL.Query().Get("uid"))
	s.UserId = userid
	return &s, nil
}

func (s *mutualFundSell) Validate() error {

	if s.AccountId == 0 {
		return errors.New("invalid account id")
	}
	if s.UserId == 0 {
		return errors.New("invalid user id")
	}

	if s.MutualFundId == 0 {
		return errors.New("invalid mutual fund id")
	}

	if s.Quantity == 0 {
		return errors.New("invalid quantity")
	}

	if s.AmountPerQuantity == 0 {
		return errors.New("invalid amount per quantity")
	}
	if s.TaxAmount == 0 {
		return errors.New("invalid tax amount")
	}

	return nil
}

package request

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"
	"strings"
)

func NewStockBuy(r *http.Request) (*stockBuy, error) {
	var s stockBuy
	if r.Method == http.MethodPost {
		decoder := json.NewDecoder(r.Body)
		err := decoder.Decode(&s)
		if err != nil {
			return &s, err
		}
	} else {
		accountId, _ := strconv.Atoi(r.URL.Query().Get("account_id"))
		s.AccountId = accountId

		stockId, _ := strconv.Atoi(r.URL.Query().Get("stock_id"))
		s.StockId = stockId

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

func (s *stockBuy) Validate() error {

	if s.AccountId == 0 {
		return errors.New("invalid account id")
	}
	if s.UserId == 0 {
		return errors.New("invalid user id")
	}

	if s.StockId == 0 {
		return errors.New("invalid stock id")
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

func NewStockSell(r *http.Request) (*stockSell, error) {
	var s stockSell
	if r.Method == http.MethodPost {
		decoder := json.NewDecoder(r.Body)
		err := decoder.Decode(&s)
		if err != nil {
			return &s, err
		}
	} else {
		accountId, _ := strconv.Atoi(r.URL.Query().Get("account_id"))
		s.AccountId = accountId

		stockId, _ := strconv.Atoi(r.URL.Query().Get("stock_id"))
		s.StockId = stockId

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

func (s *stockSell) Validate() error {

	if s.AccountId == 0 {
		return errors.New("invalid account id")
	}
	if s.UserId == 0 {
		return errors.New("invalid user id")
	}

	if s.StockId == 0 {
		return errors.New("invalid stock id")
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

func NewStockDividendAdd(r *http.Request) (*stockDividendAdd, error) {
	var s stockDividendAdd
	if r.Method == http.MethodPost {
		decoder := json.NewDecoder(r.Body)
		err := decoder.Decode(&s)
		if err != nil {
			return &s, err
		}
	} else {
		accountId, _ := strconv.Atoi(r.URL.Query().Get("account_id"))
		s.AccountId = accountId

		stockId, _ := strconv.Atoi(r.URL.Query().Get("stock_id"))
		s.StockId = stockId

		inventoryId, _ := strconv.Atoi(r.URL.Query().Get("inventory_id"))
		s.InventoryId = inventoryId

		quantity, _ := strconv.Atoi(r.URL.Query().Get("quantity"))
		s.Quantity = quantity

		amountPerQuantity, _ := strconv.ParseFloat(strings.TrimSpace(r.URL.Query().Get("amount_per_quantity")), 64)
		s.AmountPerQuantity = amountPerQuantity

	}

	userid, _ := strconv.Atoi(r.URL.Query().Get("uid"))
	s.UserId = userid
	return &s, nil
}

func (s *stockDividendAdd) Validate() error {

	if s.AccountId == 0 {
		return errors.New("invalid account id")
	}
	if s.UserId == 0 {
		return errors.New("invalid user id")
	}

	if s.StockId == 0 {
		return errors.New("invalid stock id")
	}

	if s.Quantity == 0 {
		return errors.New("invalid quantity")
	}

	if s.AmountPerQuantity == 0 {
		return errors.New("invalid amount per quantity")
	}

	return nil
}

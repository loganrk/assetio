package v1

import (
	"assetio/internal/adapters/handler/http/v1/request"
	"assetio/internal/adapters/handler/http/v1/response"
	"assetio/internal/port"
	"context"
	"net/http"
)

func (h *handler) StockBuy(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	res := response.New()

	req, err := request.NewStockBuy(r)
	if err != nil {
		res.SetStatus(http.StatusBadRequest)
		res.SetError(ERROR_CODE_REQUEST_INVALID, "invalid request parameters")
		res.Send(w)
		return
	}

	err = req.Validate()
	if err != nil {
		res.SetStatus(http.StatusBadRequest)
		res.SetError(ERROR_CODE_REQUEST_INVALID, err.Error())
		res.Send(w)
		return
	}

	err = h.usecases.Stock.BuyStock(ctx, req.UserId, req.AccountId, req.InventoryId, req.StockId, req.Quantity, req.AmountPerQuantity, req.TaxAmount)

	if err != nil {
		res.SetStatus(http.StatusBadRequest)
		res.SetError(ERROR_CODE_REQUEST_INVALID, err.Error())
		res.Send(w)
		return
	}

	resData := port.SecurityCreateClientResponse{
		Message: "stock buy successfully",
	}
	res.SetData(resData)
	res.Send(w)
}

func (h *handler) StockSell(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	res := response.New()

	req, err := request.NewStockSell(r)
	if err != nil {
		res.SetStatus(http.StatusBadRequest)
		res.SetError(ERROR_CODE_REQUEST_INVALID, "invalid request parameters")
		res.Send(w)
		return
	}

	err = req.Validate()
	if err != nil {
		res.SetStatus(http.StatusBadRequest)
		res.SetError(ERROR_CODE_REQUEST_INVALID, err.Error())
		res.Send(w)
		return
	}

	err = h.usecases.Stock.SellStock(ctx, req.UserId, req.AccountId, req.InventoryId, req.StockId, req.Quantity, req.AmountPerQuantity, req.TaxAmount)

	if err != nil {
		res.SetStatus(http.StatusBadRequest)
		res.SetError(ERROR_CODE_REQUEST_INVALID, err.Error())
		res.Send(w)
		return
	}

	resData := port.SecurityCreateClientResponse{
		Message: "stock sell successfully",
	}
	res.SetData(resData)
	res.Send(w)
}

func (h *handler) StockDividendAdd(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	res := response.New()

	req, err := request.NewStockDividendAdd(r)
	if err != nil {
		res.SetStatus(http.StatusBadRequest)
		res.SetError(ERROR_CODE_REQUEST_INVALID, "invalid request parameters")
		res.Send(w)
		return
	}

	err = req.Validate()
	if err != nil {
		res.SetStatus(http.StatusBadRequest)
		res.SetError(ERROR_CODE_REQUEST_INVALID, err.Error())
		res.Send(w)
		return
	}

	err = h.usecases.Stock.StockDividendAdd(ctx, req.UserId, req.AccountId, req.InventoryId, req.StockId, req.Quantity, req.AmountPerQuantity)

	if err != nil {
		res.SetStatus(http.StatusBadRequest)
		res.SetError(ERROR_CODE_REQUEST_INVALID, err.Error())
		res.Send(w)
		return
	}

	resData := port.SecurityCreateClientResponse{
		Message: "stock dividend successfully",
	}
	res.SetData(resData)
	res.Send(w)
}

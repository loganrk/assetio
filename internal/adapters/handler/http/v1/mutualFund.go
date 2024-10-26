package v1

import (
	"assetio/internal/adapters/handler/http/v1/request"
	"assetio/internal/adapters/handler/http/v1/response"
	"assetio/internal/port"
	"context"
	"net/http"
)

func (h *handler) MutualFundBuy(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	res := response.New()

	req, err := request.NewMutualFundBuy(r)
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

	err = h.usecases.MutualFund.BuyMutualFund(ctx, req.UserId, req.AccountId, req.InventoryId, req.MutualFundId, req.Quantity, req.AmountPerQuantity, req.TaxAmount)

	if err != nil {
		res.SetStatus(http.StatusBadRequest)
		res.SetError(ERROR_CODE_REQUEST_INVALID, err.Error())
		res.Send(w)
		return
	}

	resData := port.SecurityCreateClientResponse{
		Message: "mutual fund buy successfully",
	}
	res.SetData(resData)
	res.Send(w)
}

func (h *handler) MutualFundSell(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	res := response.New()

	req, err := request.NewMutualFundSell(r)
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

	err = h.usecases.MutualFund.SellMutualFund(ctx, req.UserId, req.AccountId, req.InventoryId, req.MutualFundId, req.Quantity, req.AmountPerQuantity, req.TaxAmount)

	if err != nil {
		res.SetStatus(http.StatusBadRequest)
		res.SetError(ERROR_CODE_REQUEST_INVALID, err.Error())
		res.Send(w)
		return
	}

	resData := port.SecurityCreateClientResponse{
		Message: "mutual fund sell successfully",
	}
	res.SetData(resData)
	res.Send(w)
}
func (h *handler) MutualFundSummary(w http.ResponseWriter, r *http.Request)     {}
func (h *handler) MutualFundInventory(w http.ResponseWriter, r *http.Request)   {}
func (h *handler) MutualFundTransaction(w http.ResponseWriter, r *http.Request) {}

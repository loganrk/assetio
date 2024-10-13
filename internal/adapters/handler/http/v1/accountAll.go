package v1

import (
	"assetio/internal/adapters/handler/http/v1/request"
	"assetio/internal/adapters/handler/http/v1/response"
	"assetio/internal/port"
	"context"

	"net/http"
)

func (h *handler) AccountAll(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	res := response.New()

	req, err := request.NewAccountAll(r)
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

	accounts, err := h.usecases.Account.GetAccounts(ctx, req.GetUserId())

	if err != nil {
		res.SetStatus(http.StatusInternalServerError)
		res.SetError(ERROR_CODE_INTERNAL_SERVER, "internal server error")
		res.Send(w)
		return
	}

	if len(accounts) == 0 {
		res.SetData(nil)
		res.Send(w)
		return

	}

	var resData []port.AccountAllClientResponse
	for _, account := range accounts {
		resData = append(resData, port.AccountAllClientResponse{
			Id:     account.Id,
			Name:   account.Name,
			Status: h.usecases.Account.GetStatusString(account.Status),
		})
	}

	res.SetData(resData)
	res.Send(w)

}

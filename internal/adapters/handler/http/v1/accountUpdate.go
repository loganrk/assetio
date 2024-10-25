package v1

import (
	"assetio/internal/adapters/handler/http/v1/request"
	"assetio/internal/adapters/handler/http/v1/response"
	"assetio/internal/port"
	"context"
	"net/http"
)

func (h *handler) AccountUpdate(w http.ResponseWriter, r *http.Request) {

	ctx := context.Background()
	res := response.New()

	req, err := request.NewAccountUpdate(r)
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

	err = h.usecases.Account.UpdateAccount(ctx, req.AccountId, req.UserId, req.Name)
	if err != nil {
		res.SetStatus(http.StatusInternalServerError)
		res.SetError(ERROR_CODE_INTERNAL_SERVER, "internal server error")
		res.Send(w)
		return
	}

	resData := port.AccountUpdateClientResponse{
		Message: "account updated successfully",
	}
	res.SetData(resData)
	res.Send(w)
}

package v1

import (
	"assetio/internal/adapters/handler/http/v1/request"
	"assetio/internal/adapters/handler/http/v1/response"
	"assetio/internal/port"
	"context"
	"net/http"
)

func (h *handler) SecurityAll(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	res := response.New()

	req, err := request.NewSecurityAll(r)
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
	securityType := h.usecases.Security.GetType(req.Type)
	securityExchange := h.usecases.Security.GetExchange(req.Exchange)

	if securityType == 0 {
		res.SetStatus(http.StatusBadRequest)
		res.SetError(ERROR_CODE_REQUEST_INVALID, "invalid type")
		res.Send(w)
		return
	}

	if securityExchange == 0 {
		res.SetStatus(http.StatusBadRequest)
		res.SetError(ERROR_CODE_REQUEST_INVALID, "invalid exchange")
		res.Send(w)
		return
	}

	securities, err := h.usecases.Security.GetSecurities(ctx, securityType, securityExchange)
	if err != nil {
		res.SetStatus(http.StatusInternalServerError)
		res.SetError(ERROR_CODE_INTERNAL_SERVER, "internal server error")
		res.Send(w)
		return
	}

	if len(securities) == 0 {
		res.SetData(nil)
		res.Send(w)
		return

	}

	var resData []port.SecurityAllClientResponse
	for _, security := range securities {
		resData = append(resData, port.SecurityAllClientResponse{
			Id:       security.Id,
			Type:     h.usecases.Security.GetTypeString(security.Type),
			Exchange: h.usecases.Security.GetExchangeString(security.Exchange),
			Symbol:   security.Symbol,
			Name:     security.Name,
		})
	}

	res.SetData(resData)
	res.Send(w)
}

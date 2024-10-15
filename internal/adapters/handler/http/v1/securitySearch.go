package v1

import (
	"assetio/internal/adapters/handler/http/v1/request"
	"assetio/internal/adapters/handler/http/v1/response"
	"assetio/internal/port"
	"context"
	"net/http"
)

func (h *handler) SecuritySearch(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	res := response.New()

	req, err := request.NewSecuritySearch(r)
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
	securityType := h.usecases.Security.GetType(req.GetType())
	securityExchange := h.usecases.Security.GetExchange(req.GetExchange())

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

	securities, err := h.usecases.Security.SearchSecurities(ctx, securityType, securityExchange, req.GetSearch())
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

	var resData []port.SecuritySearchClientResponse
	for _, security := range securities {
		resData = append(resData, port.SecuritySearchClientResponse{
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

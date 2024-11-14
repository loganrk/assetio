package v1

import (
	"assetio/internal/adapters/handler/response"
	"assetio/internal/domain"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/schema"
)

func (h *handler) MutualFundBuy(w http.ResponseWriter, r *http.Request) {

	var request domain.ClientMutualFundBuyRequest
	res := response.New()

	if r.Method == http.MethodPost {
		decoder := json.NewDecoder(r.Body)
		err := decoder.Decode(&request)
		if err != nil {
			res.SetStatus(http.StatusBadRequest)
			res.SetError(ERROR_CODE_REQUEST_INVALID, err.Error())
			res.Send(w)
			return
		}
	} else {
		var decoder = schema.NewDecoder()
		decoder.IgnoreUnknownKeys(true)
		if err := decoder.Decode(&request, r.URL.Query()); err != nil {
			res.SetStatus(http.StatusBadRequest)
			res.SetError(ERROR_CODE_REQUEST_INVALID, err.Error())
			res.Send(w)
			return
		}
	}

	userid, _ := strconv.Atoi(r.URL.Query().Get("uid"))
	request.UserId = userid

	err := h.validator.MutualFundBuy(request)
	if err != nil {
		res.SetStatus(http.StatusBadRequest)
		res.SetError(ERROR_CODE_REQUEST_INVALID, err.Error())
		res.Send(w)
		return
	}

	resData := h.usecases.MutualFund.MutualFundBuy(request)
	resData.Send(w)

}
func (h *handler) MutualFundAdd(w http.ResponseWriter, r *http.Request) {}
func (h *handler) MutualFundSell(w http.ResponseWriter, r *http.Request) {
	var request domain.ClientMutualFundSellRequest
	res := response.New()

	if r.Method == http.MethodPost {
		decoder := json.NewDecoder(r.Body)
		err := decoder.Decode(&request)
		if err != nil {
			res.SetStatus(http.StatusBadRequest)
			res.SetError(ERROR_CODE_REQUEST_INVALID, err.Error())
			res.Send(w)
			return
		}
	} else {
		var decoder = schema.NewDecoder()
		decoder.IgnoreUnknownKeys(true)
		if err := decoder.Decode(&request, r.URL.Query()); err != nil {
			res.SetStatus(http.StatusBadRequest)
			res.SetError(ERROR_CODE_REQUEST_INVALID, err.Error())
			res.Send(w)
			return
		}
	}

	userid, _ := strconv.Atoi(r.URL.Query().Get("uid"))
	request.UserId = userid

	err := h.validator.MutualFundSell(request)
	if err != nil {
		res.SetStatus(http.StatusBadRequest)
		res.SetError(ERROR_CODE_REQUEST_INVALID, err.Error())
		res.Send(w)
		return
	}

	resData := h.usecases.MutualFund.MutualFundSell(request)
	resData.Send(w)
}
func (h *handler) MutualFundSummary(w http.ResponseWriter, r *http.Request) {
	var request domain.ClientMutualFundSummaryRequest
	res := response.New()

	if r.Method == http.MethodPost {
		decoder := json.NewDecoder(r.Body)
		err := decoder.Decode(&request)
		if err != nil {
			res.SetStatus(http.StatusBadRequest)
			res.SetError(ERROR_CODE_REQUEST_INVALID, err.Error())
			res.Send(w)
			return
		}
	} else {
		var decoder = schema.NewDecoder()
		decoder.IgnoreUnknownKeys(true)
		if err := decoder.Decode(&request, r.URL.Query()); err != nil {
			res.SetStatus(http.StatusBadRequest)
			res.SetError(ERROR_CODE_REQUEST_INVALID, err.Error())
			res.Send(w)
			return
		}
	}

	userid, _ := strconv.Atoi(r.URL.Query().Get("uid"))
	request.UserId = userid

	err := h.validator.MutualFundSummary(request)
	if err != nil {
		res.SetStatus(http.StatusBadRequest)
		res.SetError(ERROR_CODE_REQUEST_INVALID, err.Error())
		res.Send(w)
		return
	}

	resData := h.usecases.MutualFund.MutualFundSummary(request)
	resData.Send(w)
}
func (h *handler) MutualFundInventory(w http.ResponseWriter, r *http.Request) {
	var request domain.ClientMutualFundInventoryRequest
	res := response.New()

	if r.Method == http.MethodPost {
		decoder := json.NewDecoder(r.Body)
		err := decoder.Decode(&request)
		if err != nil {
			res.SetStatus(http.StatusBadRequest)
			res.SetError(ERROR_CODE_REQUEST_INVALID, err.Error())
			res.Send(w)
			return
		}
	} else {
		var decoder = schema.NewDecoder()
		decoder.IgnoreUnknownKeys(true)
		if err := decoder.Decode(&request, r.URL.Query()); err != nil {
			res.SetStatus(http.StatusBadRequest)
			res.SetError(ERROR_CODE_REQUEST_INVALID, err.Error())
			res.Send(w)
			return
		}
	}

	userid, _ := strconv.Atoi(r.URL.Query().Get("uid"))
	request.UserId = userid

	err := h.validator.MutualFundInventory(request)
	if err != nil {
		res.SetStatus(http.StatusBadRequest)
		res.SetError(ERROR_CODE_REQUEST_INVALID, err.Error())
		res.Send(w)
		return
	}

	resData := h.usecases.MutualFund.MutualFundInventory(request)
	resData.Send(w)
}
func (h *handler) MutualFundTransaction(w http.ResponseWriter, r *http.Request) {
	var request domain.ClientMutualFundInventoryLedgersRequest
	res := response.New()

	if r.Method == http.MethodPost {
		decoder := json.NewDecoder(r.Body)
		err := decoder.Decode(&request)
		if err != nil {
			res.SetStatus(http.StatusBadRequest)
			res.SetError(ERROR_CODE_REQUEST_INVALID, err.Error())
			res.Send(w)
			return
		}
	} else {
		var decoder = schema.NewDecoder()
		decoder.IgnoreUnknownKeys(true)
		if err := decoder.Decode(&request, r.URL.Query()); err != nil {
			res.SetStatus(http.StatusBadRequest)
			res.SetError(ERROR_CODE_REQUEST_INVALID, err.Error())
			res.Send(w)
			return
		}
	}

	userid, _ := strconv.Atoi(r.URL.Query().Get("uid"))
	request.UserId = userid

	err := h.validator.MutualFundInventoryLedgers(request)
	if err != nil {
		res.SetStatus(http.StatusBadRequest)
		res.SetError(ERROR_CODE_REQUEST_INVALID, err.Error())
		res.Send(w)
		return
	}

	resData := h.usecases.MutualFund.MutualFundInventoryLedgers(request)
	resData.Send(w)

}

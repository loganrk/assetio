package v1

import (
	"assetio/internal/adapters/handler/response"
	"assetio/internal/domain"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/schema"
)

func (h *handler) StockBuy(w http.ResponseWriter, r *http.Request) {

	var request domain.ClientStockBuyRequest
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

	err := h.validator.StockBuy(request)
	if err != nil {
		res.SetStatus(http.StatusBadRequest)
		res.SetError(ERROR_CODE_REQUEST_INVALID, err.Error())
		res.Send(w)
		return
	}

	resData := h.usecases.Stock.StockBuy(request)
	resData.Send(w)

}

func (h *handler) StockSell(w http.ResponseWriter, r *http.Request) {
	var request domain.ClientStockSellRequest
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

	err := h.validator.StockSell(request)
	if err != nil {
		res.SetStatus(http.StatusBadRequest)
		res.SetError(ERROR_CODE_REQUEST_INVALID, err.Error())
		res.Send(w)
		return
	}

	resData := h.usecases.Stock.StockSell(request)
	resData.Send(w)

}

func (h *handler) StockDividendAdd(w http.ResponseWriter, r *http.Request) {
	var request domain.ClientStockDividendAddRequest
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

	err := h.validator.StockDividendAdd(request)
	if err != nil {
		res.SetStatus(http.StatusBadRequest)
		res.SetError(ERROR_CODE_REQUEST_INVALID, err.Error())
		res.Send(w)
		return
	}

	resData := h.usecases.Stock.StockDividendAdd(request)
	resData.Send(w)

}
func (h *handler) StockSplit(w http.ResponseWriter, r *http.Request) {
	var request domain.ClientStockSplitRequest
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

	err := h.validator.StockSplit(request)
	if err != nil {
		res.SetStatus(http.StatusBadRequest)
		res.SetError(ERROR_CODE_REQUEST_INVALID, err.Error())
		res.Send(w)
		return
	}

	resData := h.usecases.Stock.StockSplit(request)
	resData.Send(w)

}

func (h *handler) StockSummary(w http.ResponseWriter, r *http.Request) {
	var request domain.ClientStockSummaryRequest
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

	err := h.validator.StockSummary(request)
	if err != nil {
		res.SetStatus(http.StatusBadRequest)
		res.SetError(ERROR_CODE_REQUEST_INVALID, err.Error())
		res.Send(w)
		return
	}

	resData := h.usecases.Stock.StockSummary(request)
	resData.Send(w)

}

func (h *handler) StockInventories(w http.ResponseWriter, r *http.Request) {
	var request domain.ClientStockInventoriesRequest
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

	err := h.validator.StockInventories(request)
	if err != nil {
		res.SetStatus(http.StatusBadRequest)
		res.SetError(ERROR_CODE_REQUEST_INVALID, err.Error())
		res.Send(w)
		return
	}

	resData := h.usecases.Stock.StockInventories(request)
	resData.Send(w)

}

func (h *handler) StockInventoryLedgers(w http.ResponseWriter, r *http.Request) {
	var request domain.ClientStockInventoryLedgersRequest
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

	err := h.validator.StockInventoryLedgers(request)
	if err != nil {
		res.SetStatus(http.StatusBadRequest)
		res.SetError(ERROR_CODE_REQUEST_INVALID, err.Error())
		res.Send(w)
		return
	}

	resData := h.usecases.Stock.StockInventoryLedgers(request)
	resData.Send(w)

}

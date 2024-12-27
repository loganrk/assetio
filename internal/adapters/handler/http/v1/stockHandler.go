package v1

import (
	"assetio/internal/adapters/handler/response"
	"assetio/internal/domain"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/schema"
)

// StockBuy handles the request to buy stocks for a user
func (h *handler) StockBuy(w http.ResponseWriter, r *http.Request) {
	var request domain.ClientStockBuyRequest
	res := response.New()

	// Decode request based on HTTP method (POST or GET)
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

	// Assign user ID from URL query to the request
	userid, _ := strconv.Atoi(r.URL.Query().Get("uid"))
	request.UserId = userid

	// Validate the buy request
	err := h.validator.StockBuy(request)
	if err != nil {
		res.SetStatus(http.StatusBadRequest)
		res.SetError(ERROR_CODE_REQUEST_INVALID, err.Error())
		res.Send(w)
		return
	}

	// Call the usecase to perform the stock buy action
	resData := h.usecases.Stock.StockBuy(request)
	resData.Send(w)
}

// StockSell handles the request to sell stocks for a user
func (h *handler) StockSell(w http.ResponseWriter, r *http.Request) {
	var request domain.ClientStockSellRequest
	res := response.New()

	// Decode request based on HTTP method (POST or GET)
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

	// Assign user ID from URL query to the request
	userid, _ := strconv.Atoi(r.URL.Query().Get("uid"))
	request.UserId = userid

	// Validate the sell request
	err := h.validator.StockSell(request)
	if err != nil {
		res.SetStatus(http.StatusBadRequest)
		res.SetError(ERROR_CODE_REQUEST_INVALID, err.Error())
		res.Send(w)
		return
	}

	// Call the usecase to perform the stock sell action
	resData := h.usecases.Stock.StockSell(request)
	resData.Send(w)
}

// StockDividendAdd handles the request to add a dividend for a stock
func (h *handler) StockDividendAdd(w http.ResponseWriter, r *http.Request) {
	var request domain.ClientStockDividendAddRequest
	res := response.New()

	// Decode request based on HTTP method (POST or GET)
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

	// Assign user ID from URL query to the request
	userid, _ := strconv.Atoi(r.URL.Query().Get("uid"))
	request.UserId = userid

	// Validate the dividend add request
	err := h.validator.StockDividendAdd(request)
	if err != nil {
		res.SetStatus(http.StatusBadRequest)
		res.SetError(ERROR_CODE_REQUEST_INVALID, err.Error())
		res.Send(w)
		return
	}

	// Call the usecase to perform the stock dividend add action
	resData := h.usecases.Stock.StockDividendAdd(request)
	resData.Send(w)
}

// StockSplit handles the request to perform a stock split
func (h *handler) StockSplit(w http.ResponseWriter, r *http.Request) {
	var request domain.ClientStockSplitRequest
	res := response.New()

	// Decode request based on HTTP method (POST or GET)
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

	// Assign user ID from URL query to the request
	userid, _ := strconv.Atoi(r.URL.Query().Get("uid"))
	request.UserId = userid

	// Validate the stock split request
	err := h.validator.StockSplit(request)
	if err != nil {
		res.SetStatus(http.StatusBadRequest)
		res.SetError(ERROR_CODE_REQUEST_INVALID, err.Error())
		res.Send(w)
		return
	}

	// Call the usecase to perform the stock split action
	resData := h.usecases.Stock.StockSplit(request)
	resData.Send(w)
}

// StockSummary handles the request to get a summary of a user's stock holdings
func (h *handler) StockSummary(w http.ResponseWriter, r *http.Request) {
	var request domain.ClientStockSummaryRequest
	res := response.New()

	// Decode request based on HTTP method (POST or GET)
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

	// Assign user ID from URL query to the request
	userid, _ := strconv.Atoi(r.URL.Query().Get("uid"))
	request.UserId = userid

	// Validate the stock summary request
	err := h.validator.StockSummary(request)
	if err != nil {
		res.SetStatus(http.StatusBadRequest)
		res.SetError(ERROR_CODE_REQUEST_INVALID, err.Error())
		res.Send(w)
		return
	}

	// Call the usecase to get the stock summary
	resData := h.usecases.Stock.StockSummary(request)
	resData.Send(w)
}

// StockInventories handles the request to get the list of stock inventories for a user
func (h *handler) StockInventories(w http.ResponseWriter, r *http.Request) {
	var request domain.ClientStockInventoriesRequest
	res := response.New()

	// Decode request based on HTTP method (POST or GET)
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

	// Assign user ID from URL query to the request
	userid, _ := strconv.Atoi(r.URL.Query().Get("uid"))
	request.UserId = userid

	// Validate the stock inventories request
	err := h.validator.StockInventories(request)
	if err != nil {
		res.SetStatus(http.StatusBadRequest)
		res.SetError(ERROR_CODE_REQUEST_INVALID, err.Error())
		res.Send(w)
		return
	}

	// Call the usecase to get the list of stock inventories
	resData := h.usecases.Stock.StockInventories(request)
	resData.Send(w)
}

// StockInventoryLedgers handles the request to get the inventory ledger of stocks for a user
func (h *handler) StockInventoryLedgers(w http.ResponseWriter, r *http.Request) {
	var request domain.ClientStockInventoryLedgersRequest
	res := response.New()

	// Decode request based on HTTP method (POST or GET)
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

	// Assign user ID from URL query to the request
	userid, _ := strconv.Atoi(r.URL.Query().Get("uid"))
	request.UserId = userid

	// Validate the inventory ledger request
	err := h.validator.StockInventoryLedgers(request)
	if err != nil {
		res.SetStatus(http.StatusBadRequest)
		res.SetError(ERROR_CODE_REQUEST_INVALID, err.Error())
		res.Send(w)
		return
	}

	// Call the usecase to get the stock inventory ledger
	resData := h.usecases.Stock.StockInventoryLedgers(request)
	resData.Send(w)
}

func (h *handler) StockDividends(w http.ResponseWriter, r *http.Request) {
	var request domain.ClientStockDividendsRequest
	res := response.New()

	// Decode request based on HTTP method (POST or GET)
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

	// Assign user ID from URL query to the request
	userid, _ := strconv.Atoi(r.URL.Query().Get("uid"))
	request.UserId = userid

	// Validate the inventory ledger request
	err := h.validator.StockDividends(request)
	if err != nil {
		res.SetStatus(http.StatusBadRequest)
		res.SetError(ERROR_CODE_REQUEST_INVALID, err.Error())
		res.Send(w)
		return
	}

	// Call the usecase to get the stock inventory ledger
	resData := h.usecases.Stock.StockDividends(request)
	resData.Send(w)
}

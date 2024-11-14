package v1

import (
	"assetio/internal/adapters/handler/response"
	"assetio/internal/domain"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/schema"
)

// MutualFundBuy handles the buying of mutual funds. It decodes the request, validates the input data,
// and calls the corresponding use case to perform the buy operation.
func (h *handler) MutualFundBuy(w http.ResponseWriter, r *http.Request) {
	var request domain.ClientMutualFundBuyRequest
	res := response.New()

	// If the request method is POST, decode the body into the request struct
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
		// Otherwise, decode the URL query parameters into the request struct
		var decoder = schema.NewDecoder()
		decoder.IgnoreUnknownKeys(true)
		if err := decoder.Decode(&request, r.URL.Query()); err != nil {
			res.SetStatus(http.StatusBadRequest)
			res.SetError(ERROR_CODE_REQUEST_INVALID, err.Error())
			res.Send(w)
			return
		}
	}

	// Extract the user ID from the URL query parameters and set it in the request
	userid, _ := strconv.Atoi(r.URL.Query().Get("uid"))
	request.UserId = userid

	// Validate the request data
	err := h.validator.MutualFundBuy(request)
	if err != nil {
		res.SetStatus(http.StatusBadRequest)
		res.SetError(ERROR_CODE_REQUEST_INVALID, err.Error())
		res.Send(w)
		return
	}

	// Call the use case to process the mutual fund buy request and send the response
	resData := h.usecases.MutualFund.MutualFundBuy(request)
	resData.Send(w)
}

// MutualFundAdd handles the adding of mutual fund details. The method is currently empty and needs to be implemented.
func (h *handler) MutualFundAdd(w http.ResponseWriter, r *http.Request) {}

// MutualFundSell handles the selling of mutual funds. It decodes the request, validates the input data,
// and calls the corresponding use case to perform the sell operation.
func (h *handler) MutualFundSell(w http.ResponseWriter, r *http.Request) {
	var request domain.ClientMutualFundSellRequest
	res := response.New()

	// If the request method is POST, decode the body into the request struct
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
		// Otherwise, decode the URL query parameters into the request struct
		var decoder = schema.NewDecoder()
		decoder.IgnoreUnknownKeys(true)
		if err := decoder.Decode(&request, r.URL.Query()); err != nil {
			res.SetStatus(http.StatusBadRequest)
			res.SetError(ERROR_CODE_REQUEST_INVALID, err.Error())
			res.Send(w)
			return
		}
	}

	// Extract the user ID from the URL query parameters and set it in the request
	userid, _ := strconv.Atoi(r.URL.Query().Get("uid"))
	request.UserId = userid

	// Validate the request data
	err := h.validator.MutualFundSell(request)
	if err != nil {
		res.SetStatus(http.StatusBadRequest)
		res.SetError(ERROR_CODE_REQUEST_INVALID, err.Error())
		res.Send(w)
		return
	}

	// Call the use case to process the mutual fund sell request and send the response
	resData := h.usecases.MutualFund.MutualFundSell(request)
	resData.Send(w)
}

// MutualFundSummary handles fetching the mutual fund summary for a user. It decodes the request, validates the input data,
// and calls the corresponding use case to fetch the summary.
func (h *handler) MutualFundSummary(w http.ResponseWriter, r *http.Request) {
	var request domain.ClientMutualFundSummaryRequest
	res := response.New()

	// If the request method is POST, decode the body into the request struct
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
		// Otherwise, decode the URL query parameters into the request struct
		var decoder = schema.NewDecoder()
		decoder.IgnoreUnknownKeys(true)
		if err := decoder.Decode(&request, r.URL.Query()); err != nil {
			res.SetStatus(http.StatusBadRequest)
			res.SetError(ERROR_CODE_REQUEST_INVALID, err.Error())
			res.Send(w)
			return
		}
	}

	// Extract the user ID from the URL query parameters and set it in the request
	userid, _ := strconv.Atoi(r.URL.Query().Get("uid"))
	request.UserId = userid

	// Validate the request data
	err := h.validator.MutualFundSummary(request)
	if err != nil {
		res.SetStatus(http.StatusBadRequest)
		res.SetError(ERROR_CODE_REQUEST_INVALID, err.Error())
		res.Send(w)
		return
	}

	// Call the use case to fetch the mutual fund summary and send the response
	resData := h.usecases.MutualFund.MutualFundSummary(request)
	resData.Send(w)
}

// MutualFundInventory handles fetching the mutual fund inventory for a user. It decodes the request, validates the input data,
// and calls the corresponding use case to fetch the inventory.
func (h *handler) MutualFundInventory(w http.ResponseWriter, r *http.Request) {
	var request domain.ClientMutualFundInventoryRequest
	res := response.New()

	// If the request method is POST, decode the body into the request struct
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
		// Otherwise, decode the URL query parameters into the request struct
		var decoder = schema.NewDecoder()
		decoder.IgnoreUnknownKeys(true)
		if err := decoder.Decode(&request, r.URL.Query()); err != nil {
			res.SetStatus(http.StatusBadRequest)
			res.SetError(ERROR_CODE_REQUEST_INVALID, err.Error())
			res.Send(w)
			return
		}
	}

	// Extract the user ID from the URL query parameters and set it in the request
	userid, _ := strconv.Atoi(r.URL.Query().Get("uid"))
	request.UserId = userid

	// Validate the request data
	err := h.validator.MutualFundInventory(request)
	if err != nil {
		res.SetStatus(http.StatusBadRequest)
		res.SetError(ERROR_CODE_REQUEST_INVALID, err.Error())
		res.Send(w)
		return
	}

	// Call the use case to fetch the mutual fund inventory and send the response
	resData := h.usecases.MutualFund.MutualFundInventory(request)
	resData.Send(w)
}

// MutualFundTransaction handles fetching mutual fund transaction records. It decodes the request, validates the input data,
// and calls the corresponding use case to fetch the transaction details.
func (h *handler) MutualFundTransaction(w http.ResponseWriter, r *http.Request) {
	var request domain.ClientMutualFundInventoryLedgersRequest
	res := response.New()

	// If the request method is POST, decode the body into the request struct
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
		// Otherwise, decode the URL query parameters into the request struct
		var decoder = schema.NewDecoder()
		decoder.IgnoreUnknownKeys(true)
		if err := decoder.Decode(&request, r.URL.Query()); err != nil {
			res.SetStatus(http.StatusBadRequest)
			res.SetError(ERROR_CODE_REQUEST_INVALID, err.Error())
			res.Send(w)
			return
		}
	}

	// Extract the user ID from the URL query parameters and set it in the request
	userid, _ := strconv.Atoi(r.URL.Query().Get("uid"))
	request.UserId = userid

	// Validate the request data
	err := h.validator.MutualFundInventoryLedgers(request)
	if err != nil {
		res.SetStatus(http.StatusBadRequest)
		res.SetError(ERROR_CODE_REQUEST_INVALID, err.Error())
		res.Send(w)
		return
	}

	// Call the use case to fetch the mutual fund transactions and send the response
	resData := h.usecases.MutualFund.MutualFundInventoryLedgers(request)
	resData.Send(w)
}

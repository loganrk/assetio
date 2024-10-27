package v1

import (
	"net/http"
)

func (h *handler) MutualFundBuy(w http.ResponseWriter, r *http.Request)         {}
func (h *handler) MutualFundSell(w http.ResponseWriter, r *http.Request)        {}
func (h *handler) MutualFundSummary(w http.ResponseWriter, r *http.Request)     {}
func (h *handler) MutualFundInventory(w http.ResponseWriter, r *http.Request)   {}
func (h *handler) MutualFundTransaction(w http.ResponseWriter, r *http.Request) {}

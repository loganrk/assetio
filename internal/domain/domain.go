package domain

import (
	"net/http"
)

type List struct {
	Account    AccountSvr
	Security   SecuritySvr
	Stock      StockSvr
	MutualFund MutualFundSvr
}

type AccountSvr interface {
	AccountCreate(request ClientAccountCreateRequest) Response
	AccountAll(request ClientAccountAllRequest) Response
	AccountGet(request ClientAccountGetRequest) Response
	AccountActivate(request ClientAccountActivateRequest) Response
	AccountInactivate(request ClientAccountInactivateRequest) Response
	AccountUpdate(request ClientAccountUpdateRequest) Response
}

type SecuritySvr interface {
	SecurityCreate(request ClientSecurityCreateRequest) Response
	SecurityAll(request ClientSecurityAllRequest) Response
	SecurityGet(request ClientSecurityGetRequest) Response
	SecuritySearch(request ClientSecuritySearchRequest) Response
	SecurityUpdate(request ClientSecurityUpdateRequest) Response
}

type StockSvr interface {
	StockBuy(request ClientStockBuyRequest) Response
	StockSell(request ClientStockSellRequest) Response
	StockDividendAdd(request ClientStockDividendAddRequest) Response
	StockSplit(request ClientStockSplitRequest) Response
	StockSummary(request ClientStockSummaryRequest) Response
	StockInventories(request ClientStockInventoriesRequest) Response
	StockInventoryLedgers(request ClientStockInventoryLedgersRequest) Response
}

type MutualFundSvr interface {
	MutualFundBuy(request ClientMutualFundBuyRequest) Response
	MutualFundAdd(request ClientMutualFundAddRequest) Response
	MutualFundSell(request ClientMutualFundSellRequest) Response
	MutualFundSummary(request ClientMutualFundSummaryRequest) Response
	MutualFundInventory(request ClientMutualFundInventoryRequest) Response
	MutualFundInventoryLedgers(request ClientMutualFundInventoryLedgersRequest) Response
}

type Response interface {
	SetError(errCode string, errMsg string)
	SetStatus(status int)
	SetData(data any)
	Send(w http.ResponseWriter)
}

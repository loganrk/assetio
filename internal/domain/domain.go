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
	StockSummary(request ClientStockSummaryRequest) Response
	StockInventory(request ClientStockInventoryRequest) Response
	StockInventoryTransactions(request ClientStockInventoryTransactionsRequest) Response
}

type MutualFundSvr interface {
}

type Response interface {
	SetError(errCode string, errMsg string)
	SetStatus(status int)
	SetData(data any)
	Send(w http.ResponseWriter)
}

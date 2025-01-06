package domain

import (
	"net/http"
)

// List holds the different services available for managing accounts, securities, stocks.
// It serves as a container for these services, each implementing its own interface for specific operations.
type List struct {
	Account  AccountSvr  // Service for account-related operations
	Security SecuritySvr // Service for security-related operations
	Stock    StockSvr    // Service for stock-related operations
}

// AccountSvr defines the interface for account-related service operations.
// It includes methods to create, retrieve, update, activate, inactivate, and fetch accounts.
type AccountSvr interface {
	// AccountCreate creates a new client account based on the provided request data.
	AccountCreate(request ClientAccountCreateRequest) Response

	// AccountAll retrieves all client accounts.
	AccountAll(request ClientAccountAllRequest) Response

	// AccountGet retrieves a specific client account using a unique identifier.
	AccountGet(request ClientAccountGetRequest) Response

	// AccountActivate activates a specific account based on the provided request.
	AccountActivate(request ClientAccountActivateRequest) Response

	// AccountInactivate inactivates a specific account.
	AccountInactivate(request ClientAccountInactivateRequest) Response

	// AccountUpdate updates the information of an existing account.
	AccountUpdate(request ClientAccountUpdateRequest) Response
}

// SecuritySvr defines the interface for security-related service operations.
// It includes methods to create, retrieve, update, search, and fetch all securities.
type SecuritySvr interface {
	// SecurityCreate creates a new security based on the provided request data.
	SecurityCreate(request ClientSecurityCreateRequest) Response

	// SecurityAll retrieves all available securities.
	SecurityAll(request ClientSecurityAllRequest) Response

	// SecurityGet retrieves a specific security using a unique identifier.
	SecurityGet(request ClientSecurityGetRequest) Response

	// SecuritySearch searches for securities based on the provided criteria.
	SecuritySearch(request ClientSecuritySearchRequest) Response

	// SecurityUpdate updates an existing security's information.
	SecurityUpdate(request ClientSecurityUpdateRequest) Response
}

// StockSvr defines the interface for stock-related service operations.
// It includes methods to buy, sell, add dividends, split stocks, retrieve stock summaries, inventories, and inventory ledgers.
type StockSvr interface {
	// StockBuy facilitates the purchase of stocks based on the provided request data.
	StockBuy(request ClientStockBuyRequest) Response

	// StockSell facilitates the sale of stocks based on the provided request data.
	StockSell(request ClientStockSellRequest) Response

	// StockDividendAdd adds a dividend to a stock.
	StockDividendAdd(request ClientStockDividendAddRequest) Response

	// StockSplit performs a stock split for the given stock data.
	StockSplit(request ClientStockSplitRequest) Response

	// StockSplit performs a stock bonus for the given stock data.
	StockBonus(request ClientStockBonusRequest) Response

	// StockMerge performs a stock merge for the given stock data.
	StockMerge(request ClientStockMergeRequest) Response

	// StockDemerge performs a stock demerge for the given stock data.
	StockDemerge(request ClientStockDemergeRequest) Response

	// StockSummary retrieves a summary of stock-related information.
	StockSummary(request ClientStockSummaryRequest) Response

	// StockInventories retrieves the inventory of stocks held.
	StockInventories(request ClientStockInventoriesRequest) Response

	// StockInventoryLedgers retrieves the ledger information of stock inventories.
	StockInventoryLedgers(request ClientStockInventoryLedgersRequest) Response

	StockDividends(request ClientStockDividendsRequest) Response
}

// Response defines the interface for a service response.
// It allows setting error codes, statuses, and data, and provides a method to send the response via HTTP.
type Response interface {
	// SetError sets the error code and message for the response.
	SetError(errCode string, errMsg string)

	// SetStatus sets the status code for the response.
	SetStatus(status int)

	// SetData sets the data payload for the response.
	SetData(data any)

	// Send sends the response data via the provided HTTP writer.
	Send(w http.ResponseWriter)
}

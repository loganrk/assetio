package stock

import (
	"assetio/internal/adapters/handler/response"
	"assetio/internal/constant"
	"assetio/internal/domain"
	"context"
	"fmt"
	"net/http"
	"sync"
)

// StockSummary generates a summary of the client's stock holdings based on their account ID and security type.
//
// Parameters:
//   - request: domain.ClientStockSummaryRequest - Contains the account ID and security type for which to fetch the stock summary.
//
// Returns:
//   - domain.Response - Includes a list of stock summary details such as stock ID, symbol, exchange, name, quantity, and amount.
//     Returns an error message if any issue is encountered during data retrieval.
func (s *stockUsecase) StockSummary(request domain.ClientStockSummaryRequest) domain.Response {
	// Create a new context for managing request lifecycle.
	ctx := context.Background()

	// Initialize a response object to store the result of the request.
	res := response.New()

	// Fetch inventory data for the specified account ID and security type (e.g., stocks).
	// If the retrieval fails, log the error and return an internal server error.
	inventoriesData, err := s.mysql.GetInvertriesSummaryByAccountIdAndSecurityType(ctx, request.AccountId, constant.SECURITY_TYPE_STOCK)
	if err != nil {
		// Log the error with request context, error type, and message for troubleshooting.
		s.logger.Errorw(ctx, "inventoriesData failed",
			constant.ERROR_TYPE, constant.ERROR_TYPE_DBEXECUTION,
			constant.ERROR_MESSAGE, err.Error(),
			constant.REQUEST, request,
		)

		// Set the response status to HTTP 500 and provide a generic error message to the client.
		res.SetStatus(http.StatusInternalServerError)
		res.SetError(constant.ERROR_CODE_INTERNAL_SERVER, "internal server error")
		return res
	}

	// Prepare a slice to accumulate the stock summary data to be sent in the response.
	var resData []domain.ClientStockSummaryResponse

	var wg sync.WaitGroup

	// Iterate over each inventory record and transform it into a stock summary format.
	for _, inventoryData := range inventoriesData {
		wg.Add(1)
		go func(inventoryData domain.InventorySummary) {
			metaData := domain.ClientStockSummaryResponse{
				StockId:       inventoryData.SecurityId,
				StockSymbol:   inventoryData.SecuritySymbol,
				StockExchange: inventoryData.SecurityExchange,
				StockName:     inventoryData.SecurityName,
				Quantity:      int(inventoryData.AvailableQuantity),
				Amount:        inventoryData.TotalValue,
			}

			markerData, err := s.marketer.Query(inventoryData.SecuritySymbol, "NSE")
			if err == nil {
				metaData.MarketPrice = markerData.GetMarketPrice()
				metaData.MarketChange = markerData.GetMarketChange()
				metaData.MarketChangePercent = markerData.GetMarketChangePercent()
			}

			resData = append(resData, metaData)
			wg.Done()
		}(inventoryData)

	}
	wg.Wait()

	// Set the formatted stock summary data in the response to be returned to the client.
	res.SetData(resData)
	return res
}

// StockInventories retrieves the inventory details of a client's specific stock holdings.
//
// Parameters:
//   - request: domain.ClientStockInventoriesRequest - contains details of the stock inventory request,
//     including the stock ID and account ID for which to fetch inventory data.
//
// Returns:
//   - domain.Response - includes a list of inventory details such as inventory ID, amount, and available quantity.
//     Returns an error message if any issue is encountered during data retrieval.
func (s *stockUsecase) StockInventories(request domain.ClientStockInventoriesRequest) domain.Response {
	// Create a new background context to manage the request lifecycle.
	ctx := context.Background()

	// Initialize a new response object to hold the result of the request.
	res := response.New()

	// Fetch the security data for the specified stock ID to verify its existence and type.
	secuirityData, err := s.mysql.GetSecurityDataById(ctx, request.StockId)
	if err != nil {
		// Log an error message with the context, error type, and error details for troubleshooting.
		s.logger.Errorw(ctx, "secuirityData failed",
			constant.ERROR_TYPE, constant.ERROR_TYPE_DBEXECUTION,
			constant.ERROR_MESSAGE, err.Error(),
			constant.REQUEST, request,
		)

		// Set HTTP status to 500 and include an internal server error message in the response.
		res.SetStatus(http.StatusInternalServerError)
		res.SetError(constant.ERROR_CODE_INTERNAL_SERVER, "internal server error")
		return res
	}

	// Validate that the retrieved security data corresponds to a stock type.
	// If not, return an error message indicating an invalid stock request.
	if secuirityData.Type != constant.SECURITY_TYPE_STOCK {
		res.SetStatus(http.StatusBadRequest)
		res.SetError(constant.ERROR_CODE_REQUEST_INVALID, "incorrect stock")
		return res
	}

	// Retrieve the inventory data for the specified account and stock ID.
	// If an error occurs, log it and return an internal server error.
	inventoriesData, err := s.mysql.GetInvertriesByAccountIdAndSecurityId(ctx, request.AccountId, request.StockId)
	if err != nil {
		// Log an error message for debugging database execution issues.
		s.logger.Errorw(ctx, "GetInvertriesByAccountIdAndSecurityId failed",
			constant.ERROR_TYPE, constant.ERROR_TYPE_DBEXECUTION,
			constant.ERROR_MESSAGE, err.Error(),
			constant.REQUEST, request,
		)

		// Set HTTP status to 500 and provide an internal server error message to the client.
		res.SetStatus(http.StatusInternalServerError)
		res.SetError(constant.ERROR_CODE_INTERNAL_SERVER, "internal server error")
		return res
	}
	var marketPrice, marketChange, marketChangePercent float64
	markerData, err := s.marketer.Query(secuirityData.Symbol, "NSE")
	if err == nil {
		marketPrice = markerData.GetMarketPrice()
		marketChange = markerData.GetMarketChange()
		marketChangePercent = markerData.GetMarketChangePercent()
	}

	// Initialize a slice to store the inventory details for the response.
	var resData []domain.ClientStockInventoriesResponse

	// Process each inventory record and transform it into a client-specific response format.
	for _, inventoryData := range inventoriesData {
		resData = append(resData, domain.ClientStockInventoriesResponse{
			InventoryId:         inventoryData.Id,
			Amount:              (inventoryData.TotalValue / inventoryData.AvailableQuantity),
			Quantity:            inventoryData.AvailableQuantity,
			MarketPrice:         marketPrice,
			MarketChange:        marketChange,
			MarketChangePercent: marketChangePercent,
			Date:                inventoryData.Date.Format("02-01-2006"),
		})
	}

	// Set the processed inventory data in the response to be returned to the client.
	res.SetData(resData)
	return res
}

// StockInventoryLedgers retrieves the inventory ledger details for a specific client's stock inventory.
//
// Parameters:
//   - request: domain.ClientStockInventoryLedgersRequest - contains details of the inventory ledger request,
//     including the inventory ID and account ID for which to fetch ledger data.
//
// Returns:
//   - domain.Response - includes a list of ledger details such as transaction ID, type, amount, fee, quantity, and date.
//     Returns an error message if any issue is encountered during data retrieval.
func (s *stockUsecase) StockInventoryLedgers(request domain.ClientStockInventoryLedgersRequest) domain.Response {

	// Create a new background context to manage the request lifecycle.
	ctx := context.Background()

	// Initialize a new response object to hold the result of the request.
	res := response.New()

	// Retrieve inventory data based on the provided inventory ID to validate the request.
	inventoryData, err := s.mysql.GetInventoryDataById(ctx, request.InventoryId)
	if err != nil {
		// Log an error if inventory data retrieval fails.
		s.logger.Errorw(ctx, "GetInventoryDataById failed",
			constant.ERROR_TYPE, constant.ERROR_TYPE_DBEXECUTION,
			constant.ERROR_MESSAGE, err.Error(),
			constant.REQUEST, request,
		)

		// Set HTTP status to 500 and include an internal server error message in the response.
		res.SetStatus(http.StatusInternalServerError)
		res.SetError(constant.ERROR_CODE_INTERNAL_SERVER, "internal server error")
		return res
	}

	if inventoryData.Id == 0 {
		res.SetStatus(http.StatusBadRequest)
		res.SetError(constant.ERROR_CODE_REQUEST_INVALID, "incorrect inventory")
		return res
	}

	// Retrieve the security data associated with the inventory's security ID to validate its type.
	secuirityData, err := s.mysql.GetSecurityDataById(ctx, inventoryData.SecurityId)
	if err != nil {
		// Log an error if security data retrieval fails.
		s.logger.Errorw(ctx, "GetSecurityDataById failed",
			constant.ERROR_TYPE, constant.ERROR_TYPE_DBEXECUTION,
			constant.ERROR_MESSAGE, err.Error(),
			constant.REQUEST, request,
		)

		// Set HTTP status to 500 and include an internal server error message in the response.
		res.SetStatus(http.StatusInternalServerError)
		res.SetError(constant.ERROR_CODE_INTERNAL_SERVER, "internal server error")
		return res
	}

	// Ensure that the retrieved security data corresponds to a stock type; otherwise, return a bad request error.
	if secuirityData.Type != constant.SECURITY_TYPE_STOCK {
		res.SetStatus(http.StatusBadRequest)
		res.SetError(constant.ERROR_CODE_REQUEST_INVALID, "incorrect stock")
		return res
	}

	// Retrieve the inventory ledger data for the specified inventory Id
	ledgersData, err := s.mysql.GetInventoryLedgersByInventoryId(ctx, inventoryData.Id)
	if err != nil {
		// Log an error if the ledger data retrieval fails.
		s.logger.Errorw(ctx, "GetInventoryLedgersByInventoryIdAndAccountId failed",
			constant.ERROR_TYPE, constant.ERROR_TYPE_DBEXECUTION,
			constant.ERROR_MESSAGE, err.Error(),
			constant.REQUEST, request,
		)

		// Set HTTP status to 500 and provide an internal server error message to the client.
		res.SetStatus(http.StatusInternalServerError)
		res.SetError(constant.ERROR_CODE_INTERNAL_SERVER, "internal server error")
		return res
	}

	// Initialize a slice to store the ledger details for the response.
	var resData []domain.ClientStockInventoryLedgersResponse

	fmt.Println("ledgersData", ledgersData)
	// Process each transaction ledger record and format it for the response.
	for _, ledgerData := range ledgersData {
		if ledgerData.Quantity > 0 {
			resData = append(resData, domain.ClientStockInventoryLedgersResponse{

				LedgerId: ledgerData.Id,
				Type:     string(ledgerData.Type),
				Amount:   (ledgerData.TotalValue / ledgerData.Quantity),
				Quantity: ledgerData.Quantity,
				Date:     ledgerData.Date.Format("02-01-2006"),
			})
		} else {
			resData = append(resData, domain.ClientStockInventoryLedgersResponse{

				LedgerId:    ledgerData.Id,
				Type:        string(ledgerData.Type),
				TotalAmount: ledgerData.TotalValue,
				Date:        ledgerData.Date.Format("02-01-2006"),
			})
		}

	}

	// Set the formatted transaction ledger data in the response.
	res.SetData(resData)
	return res
}

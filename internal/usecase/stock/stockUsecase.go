package stock

import (
	"assetio/internal/adapters/handler/response"
	"assetio/internal/constant"
	"assetio/internal/domain"
	"assetio/internal/port"
	"context"
	"net/http"
	"sync"
	"time"
)

type stockUsecase struct {
	logger   port.Logger
	mysql    port.RepositoryStore
	marketer port.Marketer
}

func New(loggerIns port.Logger, mysqlIns port.RepositoryStore, marketerIns port.Marketer) domain.StockSvr {
	return &stockUsecase{
		mysql:    mysqlIns,
		logger:   loggerIns,
		marketer: marketerIns,
	}
}

// StockBuy processes a stock purchase for a client by validating security information,
// managing inventory records, recording ledger entries, and updating transaction details.
//
// Parameters:
//   - request: domain.ClientStockBuyRequest - contains details of the stock purchase request,
//     including stock ID, account ID, quantity, price, and fees.
//
// Returns:
//   - domain.Response - contains the outcome of the stock purchase request,
//     either confirming success or detailing any encountered error.
func (s *stockUsecase) StockBuy(request domain.ClientStockBuyRequest) domain.Response {
	// Create a new background context to manage the request lifecycle.
	ctx := context.Background()

	// Initialize a new response object to hold the result of the request.
	res := response.New()

	// Validate security data for the stock ID.
	secuirity, err := s.mysql.GetSecurityDataById(ctx, request.StockId)
	if err != nil {
		s.logger.Errorw(ctx, "GetSecurityDataById failed",
			constant.ERROR_TYPE, constant.ERROR_TYPE_DBEXECUTION,
			constant.ERROR_MESSAGE, err.Error(),
			constant.REQUEST, request,
		)
		res.SetStatus(http.StatusInternalServerError)
		res.SetError(constant.ERROR_CODE_INTERNAL_SERVER, "internal server error")
		return res
	}

	// Check if security type is valid for stock.
	if secuirity.Type != constant.SECURITY_TYPE_STOCK {
		res.SetStatus(http.StatusBadRequest)
		res.SetError(constant.ERROR_CODE_REQUEST_INVALID, "invalid stock")
		return res
	}

	var date = time.Now()

	if request.Date != "" {
		parsedDate, err := time.Parse(constant.DATE_LAYOUT, request.Date)
		if err == nil {
			date = parsedDate
		}
	}

	var inventory domain.Inventories
	// If inventory ID is provided, fetch existing inventory; otherwise, create new.
	if request.InventoryId != 0 {
		inventory, err = s.mysql.GetInventoryDataById(ctx, request.InventoryId)
		if err != nil {
			s.logger.Errorw(ctx, "GetInventoryDataById failed",
				constant.ERROR_TYPE, constant.ERROR_TYPE_DBEXECUTION,
				constant.ERROR_MESSAGE, err.Error(),
				constant.REQUEST, request,
			)
			res.SetStatus(http.StatusInternalServerError)
			res.SetError(constant.ERROR_CODE_INTERNAL_SERVER, "internal server error")
			return res
		}

		// Validate account and stock match for inventory.
		if request.AccountId != inventory.AccountId {
			res.SetStatus(http.StatusBadRequest)
			res.SetError(constant.ERROR_CODE_REQUEST_INVALID, "incorrect inventory")
			return res
		}
		if secuirity.Id != inventory.SecurityId {
			res.SetStatus(http.StatusBadRequest)
			res.SetError(constant.ERROR_CODE_REQUEST_INVALID, "incorrect stock")
			return res
		}

	} else {
		// Insert new inventory if ID is not provided.
		inventory, err = s.mysql.InsertInventoryData(ctx, domain.Inventories{
			AccountId:  request.AccountId,
			SecurityId: secuirity.Id,
			Date:       date,
		})
		if err != nil {
			s.logger.Errorw(ctx, "InsertInventoryData failed",
				constant.ERROR_TYPE, constant.ERROR_TYPE_DBEXECUTION,
				constant.ERROR_MESSAGE, err.Error(),
				constant.REQUEST, request,
			)
			// Set HTTP status to 500 and include an internal server error message in the response.
			res.SetStatus(http.StatusInternalServerError)
			res.SetError(constant.ERROR_CODE_INTERNAL_SERVER, "internal server error")
			return res
		}
	}

	// Record ledger entry for buy transaction.
	inventoryLedgerData, err := s.mysql.InsertInventoryLedger(ctx, domain.InventoryLedger{
		InventoryId:  inventory.Id,
		Type:         domain.BUY,
		Quantity:     request.Quantity,
		AveragePrice: request.AveragePrice,
		Fee:          request.FeeAmount,
		TotalValue:   request.Quantity * request.AveragePrice,
		Date:         date,
	})

	if err != nil {
		s.logger.Errorw(ctx, "InsertInventoryLedger failed",
			constant.ERROR_TYPE, constant.ERROR_TYPE_DBEXECUTION,
			constant.ERROR_MESSAGE, err.Error(),
			constant.REQUEST, request,
		)
		res.SetStatus(http.StatusInternalServerError)
		res.SetError(constant.ERROR_CODE_INTERNAL_SERVER, "internal server error")
		return res
	}

	// Update inventory with new quantity, value, and average price.
	inventory.AvailableQuantity += inventoryLedgerData.Quantity
	inventory.TotalValue += inventoryLedgerData.TotalValue
	inventory.AveragePrice = inventory.TotalValue / inventory.AvailableQuantity

	// Update inventory data in database.
	err = s.mysql.UpdateInventoryDetailsById(ctx, inventory.Id, inventory.AvailableQuantity, inventory.AveragePrice, inventory.TotalValue)
	if err != nil {
		s.logger.Errorw(ctx, "UpdateInventoryDataById failed",
			constant.ERROR_TYPE, constant.ERROR_TYPE_DBEXECUTION,
			constant.ERROR_MESSAGE, err.Error(),
			constant.REQUEST, request,
		)
		res.SetStatus(http.StatusInternalServerError)
		res.SetError(constant.ERROR_CODE_INTERNAL_SERVER, "internal server error")
		return res
	}

	// Insert transaction record for buy operation.
	transactionData, err := s.mysql.InsertTransaction(ctx, domain.Transactions{
		AccountId:    request.AccountId,
		Type:         domain.BUY,
		Quantity:     request.Quantity,
		AveragePrice: request.AveragePrice,
		Fee:          request.FeeAmount,
		TotalValue:   request.Quantity * request.AveragePrice,
		Date:         date,
	})

	if err != nil {
		s.logger.Errorw(ctx, "InsertTransaction failed",
			constant.ERROR_TYPE, constant.ERROR_TYPE_DBEXECUTION,
			constant.ERROR_MESSAGE, err.Error(),
			constant.REQUEST, request,
		)
		res.SetStatus(http.StatusInternalServerError)
		res.SetError(constant.ERROR_CODE_INTERNAL_SERVER, "internal server error")
		return res
	}

	// Link ledger entry to transaction by updating ledger with transaction ID.
	err = s.mysql.UpdateInventoryLedgerTransactionIdById(ctx, inventoryLedgerData.Id, transactionData.Id)
	if err != nil {
		s.logger.Errorw(ctx, "UpdateInventoryLedgerTransactionIdById failed",
			constant.ERROR_TYPE, constant.ERROR_TYPE_DBEXECUTION,
			constant.ERROR_MESSAGE, err.Error(),
			constant.REQUEST, request,
		)
		res.SetStatus(http.StatusInternalServerError)
		res.SetError(constant.ERROR_CODE_INTERNAL_SERVER, "internal server error")
		return res
	}

	// Set success response message.
	resData := domain.ClientStockBuyResponse{
		Message: "stock buy successfully",
	}

	res.SetData(resData)
	return res
}

// StockSell handles the sale of stocks by validating security information,
// managing inventory records, and recording ledger and transaction details.
//
// Parameters:
//   - request: domain.ClientStockSellRequest - contains details of the stock sale request,
//     including stock ID, account ID, quantity, price, and fees.
//
// Returns:
//   - domain.Response - returns the outcome of the stock sale request,
//     including a success message or an error if any issue is encountered

func (s *stockUsecase) StockSell(request domain.ClientStockSellRequest) domain.Response {
	ctx := context.Background()
	res := response.New()

	// Validate security data for the stock ID.
	secuirity, err := s.mysql.GetSecurityDataById(ctx, request.StockId)
	if err != nil {
		s.logger.Errorw(ctx, "GetSecurityDataById failed",
			constant.ERROR_TYPE, constant.ERROR_TYPE_DBEXECUTION,
			constant.ERROR_MESSAGE, err.Error(),
			constant.REQUEST, request,
		)
		res.SetStatus(http.StatusInternalServerError)
		res.SetError(constant.ERROR_CODE_INTERNAL_SERVER, "internal server error")
		return res
	}
	if secuirity.Type != constant.SECURITY_TYPE_STOCK {
		res.SetStatus(http.StatusBadRequest)
		res.SetError(constant.ERROR_CODE_REQUEST_INVALID, "incorrect stock")
		return res
	}

	var date = time.Now()

	if request.Date != "" {
		parsedDate, err := time.Parse(constant.DATE_LAYOUT, request.Date)
		if err == nil {
			date = parsedDate
		}
	}

	var inventories []domain.Inventories

	// Retrieve inventory based on InventoryId, or get all active inventories for the account and stock.
	if request.InventoryId != 0 {
		inventory, err := s.mysql.GetInventoryDataById(ctx, request.InventoryId)
		if err != nil {
			s.logger.Errorw(ctx, "GetInventoryDataById failed",
				constant.ERROR_TYPE, constant.ERROR_TYPE_DBEXECUTION,
				constant.ERROR_MESSAGE, err.Error(),
				constant.REQUEST, request,
			)
			res.SetStatus(http.StatusInternalServerError)
			res.SetError(constant.ERROR_CODE_INTERNAL_SERVER, "internal server error")
			return res
		}

		// Validate account and stock match for inventory and ensure sufficient quantity.

		if request.AccountId != inventory.AccountId {
			res.SetStatus(http.StatusBadRequest)
			res.SetError(constant.ERROR_CODE_REQUEST_INVALID, "incorrect inventory")
			return res
		}
		if secuirity.Id != inventory.SecurityId {
			res.SetStatus(http.StatusBadRequest)
			res.SetError(constant.ERROR_CODE_REQUEST_INVALID, "incorrect stock")
			return res
		}

		if inventory.AvailableQuantity < request.Quantity {
			res.SetStatus(http.StatusBadRequest)
			res.SetError(constant.ERROR_CODE_REQUEST_INVALID, "requested stock not available to sell")
			return res
		}

		inventories = append(inventories, inventory)
	} else {
		// If no InventoryId, fetch all active inventories for this stock and account.
		inventories, err = s.mysql.GetActiveInventoriesByAccountIdAndSecurityId(ctx, request.AccountId, request.StockId)
		if err != nil {
			s.logger.Errorw(ctx, "GetActiveInventoriesByAccountIdAndSecurityId failed",
				constant.ERROR_TYPE, constant.ERROR_TYPE_DBEXECUTION,
				constant.ERROR_MESSAGE, err.Error(),
				constant.REQUEST, request,
			)
			res.SetStatus(http.StatusInternalServerError)
			res.SetError(constant.ERROR_CODE_INTERNAL_SERVER, "internal server error")
			return res
		}

		// Ensure available quantity across all inventories is sufficient.
		var availabletoSell float64
		for _, inventory := range inventories {
			availabletoSell += inventory.AvailableQuantity
		}

		if availabletoSell < request.Quantity {
			res.SetStatus(http.StatusBadRequest)
			res.SetError(constant.ERROR_CODE_REQUEST_INVALID, "requested stock not available to sell")
			return res
		}
	}

	var inventoryLedgerIds []int
	quantity := request.Quantity

	// Process each inventory to fulfill the sale quantity.
	for _, inventory := range inventories {
		if quantity <= 0 {
			break
		}

		// Determine ledger quantity to deduct from each inventory.
		var ledgerQuanity float64
		if inventory.AvailableQuantity < quantity {
			ledgerQuanity = inventory.AvailableQuantity
		} else {
			ledgerQuanity = quantity
		}

		// Record ledger entry for sell transaction.
		inventoryLedgerData, err := s.mysql.InsertInventoryLedger(ctx, domain.InventoryLedger{
			InventoryId:  inventory.Id,
			Type:         domain.SELL,
			Quantity:     ledgerQuanity,
			AveragePrice: request.AveragePrice,
			Fee:          request.FeeAmount,
			TotalValue:   ledgerQuanity * request.AveragePrice,
			Date:         date,
		})

		if err != nil {
			s.logger.Errorw(ctx, "InsertInventoryLedger failed",
				constant.ERROR_TYPE, constant.ERROR_TYPE_DBEXECUTION,
				constant.ERROR_MESSAGE, err.Error(),
				constant.REQUEST, request,
			)
			res.SetStatus(http.StatusInternalServerError)
			res.SetError(constant.ERROR_CODE_INTERNAL_SERVER, "internal server error")
			return res
		}

		inventoryLedgerIds = append(inventoryLedgerIds, inventoryLedgerData.Id)
		inventory, err := s.mysql.GetInventoryDataById(ctx, inventory.Id)
		if err != nil {
			s.logger.Errorw(ctx, "InsertInventoryLedger failed",
				constant.ERROR_TYPE, constant.ERROR_TYPE_DBEXECUTION,
				constant.ERROR_MESSAGE, err.Error(),
				constant.REQUEST, request,
			)
			res.SetStatus(http.StatusInternalServerError)
			res.SetError(constant.ERROR_CODE_INTERNAL_SERVER, "internal server error")
			return res
		}
		// Update inventory data with reduced quantity and recalculated total value.
		inventory.AvailableQuantity -= ledgerQuanity
		inventory.TotalValue = inventory.AvailableQuantity * inventory.AveragePrice
		err = s.mysql.UpdateInventoryDetailsById(ctx, inventory.Id, inventory.AvailableQuantity, inventory.AveragePrice, inventory.TotalValue)
		if err != nil {
			s.logger.Errorw(ctx, "UpdateInventoryDataById failed",
				constant.ERROR_TYPE, constant.ERROR_TYPE_DBEXECUTION,
				constant.ERROR_MESSAGE, err.Error(),
				constant.REQUEST, request,
			)
			res.SetStatus(http.StatusInternalServerError)
			res.SetError(constant.ERROR_CODE_INTERNAL_SERVER, "internal server error")
			return res
		}

		quantity -= ledgerQuanity
	}

	// Insert transaction record for the sell operation.
	transactionData, err := s.mysql.InsertTransaction(ctx, domain.Transactions{
		AccountId:    request.AccountId,
		Type:         domain.SELL,
		Quantity:     request.Quantity,
		AveragePrice: request.AveragePrice,
		Fee:          request.FeeAmount,
		TotalValue:   request.Quantity * request.AveragePrice,
		Date:         date,
	})
	if err != nil {
		s.logger.Errorw(ctx, "InsertTransaction failed",
			constant.ERROR_TYPE, constant.ERROR_TYPE_DBEXECUTION,
			constant.ERROR_MESSAGE, err.Error(),
			constant.REQUEST, request,
		)
		res.SetStatus(http.StatusInternalServerError)
		res.SetError(constant.ERROR_CODE_INTERNAL_SERVER, "internal server error")
		return res
	}

	// Update ledger entries with the transaction ID for tracking purposes.
	err = s.mysql.UpdateInventoryLedgerTransactionIdByIds(ctx, inventoryLedgerIds, transactionData.Id)
	if err != nil {
		s.logger.Errorw(ctx, "UpdateInventoryLedgerTransactionIdById failed",
			constant.ERROR_TYPE, constant.ERROR_TYPE_DBEXECUTION,
			constant.ERROR_MESSAGE, err.Error(),
			constant.REQUEST, request,
		)
		res.SetStatus(http.StatusInternalServerError)
		res.SetError(constant.ERROR_CODE_INTERNAL_SERVER, "internal server error")
		return res
	}

	// Set success response message.
	resData := domain.ClientStockSellResponse{
		Message: "stock sell successfully",
	}

	res.SetData(resData)
	return res
}

// StockSplit processes a stock split for a client by validating security information,
// managing inventory records, recording ledger entries, and updating transaction details.
//
// Parameters:
//   - request: domain.ClientStockSplitRequest - contains details of the stock purchase request,
//     including stock ID, account ID, quantity, price, and fees.
//
// Returns:
//   - domain.Response - contains the outcome of the stock purchase request,
//     either confirming success or detailing any encountered error.
func (s *stockUsecase) StockSplit(request domain.ClientStockSplitRequest) domain.Response {
	// Create a new background context to manage the request lifecycle.
	ctx := context.Background()

	// Initialize a new response object to hold the result of the request.
	res := response.New()

	// Validate security data for the stock ID.
	secuirity, err := s.mysql.GetSecurityDataById(ctx, request.StockId)
	if err != nil {
		s.logger.Errorw(ctx, "GetSecurityDataById failed",
			constant.ERROR_TYPE, constant.ERROR_TYPE_DBEXECUTION,
			constant.ERROR_MESSAGE, err.Error(),
			constant.REQUEST, request,
		)
		res.SetStatus(http.StatusInternalServerError)
		res.SetError(constant.ERROR_CODE_INTERNAL_SERVER, "internal server error")
		return res
	}

	// Check if security type is valid for stock.
	if secuirity.Type != constant.SECURITY_TYPE_STOCK {
		res.SetStatus(http.StatusBadRequest)
		res.SetError(constant.ERROR_CODE_REQUEST_INVALID, "invalid stock")
		return res
	}

	var inventory domain.Inventories
	// If inventory ID is provided, fetch existing inventory; otherwise, create new.
	if request.InventoryId != 0 {
		inventory, err = s.mysql.GetInventoryDataById(ctx, request.InventoryId)
		if err != nil {
			s.logger.Errorw(ctx, "GetInventoryDataById failed",
				constant.ERROR_TYPE, constant.ERROR_TYPE_DBEXECUTION,
				constant.ERROR_MESSAGE, err.Error(),
				constant.REQUEST, request,
			)
			res.SetStatus(http.StatusInternalServerError)
			res.SetError(constant.ERROR_CODE_INTERNAL_SERVER, "internal server error")
			return res
		}

		// Validate account and stock match for inventory.
		if request.AccountId != inventory.AccountId {
			res.SetStatus(http.StatusBadRequest)
			res.SetError(constant.ERROR_CODE_REQUEST_INVALID, "incorrect inventory")
			return res
		}
		if secuirity.Id != inventory.SecurityId {
			res.SetStatus(http.StatusBadRequest)
			res.SetError(constant.ERROR_CODE_REQUEST_INVALID, "incorrect stock")
			return res
		}

	} else {
		// Insert new inventory if ID is not provided.
		inventory, err = s.mysql.InsertInventoryData(ctx, domain.Inventories{
			AccountId:  request.AccountId,
			SecurityId: secuirity.Id,
		})
		if err != nil {
			s.logger.Errorw(ctx, "InsertInventoryData failed",
				constant.ERROR_TYPE, constant.ERROR_TYPE_DBEXECUTION,
				constant.ERROR_MESSAGE, err.Error(),
				constant.REQUEST, request,
			)
			// Set HTTP status to 500 and include an internal server error message in the response.
			res.SetStatus(http.StatusInternalServerError)
			res.SetError(constant.ERROR_CODE_INTERNAL_SERVER, "internal server error")
			return res
		}
	}

	// Record ledger entry for split transaction.
	inventoryLedgerData, err := s.mysql.InsertInventoryLedger(ctx, domain.InventoryLedger{
		InventoryId: inventory.Id,
		Type:        domain.SPLIT,
		Quantity:    request.Quantity,
		Fee:         request.FeeAmount,
		Date:        time.Now(),
	})

	if err != nil {
		s.logger.Errorw(ctx, "InsertInventoryLedger failed",
			constant.ERROR_TYPE, constant.ERROR_TYPE_DBEXECUTION,
			constant.ERROR_MESSAGE, err.Error(),
			constant.REQUEST, request,
		)
		res.SetStatus(http.StatusInternalServerError)
		res.SetError(constant.ERROR_CODE_INTERNAL_SERVER, "internal server error")
		return res
	}

	// Update inventory with new quantity, value, and average price.
	inventory.AvailableQuantity += inventoryLedgerData.Quantity
	inventory.TotalValue += inventoryLedgerData.TotalValue
	inventory.AveragePrice = inventory.TotalValue / inventory.AvailableQuantity

	// Update inventory data in database.
	err = s.mysql.UpdateInventoryDetailsById(ctx, inventory.Id, inventory.AvailableQuantity, inventory.AveragePrice, inventory.TotalValue)
	if err != nil {
		s.logger.Errorw(ctx, "UpdateInventoryDataById failed",
			constant.ERROR_TYPE, constant.ERROR_TYPE_DBEXECUTION,
			constant.ERROR_MESSAGE, err.Error(),
			constant.REQUEST, request,
		)
		res.SetStatus(http.StatusInternalServerError)
		res.SetError(constant.ERROR_CODE_INTERNAL_SERVER, "internal server error")
		return res
	}

	// Insert transaction record for buy operation.
	transactionData, err := s.mysql.InsertTransaction(ctx, domain.Transactions{
		AccountId: request.AccountId,
		Type:      domain.SPLIT,
		Quantity:  request.Quantity,
		Fee:       request.FeeAmount,
		Date:      time.Now(),
	})

	if err != nil {
		s.logger.Errorw(ctx, "InsertTransaction failed",
			constant.ERROR_TYPE, constant.ERROR_TYPE_DBEXECUTION,
			constant.ERROR_MESSAGE, err.Error(),
			constant.REQUEST, request,
		)
		res.SetStatus(http.StatusInternalServerError)
		res.SetError(constant.ERROR_CODE_INTERNAL_SERVER, "internal server error")
		return res
	}

	// Link ledger entry to transaction by updating ledger with transaction ID.
	err = s.mysql.UpdateInventoryLedgerTransactionIdById(ctx, inventoryLedgerData.Id, transactionData.Id)
	if err != nil {
		s.logger.Errorw(ctx, "UpdateInventoryLedgerTransactionIdById failed",
			constant.ERROR_TYPE, constant.ERROR_TYPE_DBEXECUTION,
			constant.ERROR_MESSAGE, err.Error(),
			constant.REQUEST, request,
		)
		res.SetStatus(http.StatusInternalServerError)
		res.SetError(constant.ERROR_CODE_INTERNAL_SERVER, "internal server error")
		return res
	}

	// Set success response message.
	resData := domain.ClientStockBuyResponse{
		Message: "stock split successfully",
	}

	res.SetData(resData)
	return res
}

// StockDividendAdd handles the addition of dividends for a client's stock holdings.
//
// Parameters:
//   - request: domain.ClientStockDividendAddRequest - contains details of the stock dividend request,
//     including stock ID, account ID, dividend quantity, and average price.
//
// Returns:
//   - domain.Response - returns the outcome of the dividend addition request,
//     including a success message or an error if an issue is encountered.
func (s *stockUsecase) StockDividendAdd(request domain.ClientStockDividendAddRequest) domain.Response {
	ctx := context.Background()
	res := response.New()

	// Retrieve and validate the security data for the provided stock ID.
	secuirity, err := s.mysql.GetSecurityDataById(ctx, request.StockId)
	if err != nil {
		s.logger.Errorw(ctx, "GetSecurityDataById failed",
			constant.ERROR_TYPE, constant.ERROR_TYPE_DBEXECUTION,
			constant.ERROR_MESSAGE, err.Error(),
			constant.REQUEST, request,
		)
		res.SetStatus(http.StatusInternalServerError)
		res.SetError(constant.ERROR_CODE_INTERNAL_SERVER, "internal server error")
		return res
	}
	if secuirity.Type != constant.SECURITY_TYPE_STOCK {
		res.SetStatus(http.StatusBadRequest)
		res.SetError(constant.ERROR_CODE_REQUEST_INVALID, "incorrect stock")
		return res
	}

	var inventories []domain.Inventories

	// If InventoryId is provided, fetch and validate the specific inventory.
	if request.InventoryId != 0 {
		inventory, err := s.mysql.GetInventoryDataById(ctx, request.InventoryId)
		if err != nil {
			s.logger.Errorw(ctx, "GetInventoryDataById failed",
				constant.ERROR_TYPE, constant.ERROR_TYPE_DBEXECUTION,
				constant.ERROR_MESSAGE, err.Error(),
				constant.REQUEST, request,
			)
			res.SetStatus(http.StatusInternalServerError)
			res.SetError(constant.ERROR_CODE_INTERNAL_SERVER, "internal server error")
			return res
		}

		// Validate account, stock match, and check availability of inventory quantity.
		if request.AccountId != inventory.AccountId {
			res.SetStatus(http.StatusBadRequest)
			res.SetError(constant.ERROR_CODE_REQUEST_INVALID, "incorrect inventory")
			return res
		}
		if secuirity.Id != inventory.SecurityId {
			res.SetStatus(http.StatusBadRequest)
			res.SetError(constant.ERROR_CODE_REQUEST_INVALID, "incorrect stock")
			return res
		}

		inventories = append(inventories, inventory)
	} else {
		// If no InventoryId is provided, retrieve all active inventories for the stock and account.
		inventories, err = s.mysql.GetActiveInventoriesByAccountIdAndSecurityId(ctx, request.AccountId, request.StockId)
		if err != nil {
			s.logger.Errorw(ctx, "GetActiveInventoriesByAccountIdAndSecurityId failed",
				constant.ERROR_TYPE, constant.ERROR_TYPE_DBEXECUTION,
				constant.ERROR_MESSAGE, err.Error(),
				constant.REQUEST, request,
			)
			res.SetStatus(http.StatusInternalServerError)
			res.SetError(constant.ERROR_CODE_INTERNAL_SERVER, "internal server error")
			return res
		}

	}

	var inventoryLedgerIds []int
	var totalQuantity float64
	// Process each inventory to apply the dividend for the specified quantity.
	for _, inventory := range inventories {
		if inventory.AvailableQuantity <= 0 {
			continue
		}

		inventory, err := s.mysql.GetInventoryDataById(ctx, inventory.Id)
		if err != nil {
			s.logger.Errorw(ctx, "InsertInventoryLedger failed",
				constant.ERROR_TYPE, constant.ERROR_TYPE_DBEXECUTION,
				constant.ERROR_MESSAGE, err.Error(),
				constant.REQUEST, request,
			)
			res.SetStatus(http.StatusInternalServerError)
			res.SetError(constant.ERROR_CODE_INTERNAL_SERVER, "internal server error")
			return res
		}

		// Insert a ledger entry to record the dividend addition.
		inventoryLedgerData, err := s.mysql.InsertInventoryLedger(ctx, domain.InventoryLedger{
			InventoryId:  inventory.Id,
			Type:         domain.DIVIDEND,
			Quantity:     inventory.AvailableQuantity,
			AveragePrice: request.AmountPerQuantity,
			TotalValue:   inventory.AvailableQuantity * request.AmountPerQuantity,
			Date:         time.Now(),
		})

		if err != nil {
			s.logger.Errorw(ctx, "InsertInventoryLedger failed",
				constant.ERROR_TYPE, constant.ERROR_TYPE_DBEXECUTION,
				constant.ERROR_MESSAGE, err.Error(),
				constant.REQUEST, request,
			)
			res.SetStatus(http.StatusInternalServerError)
			res.SetError(constant.ERROR_CODE_INTERNAL_SERVER, "internal server error")
			return res
		}

		// Append ledger ID to track this entry and decrement remaining quantity.
		inventoryLedgerIds = append(inventoryLedgerIds, inventoryLedgerData.Id)
		totalQuantity += inventory.AvailableQuantity
	}

	// Insert a transaction entry to log the dividend distribution.
	transactionData, err := s.mysql.InsertTransaction(ctx, domain.Transactions{
		AccountId:    request.AccountId,
		Type:         domain.DIVIDEND,
		Quantity:     totalQuantity,
		AveragePrice: request.AmountPerQuantity,
		TotalValue:   totalQuantity * request.AmountPerQuantity,
		Date:         time.Now(),
	})

	if err != nil {
		s.logger.Errorw(ctx, "InsertTransaction failed",
			constant.ERROR_TYPE, constant.ERROR_TYPE_DBEXECUTION,
			constant.ERROR_MESSAGE, err.Error(),
			constant.REQUEST, request,
		)
		res.SetStatus(http.StatusInternalServerError)
		res.SetError(constant.ERROR_CODE_INTERNAL_SERVER, "internal server error")
		return res
	}

	// Update the inventory ledger entries with the transaction ID for tracking.
	err = s.mysql.UpdateInventoryLedgerTransactionIdByIds(ctx, inventoryLedgerIds, transactionData.Id)
	if err != nil {
		s.logger.Errorw(ctx, "UpdateInventoryLedgerTransactionIdById failed",
			constant.ERROR_TYPE, constant.ERROR_TYPE_DBEXECUTION,
			constant.ERROR_MESSAGE, err.Error(),
			constant.REQUEST, request,
		)
		res.SetStatus(http.StatusInternalServerError)
		res.SetError(constant.ERROR_CODE_INTERNAL_SERVER, "internal server error")
		return res
	}

	// Set success message and response data.
	resData := domain.ClientStockDividendResponse{
		Message: "stock dividend successfully",
	}

	res.SetData(resData)
	return res
}

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
			Quantity:            int(inventoryData.AvailableQuantity),
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

	// Retrieve the inventory ledger data for the specified inventory and account IDs.
	transactionsData, err := s.mysql.GetInventoryLedgersByInventoryIdAndAccountId(ctx, request.AccountId, inventoryData.Id)
	if err != nil {
		// Log an error if the ledger data retrieval fails.
		s.logger.Errorw(ctx, "GetInventoryLedgersByIdAndAccountId failed",
			constant.ERROR_TYPE, constant.ERROR_TYPE_DBEXECUTION,
			constant.ERROR_MESSAGE, err.Error(),
			constant.REQUEST, request,
		)

		// Set HTTP status to 500 and provide an internal server error message to the client.
		res.SetStatus(http.StatusInternalServerError)
		res.SetError(constant.ERROR_CODE_INTERNAL_SERVER, "internal server error")
		return res
	}

	// Initialize a slice to store the transaction ledger details for the response.
	var resData []domain.ClientStockInventoryLedgersResponse

	// Process each transaction ledger record and format it for the response.
	for _, transactionData := range transactionsData {
		resData = append(resData, domain.ClientStockInventoryLedgersResponse{
			TransactionId:   transactionData.Id,
			TransactionType: string(transactionData.Type),
			Amount:          (transactionData.TotalValue / transactionData.Quantity),
			Quantity:        int(transactionData.Quantity),
			Date:            transactionData.Date.Format("02-01-2006"),
		})
	}

	// Set the formatted transaction ledger data in the response.
	res.SetData(resData)
	return res
}

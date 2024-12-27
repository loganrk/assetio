package stock

import (
	"assetio/internal/adapters/handler/response"
	"assetio/internal/constant"
	"assetio/internal/domain"
	"context"
	"net/http"
	"time"
)

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
		AccountId:  request.AccountId,
		SecurityId: secuirity.Id,
		Type:       domain.SPLIT,
		Quantity:   request.Quantity,
		Fee:        request.FeeAmount,
		Date:       time.Now(),
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

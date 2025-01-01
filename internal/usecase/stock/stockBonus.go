package stock

import (
	"assetio/internal/adapters/handler/response"
	"assetio/internal/constant"
	"assetio/internal/domain"
	"context"
	"net/http"
	"time"
)

// StockBonus processes a stock bonus for a client by validating security information,
// managing inventory records, recording ledger entries, and updating transaction details.
//
// Parameters:
//   - request: domain.ClientStockBonusRequest - contains details of the stock purchase request,
//     including stock ID, account ID, quantity, price, and fees.
//
// Returns:
//   - domain.Response - contains the outcome of the stock purchase request,
//     either confirming success or detailing any encountered error.
func (s *stockUsecase) StockBonus(request domain.ClientStockBonusRequest) domain.Response {
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

	inventories, err := s.mysql.GetActiveInventoriesByAccountIdAndSecurityId(ctx, request.AccountId, request.StockId)
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

	// Insert transaction record for buy operation.
	transactionData, err := s.mysql.InsertTransaction(ctx, domain.Transactions{
		AccountId:  request.AccountId,
		SecurityId: secuirity.Id,
		Type:       domain.BONUS,
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

	var availableQuantities float64
	for _, inventory := range inventories {
		availableQuantities += inventory.AvailableQuantity
	}

	newlyAddedRatio := float64(request.Quantity) / availableQuantities

	for _, inventory := range inventories {
		newStockForInv := newlyAddedRatio * inventory.AvailableQuantity
		inventoryLedgerData, err := s.mysql.InsertInventoryLedger(ctx, domain.InventoryLedger{
			InventoryId:   inventory.Id,
			TransactionId: transactionData.Id,
			Type:          domain.BONUS,
			Quantity:      newStockForInv,
			Fee:           request.FeeAmount,
			Date:          time.Now(),
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
	}

	// Set success response message.
	resData := domain.ClientStockBonusResponse{
		Message: "stock bonus added successfully",
	}

	res.SetData(resData)
	return res
}

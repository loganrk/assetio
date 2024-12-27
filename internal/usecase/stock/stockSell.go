package stock

import (
	"assetio/internal/adapters/handler/response"
	"assetio/internal/constant"
	"assetio/internal/domain"
	"context"
	"net/http"
	"time"
)

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
		SecurityId:   secuirity.Id,
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

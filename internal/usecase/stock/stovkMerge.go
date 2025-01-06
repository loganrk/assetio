package stock

import (
	"assetio/internal/adapters/handler/response"
	"assetio/internal/constant"
	"assetio/internal/domain"
	"context"
	"net/http"
	"time"
)

func (s *stockUsecase) StockMerge(request domain.ClientStockMergeRequest) domain.Response {
	// Create a new background context to manage the request lifecycle.
	ctx := context.Background()

	// Initialize a new response object to hold the result of the request.
	res := response.New()

	// Validate security data for the stock ID.
	parrentSecuirity, err := s.mysql.GetSecurityDataById(ctx, request.ParentStockId)
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
	if parrentSecuirity.Id == 0 || parrentSecuirity.Type != constant.SECURITY_TYPE_STOCK {
		res.SetStatus(http.StatusBadRequest)
		res.SetError(constant.ERROR_CODE_REQUEST_INVALID, "invalid parrent stock")
		return res
	}

	// Validate security data for the stock ID.
	newSecuirity, err := s.mysql.GetSecurityDataById(ctx, request.NewStockId)
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
	if newSecuirity.Id == 0 || newSecuirity.Type != constant.SECURITY_TYPE_STOCK {
		res.SetStatus(http.StatusBadRequest)
		res.SetError(constant.ERROR_CODE_REQUEST_INVALID, "invalid new stock")
		return res
	}

	var date = time.Now()

	if request.Date != "" {
		parsedDate, err := time.Parse(constant.DATE_LAYOUT, request.Date)
		if err == nil {
			date = parsedDate
		}
	}

	inventories, err := s.mysql.GetActiveInventoriesByAccountIdAndSecurityId(ctx, request.AccountId, request.ParentStockId)
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

	var availableQuantities float64
	var totalAmount float64
	var inventoryLedgerIds []int

	for _, inventory := range inventories {
		availableQuantities += inventory.AvailableQuantity
		totalAmount += inventory.TotalValue
	}

	averagePrice := totalAmount / availableQuantities

	for _, inventory := range inventories {

		inventoryLedgerData, err := s.mysql.InsertInventoryLedger(ctx, domain.InventoryLedger{
			InventoryId:  inventory.Id,
			Type:         domain.MERGER_TRANSFER,
			Quantity:     inventory.AvailableQuantity,
			AveragePrice: inventory.AveragePrice,
			TotalValue:   inventory.TotalValue,
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

		inventoryLedgerIds = append(inventoryLedgerIds, inventoryLedgerData.Id)
		inventory.AvailableQuantity -= inventoryLedgerData.Quantity
		inventory.TotalValue -= inventoryLedgerData.TotalValue
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

	// Insert transaction record for buy operation.
	transactionData, err := s.mysql.InsertTransaction(ctx, domain.Transactions{
		AccountId:    request.AccountId,
		SecurityId:   parrentSecuirity.Id,
		Type:         domain.MERGER_TRANSFER,
		AveragePrice: averagePrice,
		Quantity:     availableQuantities,
		TotalValue:   averagePrice * availableQuantities,
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

	// Insert new inventory .
	inventory, err := s.mysql.InsertInventoryData(ctx, domain.Inventories{
		AccountId:  request.AccountId,
		SecurityId: newSecuirity.Id,
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

	// Record ledger entry for buy transaction.
	inventoryLedgerData, err := s.mysql.InsertInventoryLedger(ctx, domain.InventoryLedger{
		InventoryId:  inventory.Id,
		Type:         domain.MERGER,
		Quantity:     request.Quantity,
		AveragePrice: totalAmount / request.Quantity,
		TotalValue:   totalAmount,
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
	transactionData, err = s.mysql.InsertTransaction(ctx, domain.Transactions{
		AccountId:    request.AccountId,
		SecurityId:   newSecuirity.Id,
		Type:         domain.MERGER,
		Quantity:     request.Quantity,
		AveragePrice: totalAmount / request.Quantity,
		TotalValue:   totalAmount,
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
	resData := domain.ClientStockMergeResponse{
		Message: "stock merge successfully",
	}

	res.SetData(resData)
	return res
}

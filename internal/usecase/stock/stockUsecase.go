package stock

import (
	"assetio/internal/adapters/handler/response"
	"assetio/internal/constant"
	"assetio/internal/domain"
	"assetio/internal/port"
	"context"
	"net/http"
	"time"
)

type stockUsecase struct {
	logger port.Logger
	mysql  port.RepositoryStore
}

func New(loggerIns port.Logger, mysqlIns port.RepositoryStore) domain.StockSvr {
	return &stockUsecase{
		mysql:  mysqlIns,
		logger: loggerIns,
	}
}

func (s *stockUsecase) StockBuy(request domain.ClientStockBuyRequest) domain.Response {
	ctx := context.Background()
	res := response.New()

	secuirity, err := s.mysql.GetSecuriryById(ctx, request.StockId)
	if err != nil {
		res.SetStatus(http.StatusInternalServerError)
		res.SetError(constant.ERROR_CODE_INTERNAL_SERVER, "internal server error")
		return res
	}

	if secuirity.Type != constant.SECURITY_TYPE_STOCK {
		res.SetStatus(http.StatusBadRequest)
		res.SetError(constant.ERROR_CODE_REQUEST_INVALID, "invalid stock")
		return res
	}

	inventory, err := s.mysql.InsertInventory(ctx, domain.Inventories{
		AccountId:         request.AccountId,
		SecurityId:        secuirity.Id,
		AvailableQuantiry: request.Quantity,
		Price:             request.AmountPerQuantity,
	})
	if err != nil {
		res.SetStatus(http.StatusInternalServerError)
		res.SetError(constant.ERROR_CODE_INTERNAL_SERVER, "internal server error")
		return res
	}

	_, err = s.mysql.InsertTransaction(ctx, domain.Transactions{
		AccountId:   request.AccountId,
		InventoryId: inventory.Id,
		Type:        domain.Buy,
		Quantity:    request.Quantity,
		Price:       request.AmountPerQuantity,
		Fee:         request.FeeAmount,
		Date:        time.Now(),
	})

	if err != nil {
		res.SetStatus(http.StatusInternalServerError)
		res.SetError(constant.ERROR_CODE_INTERNAL_SERVER, "internal server error")
		return res
	}

	resData := domain.ClientStockBuyResponse{
		Message: "stock buy successfully",
	}

	res.SetData(resData)
	return res
}

func (s *stockUsecase) StockSell(request domain.ClientStockSellRequest) domain.Response {
	ctx := context.Background()
	res := response.New()

	secuirity, err := s.mysql.GetSecuriryById(ctx, request.StockId)
	if err != nil {
		res.SetStatus(http.StatusInternalServerError)
		res.SetError(constant.ERROR_CODE_INTERNAL_SERVER, "internal server error")
		return res
	}
	if secuirity.Type != constant.SECURITY_TYPE_STOCK {
		res.SetStatus(http.StatusBadRequest)
		res.SetError(constant.ERROR_CODE_REQUEST_INVALID, "incorrect stock")
		return res
	}

	if request.InventoryId != 0 {
		inventory, err := s.mysql.GetInventoryById(ctx, request.InventoryId)
		if err != nil {
			res.SetStatus(http.StatusInternalServerError)
			res.SetError(constant.ERROR_CODE_INTERNAL_SERVER, "internal server error")
			return res
		}

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

		if inventory.AvailableQuantiry < request.Quantity {
			res.SetStatus(http.StatusBadRequest)
			res.SetError(constant.ERROR_CODE_REQUEST_INVALID, "no stock available to sell")
			return res
		}

		err = s.mysql.UpdateAvailableQuanityToInventoryById(ctx, inventory.Id, inventory.AvailableQuantiry-request.Quantity)
		if err != nil {
			res.SetStatus(http.StatusInternalServerError)
			res.SetError(constant.ERROR_CODE_INTERNAL_SERVER, "internal server error")
			return res
		}
		_, err = s.mysql.InsertTransaction(ctx, domain.Transactions{
			AccountId:   request.AccountId,
			InventoryId: inventory.Id,
			Type:        domain.Sell,
			Quantity:    inventory.AvailableQuantiry - request.Quantity,
			Price:       request.AmountPerQuantity,
			Fee:         request.FeeAmount,
			Date:        time.Now(),
		})

		if err != nil {
			res.SetStatus(http.StatusInternalServerError)
			res.SetError(constant.ERROR_CODE_INTERNAL_SERVER, "internal server error")
			return res
		}

	} else {
		inventories, err := s.mysql.GetActiveInventoriesByAccountIdAndSecurityId(ctx, request.AccountId, request.StockId)
		if err != nil {
			res.SetStatus(http.StatusInternalServerError)
			res.SetError(constant.ERROR_CODE_INTERNAL_SERVER, "internal server error")
			return res
		}
		availQuanitiy := request.Quantity
		for _, inventory := range inventories {
			if availQuanitiy != 0 {
				if inventory.AvailableQuantiry <= availQuanitiy {
					err := s.mysql.UpdateAvailableQuanityToInventoryById(ctx, inventory.Id, 0)
					if err != nil {
						res.SetStatus(http.StatusInternalServerError)
						res.SetError(constant.ERROR_CODE_INTERNAL_SERVER, "internal server error")
						return res
					}
					_, err = s.mysql.InsertTransaction(ctx, domain.Transactions{
						AccountId:   request.AccountId,
						InventoryId: inventory.Id,
						Type:        domain.Sell,
						Quantity:    inventory.AvailableQuantiry,
						Price:       request.AmountPerQuantity,
						Fee:         request.FeeAmount,
						Date:        time.Now(),
					})

					if err != nil {
						res.SetStatus(http.StatusInternalServerError)
						res.SetError(constant.ERROR_CODE_INTERNAL_SERVER, "internal server error")
						return res
					}

					availQuanitiy = availQuanitiy - inventory.AvailableQuantiry
				} else {
					err := s.mysql.UpdateAvailableQuanityToInventoryById(ctx, inventory.Id, inventory.AvailableQuantiry-availQuanitiy)
					if err != nil {
						res.SetStatus(http.StatusInternalServerError)
						res.SetError(constant.ERROR_CODE_INTERNAL_SERVER, "internal server error")
						return res
					}

					_, err = s.mysql.InsertTransaction(ctx, domain.Transactions{
						AccountId:   request.AccountId,
						InventoryId: inventory.Id,
						Type:        domain.Sell,
						Quantity:    request.Quantity,
						Price:       request.AmountPerQuantity,
						Fee:         request.FeeAmount,
						Date:        time.Now(),
					})

					if err != nil {
						res.SetStatus(http.StatusInternalServerError)
						res.SetError(constant.ERROR_CODE_INTERNAL_SERVER, "internal server error")
						return res
					}

					availQuanitiy = 0
				}
			} else {
				break
			}
		}

		if availQuanitiy != 0 {
			res.SetStatus(http.StatusBadRequest)
			res.SetError(constant.ERROR_CODE_REQUEST_INVALID, "some of stock quanity not available to sell")
			return res
		}

	}
	resData := domain.ClientStockSellResponse{
		Message: "stock sell successfully",
	}

	res.SetData(resData)
	return res
}

func (s *stockUsecase) StockDividendAdd(request domain.ClientStockDividendAddRequest) domain.Response {
	ctx := context.Background()
	res := response.New()

	secuirity, err := s.mysql.GetSecuriryById(ctx, request.StockId)
	if err != nil {
		res.SetStatus(http.StatusInternalServerError)
		res.SetError(constant.ERROR_CODE_INTERNAL_SERVER, "internal server error")
		return res
	}
	if secuirity.Type != constant.SECURITY_TYPE_STOCK {
		res.SetStatus(http.StatusBadRequest)
		res.SetError(constant.ERROR_CODE_REQUEST_INVALID, "incorrect stock")
		return res

	}

	if request.InventoryId != 0 {
		inventory, err := s.mysql.GetInventoryById(ctx, request.InventoryId)
		if err != nil {
			res.SetStatus(http.StatusInternalServerError)
			res.SetError(constant.ERROR_CODE_INTERNAL_SERVER, "internal server error")
			return res
		}

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

		if inventory.AvailableQuantiry < request.Quantity {
			res.SetStatus(http.StatusBadRequest)
			res.SetError(constant.ERROR_CODE_REQUEST_INVALID, "stock quanity not available to dividend")
			return res
		}
		_, err = s.mysql.InsertTransaction(ctx, domain.Transactions{
			AccountId:   request.AccountId,
			InventoryId: inventory.Id,
			Type:        domain.Dividend,
			Quantity:    request.Quantity,
			Price:       request.AmountPerQuantity,
			Date:        time.Now(),
		})

		if err != nil {
			res.SetStatus(http.StatusInternalServerError)
			res.SetError(constant.ERROR_CODE_INTERNAL_SERVER, "internal server error")
			return res
		}

	} else {
		inventories, err := s.mysql.GetActiveInventoriesByAccountIdAndSecurityId(ctx, request.AccountId, request.StockId)
		if err != nil {
			res.SetStatus(http.StatusInternalServerError)
			res.SetError(constant.ERROR_CODE_INTERNAL_SERVER, "internal server error")
			return res
		}

		availQuanitiy := request.Quantity

		for _, inventory := range inventories {
			if availQuanitiy != 0 {
				if inventory.AvailableQuantiry <= availQuanitiy {
					_, err = s.mysql.InsertTransaction(ctx, domain.Transactions{
						AccountId:   request.AccountId,
						InventoryId: inventory.Id,
						Type:        domain.Dividend,
						Quantity:    inventory.AvailableQuantiry,
						Price:       request.AmountPerQuantity,
						Date:        time.Now(),
					})

					if err != nil {
						res.SetStatus(http.StatusInternalServerError)
						res.SetError(constant.ERROR_CODE_INTERNAL_SERVER, "internal server error")
						return res
					}

					availQuanitiy = availQuanitiy - inventory.AvailableQuantiry

				} else {
					_, err = s.mysql.InsertTransaction(ctx, domain.Transactions{
						AccountId:   request.AccountId,
						InventoryId: inventory.Id,
						Type:        domain.Dividend,
						Quantity:    availQuanitiy,
						Price:       request.AmountPerQuantity,
						Date:        time.Now(),
					})

					if err != nil {
						res.SetStatus(http.StatusInternalServerError)
						res.SetError(constant.ERROR_CODE_INTERNAL_SERVER, "internal server error")
						return res
					}

					availQuanitiy = 0
				}
			} else {
				break
			}
		}

		if availQuanitiy != 0 {
			res.SetStatus(http.StatusBadRequest)
			res.SetError(constant.ERROR_CODE_REQUEST_INVALID, "some of stock quanity not available to dividend")
			return res
		}
	}

	resData := domain.ClientStockDividendResponse{
		Message: "stock dividend added successfully",
	}

	res.SetData(resData)
	return res
}

func (s *stockUsecase) StockSummary(request domain.ClientStockSummaryRequest) domain.Response {
	ctx := context.Background()
	res := response.New()

	inventoriesData, err := s.mysql.SelectInvertriesSummaryByAccountIdAndSecurityType(ctx, request.AccountId, constant.SECURITY_TYPE_STOCK)
	if err != nil {
		res.SetStatus(http.StatusInternalServerError)
		res.SetError(constant.ERROR_CODE_INTERNAL_SERVER, "internal server error")
		return res
	}

	var resData []domain.ClientStockSummaryResponse

	for _, inventoryData := range inventoriesData {
		resData = append(resData, domain.ClientStockSummaryResponse{
			StockId:       inventoryData.SecurityId,
			StockSymbol:   inventoryData.SecuritySymbol,
			StockExchange: inventoryData.SecurityExchange,
			StockName:     inventoryData.SecurityName,
			Quantity:      int(inventoryData.Quantity),
			Amount:        inventoryData.Amount,
		})
	}

	res.SetData(resData)
	return res

}

func (s *stockUsecase) StockInventory(request domain.ClientStockInventoryRequest) domain.Response {
	ctx := context.Background()
	res := response.New()

	secuirityData, err := s.mysql.GetSecuriryById(ctx, request.StockId)
	if err != nil {
		res.SetStatus(http.StatusInternalServerError)
		res.SetError(constant.ERROR_CODE_INTERNAL_SERVER, "internal server error")
		return res
	}

	if secuirityData.Type != constant.SECURITY_TYPE_STOCK {
		res.SetStatus(http.StatusBadRequest)
		res.SetError(constant.ERROR_CODE_REQUEST_INVALID, "incorrect stock")
		return res
	}

	inventoriesData, err := s.mysql.SelectInvertriesByAccountIdAndStockId(ctx, request.AccountId, request.StockId)
	if err != nil {
		res.SetStatus(http.StatusInternalServerError)
		res.SetError(constant.ERROR_CODE_INTERNAL_SERVER, "internal server error")
		return res
	}
	var resData []domain.ClientStockInventoryResponse

	for _, inventoryData := range inventoriesData {
		resData = append(resData, domain.ClientStockInventoryResponse{
			InventoryId: inventoryData.Id,
			Amount:      inventoryData.Price,
			Quantity:    int(inventoryData.AvailableQuantiry),
			Date:        inventoryData.Date.Format("02-01-2006"),
		})
	}

	res.SetData(resData)
	return res
}

func (s *stockUsecase) StockInventoryTransactions(request domain.ClientStockInventoryTransactionsRequest) domain.Response {

	ctx := context.Background()
	res := response.New()

	inventoryData, err := s.mysql.GetInventoryById(ctx, request.InventoryId)
	if err != nil {
		res.SetStatus(http.StatusInternalServerError)
		res.SetError(constant.ERROR_CODE_INTERNAL_SERVER, "internal server error")
		return res
	}

	secuirityData, err := s.mysql.GetSecuriryById(ctx, inventoryData.SecurityId)
	if err != nil {
		res.SetStatus(http.StatusInternalServerError)
		res.SetError(constant.ERROR_CODE_INTERNAL_SERVER, "internal server error")
		return res
	}

	if secuirityData.Type != constant.SECURITY_TYPE_STOCK {
		res.SetStatus(http.StatusBadRequest)
		res.SetError(constant.ERROR_CODE_REQUEST_INVALID, "incorrect stock")
		return res
	}

	transactionsData, err := s.mysql.SelectInvertriesTransactionByIdAndAccountId(ctx, request.AccountId, inventoryData.Id)

	if err != nil {
		res.SetStatus(http.StatusInternalServerError)
		res.SetError(constant.ERROR_CODE_INTERNAL_SERVER, "internal server error")
		return res
	}

	var resData []domain.ClientStockInventoryTransactionsResponse

	for _, transactionData := range transactionsData {
		resData = append(resData, domain.ClientStockInventoryTransactionsResponse{
			TransactionId:   transactionData.Id,
			TransactionType: string(transactionData.Type),
			Amount:          transactionData.Price,
			Fee:             transactionData.Fee,
			Quantity:        int(transactionData.Quantity),
			Date:            transactionData.Date.Format("02-01-2006"),
		})
	}

	res.SetData(resData)
	return res
}

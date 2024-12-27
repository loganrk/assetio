package stock

import (
	"assetio/internal/adapters/handler/response"
	"assetio/internal/constant"
	"assetio/internal/domain"
	"context"
	"net/http"
	"time"
)

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
	date := time.Now()

	if request.Date != "" {
		parsedDate, err := time.Parse(constant.DATE_LAYOUT, request.Date)
		if err == nil {
			date = parsedDate
		}
	}

	availableQuanity, err := s.mysql.GetInventoryAvailableQuanitityBySecurityIdAndDate(ctx, request.AccountId, request.StockId, date)

	if err != nil {
		s.logger.Errorw(ctx, "GetInventoryAvailableQuanitityBySecurityIdAndDate failed",
			constant.ERROR_TYPE, constant.ERROR_TYPE_DBEXECUTION,
			constant.ERROR_MESSAGE, err.Error(),
			constant.REQUEST, request,
		)
		res.SetStatus(http.StatusInternalServerError)
		res.SetError(constant.ERROR_CODE_INTERNAL_SERVER, "internal server error")
		return res
	}

	if availableQuanity <= 0 {
		res.SetStatus(http.StatusConflict)
		res.SetError(constant.ERROR_CODE_DATA_EXISTS, "no stocks available to add dividend")
		return res
	}

	// Insert a transaction entry to log the dividend distribution.
	_, err = s.mysql.InsertTransaction(ctx, domain.Transactions{
		AccountId:    request.AccountId,
		SecurityId:   secuirity.Id,
		Type:         domain.DIVIDEND,
		Quantity:     availableQuanity,
		AveragePrice: request.AmountPerQuantity,
		TotalValue:   availableQuanity * request.AmountPerQuantity,
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

	// Set success message and response data.
	resData := domain.ClientStockDividendResponse{
		Message: "stock dividend successfully",
	}

	res.SetData(resData)
	return res
}

func (s *stockUsecase) StockDividends(request domain.ClientStockDividendsRequest) domain.Response {

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

	transactionsData, err := s.mysql.GetDividendTransactionsByAccountIdAndSecurityId(ctx, request.AccountId, request.StockId)

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

	var resData []domain.ClientStockDividendsResponse

	for _, transactionData := range transactionsData {
		resData = append(resData, domain.ClientStockDividendsResponse{
			Quantity: int(transactionData.Quantity),
			Amount:   transactionData.Price,
			Date:     transactionData.Date.Format("02-01-2006"),
		})
	}

	// Set the processed inventory data in the response to be returned to the client.
	res.SetData(resData)
	return res

}

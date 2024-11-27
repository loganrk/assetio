package usecases

import (
	"assetio/internal/adapters/handler/response"
	"assetio/internal/constant"
	"assetio/internal/domain"
	"assetio/internal/port"
	"context"
	"net/http"
)

type securityUsecase struct {
	logger port.Logger
	mysql  port.RepositoryStore
}

func New(loggerIns port.Logger, mysqlIns port.RepositoryStore) domain.SecuritySvr {
	return &securityUsecase{
		mysql:  mysqlIns,
		logger: loggerIns,
	}
}

// SecurityCreate processes the creation of a new security by validating the type and exchange,
// checking if the security already exists, and then inserting a new record into the database.
//
// Parameters:
//   - request: domain.ClientSecurityCreateRequest - contains the details of the security creation request,
//     including the type, exchange, name, and symbol of the security.
//
// Returns:
//   - domain.Response - contains the outcome of the security creation request,
//     either confirming success or detailing any encountered error.
func (s *securityUsecase) SecurityCreate(request domain.ClientSecurityCreateRequest) domain.Response {
	// Create a new background context to manage the request lifecycle.
	ctx := context.Background()

	// Initialize a new response object to hold the result of the request.
	res := response.New()

	// Get the security type based on the provided request type.
	securityType := s.getType(request.Type)
	// If the security type is invalid, return a bad request error.
	if securityType == 0 {
		res.SetStatus(http.StatusBadRequest)
		res.SetError(constant.ERROR_CODE_REQUEST_INVALID, "invalid type")
		return res
	}

	// Get the security exchange based on the provided request exchange.
	securityExchange := s.getExchange(request.Exchange)
	// If the security exchange is invalid, return a bad request error.
	if securityExchange == 0 {
		res.SetStatus(http.StatusBadRequest)
		res.SetError(constant.ERROR_CODE_REQUEST_INVALID, "invalid exchange")
		return res
	}

	// Check if the security already exists in the database by type, exchange, and symbol.
	securityData, err := s.mysql.GetSecurityDataByTypeAndExchangeAndSymbol(ctx, securityType, securityExchange, request.Symbol)

	if err != nil {
		// Log error and return internal server error response if the database query fails.
		s.logger.Errorw(ctx, "GetSecurityDataByTypeAndExchangeAndSymbol failed",
			constant.ERROR_TYPE, constant.ERROR_TYPE_DBEXECUTION,
			constant.ERROR_MESSAGE, err.Error(),
			constant.REQUEST, request,
		)

		res.SetStatus(http.StatusInternalServerError)
		res.SetError(constant.ERROR_CODE_INTERNAL_SERVER, "internal server error")
		return res
	}

	// If the security symbol already exists, return a response indicating that.
	if securityData.Id != 0 {
		res.SetStatus(http.StatusConflict)
		res.SetError(constant.ERROR_CODE_DATA_EXISTS, "symbol already available")
		return res
	}

	// Insert the new security data into the database.
	securityData, err = s.mysql.InsertSecurityData(ctx, domain.Securities{
		Type:     securityType,
		Exchange: securityExchange,
		Name:     request.Name,
		Symbol:   request.Symbol,
	})

	if err != nil {
		// Log error and return internal server error response if the insert fails.
		s.logger.Errorw(ctx, "InsertSecurityData failed",
			constant.ERROR_TYPE, constant.ERROR_TYPE_DBEXECUTION,
			constant.ERROR_MESSAGE, err.Error(),
			constant.REQUEST, request,
		)

		res.SetStatus(http.StatusInternalServerError)
		res.SetError(constant.ERROR_CODE_INTERNAL_SERVER, "internal server error")
		return res
	}

	// Set success response message upon successful security creation.
	resData := domain.ClientSecurityCreateResponse{
		SecurityId: securityData.Id,
		Message:    "security created successfully",
	}

	// Set the response data and return.
	res.SetData(resData)
	return res
}

// SecurityAll retrieves a list of all securities for a specific type and exchange,
// and returns the data in a structured format for the client.
//
// Parameters:
//   - request: domain.ClientSecurityAllRequest - contains the details of the security request,
//     including the type and exchange of the securities to be fetched.
//
// Returns:
//   - domain.Response - contains the list of securities for the specified type and exchange,
//     or an error message if an error occurs during the retrieval process.
func (s *securityUsecase) SecurityAll(request domain.ClientSecurityAllRequest) domain.Response {
	// Create a new background context to manage the request lifecycle.
	ctx := context.Background()

	// Initialize a new response object to hold the result of the request.
	res := response.New()

	// Get the security type based on the provided request type.
	securityType := s.getType(request.Type)
	// If the security type is invalid, return a bad request error.
	if securityType == 0 {
		res.SetStatus(http.StatusBadRequest)
		res.SetError(constant.ERROR_CODE_REQUEST_INVALID, "invalid type")
		return res
	}

	// Retrieve the list of securities data from the database by type
	securitiesData, err := s.mysql.GetSecuritiesDataByType(ctx, securityType)

	if err != nil {
		// Log error and return internal server error response if the database query fails.
		s.logger.Errorw(ctx, "GetSecuritiesDataByType failed",
			constant.ERROR_TYPE, constant.ERROR_TYPE_DBEXECUTION,
			constant.ERROR_MESSAGE, err.Error(),
			constant.REQUEST, request,
		)

		res.SetStatus(http.StatusInternalServerError)
		res.SetError(constant.ERROR_CODE_INTERNAL_SERVER, "internal server error")
		return res
	}

	// If no securities were found, return an empty response.
	if len(securitiesData) == 0 {
		res.SetData(nil)
		return res
	}

	// Prepare the response data by transforming the database data into the response format.
	var resData []domain.ClientSecurityAllResponse
	for _, securityData := range securitiesData {
		resData = append(resData, domain.ClientSecurityAllResponse{
			Id:       securityData.Id,
			Type:     s.getTypeString(securityData.Type),
			Exchange: s.getExchangeString(securityData.Exchange),
			Symbol:   securityData.Symbol,
			Name:     securityData.Name,
		})
	}

	// Set the formatted response data.
	res.SetData(resData)
	return res
}

// SecurityGet retrieves the details of a specific security by its ID.
// It returns the security information if found, or an error if not.
//
// Parameters:
//   - request: domain.ClientSecurityGetRequest - contains the ID of the security to be fetched.
//
// Returns:
//   - domain.Response - contains the details of the security or an error message if the security ID is invalid or the retrieval fails.
func (s *securityUsecase) SecurityGet(request domain.ClientSecurityGetRequest) domain.Response {
	// Create a new background context to manage the request lifecycle.
	ctx := context.Background()

	// Initialize a new response object to hold the result of the request.
	res := response.New()

	// Retrieve the security data by its ID from the database.
	securityData, err := s.mysql.GetSecurityDataById(ctx, request.SecurityId)
	if err != nil {
		// Log error and return internal server error response if the database query fails.
		s.logger.Errorw(ctx, "GetSecurityDataById failed",
			constant.ERROR_TYPE, constant.ERROR_TYPE_DBEXECUTION,
			constant.ERROR_MESSAGE, err.Error(),
			constant.REQUEST, request,
		)

		res.SetStatus(http.StatusInternalServerError)
		res.SetError(constant.ERROR_CODE_INTERNAL_SERVER, "internal server error")
		return res
	}

	// If no security was found with the provided ID, return a bad request error.
	if securityData.Id == 0 {
		res.SetStatus(http.StatusBadRequest)
		res.SetError(constant.ERROR_CODE_REQUEST_INVALID, "invalid security id")
		return res
	}

	// Prepare the response data with the security details.
	resData := domain.ClientSecurityGetResponse{
		Id:       securityData.Id,
		Type:     s.getTypeString(securityData.Type),
		Exchange: s.getExchangeString(securityData.Exchange),
		Symbol:   securityData.Symbol,
		Name:     securityData.Name,
	}

	// Set the response data.
	res.SetData(resData)
	return res
}

// SecuritySearch searches for securities based on the provided type, exchange, and search query.
// It returns a list of matching securities or an error message if the search fails.
//
// Parameters:
//   - request: domain.ClientSecuritySearchRequest - contains the details of the search criteria,
//     including type, exchange, and search query.
//
// Returns:
//   - domain.Response - contains the list of matching securities or an error message if the search fails.
func (s *securityUsecase) SecuritySearch(request domain.ClientSecuritySearchRequest) domain.Response {
	// Create a new background context to manage the request lifecycle.
	ctx := context.Background()

	// Initialize a new response object to hold the result of the request.
	res := response.New()

	// Get the security type based on the provided request type.
	securityType := s.getType(request.Type)
	// If the security type is invalid, return a bad request error.
	if securityType == 0 {
		res.SetStatus(http.StatusBadRequest)
		res.SetError(constant.ERROR_CODE_REQUEST_INVALID, "invalid type")
		return res
	}

	// Get the security exchange based on the provided request exchange.
	securityExchange := s.getExchange(request.Exchange)
	// If the security exchange is invalid, return a bad request error.
	if securityExchange == 0 {
		res.SetStatus(http.StatusBadRequest)
		res.SetError(constant.ERROR_CODE_REQUEST_INVALID, "invalid exchange")
		return res
	}

	// Perform the search for securities matching the provided type, exchange, and search query.
	securitiesData, err := s.mysql.SearchSecuritiesDataByTypeAndExchange(ctx, securityType, securityExchange, request.Search)
	if err != nil {
		// Log error and return internal server error response if the search query fails.
		s.logger.Errorw(ctx, "SearchSecuritiesDataByTypeAndExchange failed",
			constant.ERROR_TYPE, constant.ERROR_TYPE_DBEXECUTION,
			constant.ERROR_MESSAGE, err.Error(),
			constant.REQUEST, request,
		)

		res.SetStatus(http.StatusInternalServerError)
		res.SetError(constant.ERROR_CODE_INTERNAL_SERVER, "internal server error")
		return res
	}

	// If no matching securities were found, return an empty response.
	if len(securitiesData) == 0 {
		res.SetData(nil)
		return res
	}

	// Prepare the response data by transforming the database results into a structured format.
	var resData []domain.ClientSecuritySearchResponse
	for _, securityData := range securitiesData {
		resData = append(resData, domain.ClientSecuritySearchResponse{
			Id:       securityData.Id,
			Type:     s.getTypeString(securityData.Type),
			Exchange: s.getExchangeString(securityData.Exchange),
			Symbol:   securityData.Symbol,
			Name:     securityData.Name,
		})
	}

	// Set the formatted response data.
	res.SetData(resData)
	return res
}

// SecurityUpdate updates the details of an existing security.
// It checks for duplicate symbols and validates the security ID, type, and exchange.
//
// Parameters:
//   - request: domain.ClientSecurityUpdateRequest - contains the security ID and updated details for the security.
//
// Returns:
//   - domain.Response - contains the result of the update operation, including a success message or error.
func (s *securityUsecase) SecurityUpdate(request domain.ClientSecurityUpdateRequest) domain.Response {
	// Create a new background context to manage the request lifecycle.
	ctx := context.Background()

	// Initialize a new response object to hold the result of the request.
	res := response.New()

	// Validate the provided security type.
	securityType := s.getType(request.Type)
	if securityType == 0 {
		res.SetStatus(http.StatusBadRequest)
		res.SetError(constant.ERROR_CODE_REQUEST_INVALID, "invalid type")
		return res
	}

	// Validate the provided security exchange.
	securityExchange := s.getExchange(request.Exchange)
	if securityExchange == 0 {
		res.SetStatus(http.StatusBadRequest)
		res.SetError(constant.ERROR_CODE_REQUEST_INVALID, "invalid exchange")
		return res
	}

	// Retrieve the security data by its ID to ensure it exists in the system.
	securityData, err := s.mysql.GetSecurityDataById(ctx, request.SecurityId)
	if err != nil {
		// Log the error and return a generic internal server error if the retrieval fails.
		s.logger.Errorw(ctx, "GetSecurityDataById failed",
			constant.ERROR_TYPE, constant.ERROR_TYPE_DBEXECUTION,
			constant.ERROR_MESSAGE, err.Error(),
			constant.REQUEST, request,
		)

		res.SetStatus(http.StatusInternalServerError)
		res.SetError(constant.ERROR_CODE_INTERNAL_SERVER, "internal server error")
		return res
	}

	// If the security ID does not exist, return an error.
	if securityData.Id == 0 {
		res.SetStatus(http.StatusBadRequest)
		res.SetError(constant.ERROR_CODE_REQUEST_INVALID, "invalid security id")
		return res
	}

	// Check if a security with the same type, exchange, and symbol already exists.
	securityData, err = s.mysql.GetSecurityDataByTypeAndExchangeAndSymbol(ctx, securityType, securityExchange, request.Symbol)
	if err != nil {
		// Log the error and return a generic internal server error if the database query fails.
		s.logger.Errorw(ctx, "GetSecurityDataByTypeAndExchangeAndSymbol failed",
			constant.ERROR_TYPE, constant.ERROR_TYPE_DBEXECUTION,
			constant.ERROR_MESSAGE, err.Error(),
			constant.REQUEST, request,
		)

		res.SetStatus(http.StatusInternalServerError)
		res.SetError(constant.ERROR_CODE_INTERNAL_SERVER, "internal server error")
		return res
	}

	// If a security with the same symbol already exists but with a different ID, return a message indicating a conflict.
	if securityData.Id != 0 && securityData.Id != request.SecurityId {
		res.SetStatus(http.StatusConflict)
		res.SetError(constant.ERROR_CODE_DATA_EXISTS, "symbol already available")
		return res
	}

	// Update the security data in the database.
	securityData = domain.Securities{
		Type:     securityType,
		Exchange: securityExchange,
		Name:     request.Name,
		Symbol:   request.Symbol,
	}

	err = s.mysql.UpdateSecurityData(ctx, request.SecurityId, securityData)
	if err != nil {
		// Log the error and return a generic internal server error if the update fails.
		s.logger.Errorw(ctx, "UpdateSecurityData failed",
			constant.ERROR_TYPE, constant.ERROR_TYPE_DBEXECUTION,
			constant.ERROR_MESSAGE, err.Error(),
			constant.REQUEST, request,
		)

		res.SetStatus(http.StatusInternalServerError)
		res.SetError(constant.ERROR_CODE_INTERNAL_SERVER, "internal server error")
		return res
	}

	// Return a success message indicating that the security has been updated successfully.
	resData := domain.ClientSecurityUpdateResponse{
		Message: "security updated successfully",
	}

	res.SetData(resData)
	return res
}

// getType converts a string representation of a security type to its corresponding integer constant.
// Returns the corresponding integer constant for stock security type or 0 if invalid.
func (s *securityUsecase) getType(typeData string) int {
	// Check if the provided type is "stock" and return the corresponding constant
	if typeData == constant.SECURITY_TYPE_STOCK_STRING {
		return constant.SECURITY_TYPE_STOCK
	}
	// Return 0 if the type is invalid
	return 0
}

// getExchange converts a string representation of an exchange to its corresponding integer constant.
// Returns the corresponding constant for NSE or BSE exchange or 0 if invalid.
func (s *securityUsecase) getExchange(exchange string) int {
	// Check if the provided exchange is "NSE" and return the corresponding constant
	if exchange == constant.EXCHANGE_TYPE_NSE_STRING {
		return constant.EXCHANGE_TYPE_NSE
	} else if exchange == constant.EXCHANGE_TYPE_BSE_STRING {
		// Check if the provided exchange is "BSE" and return the corresponding constant
		return constant.EXCHANGE_TYPE_BSE
	}
	// Return 0 if the exchange is invalid
	return 0
}

// getTypeString converts an integer security type constant to its corresponding string representation.
// Returns the string representation of a stock security type or an empty string if invalid.
func (s *securityUsecase) getTypeString(typeData int) string {
	// Check if the provided type matches the stock security type and return the corresponding string
	if typeData == constant.SECURITY_TYPE_STOCK {
		return constant.SECURITY_TYPE_STOCK_STRING
	}
	// Return an empty string if the type is invalid
	return ""
}

// getExchangeString converts an integer exchange constant to its corresponding string representation.
// Returns the string representation of NSE or BSE exchange or an empty string if invalid.
func (s *securityUsecase) getExchangeString(exchange int) string {
	// Check if the provided exchange matches NSE and return the corresponding string
	if exchange == constant.EXCHANGE_TYPE_NSE {
		return constant.EXCHANGE_TYPE_NSE_STRING
	} else if exchange == constant.EXCHANGE_TYPE_BSE {
		// Check if the provided exchange matches BSE and return the corresponding string
		return constant.EXCHANGE_TYPE_BSE_STRING
	}
	// Return an empty string if the exchange is invalid
	return ""
}

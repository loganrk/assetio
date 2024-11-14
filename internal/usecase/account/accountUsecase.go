package account

import (
	"assetio/internal/adapters/handler/response"
	"assetio/internal/constant"
	"assetio/internal/domain"
	"assetio/internal/port"
	"context"
	"net/http"
)

type accountUsecase struct {
	logger port.Logger
	mysql  port.RepositoryStore
}

func New(loggerIns port.Logger, mysqlIns port.RepositoryStore) domain.AccountSvr {
	return &accountUsecase{
		mysql:  mysqlIns,
		logger: loggerIns,
	}
}

// AccountCreate processes the creation of a new account by inserting account data
// into the database and returning a response indicating the result of the operation.
//
// Parameters:
//   - request: domain.ClientAccountCreateRequest - contains the details of the account creation request,
//     including the account name, user ID, and status.
//
// Returns:
//   - domain.Response - contains the outcome of the account creation request,
//     either confirming success or detailing any encountered error.
func (a *accountUsecase) AccountCreate(request domain.ClientAccountCreateRequest) domain.Response {
	// Create a new background context to manage the request lifecycle.
	ctx := context.Background()

	// Initialize a new response object to hold the result of the request.
	res := response.New()

	// Insert the new account data into the database.
	_, err := a.mysql.InsertAccountData(ctx, domain.Accounts{
		Name:   request.Name,                   // Assign the account name from the request.
		UserId: request.UserId,                 // Assign the user ID from the request.
		Status: constant.ACCOUNT_STATUS_ACTIVE, // Set the account status to active by default.
	})

	// If there is an error during the database insert operation.
	if err != nil {
		// Log the error along with details about the request and error message.
		a.logger.Errorw(ctx, "InsertAccountData failed",
			constant.ERROR_TYPE, constant.ERROR_TYPE_DBEXECUTION,
			constant.ERROR_MESSAGE, err.Error(),
			constant.REQUEST, request,
		)

		// Set the response status to internal server error.
		res.SetStatus(http.StatusInternalServerError)
		res.SetError(constant.ERROR_CODE_INTERNAL_SERVER, "internal server error")
		return res
	}

	// If the account is created successfully, set the success message.
	resData := domain.ClientAccountCreateResponse{
		Message: "account created successfully", // Inform the user that the account was created.
	}

	// Set the response data and return.
	res.SetData(resData)
	return res
}

// AccountAll retrieves all accounts for a given user by querying the database and
// returns a response containing the list of accounts or an error if the operation fails.
//
// Parameters:
//   - request: domain.ClientAccountAllRequest - contains the user ID to fetch their associated accounts.
//
// Returns:
//   - domain.Response - contains the outcome of the account retrieval request,
//     either returning a list of accounts or detailing any encountered error.
func (a *accountUsecase) AccountAll(request domain.ClientAccountAllRequest) domain.Response {

	// Create a new background context to manage the request lifecycle.
	ctx := context.Background()

	// Initialize a new response object to hold the result of the request.
	res := response.New()

	// Fetch all account data from the database for the specified user ID.
	accounts, err := a.mysql.GetAccountsData(ctx, request.UserId)

	// If an error occurs while retrieving the accounts, log the error and return an internal server error response.
	if err != nil {
		a.logger.Errorw(ctx, "GetAccountsData failed",
			constant.ERROR_TYPE, constant.ERROR_TYPE_DBEXECUTION,
			constant.ERROR_MESSAGE, err.Error(),
			constant.REQUEST, request,
		)

		// Set the response status to internal server error.
		res.SetStatus(http.StatusInternalServerError)
		res.SetError(constant.ERROR_CODE_INTERNAL_SERVER, "internal server error")
		return res
	}

	// If no accounts are found for the user, return a nil response data.
	if len(accounts) == 0 {
		res.SetData(nil) // No accounts found, so the response data is set to nil.
		return res
	}

	// Create a slice to hold the formatted response data.
	var resData []domain.ClientAccountAllResponse

	// Iterate through the accounts and format the data for the response.
	for _, account := range accounts {
		resData = append(resData, domain.ClientAccountAllResponse{
			Id:     account.Id,                        // Include the account ID in the response.
			Name:   account.Name,                      // Include the account name in the response.
			Status: a.getStatusString(account.Status), // Get the status string and include it in the response.
		})
	}

	// Set the response data containing the formatted account information.
	res.SetData(resData)
	return res
}

// AccountGet retrieves a specific account for a given user by querying the database
// using the provided account ID and user ID, and returns the account details or an error.
//
// Parameters:
//   - request: domain.ClientAccountGetRequest - contains the account ID and user ID to fetch the specified account.
//
// Returns:
//   - domain.Response - contains the outcome of the account retrieval request,
//     either returning the account details or detailing any encountered error.
func (a *accountUsecase) AccountGet(request domain.ClientAccountGetRequest) domain.Response {
	// Create a new background context to manage the request lifecycle.
	ctx := context.Background()

	// Initialize a new response object to hold the result of the request.
	res := response.New()

	// Fetch the account data from the database by account ID and user ID.
	account, err := a.mysql.GetAccountDataByIdAndUserId(ctx, request.AccountId, request.UserId)

	// If an error occurs while retrieving the account, log the error and return an internal server error response.
	if err != nil {
		a.logger.Errorw(ctx, "GetAccountDataByIdAndUserId failed",
			constant.ERROR_TYPE, constant.ERROR_TYPE_DBEXECUTION,
			constant.ERROR_MESSAGE, err.Error(),
			constant.REQUEST, request,
		)

		// Set the response status to internal server error.
		res.SetStatus(http.StatusInternalServerError)
		res.SetError(constant.ERROR_CODE_INTERNAL_SERVER, "internal server error")
		return res
	}

	// If no account is found (i.e., account ID is 0), return a bad request error response.
	if account.Id == 0 {
		res.SetStatus(http.StatusBadRequest)
		res.SetError(constant.ERROR_CODE_REQUEST_INVALID, "incorrect account id") // Inform the user that the account ID is invalid.
		return res
	}

	// Create the response data object with the account details.
	resData := domain.ClientAccountGetResponse{
		Id:     account.Id,                        // Include the account ID in the response.
		Name:   account.Name,                      // Include the account name in the response.
		Status: a.getStatusString(account.Status), // Convert the account status to a string and include it in the response.
	}

	// Set the response data containing the account details.
	res.SetData(resData)
	return res
}

// AccountActivate activates a specific account for a given user by updating the account's status in the database,
// and returns a response indicating whether the account was successfully activated or if an error occurred.
//
// Parameters:
//   - request: domain.ClientAccountActivateRequest - contains the account ID and user ID to activate the account.
//
// Returns:
//   - domain.Response - contains the outcome of the account activation request,
//     either confirming success, stating the account is already active, or detailing any encountered error.
func (a *accountUsecase) AccountActivate(request domain.ClientAccountActivateRequest) domain.Response {
	// Create a new background context to manage the request lifecycle.
	ctx := context.Background()

	// Initialize a new response object to hold the result of the request.
	res := response.New()

	// Fetch the account data from the database by account ID and user ID.
	account, err := a.mysql.GetAccountDataByIdAndUserId(ctx, request.AccountId, request.UserId)

	// If an error occurs while retrieving the account, log the error and return an internal server error response.
	if err != nil {
		a.logger.Errorw(ctx, "GetAccountDataByIdAndUserId failed",
			constant.ERROR_TYPE, constant.ERROR_TYPE_DBEXECUTION,
			constant.ERROR_MESSAGE, err.Error(),
			constant.REQUEST, request,
		)
		// Set the response status to internal server error.
		res.SetStatus(http.StatusInternalServerError)
		res.SetError(constant.ERROR_CODE_INTERNAL_SERVER, "internal server error")
		return res
	}

	// If no account is found (i.e., account ID is 0), return a bad request error response.
	if account.Id == 0 {
		res.SetStatus(http.StatusBadRequest)
		res.SetError(constant.ERROR_CODE_REQUEST_INVALID, "incorrect account id") // Inform the user that the account ID is invalid.
		return res
	}

	// If the account is already active, return a response indicating no action is needed.
	if account.Status == constant.ACCOUNT_STATUS_ACTIVE {
		resData := domain.ClientAccountActivateResponse{
			Message: "account already active", // Inform the user that the account is already active.
		}
		res.SetData(resData)
		return res
	}

	// Prepare the account data to update the status to active.
	accountData := domain.Accounts{
		Status: constant.ACCOUNT_STATUS_ACTIVE, // Set the account status to active.
	}

	// Update the account data in the database.
	err = a.mysql.UpdateAccountData(ctx, request.AccountId, request.UserId, accountData)

	// If an error occurs while updating the account, log the error and return an internal server error response.
	if err != nil {
		a.logger.Errorw(ctx, "UpdateAccountData failed",
			constant.ERROR_TYPE, constant.ERROR_TYPE_DBEXECUTION,
			constant.ERROR_MESSAGE, err.Error(),
			constant.REQUEST, request,
		)

		// Set the response status to internal server error.
		res.SetStatus(http.StatusInternalServerError)
		res.SetError(constant.ERROR_CODE_INTERNAL_SERVER, "internal server error")
		return res
	}

	// If the account is successfully activated, set the success message.
	resData := domain.ClientAccountActivateResponse{
		Message: "account activated successfully", // Inform the user that the account was activated successfully.
	}

	// Set the response data containing the activation message.
	res.SetData(resData)
	return res
}

// AccountInactivate inactivates a specific account for a given user by updating the account's status in the database,
// and returns a response indicating whether the account was successfully inactivated or if an error occurred.
//
// Parameters:
//   - request: domain.ClientAccountInactivateRequest - contains the account ID and user ID to inactivate the account.
//
// Returns:
//   - domain.Response - contains the outcome of the account inactivation request,
//     either confirming success, stating the account is already inactive, or detailing any encountered error.
func (a *accountUsecase) AccountInactivate(request domain.ClientAccountInactivateRequest) domain.Response {
	// Create a new background context to manage the request lifecycle.
	ctx := context.Background()

	// Initialize a new response object to hold the result of the request.
	res := response.New()

	// Fetch the account data from the database by account ID and user ID.
	account, err := a.mysql.GetAccountDataByIdAndUserId(ctx, request.AccountId, request.UserId)

	// If an error occurs while retrieving the account, log the error and return an internal server error response.
	if err != nil {
		a.logger.Errorw(ctx, "GetAccountDataByIdAndUserId failed",
			constant.ERROR_TYPE, constant.ERROR_TYPE_DBEXECUTION,
			constant.ERROR_MESSAGE, err.Error(),
			constant.REQUEST, request,
		)

		// Set the response status to internal server error.
		res.SetStatus(http.StatusInternalServerError)
		res.SetError(constant.ERROR_CODE_INTERNAL_SERVER, "internal server error")
		return res
	}

	// If no account is found (i.e., account ID is 0), return a bad request error response.
	if account.Id == 0 {
		res.SetStatus(http.StatusBadRequest)
		res.SetError(constant.ERROR_CODE_REQUEST_INVALID, "incorrect account id") // Inform the user that the account ID is invalid.
		return res
	}

	// If the account is already inactive, return a response indicating no action is needed.
	if account.Status == constant.ACCOUNT_STATUS_INACTIVE {
		resData := domain.ClientAccountInActivateResponse{
			Message: "account already inactive", // Inform the user that the account is already inactive.
		}
		res.SetData(resData)
		return res
	}

	// Prepare the account data to update the status to inactive.
	accountData := domain.Accounts{
		Status: constant.ACCOUNT_STATUS_INACTIVE, // Set the account status to inactive.
	}

	// Update the account data in the database.
	err = a.mysql.UpdateAccountData(ctx, request.AccountId, request.UserId, accountData)

	// If an error occurs while updating the account, log the error and return an internal server error response.
	if err != nil {
		a.logger.Errorw(ctx, "UpdateAccountData failed",
			constant.ERROR_TYPE, constant.ERROR_TYPE_DBEXECUTION,
			constant.ERROR_MESSAGE, err.Error(),
			constant.REQUEST, request,
		)

		// Set the response status to internal server error.
		res.SetStatus(http.StatusInternalServerError)
		res.SetError(constant.ERROR_CODE_INTERNAL_SERVER, "internal server error")
		return res
	}

	// If the account is successfully inactivated, set the success message.
	resData := domain.ClientAccountInActivateResponse{
		Message: "account inactivated successfully", // Inform the user that the account was inactivated successfully.
	}

	// Set the response data containing the inactivation message.
	res.SetData(resData)
	return res
}

// AccountUpdate updates the details of a specific account for a given user by modifying the account's information in the database,
// and returns a response indicating whether the account was successfully updated or if an error occurred.
//
// Parameters:
//   - request: domain.ClientAccountUpdateRequest - contains the account ID, user ID, and updated account details (e.g., name).
//
// Returns:
//   - domain.Response - contains the outcome of the account update request,
//     either confirming success or detailing any encountered error.
func (a *accountUsecase) AccountUpdate(request domain.ClientAccountUpdateRequest) domain.Response {
	// Create a new background context to manage the request lifecycle.
	ctx := context.Background()

	// Initialize a new response object to hold the result of the request.
	res := response.New()

	// Prepare the account data to update the account's name.
	accountData := domain.Accounts{
		Name: request.Name, // Set the account's name from the request data.
	}

	// Update the account data in the database.
	err := a.mysql.UpdateAccountData(ctx, request.AccountId, request.UserId, accountData)

	// If an error occurs while updating the account, log the error and return an internal server error response.
	if err != nil {
		a.logger.Errorw(ctx, "UpdateAccountData failed",
			constant.ERROR_TYPE, constant.ERROR_TYPE_DBEXECUTION,
			constant.ERROR_MESSAGE, err.Error(),
			constant.REQUEST, request,
		)

		// Set the response status to internal server error.
		res.SetStatus(http.StatusInternalServerError)
		res.SetError(constant.ERROR_CODE_INTERNAL_SERVER, "internal server error")
		return res
	}

	// If the account is successfully updated, set the success message.
	resData := domain.ClientAccountUpdateResponse{
		Message: "account updated successfully", // Inform the user that the account was updated successfully.
	}

	// Set the response data containing the update message.
	res.SetData(resData)
	return res
}

// getStatusString returns a human-readable string representation of an account's status based on its integer value.
func (a *accountUsecase) getStatusString(status int) string {
	// If the account status is active, return the string "active".
	if status == constant.ACCOUNT_STATUS_ACTIVE {
		return "active"
	} else if status == constant.ACCOUNT_STATUS_INACTIVE { // If the account status is inactive, return the string "inactive".
		return "inactive"
	}

	// If the account status is neither active nor inactive, return "unknown".
	return "unkown"
}

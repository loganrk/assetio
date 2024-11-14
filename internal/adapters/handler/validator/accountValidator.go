package validator

import (
	"assetio/internal/domain"
	"errors"
)

// AccountCreate validates the fields in the ClientAccountCreateRequest object before creating an account.
// It checks if the required fields (Name, UserId) are valid (non-zero or non-empty).
func (v validation) AccountCreate(request domain.ClientAccountCreateRequest) error {
	if request.Name == "" {
		return errors.New("invalid name") // Name must be non-empty
	}

	if request.UserId == 0 {
		return errors.New("invalid user id") // UserId must be non-zero
	}
	return nil // Return nil if all validations pass
}

// AccountAll validates the fields in the ClientAccountAllRequest object before fetching all accounts.
// It checks if the UserId is valid (non-zero).
func (v validation) AccountAll(request domain.ClientAccountAllRequest) error {

	if request.UserId == 0 {
		return errors.New("invalid user id") // UserId must be non-zero
	}

	return nil // Return nil if validation passes
}

// AccountGet validates the fields in the ClientAccountGetRequest object before fetching account details.
// It checks if the required fields (AccountId, UserId) are valid (non-zero).
func (v validation) AccountGet(request domain.ClientAccountGetRequest) error {
	if request.AccountId == 0 {
		return errors.New("invalid account id ") // AccountId must be non-zero
	}

	if request.UserId == 0 {
		return errors.New("invalid user id") // UserId must be non-zero
	}

	return nil // Return nil if all validations pass
}

// AccountUpdate validates the fields in the ClientAccountUpdateRequest object before updating account details.
// It checks if the required fields (AccountId, UserId) are valid (non-zero).
func (v validation) AccountUpdate(request domain.ClientAccountUpdateRequest) error {
	if request.AccountId == 0 {
		return errors.New("invalid account id") // AccountId must be non-zero
	}
	if request.UserId == 0 {
		return errors.New("invalid user id") // UserId must be non-zero
	}
	return nil // Return nil if all validations pass
}

// AccountActivate validates the fields in the ClientAccountActivateRequest object before activating an account.
// It checks if the required fields (AccountId, UserId) are valid (non-zero).
func (v validation) AccountActivate(request domain.ClientAccountActivateRequest) error {
	if request.AccountId == 0 {
		return errors.New("invalid account id ") // AccountId must be non-zero
	}

	if request.UserId == 0 {
		return errors.New("invalid user id") // UserId must be non-zero
	}

	return nil // Return nil if all validations pass
}

// AccountInactivate validates the fields in the ClientAccountInactivateRequest object before inactivating an account.
// It checks if the required fields (AccountId, UserId) are valid (non-zero).
func (v validation) AccountInactivate(request domain.ClientAccountInactivateRequest) error {

	if request.AccountId == 0 {
		return errors.New("invalid account id ") // AccountId must be non-zero
	}

	if request.UserId == 0 {
		return errors.New("invalid user id") // UserId must be non-zero
	}

	return nil // Return nil if all validations pass
}

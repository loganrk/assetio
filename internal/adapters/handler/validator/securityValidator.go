package validator

import (
	"assetio/internal/domain"
	"errors"
)

// SecurityCreate validates the fields in the ClientSecurityCreateRequest object before creating a security.
// It checks if the required fields (Name, Symbol, Exchange, Type) are valid (non-empty).
func (v validation) SecurityCreate(request domain.ClientSecurityCreateRequest) error {
	if request.Name == "" {
		return errors.New("invalid name") // Name must be non-empty
	}

	if request.Symbol == "" {
		return errors.New("invalid symbol") // Symbol must be non-empty
	}

	if request.Exchange == "" {
		return errors.New("invalid exchange") // Exchange must be non-empty
	}
	if request.Type == "" {
		return errors.New("invalid type") // Type must be non-empty
	}

	return nil // Return nil if all validations pass
}

// SecurityUpdate validates the fields in the ClientSecurityUpdateRequest object before updating a security.
// It checks if the required fields (SecurityId, Name, Symbol, Exchange, Type) are valid (non-zero or non-empty).
func (v validation) SecurityUpdate(request domain.ClientSecurityUpdateRequest) error {
	if request.SecurityId == 0 {
		return errors.New("invalid security id") // SecurityId must be non-zero
	}
	if request.Name == "" {
		return errors.New("invalid name") // Name must be non-empty
	}

	if request.Symbol == "" {
		return errors.New("invalid symbol") // Symbol must be non-empty
	}

	if request.Exchange == "" {
		return errors.New("invalid exchange") // Exchange must be non-empty
	}
	if request.Type == "" {
		return errors.New("invalid type") // Type must be non-empty
	}

	return nil // Return nil if all validations pass
}

// SecurityAll validates the fields in the ClientSecurityAllRequest object before fetching all securities.
// It checks if the required fields (Exchange, Type) are valid (non-empty).
func (v validation) SecurityAll(request domain.ClientSecurityAllRequest) error {

	if request.Type == "" {
		return errors.New("invalid type") // Type must be non-empty
	}

	return nil // Return nil if all validations pass
}

// SecurityGet validates the fields in the ClientSecurityGetRequest object before retrieving a specific security.
// It checks if the required field (SecurityId) is valid (non-zero).
func (v validation) SecurityGet(request domain.ClientSecurityGetRequest) error {

	if request.SecurityId == 0 {
		return errors.New("invalid security id") // SecurityId must be non-zero
	}

	return nil // Return nil if validation passes
}

// SecuritySearch validates the fields in the ClientSecuritySearchRequest object before searching for securities.
// It checks if the required fields (Exchange, Type, Search) are valid (non-empty) and ensures the search keyword has at least 2 characters.
func (v validation) SecuritySearch(request domain.ClientSecuritySearchRequest) error {
	if request.Exchange == "" {
		return errors.New("invalid exchange") // Exchange must be non-empty
	}
	if request.Type == "" {
		return errors.New("invalid type") // Type must be non-empty
	}

	if request.Search == "" {
		return errors.New("invalid search keyword") // Search keyword must be non-empty
	}

	if len(request.Search) <= 1 {
		return errors.New("search keyword should be minimum 2 letters") // Search keyword should have at least 2 characters
	}

	return nil // Return nil if all validations pass
}

package validator

import (
	"assetio/internal/domain"
	"errors"
)

func (v validation) SecurityCreate(request domain.ClientSecurityCreateRequest) error {
	if request.Name == "" {
		return errors.New("invalid name")
	}

	if request.Symbol == "" {
		return errors.New("invalid symbol")
	}

	if request.Exchange == "" {
		return errors.New("invalid exchange")
	}
	if request.Type == "" {
		return errors.New("invalid type")
	}

	return nil
}
func (v validation) SecurityUpdate(request domain.ClientSecurityUpdateRequest) error {
	if request.SecurityId == 0 {
		return errors.New("invalid security id")
	}
	if request.Name == "" {
		return errors.New("invalid name")
	}

	if request.Symbol == "" {
		return errors.New("invalid symbol")
	}

	if request.Exchange == "" {
		return errors.New("invalid exchange")
	}
	if request.Type == "" {
		return errors.New("invalid type")
	}

	return nil
}
func (v validation) SecurityAll(request domain.ClientSecurityAllRequest) error {
	if request.Exchange == "" {
		return errors.New("invalid exchange")
	}
	if request.Type == "" {
		return errors.New("invalid type")
	}

	return nil
}
func (v validation) SecurityGet(request domain.ClientSecurityGetRequest) error {

	if request.SecurityId == 0 {
		return errors.New("invalid security id")
	}

	return nil
}
func (v validation) SecuritySearch(request domain.ClientSecuritySearchRequest) error {
	if request.Exchange == "" {
		return errors.New("invalid exchange")
	}
	if request.Type == "" {
		return errors.New("invalid type")
	}

	if request.Search == "" {
		return errors.New("invalid search keyword")
	}

	if len(request.Search) <= 1 {
		return errors.New("search keyword should be minimum 2 letters")
	}

	return nil
}

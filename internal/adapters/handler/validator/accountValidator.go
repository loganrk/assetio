package validator

import (
	"assetio/internal/domain"
	"errors"
)

func (v validation) AccountCreate(request domain.ClientAccountCreateRequest) error {
	if request.Name == "" {
		return errors.New("invalid name")
	}

	if request.UserId == 0 {
		return errors.New("invalid user id")
	}
	return nil
}

func (v validation) AccountAll(request domain.ClientAccountAllRequest) error {

	if request.UserId == 0 {
		return errors.New("invalid user id")
	}

	return nil
}

func (v validation) AccountGet(request domain.ClientAccountGetRequest) error {
	if request.AccountId == 0 {
		return errors.New("invalid account id ")
	}

	if request.UserId == 0 {
		return errors.New("invalid user id")
	}

	return nil
}

func (v validation) AccountUpdate(request domain.ClientAccountUpdateRequest) error {
	if request.AccountId == 0 {
		return errors.New("invalid account id")
	}
	if request.UserId == 0 {
		return errors.New("invalid user id")
	}
	return nil
}

func (v validation) AccountActivate(request domain.ClientAccountActivateRequest) error {
	if request.AccountId == 0 {
		return errors.New("invalid account id ")
	}

	if request.UserId == 0 {
		return errors.New("invalid user id")
	}

	return nil
}

func (v validation) AccountInactivate(request domain.ClientAccountInactivateRequest) error {

	if request.AccountId == 0 {
		return errors.New("invalid account id ")
	}

	if request.UserId == 0 {
		return errors.New("invalid user id")
	}

	return nil
}

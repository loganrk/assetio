package validator

import "assetio/internal/port"

type validation struct {
}

func New() port.Validator {
	return validation{}

}

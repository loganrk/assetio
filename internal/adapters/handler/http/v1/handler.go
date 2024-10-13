package v1

import (
	"assetio/internal/domain"
	"assetio/internal/port"
)

const (
	ERROR_CODE_INTERNAL_SERVER = "SE01"

	ERROR_CODE_REQUEST_INVALID        = "RE01"
	ERROR_CODE_REQUEST_PARAMS_INVALID = "RE02"

	ERROR_CODE_NO_DATA = "DA01"
)

type handler struct {
	usecases domain.List
	logger   port.Logger
}

func New(loggerIns port.Logger, svcList domain.List) port.Handler {
	return &handler{
		usecases: svcList,
		logger:   loggerIns,
	}
}

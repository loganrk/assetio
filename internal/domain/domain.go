package domain

import (
	"context"
)

type List struct {
	Account  AccountSvr
	Security SecuritySvr
	Stock    StockSvr
}

type AccountSvr interface {
	CreateAccount(ctx context.Context, userId int, name string) error
	GetAccounts(ctx context.Context, userId int) ([]Account, error)
	GetAccount(ctx context.Context, accountId, userId int) (Account, error)
	UpdateAccount(ctx context.Context, accountId, userId int, name string) error
	GetStatusString(status int) string
	AccountActivate(ctx context.Context, accountId, userId int) error
	AccountInactivate(ctx context.Context, accountId, userId int) error
}

type SecuritySvr interface {
	GetType(typeData string) int
	GetExchange(exchange string) int
	GetTypeString(typeData int) string
	GetExchangeString(exchange int) string

	CreateSecuriry(ctx context.Context, types, exchange int, symbol, name string) error
	GetSecuriry(ctx context.Context, types, exchange int, symbol string) (Security, error)
	GetSecuriryById(ctx context.Context, securityId int) (Security, error)
	UpdateSecuriry(ctx context.Context, securityId, types, exchange int, symbol string, name string) error
	GetSecurities(ctx context.Context, types, exchange int) ([]Security, error)
	SearchSecurities(ctx context.Context, types, exchange int, search string) ([]Security, error)
}

type StockSvr interface {
	BuyStock(ctx context.Context, userId int, accountId int, inventoryId int, stockId int, quantity int, amountPerQuantity float64, taxAmount float64) error
	SellStock(ctx context.Context, userId int, accountId int, inventoryId int, stockId int, quantity int, amountPerQuantity float64, taxAmount float64) error
	StockDividendAdd(ctx context.Context, userId int, accountId int, inventoryId int, stockId int, quantity int, amountPerQuantity float64) error
}

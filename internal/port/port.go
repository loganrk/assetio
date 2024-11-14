package port

import (
	"assetio/internal/domain"
	"context"
	"net/http"
	"time"
)

type Handler interface {
	AccountCreate(w http.ResponseWriter, r *http.Request)
	AccountAll(w http.ResponseWriter, r *http.Request)
	AccountGet(w http.ResponseWriter, r *http.Request)
	AccountUpdate(w http.ResponseWriter, r *http.Request)
	AccountActivate(w http.ResponseWriter, r *http.Request)
	AccountInactivate(w http.ResponseWriter, r *http.Request)

	SecurityCreate(w http.ResponseWriter, r *http.Request)
	SecurityUpdate(w http.ResponseWriter, r *http.Request)
	SecurityAll(w http.ResponseWriter, r *http.Request)
	SecurityGet(w http.ResponseWriter, r *http.Request)
	SecuritySearch(w http.ResponseWriter, r *http.Request)

	StockBuy(w http.ResponseWriter, r *http.Request)
	StockSell(w http.ResponseWriter, r *http.Request)
	StockSplit(w http.ResponseWriter, r *http.Request)
	StockDividendAdd(w http.ResponseWriter, r *http.Request)
	StockSummary(w http.ResponseWriter, r *http.Request)
	StockInventories(w http.ResponseWriter, r *http.Request)
	StockInventoryLedgers(w http.ResponseWriter, r *http.Request)

	MutualFundBuy(w http.ResponseWriter, r *http.Request)
	MutualFundAdd(w http.ResponseWriter, r *http.Request)
	MutualFundSell(w http.ResponseWriter, r *http.Request)
	MutualFundSummary(w http.ResponseWriter, r *http.Request)
	MutualFundInventory(w http.ResponseWriter, r *http.Request)
	MutualFundTransaction(w http.ResponseWriter, r *http.Request)
}

type Validator interface {
	AccountCreate(request domain.ClientAccountCreateRequest) error
	AccountAll(request domain.ClientAccountAllRequest) error
	AccountGet(request domain.ClientAccountGetRequest) error
	AccountUpdate(request domain.ClientAccountUpdateRequest) error
	AccountActivate(request domain.ClientAccountActivateRequest) error
	AccountInactivate(request domain.ClientAccountInactivateRequest) error

	SecurityCreate(request domain.ClientSecurityCreateRequest) error
	SecurityUpdate(request domain.ClientSecurityUpdateRequest) error
	SecurityAll(request domain.ClientSecurityAllRequest) error
	SecurityGet(request domain.ClientSecurityGetRequest) error
	SecuritySearch(request domain.ClientSecuritySearchRequest) error

	StockBuy(request domain.ClientStockBuyRequest) error
	StockSell(request domain.ClientStockSellRequest) error
	StockSplit(request domain.ClientStockSplitRequest) error
	StockDividendAdd(request domain.ClientStockDividendAddRequest) error
	StockSummary(request domain.ClientStockSummaryRequest) error
	StockInventories(request domain.ClientStockInventoriesRequest) error
	StockInventoryLedgers(request domain.ClientStockInventoryLedgersRequest) error

	MutualFundBuy(request domain.ClientMutualFundBuyRequest) error
	MutualFundAdd(request domain.ClientMutualFundAddRequest) error
	MutualFundSell(request domain.ClientMutualFundSellRequest) error
	MutualFundSummary(request domain.ClientMutualFundSummaryRequest) error
	MutualFundInventory(request domain.ClientMutualFundInventoryRequest) error
	MutualFundInventoryLedgers(request domain.ClientMutualFundInventoryLedgersRequest) error
}

type RepositoryStore interface {
	AutoMigrate()
	InsertAccountData(ctx context.Context, accountData domain.Accounts) (domain.Accounts, error)
	GetAccountDataByIdAndUserId(ctx context.Context, accountId int, userId int) (domain.Accounts, error)
	GetAccountsData(ctx context.Context, userId int) ([]domain.Accounts, error)
	UpdateAccountData(ctx context.Context, accountId, userId int, accountData domain.Accounts) error

	InsertSecurityData(ctx context.Context, securityData domain.Securities) (domain.Securities, error)
	GetSecurityDataById(ctx context.Context, securityId int) (domain.Securities, error)
	GetSecurityDataByTypeAndExchangeAndSymbol(ctx context.Context, types, exchange int, symbol string) (domain.Securities, error)
	UpdateSecurityData(ctx context.Context, securityId int, securityData domain.Securities) error
	GetSecuritiesDataByExchange(ctx context.Context, types, exchange int) ([]domain.Securities, error)
	SearchSecuritiesDataByTypeAndExchange(ctx context.Context, types, exchange int, search string) ([]domain.Securities, error)

	InsertInventoryLedger(ctx context.Context, inventoryLedgerData domain.InventoryLedger) (domain.InventoryLedger, error)
	UpdateInventoryDataById(ctx context.Context, inventoryId int, inventoryData domain.Inventories) error
	InsertTransaction(ctx context.Context, transactionData domain.Transactions) (domain.Transactions, error)
	UpdateInventoryLedgerTransactionIdById(ctx context.Context, ledgerId, transactionId int) error
	UpdateInventoryLedgerTransactionIdByIds(ctx context.Context, ledgerIds []int, transactionId int) error

	GetInvertriesSummaryByAccountIdAndSecurityType(ctx context.Context, accountId, securityType int) ([]domain.InventorySummary, error)
	GetInvertriesByAccountIdAndSecurityId(ctx context.Context, accountId, securityId int) ([]domain.InventoryDetails, error)

	InsertTransactionData(ctx context.Context, transactionData domain.Transactions) (domain.Transactions, error)
	InsertInventoryData(ctx context.Context, inventoryData domain.Inventories) (domain.Inventories, error)
	GetInventoryDataById(ctx context.Context, inventoryId int) (domain.Inventories, error)
	UpdateAvailableQuanityToInventoryById(ctx context.Context, inventoryId int, quantity float64) error
	GetActiveInventoriesByAccountIdAndSecurityId(ctx context.Context, accountId, securityId int) ([]domain.Inventories, error)
	GetInventoryLedgersByInventoryIdAndAccountId(ctx context.Context, accountId, inventoryId int) ([]domain.InventoryLedgers, error)
}

type Router interface {
	RegisterRoute(method, path string, handlerFunc http.HandlerFunc)
	StartServer(port string) error
	UseBefore(middlewares ...http.Handler)
	NewGroup(groupName string) RouterGroup
}

type RouterGroup interface {
	RegisterRoute(method, path string, handlerFunc http.HandlerFunc)
	UseBefore(middlewares ...http.Handler)
}

type Cipher interface {
	Encrypt(text string) (string, error)
	Decrypt(cryptoText string) (string, error)
	GetKey() string
}

type Token interface {
	GetAccessTokenData(encryptedToken string) (int, time.Time, error)
}

type Auth interface {
	ValidateApiKey() http.Handler
	ValidateAccessToken() http.Handler
}

type Logger interface {
	Debug(ctx context.Context, messages ...any)
	Info(ctx context.Context, messages ...any)
	Warn(ctx context.Context, messages ...any)
	Error(ctx context.Context, messages ...any)
	Fatal(ctx context.Context, messages ...any)
	Debugf(ctx context.Context, template string, args ...any)
	Infof(ctx context.Context, template string, args ...any)
	Warnf(ctx context.Context, template string, args ...any)
	Errorf(ctx context.Context, template string, args ...any)
	Fatalf(ctx context.Context, template string, args ...any)
	Debugw(ctx context.Context, msg string, keysAndValues ...any)
	Infow(ctx context.Context, msg string, keysAndValues ...any)
	Warnw(ctx context.Context, msg string, keysAndValues ...any)
	Errorw(ctx context.Context, msg string, keysAndValues ...any)
	Fatalw(ctx context.Context, msg string, keysAndValues ...any)
	Sync(ctx context.Context) error
}

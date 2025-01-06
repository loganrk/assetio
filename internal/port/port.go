package port

import (
	"assetio/internal/domain"
	"context"
	"net/http"
	"time"
)

// Handler defines the interface for the various API handler methods for account, security, and stock management
type Handler interface {
	// Account-related methods
	AccountCreate(w http.ResponseWriter, r *http.Request)     // Creates a new account
	AccountAll(w http.ResponseWriter, r *http.Request)        // Retrieves all accounts for a user
	AccountGet(w http.ResponseWriter, r *http.Request)        // Retrieves details of a specific account by ID
	AccountUpdate(w http.ResponseWriter, r *http.Request)     // Updates an existing account
	AccountActivate(w http.ResponseWriter, r *http.Request)   // Activates an account
	AccountInactivate(w http.ResponseWriter, r *http.Request) // Inactivates an account

	// Security-related methods
	SecurityCreate(w http.ResponseWriter, r *http.Request) // Creates a new security (e.g., stock, bond, etc.)
	SecurityUpdate(w http.ResponseWriter, r *http.Request) // Updates details of an existing security
	SecurityAll(w http.ResponseWriter, r *http.Request)    // Retrieves all securities of a specific type or exchange
	SecurityGet(w http.ResponseWriter, r *http.Request)    // Retrieves a specific security by its ID
	SecuritySearch(w http.ResponseWriter, r *http.Request) // Searches for securities based on certain criteria

	// Stock-related methods
	StockBuy(w http.ResponseWriter, r *http.Request)              // Buys a stock for a user
	StockSell(w http.ResponseWriter, r *http.Request)             // Sells a stock for a user
	StockSplit(w http.ResponseWriter, r *http.Request)            // Splits a stock (e.g., stock split action)
	StockBonus(w http.ResponseWriter, r *http.Request)            // Bonus for the  a stock
	StockMerge(w http.ResponseWriter, r *http.Request)            // merge for the  a stock
	StockDemerge(w http.ResponseWriter, r *http.Request)          // demerge for the  a stock
	StockDividendAdd(w http.ResponseWriter, r *http.Request)      // Adds a dividend for a specific stock
	StockDividends(w http.ResponseWriter, r *http.Request)        // list of dividend for a specific stock
	StockSummary(w http.ResponseWriter, r *http.Request)          // Retrieves a summary of a user's stock holdings
	StockInventories(w http.ResponseWriter, r *http.Request)      // Retrieves the stock inventory (holdings) for a user
	StockInventoryLedgers(w http.ResponseWriter, r *http.Request) // Retrieves the inventory ledger for stock transactions
}

// Validator defines the interface for validating the different requests for account, security, stock
type Validator interface {
	// Account-related validations
	AccountCreate(request domain.ClientAccountCreateRequest) error         // Validates account creation request
	AccountAll(request domain.ClientAccountAllRequest) error               // Validates request for fetching all accounts
	AccountGet(request domain.ClientAccountGetRequest) error               // Validates request for fetching a specific account
	AccountUpdate(request domain.ClientAccountUpdateRequest) error         // Validates account update request
	AccountActivate(request domain.ClientAccountActivateRequest) error     // Validates account activation request
	AccountInactivate(request domain.ClientAccountInactivateRequest) error // Validates account inactivation request

	// Security-related validations
	SecurityCreate(request domain.ClientSecurityCreateRequest) error // Validates security creation request
	SecurityUpdate(request domain.ClientSecurityUpdateRequest) error // Validates security update request
	SecurityAll(request domain.ClientSecurityAllRequest) error       // Validates request for fetching all securities
	SecurityGet(request domain.ClientSecurityGetRequest) error       // Validates request for fetching a specific security
	SecuritySearch(request domain.ClientSecuritySearchRequest) error // Validates security search request

	// Stock-related validations
	StockBuy(request domain.ClientStockBuyRequest) error                           // Validates stock buy request
	StockSell(request domain.ClientStockSellRequest) error                         // Validates stock sell request
	StockSplit(request domain.ClientStockSplitRequest) error                       // Validates stock split request
	StockBonus(request domain.ClientStockBonusRequest) error                       // Validates stock bonus request
	StockDividendAdd(request domain.ClientStockDividendAddRequest) error           // Validates stock dividend add request
	StockMerge(request domain.ClientStockMergeRequest) error                       // Validates stock merge request
	StockDemerge(request domain.ClientStockDemergeRequest) error                   // Validates stock demerge request
	StockSummary(request domain.ClientStockSummaryRequest) error                   // Validates stock summary request
	StockInventories(request domain.ClientStockInventoriesRequest) error           // Validates request for stock inventories
	StockInventoryLedgers(request domain.ClientStockInventoryLedgersRequest) error // Validates request for stock inventory ledgers
	StockDividends(request domain.ClientStockDividendsRequest) error
}

// RepositoryStore defines the interface for interacting with the database to store and retrieve various entities like accounts, securities, transactions, etc.
type RepositoryStore interface {
	// Account-related database interactions
	AutoMigrate()                                                                                        // Automatically migrate the database schema
	InsertAccountData(ctx context.Context, accountData domain.Accounts) (domain.Accounts, error)         // Inserts new account data
	GetAccountDataByIdAndUserId(ctx context.Context, accountId int, userId int) (domain.Accounts, error) // Retrieves account data by account ID and user ID
	GetAccountsData(ctx context.Context, userId int) ([]domain.Accounts, error)                          // Retrieves all accounts for a user
	UpdateAccountData(ctx context.Context, accountId, userId int, accountData domain.Accounts) error     // Updates an existing account

	// Security-related database interactions
	InsertSecurityData(ctx context.Context, securityData domain.Securities) (domain.Securities, error)                            // Inserts new security data
	GetSecurityDataById(ctx context.Context, securityId int) (domain.Securities, error)                                           // Retrieves security data by ID
	GetSecurityDataByTypeAndExchangeAndSymbol(ctx context.Context, types, exchange int, symbol string) (domain.Securities, error) // Retrieves security data based on type, exchange, and symbol
	UpdateSecurityData(ctx context.Context, securityId int, securityData domain.Securities) error                                 // Updates an existing security
	GetSecuritiesDataByType(ctx context.Context, types int) ([]domain.Securities, error)                                          // Retrieves securities data by exchange
	SearchSecuritiesDataByTypeAndExchange(ctx context.Context, types, exchange int, search string) ([]domain.Securities, error)   // Searches for securities by type, exchange, and search term

	// Inventory-related database interactions
	InsertInventoryLedger(ctx context.Context, inventoryLedgerData domain.InventoryLedger) (domain.InventoryLedger, error)      // Inserts new inventory ledger data
	UpdateInventoryDetailsById(ctx context.Context, inventoryId int, availableQuantity, averagePrice, totalValue float64) error // Updates an inventory by ID
	InsertTransaction(ctx context.Context, transactionData domain.Transactions) (domain.Transactions, error)                    // Inserts a new transaction
	UpdateInventoryLedgerTransactionIdById(ctx context.Context, ledgerId, transactionId int) error                              // Updates inventory ledger with a transaction ID
	UpdateInventoryLedgerTransactionIdByIds(ctx context.Context, ledgerIds []int, transactionId int) error                      // Updates inventory ledgers with a transaction ID

	// Additional inventory-related database interactions
	GetInvertriesSummaryByAccountIdAndSecurityType(ctx context.Context, accountId, securityType int) ([]domain.InventorySummary, error) // Retrieves inventory summary by account ID and security type
	GetInvertriesByAccountIdAndSecurityId(ctx context.Context, accountId, securityId int) ([]domain.InventoryDetails, error)            // Retrieves detailed inventory data by account and security ID

	// Additional transaction and inventory management methods
	InsertTransactionData(ctx context.Context, transactionData domain.Transactions) (domain.Transactions, error)               // Inserts new transaction data
	InsertInventoryData(ctx context.Context, inventoryData domain.Inventories) (domain.Inventories, error)                     // Inserts new inventory data
	GetInventoryDataById(ctx context.Context, inventoryId int) (domain.Inventories, error)                                     // Retrieves inventory data by ID
	UpdateAvailableQuanityToInventoryById(ctx context.Context, inventoryId int, quantity float64) error                        // Updates the available quantity of inventory by ID
	GetActiveInventoriesByAccountIdAndSecurityId(ctx context.Context, accountId, securityId int) ([]domain.Inventories, error) // Retrieves active inventories for an account and security
	GetInventoryLedgersByInventoryId(ctx context.Context, inventoryId int) ([]domain.InventoryLedgers, error)                  // Retrieves inventory ledgers by inventory and account ID
	GetInventoryAvailableQuanitityBySecurityIdAndDate(ctx context.Context, accountId, securityId int, date time.Time) (float64, error)

	GetDividendTransactionsByAccountIdAndSecurityId(ctx context.Context, accountId, securityId int) ([]domain.DividendTransaction, error)
}

// Router defines the interface for routing API requests and handling middleware
type Router interface {
	RegisterRoute(method, path string, handlerFunc http.HandlerFunc) // Registers a route with a specified HTTP method and handler
	StartServer(port string) error                                   // Starts the server on the given port
	UseBefore(middlewares ...http.Handler)                           // Applies middleware to routes before request handling
	NewGroup(groupName string) RouterGroup                           // Creates a new route group
}

// RouterGroup defines the interface for grouping related routes together
type RouterGroup interface {
	RegisterRoute(method, path string, handlerFunc http.HandlerFunc) // Registers a route within the group
	UseBefore(middlewares ...http.Handler)                           // Applies middleware to routes within the group before request handling
}

// Cipher defines the interface for encrypting and decrypting data
type Cipher interface {
	Encrypt(text string) (string, error)       // Encrypts plain text
	Decrypt(cryptoText string) (string, error) // Decrypts encrypted text
	GetKey() string                            // Retrieves the encryption key
}

// Token defines the interface for validating and retrieving access token data
type Token interface {
	GetAccessTokenData(encryptedToken string) (int, time.Time, error) // Retrieves access token data, including user ID and expiration
}

// Auth defines the interface for API key and access token validation
type Auth interface {
	ValidateApiKey() http.Handler      // Validates the API key from the request
	ValidateAccessToken() http.Handler // Validates the access token from the request
}

// Logger defines the interface for logging various levels of log messages
type Logger interface {
	Debug(ctx context.Context, messages ...any)                   // Logs debug messages
	Info(ctx context.Context, messages ...any)                    // Logs info messages
	Warn(ctx context.Context, messages ...any)                    // Logs warning messages
	Error(ctx context.Context, messages ...any)                   // Logs error messages
	Fatal(ctx context.Context, messages ...any)                   // Logs fatal messages and terminates the program
	Debugf(ctx context.Context, template string, args ...any)     // Logs debug messages with formatting
	Infof(ctx context.Context, template string, args ...any)      // Logs info messages with formatting
	Warnf(ctx context.Context, template string, args ...any)      // Logs warning messages with formatting
	Errorf(ctx context.Context, template string, args ...any)     // Logs error messages with formatting
	Fatalf(ctx context.Context, template string, args ...any)     // Logs fatal messages with formatting and terminates the program
	Debugw(ctx context.Context, msg string, keysAndValues ...any) // Logs debug messages with structured data
	Infow(ctx context.Context, msg string, keysAndValues ...any)  // Logs info messages with structured data
	Warnw(ctx context.Context, msg string, keysAndValues ...any)  // Logs warning messages with structured data
	Errorw(ctx context.Context, msg string, keysAndValues ...any) // Logs error messages with structured data
	Fatalw(ctx context.Context, msg string, keysAndValues ...any) // Logs fatal messages with structured data and terminates the program
	Sync(ctx context.Context) error                               // Ensures all logs are written to storage
}

type Marketer interface {
	Query(symbol, exchange string) (MarketerData, error)
}
type MarketerData interface {
	GetMarketPrice() float64
	GetMarketChange() float64
	GetMarketChangePercent() float64
}

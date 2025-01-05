package main

import (
	"assetio/config"
	"assetio/external/yahoo"
	"assetio/internal/domain"
	"assetio/internal/port"
	"context"
	"log"

	"assetio/internal/adapters/handler/validator"
	"assetio/internal/adapters/middleware"

	cipherAes "assetio/internal/adapters/cipher/aes"
	handler "assetio/internal/adapters/handler/http/v1"
	loggerZap "assetio/internal/adapters/logger/zapLogger"
	repositoryMysql "assetio/internal/adapters/repository/mysql"
	routerGin "assetio/internal/adapters/router/gin"
	tokenEngineJwt "assetio/internal/adapters/tokenEngine/jwt"

	accountSrv "assetio/internal/usecase/account"
	mutualFundSrv "assetio/internal/usecase/mutualFund"
	securitySrv "assetio/internal/usecase/security"
	stockSrv "assetio/internal/usecase/stock"
)

const (
	CONFIG_FILE_PATH = `../config/yaml/`
	CONFIG_FILE_NAME = `app_config`
	CONFIG_FILE_TYPE = `yaml`
)

// main function is the entry point of the application.
func main() {
	// Initialize the application configuration by loading settings from the config file.
	appConfigIns, err := config.StartConfig(CONFIG_FILE_PATH, config.File{
		Name: CONFIG_FILE_NAME,
		Ext:  CONFIG_FILE_TYPE,
	})
	if err != nil {
		// If there is an error loading the configuration, log the error and stop the execution.
		log.Println(err)
		return
	}

	// Create an app logger instance based on the app log settings from the configuration.
	appLoggerIns, err := getLogger(appConfigIns.GetAppLog())
	if err != nil {
		// If there is an error getting the logger, log the error and stop the execution.
		log.Println(err)
		return
	}

	// Create an access logger instance based on the access log settings from the configuration.
	accessLoggerIns, err := getLogger(appConfigIns.GetAccessLog())
	if err != nil {
		// If there is an error getting the access logger, log the error and stop the execution.
		log.Println(err)
		return
	}

	// Create a new validator instance for validating inputs.
	validatorIns := validator.New()

	// Get a database instance and initialize it with the app's database configuration.
	mysqlIns, err := getDatabase(appConfigIns)
	if err != nil {
		// If there is an error getting the database instance, log the error and stop the execution.
		log.Println(err)
		return
	}
	// Automatically migrate the database schema if needed.
	mysqlIns.AutoMigrate()

	marketerIns := yahoo.New(appConfigIns.GetYahooExchangeHash())

	// Create instances of different services (Account, Security, Stock, Mutual Fund).
	accountSrvIns := accountSrv.New(appLoggerIns, mysqlIns)
	securitySrvIns := securitySrv.New(appLoggerIns, mysqlIns)
	stockSrvIns := stockSrv.New(appLoggerIns, mysqlIns, marketerIns)
	mutualFundIns := mutualFundSrv.New(appLoggerIns, mysqlIns)

	// Create a service list that contains all the service instances for easy access.
	svcList := domain.List{
		Account:    accountSrvIns,
		Security:   securitySrvIns,
		Stock:      stockSrvIns,
		MutualFund: mutualFundIns,
	}

	// Get a router instance configured with middleware, validation, and logging.
	routerIns := getRouter(appConfigIns, validatorIns, appLoggerIns, accessLoggerIns, svcList)

	// Get the app's port configuration and start the application server.
	port := appConfigIns.GetAppPort()
	// Log that the app has started.
	appLoggerIns.Infow(context.Background(), "app started", "port", port)
	appLoggerIns.Sync(context.Background())

	// Start the server on the configured port.
	err = routerIns.StartServer(port)
	if err != nil {
		// If there is an error starting the server, log the error and stop the application.
		appLoggerIns.Errorw(context.Background(), "app stoped", "port", port, "error", err)
		appLoggerIns.Sync(context.Background())
		return
	}

	// Log that the app has stopped gracefully.
	appLoggerIns.Infow(context.Background(), "app stoped", "port", port, "error", nil)
	appLoggerIns.Sync(context.Background())
}

// getLogger is a helper function to create a logger instance based on the provided log configuration.
func getLogger(logConfigIns config.Logger) (port.Logger, error) {
	// Create a logger configuration using the provided settings from the configuration.
	loggerConfig := loggerZap.Config{
		Level:          logConfigIns.GetLoggerLevel(),
		Encoding:       logConfigIns.GetLoggerEncodingMethod(),
		EncodingCaller: logConfigIns.GetLoggerEncodingCaller(),
		OutputPath:     logConfigIns.GetLoggerPath(),
	}
	// Return a new logger instance based on the configuration.
	return loggerZap.New(loggerConfig)
}

// getDatabase is a helper function to set up and return a database instance.
func getDatabase(appConfigIns config.App) (port.RepositoryStore, error) {
	// Retrieve the crypto key used for decryption from the configuration.
	cipherCryptoKey := appConfigIns.GetCipherCryptoKey()
	// Initialize the AES cipher instance for decryption.
	cipherIns := cipherAes.New(cipherCryptoKey)

	// Retrieve the encrypted database connection properties.
	encryptDbHost, encryptDbPort, encryptDbUsename, encryptDbPasword, dbName, prefix := appConfigIns.GetStoreDatabaseProperties()

	// Decrypt each property using the cipher instance.
	decryptDbHost, decryptErr := cipherIns.Decrypt(encryptDbHost)
	if decryptErr != nil {
		return nil, decryptErr
	}

	decryptdbPort, decryptErr := cipherIns.Decrypt(encryptDbPort)
	if decryptErr != nil {
		return nil, decryptErr
	}

	decryptDbUsename, decryptErr := cipherIns.Decrypt(encryptDbUsename)
	if decryptErr != nil {
		return nil, decryptErr
	}

	decryptDbPasword, decryptErr := cipherIns.Decrypt(encryptDbPasword)
	if decryptErr != nil {
		return nil, decryptErr
	}

	// Return the database instance after successful decryption and initialization.
	return repositoryMysql.New(decryptDbHost, decryptdbPort, decryptDbUsename, decryptDbPasword, dbName, prefix)
}

// getRouter is a helper function to create and configure the router for handling HTTP requests.
func getRouter(appConfigIns config.App, validatorIns port.Validator, appLoggerIns, accessLoggerIns port.Logger, svcList domain.List) port.Router {
	// Initialize the AES cipher instance using the crypto key from the configuration.
	cipherCryptoKey := appConfigIns.GetCipherCryptoKey()
	cipherIns := cipherAes.New(cipherCryptoKey)
	// Get the API keys required for middleware validation.
	apiKeys := appConfigIns.GetMiddlewareApiKeys()

	// Initialize the token engine for JWT token validation.
	tokenEngineIns := tokenEngineJwt.New(cipherIns)

	// Initialize the middleware with API keys and token engine.
	middlewareIns := middleware.New(apiKeys, tokenEngineIns)

	// Create a handler instance for managing request handling logic.
	handlerIns := handler.New(validatorIns, appLoggerIns, svcList)
	// Get API configuration from the app settings.
	apiConfigIns := appConfigIns.GetApi()

	// Initialize the router for handling HTTP requests.
	routerIns := routerGin.New(accessLoggerIns)

	// Create general and access token-specific route groups for organizing routes.
	generalGr := routerIns.NewGroup("")
	generalGr.UseBefore(middlewareIns.ValidateApiKey())

	accessTokenGr := routerIns.NewGroup("")
	accessTokenGr.UseBefore(middlewareIns.ValidateAccessToken())

	// Register routes related to account management.
	updateAccountRouters(generalGr, accessTokenGr, apiConfigIns, handlerIns)

	// Register routes related to security management.
	updateSecurityRouters(generalGr, accessTokenGr, apiConfigIns, handlerIns)

	// Register routes related to stock management.
	updateStockRouters(generalGr, accessTokenGr, apiConfigIns, handlerIns)

	// Register routes related to mutual fund management.
	updateMutualFundRouters(generalGr, accessTokenGr, apiConfigIns, handlerIns)

	// Return the configured router instance.
	return routerIns
}

// Function to update routes for account management.
func updateAccountRouters(generalGr port.RouterGroup, accessTokenGr port.RouterGroup, apiConfigIns config.Api, handlerIns port.Handler) {
	// Register route for account creation if enabled in the config.
	if apiConfigIns.GetAccountCreateEnabled() {
		apiMethod, apiRoute := apiConfigIns.GetAccountCreateProperties()
		accessTokenGr.RegisterRoute(apiMethod, apiRoute, handlerIns.AccountCreate)
	}

	// Register route for fetching all accounts if enabled in the config.
	if apiConfigIns.GetAccountAllEnabled() {
		apiMethod, apiRoute := apiConfigIns.GetAccountAllProperties()
		accessTokenGr.RegisterRoute(apiMethod, apiRoute, handlerIns.AccountAll)
	}

	// Register route for fetching a specific account if enabled in the config.
	if apiConfigIns.GetAccountGetEnabled() {
		apiMethod, apiRoute := apiConfigIns.GetAccountGetProperties()
		accessTokenGr.RegisterRoute(apiMethod, apiRoute, handlerIns.AccountGet)
	}

	// Register route for updating an account if enabled in the config.
	if apiConfigIns.GetAccountUpdateEnabled() {
		apiMethod, apiRoute := apiConfigIns.GetAccountUpdateProperties()
		accessTokenGr.RegisterRoute(apiMethod, apiRoute, handlerIns.AccountUpdate)
	}

	// Register route for activating an account if enabled in the config.
	if apiConfigIns.GetAccountActivateEnabled() {
		apiMethod, apiRoute := apiConfigIns.GetAccountActivateProperties()
		accessTokenGr.RegisterRoute(apiMethod, apiRoute, handlerIns.AccountActivate)
	}

	// Register route for deactivating an account if enabled in the config.
	if apiConfigIns.GetAccountInactivateEnabled() {
		apiMethod, apiRoute := apiConfigIns.GetAccountInactivateProperties()
		accessTokenGr.RegisterRoute(apiMethod, apiRoute, handlerIns.AccountInactivate)
	}
}

// Function to update routes for security management.
func updateSecurityRouters(generalGr port.RouterGroup, accessTokenGr port.RouterGroup, apiConfigIns config.Api, handlerIns port.Handler) {
	// Register route for security creation if enabled in the config.
	if apiConfigIns.GetSecurityCreateEnabled() {
		apiMethod, apiRoute := apiConfigIns.GetSecurityCreateProperties()
		generalGr.RegisterRoute(apiMethod, apiRoute, handlerIns.SecurityCreate)
	}

	// Register route for security update if enabled in the config.
	if apiConfigIns.GetSecurityUpdateEnabled() {
		apiMethod, apiRoute := apiConfigIns.GetSecurityUpdateProperties()
		generalGr.RegisterRoute(apiMethod, apiRoute, handlerIns.SecurityUpdate)
	}

	// Register route for fetching all securities if enabled in the config.
	if apiConfigIns.GetSecurityAllEnabled() {
		apiMethod, apiRoute := apiConfigIns.GetSecurityAllProperties()
		generalGr.RegisterRoute(apiMethod, apiRoute, handlerIns.SecurityAll)
	}

	// Register route for fetching a specific security if enabled in the config.
	if apiConfigIns.GetSecurityGetEnabled() {
		apiMethod, apiRoute := apiConfigIns.GetSecurityGetProperties()
		generalGr.RegisterRoute(apiMethod, apiRoute, handlerIns.SecurityGet)
	}

	// Register route for searching securities if enabled in the config.
	if apiConfigIns.GetSecuritySearchEnabled() {
		apiMethod, apiRoute := apiConfigIns.GetSecuritySearchProperties()
		generalGr.RegisterRoute(apiMethod, apiRoute, handlerIns.SecuritySearch)
	}
}

// Function to update routes for stock management.
func updateStockRouters(generalGr port.RouterGroup, accessTokenGr port.RouterGroup, apiConfigIns config.Api, handlerIns port.Handler) {
	// Register route for stock purchase if enabled in the config.
	if apiConfigIns.GetStockBuyEnabled() {
		apiMethod, apiRoute := apiConfigIns.GetStockBuyProperties()
		accessTokenGr.RegisterRoute(apiMethod, apiRoute, handlerIns.StockBuy)
	}

	// Register route for stock sale if enabled in the config.
	if apiConfigIns.GetStockSellEnabled() {
		apiMethod, apiRoute := apiConfigIns.GetStockSellProperties()
		accessTokenGr.RegisterRoute(apiMethod, apiRoute, handlerIns.StockSell)
	}

	// Register route for adding stock dividend if enabled in the config.
	if apiConfigIns.GetStockDividendAddEnabled() {
		apiMethod, apiRoute := apiConfigIns.GetStockDividendAddProperties()
		accessTokenGr.RegisterRoute(apiMethod, apiRoute, handlerIns.StockDividendAdd)
	}

	// Register route for adding stock dividend list if enabled in the config.
	if apiConfigIns.GetStockDividendsEnabled() {
		apiMethod, apiRoute := apiConfigIns.GetStockDividendsProperties()
		accessTokenGr.RegisterRoute(apiMethod, apiRoute, handlerIns.StockDividends)
	}

	// Register route for stock split if enabled in the config.
	if apiConfigIns.GetStockSplitEnabled() {
		apiMethod, apiRoute := apiConfigIns.GetStockSplitProperties()
		accessTokenGr.RegisterRoute(apiMethod, apiRoute, handlerIns.StockSplit)
	}
	// Register route for stock bonus if enabled in the config.
	if apiConfigIns.GetStockBonusEnabled() {
		apiMethod, apiRoute := apiConfigIns.GetStockBonusProperties()
		accessTokenGr.RegisterRoute(apiMethod, apiRoute, handlerIns.StockBonus)
	}

	// Register route for stock demerge if enabled in the config.
	if apiConfigIns.GetStockMergeEnabled() {
		apiMethod, apiRoute := apiConfigIns.GetStockMergeProperties()
		accessTokenGr.RegisterRoute(apiMethod, apiRoute, handlerIns.StockMerge)
	}

	// Register route for stock demerge if enabled in the config.
	if apiConfigIns.GetStockDemergeEnabled() {
		apiMethod, apiRoute := apiConfigIns.GetStockDemergeProperties()
		accessTokenGr.RegisterRoute(apiMethod, apiRoute, handlerIns.StockDemerge)
	}

	// Register route for fetching stock summary if enabled in the config.
	if apiConfigIns.GetStockSummarylEnabled() {
		apiMethod, apiRoute := apiConfigIns.GetStockSummaryProperties()
		accessTokenGr.RegisterRoute(apiMethod, apiRoute, handlerIns.StockSummary)
	}

	// Register route for fetching stock inventories if enabled in the config.
	if apiConfigIns.GetStockInventorieslEnabled() {
		apiMethod, apiRoute := apiConfigIns.GetStockInventoriesProperties()
		accessTokenGr.RegisterRoute(apiMethod, apiRoute, handlerIns.StockInventories)
	}

	// Register route for fetching stock inventory ledgers if enabled in the config.
	if apiConfigIns.GetStockInventoryLedgerslEnabled() {
		apiMethod, apiRoute := apiConfigIns.GetStockInventoryLedgersProperties()
		accessTokenGr.RegisterRoute(apiMethod, apiRoute, handlerIns.StockInventoryLedgers)
	}
}

// Function to update routes for mutual fund management.
func updateMutualFundRouters(generalGr port.RouterGroup, accessTokenGr port.RouterGroup, apiConfigIns config.Api, handlerIns port.Handler) {
	// Register route for mutual fund purchase if enabled in the config.
	if apiConfigIns.GetMutualFundBuyEnabled() {
		apiMethod, apiRoute := apiConfigIns.GetMutualFundBuyProperties()
		accessTokenGr.RegisterRoute(apiMethod, apiRoute, handlerIns.MutualFundBuy)
	}

	// Register route for mutual fund sale if enabled in the config.
	if apiConfigIns.GetMutualFundSellEnabled() {
		apiMethod, apiRoute := apiConfigIns.GetMutualFundSellProperties()
		accessTokenGr.RegisterRoute(apiMethod, apiRoute, handlerIns.MutualFundSell)
	}

	// Register route for fetching mutual fund summary if enabled in the config.
	if apiConfigIns.GetMutualFundSummarylEnabled() {
		apiMethod, apiRoute := apiConfigIns.GetMutualFundSummaryProperties()
		accessTokenGr.RegisterRoute(apiMethod, apiRoute, handlerIns.MutualFundSummary)
	}

	// Register route for fetching mutual fund inventory if enabled in the config.
	if apiConfigIns.GetMutualFundInventorylEnabled() {
		apiMethod, apiRoute := apiConfigIns.GetMutualFundInventoryProperties()
		accessTokenGr.RegisterRoute(apiMethod, apiRoute, handlerIns.MutualFundInventory)
	}

	// Register route for mutual fund transactions if enabled in the config.
	if apiConfigIns.GetMutualFundTransactionlEnabled() {
		apiMethod, apiRoute := apiConfigIns.GetMutualFundTransactionProperties()
		accessTokenGr.RegisterRoute(apiMethod, apiRoute, handlerIns.MutualFundTransaction)
	}
}

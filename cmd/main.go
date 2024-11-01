package main

import (
	"assetio/config"
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

func main() {
	/* get the config instance */
	appConfigIns, err := config.StartConfig(CONFIG_FILE_PATH, config.File{
		Name: CONFIG_FILE_NAME,
		Ext:  CONFIG_FILE_TYPE,
	})
	if err != nil {
		log.Println(err)
		return
	}

	/* get the app logger instance */
	appLoggerIns, err := getLogger(appConfigIns.GetAppLog())
	if err != nil {
		log.Println(err)
		return
	}

	/* get the app logger instance */
	accessLoggerIns, err := getLogger(appConfigIns.GetAccessLog())
	if err != nil {
		log.Println(err)
		return
	}

	/* get the validator instance */
	validatorIns := validator.New()

	/* get the database instance */
	mysqlIns, err := getDatabase(appConfigIns)
	if err != nil {
		log.Println(err)
		return
	}
	mysqlIns.AutoMigrate()

	/* get the account instance */
	accountSrvIns := accountSrv.New(appLoggerIns, mysqlIns)
	/* get the security instance */
	securitySrvIns := securitySrv.New(appLoggerIns, mysqlIns)
	/* get the stock instance */
	stockSrvIns := stockSrv.New(appLoggerIns, mysqlIns)
	/* get the mutual fund instance */
	mutualFundIns := mutualFundSrv.New(appLoggerIns, mysqlIns)

	svcList := domain.List{
		Account:    accountSrvIns,
		Security:   securitySrvIns,
		Stock:      stockSrvIns,
		MutualFund: mutualFundIns,
	}

	/* get the router instance */
	routerIns := getRouter(appConfigIns, validatorIns, appLoggerIns, accessLoggerIns, svcList)

	/* start the app */
	port := appConfigIns.GetAppPort()
	appLoggerIns.Infow(context.Background(), "app started", "port", port)
	appLoggerIns.Sync(context.Background())

	err = routerIns.StartServer(port)
	if err != nil {
		appLoggerIns.Errorw(context.Background(), "app stoped", "port", port, "error", err)
		appLoggerIns.Sync(context.Background())
		return
	}

	appLoggerIns.Infow(context.Background(), "app stoped", "port", port, "error", nil)
	appLoggerIns.Sync(context.Background())
}

func getLogger(logConfigIns config.Logger) (port.Logger, error) {
	loggerConfig := loggerZap.Config{
		Level:          logConfigIns.GetLoggerLevel(),
		Encoding:       logConfigIns.GetLoggerEncodingMethod(),
		EncodingCaller: logConfigIns.GetLoggerEncodingCaller(),
		OutputPath:     logConfigIns.GetLoggerPath(),
	}
	return loggerZap.New(loggerConfig)
}

func getDatabase(appConfigIns config.App) (port.RepositoryStore, error) {
	cipherCryptoKey := appConfigIns.GetCipherCryptoKey()
	cipherIns := cipherAes.New(cipherCryptoKey)

	encryptDbHost, encryptDbPort, encryptDbUsename, encryptDbPasword, dbName, prefix := appConfigIns.GetStoreDatabaseProperties()

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

	return repositoryMysql.New(decryptDbHost, decryptdbPort, decryptDbUsename, decryptDbPasword, dbName, prefix)

}

func getRouter(appConfigIns config.App, validatorIns port.Validator, appLoggerIns, accessLoggerIns port.Logger, svcList domain.List) port.Router {
	cipherCryptoKey := appConfigIns.GetCipherCryptoKey()
	cipherIns := cipherAes.New(cipherCryptoKey)
	apiKeys := appConfigIns.GetMiddlewareApiKeys()

	tokenEngineIns := tokenEngineJwt.New(cipherIns)

	middlewareIns := middleware.New(apiKeys, tokenEngineIns)

	handlerIns := handler.New(validatorIns, appLoggerIns, svcList)
	apiConfigIns := appConfigIns.GetApi()

	routerIns := routerGin.New(accessLoggerIns)

	generalGr := routerIns.NewGroup("")
	generalGr.UseBefore(middlewareIns.ValidateApiKey())

	accessTokenGr := routerIns.NewGroup("")
	accessTokenGr.UseBefore(middlewareIns.ValidateAccessToken())

	updateAccountRouters(generalGr, accessTokenGr, apiConfigIns, handlerIns)

	return routerIns
}

func updateAccountRouters(generalGr port.RouterGroup, accessTokenGr port.RouterGroup, apiConfigIns config.Api, handlerIns port.Handler) {
	if apiConfigIns.GetAccountCreateEnabled() {
		apiMethod, apiRoute := apiConfigIns.GetAccountCreateProperties()
		accessTokenGr.RegisterRoute(apiMethod, apiRoute, handlerIns.AccountCreate)
	}

	if apiConfigIns.GetAccountAllEnabled() {
		apiMethod, apiRoute := apiConfigIns.GetAccountAllProperties()
		accessTokenGr.RegisterRoute(apiMethod, apiRoute, handlerIns.AccountAll)
	}

	if apiConfigIns.GetAccountGetEnabled() {
		apiMethod, apiRoute := apiConfigIns.GetAccountGetProperties()
		accessTokenGr.RegisterRoute(apiMethod, apiRoute, handlerIns.AccountGet)
	}

	if apiConfigIns.GetAccountUpdateEnabled() {
		apiMethod, apiRoute := apiConfigIns.GetAccountUpdateProperties()
		accessTokenGr.RegisterRoute(apiMethod, apiRoute, handlerIns.AccountUpdate)
	}

	if apiConfigIns.GetAccountActivateEnabled() {
		apiMethod, apiRoute := apiConfigIns.GetAccountActivateProperties()
		accessTokenGr.RegisterRoute(apiMethod, apiRoute, handlerIns.AccountActivate)
	}

	if apiConfigIns.GetAccountInactivateEnabled() {
		apiMethod, apiRoute := apiConfigIns.GetAccountInactivateProperties()
		accessTokenGr.RegisterRoute(apiMethod, apiRoute, handlerIns.AccountInactivate)
	}

	if apiConfigIns.GetSecurityCreateEnabled() {
		apiMethod, apiRoute := apiConfigIns.GetSecurityCreateProperties()
		generalGr.RegisterRoute(apiMethod, apiRoute, handlerIns.SecurityCreate)
	}

	if apiConfigIns.GetSecurityUpdateEnabled() {
		apiMethod, apiRoute := apiConfigIns.GetSecurityUpdateProperties()
		generalGr.RegisterRoute(apiMethod, apiRoute, handlerIns.SecurityUpdate)
	}

	if apiConfigIns.GetSecurityAllEnabled() {
		apiMethod, apiRoute := apiConfigIns.GetSecurityAllProperties()
		generalGr.RegisterRoute(apiMethod, apiRoute, handlerIns.SecurityAll)
	}

	if apiConfigIns.GetSecurityGetEnabled() {
		apiMethod, apiRoute := apiConfigIns.GetSecurityGetProperties()
		generalGr.RegisterRoute(apiMethod, apiRoute, handlerIns.SecurityGet)
	}

	if apiConfigIns.GetSecuritySearchEnabled() {
		apiMethod, apiRoute := apiConfigIns.GetSecuritySearchProperties()
		generalGr.RegisterRoute(apiMethod, apiRoute, handlerIns.SecuritySearch)
	}

	if apiConfigIns.GetStockBuyEnabled() {
		apiMethod, apiRoute := apiConfigIns.GetStockBuyProperties()
		accessTokenGr.RegisterRoute(apiMethod, apiRoute, handlerIns.StockBuy)
	}

	if apiConfigIns.GetStockSellEnabled() {
		apiMethod, apiRoute := apiConfigIns.GetStockSellProperties()
		accessTokenGr.RegisterRoute(apiMethod, apiRoute, handlerIns.StockSell)
	}

	if apiConfigIns.GetStockDividendAddEnabled() {
		apiMethod, apiRoute := apiConfigIns.GetStockDividendAddProperties()
		accessTokenGr.RegisterRoute(apiMethod, apiRoute, handlerIns.StockDividendAdd)
	}

	if apiConfigIns.GetStockSummarylEnabled() {
		apiMethod, apiRoute := apiConfigIns.GetStockSummaryProperties()
		accessTokenGr.RegisterRoute(apiMethod, apiRoute, handlerIns.StockSummary)
	}

	if apiConfigIns.GetStockInventorylEnabled() {
		apiMethod, apiRoute := apiConfigIns.GetStockInventoryProperties()
		accessTokenGr.RegisterRoute(apiMethod, apiRoute, handlerIns.StockInventory)
	}
	if apiConfigIns.GetStockInventoryTransactionslEnabled() {
		apiMethod, apiRoute := apiConfigIns.GetStockInventoryTransactionsProperties()
		accessTokenGr.RegisterRoute(apiMethod, apiRoute, handlerIns.StockInventoryTransactions)
	}

	if apiConfigIns.GetMutualFundBuyEnabled() {
		apiMethod, apiRoute := apiConfigIns.GetMutualFundBuyProperties()
		accessTokenGr.RegisterRoute(apiMethod, apiRoute, handlerIns.MutualFundBuy)
	}

	if apiConfigIns.GetMutualFundSellEnabled() {
		apiMethod, apiRoute := apiConfigIns.GetMutualFundSellProperties()
		accessTokenGr.RegisterRoute(apiMethod, apiRoute, handlerIns.MutualFundSell)
	}

	if apiConfigIns.GetMutualFundSummarylEnabled() {
		apiMethod, apiRoute := apiConfigIns.GetMutualFundSummaryProperties()
		accessTokenGr.RegisterRoute(apiMethod, apiRoute, handlerIns.MutualFundSummary)
	}

	if apiConfigIns.GetMutualFundInventorylEnabled() {
		apiMethod, apiRoute := apiConfigIns.GetMutualFundInventoryProperties()
		accessTokenGr.RegisterRoute(apiMethod, apiRoute, handlerIns.MutualFundInventory)
	}
	if apiConfigIns.GetMutualFundTransactionlEnabled() {
		apiMethod, apiRoute := apiConfigIns.GetMutualFundTransactionProperties()
		accessTokenGr.RegisterRoute(apiMethod, apiRoute, handlerIns.MutualFundTransaction)
	}
}

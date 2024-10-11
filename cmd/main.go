package main

import (
	"assetio/config"
	"assetio/internal/domain"
	"assetio/internal/port"
	"context"
	"log"

	cipherAes "assetio/internal/adapters/cipher/aes"
	handler "assetio/internal/adapters/handler/http/v1"
	loggerZap "assetio/internal/adapters/logger/zapLogger"
	middlewareAuth "assetio/internal/adapters/middleware/auth"
	repositoryMysql "assetio/internal/adapters/repository/mysql"
	routerGin "assetio/internal/adapters/router/gin"
	tokenEngineJwt "assetio/internal/adapters/tokenEngine/jwt"

	accountSrv "assetio/internal/usecase/account"
)

const (
	CONFIG_FILE_PATH = `C:\xampp\htdocs\pro\assetio\config\yaml`
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

	/* get the logger instance */
	loggerIns, err := getLogger(appConfigIns.GetLogger())
	if err != nil {
		log.Println(err)
		return
	}

	/* get the database instance */
	mysqlIns, err := getDatabase(appConfigIns)
	if err != nil {
		log.Println(err)
		return
	}
	mysqlIns.AutoMigrate()

	/* get the user usecase instance */

	/* get the user service instance */
	accountSrvIns := accountSrv.New(loggerIns, mysqlIns)

	svcList := domain.List{
		Account: accountSrvIns,
	}

	/* get the router instance */
	routerIns := getRouter(appConfigIns, loggerIns, svcList)

	/* start the app */
	port := appConfigIns.GetAppPort()
	loggerIns.Infow(context.Background(), "app started", "port", port)
	loggerIns.Sync(context.Background())

	err = routerIns.StartServer(port)
	if err != nil {
		loggerIns.Errorw(context.Background(), "app stoped", "port", port, "error", err)
		loggerIns.Sync(context.Background())
		return
	}

	loggerIns.Infow(context.Background(), "app stoped", "port", port, "error", nil)
	loggerIns.Sync(context.Background())
}

func getLogger(logConfigIns config.Logger) (port.Logger, error) {
	loggerConfig := loggerZap.Config{
		Level:           logConfigIns.GetLoggerLevel(),
		Encoding:        logConfigIns.GetLoggerEncodingMethod(),
		EncodingCaller:  logConfigIns.GetLoggerEncodingCaller(),
		OutputPath:      logConfigIns.GetLoggerPath(),
		ErrorOutputPath: logConfigIns.GetLoggerErrorPath(),
	}
	return loggerZap.New(loggerConfig)
}

func getDatabase(appConfigIns config.App) (port.RepositoryMySQL, error) {
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

func getRouter(appConfigIns config.App, loggerIns port.Logger, svcList domain.List) port.Router {
	cipherCryptoKey := appConfigIns.GetCipherCryptoKey()
	cipherIns := cipherAes.New(cipherCryptoKey)
	apiKeys := appConfigIns.GetMiddlewareApiKeys()

	tokenEngineIns := tokenEngineJwt.New(cipherIns)

	middlewareAuthIns := middlewareAuth.New(apiKeys, tokenEngineIns)

	handlerIns := handler.New(loggerIns, svcList)
	apiConfigIns := appConfigIns.GetApi()

	routerIns := routerGin.New()

	accessTokenGr := routerIns.NewGroup("")
	accessTokenGr.UseBefore(middlewareAuthIns.ValidateAccessToken())

	if apiConfigIns.GetAccountNewEnabled() {
		apiMethod, apiRoute := apiConfigIns.GetAccountNewProperties()
		accessTokenGr.RegisterRoute(apiMethod, apiRoute, handlerIns.AccountAdd)
	}

	if apiConfigIns.GetAccountUpdateEnabled() {
		apiMethod, apiRoute := apiConfigIns.GetAccountUpdateProperties()
		accessTokenGr.RegisterRoute(apiMethod, apiRoute, handlerIns.AccountUpdate)
	}

	return routerIns
}

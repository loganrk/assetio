package main

import (
	"assetio/pkg/config"
	"assetio/pkg/http/v1/handler"
	"assetio/pkg/lib/logger"
	"assetio/pkg/middleware"
	"context"
	"log"

	"github.com/loganrk/go-db"
	"github.com/loganrk/go-router"

	cipher "github.com/loganrk/go-cipher"
)

const (
	CONFIG_FILE_PATH = ``
	CONFIG_FILE_NAME = `app_config`
	CONFIG_FILE_TYPE = `yaml`
)

func main() {

	appConfigIns, err := config.StartConfig(CONFIG_FILE_PATH, config.File{
		Name: CONFIG_FILE_NAME,
		Ext:  CONFIG_FILE_TYPE,
	})

	if err != nil {
		log.Println(err)
		return
	}

	loggerIns, err := createLogger(appConfigIns.GetLogger())

	if err != nil {
		log.Println(err)
		return
	}

	cipherCryptoKey := appConfigIns.GetCipherCryptoKey()
	cipherIns := cipher.New(cipherCryptoKey)

	encryptDbHost, encryptDbPort, encryptDbUsename, encryptDbPasword, dbName := appConfigIns.GetStoreDatabaseProperties()

	decryptDbHost, decryptErr := cipherIns.Decrypt(encryptDbHost)
	if decryptErr != nil {
		log.Println(decryptErr)
		return
	}

	decryptdbPort, decryptErr := cipherIns.Decrypt(encryptDbPort)
	if decryptErr != nil {
		log.Println(decryptErr)
		return
	}

	decryptDbUsename, decryptErr := cipherIns.Decrypt(encryptDbUsename)
	if decryptErr != nil {
		log.Println(decryptErr)
		return
	}

	decryptDbPasword, decryptErr := cipherIns.Decrypt(encryptDbPasword)
	if decryptErr != nil {
		log.Println(decryptErr)
		return
	}

	_, err = db.New(db.Config{
		Host:     decryptDbHost,
		Port:     decryptdbPort,
		Username: decryptDbUsename,
		Password: decryptDbPasword,
		Name:     dbName,
	})

	if err != nil {
		log.Println(err)
		return
	}

	routerIns := router.New()

	authzMiddlewareEnabled, authzMiddlewareToken := appConfigIns.GetMiddlewareAuthorizationProperties()
	if authzMiddlewareEnabled {
		authzMiddlewareIns := middleware.NewAuthz(authzMiddlewareToken)
		routerIns.UseBefore(authzMiddlewareIns.Use())
	}

	handlerIns := handler.New(loggerIns)
	apiConfigIns := appConfigIns.GetApi()

	if apiConfigIns.GetInventoryNewEnabled() {
		apiMethod, apiRoute := apiConfigIns.GetInventoryNewProperties()
		routerIns.RegisterRoute(apiMethod, apiRoute, handlerIns.InventoryAdd)
	}

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

func createLogger(logConfigIns config.Logger) (logger.Logger, error) {
	loggerConfig := logger.Config{
		Level:           logConfigIns.GetLoggerLevel(),
		Encoding:        logConfigIns.GetLoggerEncodingMethod(),
		EncodingCaller:  logConfigIns.GetLoggerEncodingCaller(),
		OutputPath:      logConfigIns.GetLoggerPath(),
		ErrorOutputPath: logConfigIns.GetLoggerErrorPath(),
	}
	return logger.New(loggerConfig)
}

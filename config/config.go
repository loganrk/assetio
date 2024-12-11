package config

import (
	"fmt"

	"github.com/spf13/viper"
)

type File struct {
	Name string
	Ext  string
}

// App defines the interface for the application configuration, including methods for retrieving application details, logging, encryption, database, cache, and API settings.
type App interface {
	// GetAppName returns the name of the application.
	GetAppName() string

	// GetAppPort returns the port on which the application runs.
	GetAppPort() string

	// GetCipherCryptoKey returns the encryption key used for cipher operations.
	GetCipherCryptoKey() string

	// GetMiddlewareApiKeys returns the list of API keys used by the middleware for authorization.
	GetMiddlewareApiKeys() []string

	// GetStoreDatabaseProperties returns the database connection details (host, port, username, password, database name, and prefix).
	GetStoreDatabaseProperties() (string, string, string, string, string, string)

	// GetStoreCacheHeapProperties returns cache heap properties (enabled status and expiry time).
	GetStoreCacheHeapProperties() (bool, int)

	// GetAppLog returns the logger used for general application logging.
	GetAppLog() Logger

	// GetAccessLog returns the logger used for access logging.
	GetAccessLog() Logger

	// GetApi returns the API configuration for the application.
	GetApi() Api

	GetYahooExchangeHash() map[string]string
}

// StartConfig reads and processes the configuration file and returns an App instance or an error.
func StartConfig(path string, file File) (App, error) {
	var appConfig app // appConfig will hold the application configuration

	var viperIns = viper.New() // Creates a new instance of viper for configuration management

	// Set the config file's directory and name for viper to read from
	viperIns.AddConfigPath(path)
	viperIns.SetConfigName(file.Name)
	viperIns.AddConfigPath(".")
	viperIns.SetConfigType(file.Ext)

	// Read the config file and handle any errors
	if err := viperIns.ReadInConfig(); err != nil {
		return nil, fmt.Errorf("error reading config file: %w", err)
	}

	// Unmarshal the config file into the appConfig struct
	err := viperIns.Unmarshal(&appConfig)
	if err != nil {
		return nil, fmt.Errorf("unable to decode into struct, %v", err)
	}

	// Return the populated appConfig or an error if something went wrong
	return appConfig, nil
}

// GetAppLog returns the application log instance for general logging purposes.
func (a app) GetAppLog() Logger {
	return a.AppLog
}

// GetAccessLog returns the access log instance for logging access-related events.
func (a app) GetAccessLog() Logger {
	return a.AccessLog
}

// GetAppName returns the name of the application from the application configuration.
func (a app) GetAppName() string {
	return a.Application.Name
}

// GetAppPort returns the port on which the application is configured to run.
func (a app) GetAppPort() string {
	return a.Application.Port
}

// GetCipherCryptoKey returns the crypto key used for encryption/decryption in the app.
func (a app) GetCipherCryptoKey() string {
	return a.Cipher.CryptoKey
}

// GetMiddlewareApiKeys returns the list of API keys that are used in the application's middleware for authentication/authorization.
func (a app) GetMiddlewareApiKeys() []string {
	return a.Middleware.Keys
}

// GetStoreDatabaseProperties returns the database connection properties (host, port, username, password, database name, and prefix).
func (a app) GetStoreDatabaseProperties() (string, string, string, string, string, string) {
	database := a.Store.Database // Retrieves the database configuration from the app's store settings

	// Returns the database connection details
	return database.Host, database.Port, database.Username, database.Password, database.Name, database.Prefix
}

// GetStoreCacheHeapProperties returns the cache heap properties, such as whether the heap cache is enabled and its expiry time.
func (a app) GetStoreCacheHeapProperties() (bool, int) {
	heapCache := a.Store.Cache.Heap // Retrieves the heap cache configuration from the app's store settings

	// Returns whether the heap cache is enabled and the expiry time in seconds
	return heapCache.Enabled, heapCache.Expiry
}

// GetApi returns the API configuration for the application.
func (a app) GetApi() Api {
	return a.Api
}

func (a app) GetYahooExchangeHash() map[string]string {
	return a.Yahoo.ExchangeHash
}

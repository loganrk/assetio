package config

// app struct represents the main application configuration, which includes various settings
// related to the application, logging, encryption, middleware, data store, and APIs.
type app struct {
	// Application contains the name and port configuration for the application.
	Application struct {
		Name string `mapstructure:"name"` // Name of the application (e.g., "MyApp").
		Port string `mapstructure:"port"` // Port number the application will listen on (e.g., "8080").
	} `mapstructure:"application"`

	// AccessLog and AppLog are the logger configurations for different types of logs.
	AccessLog logger `mapstructure:"accessLog"` // Access log configuration.
	AppLog    logger `mapstructure:"appLog"`    // Application log configuration.

	// Cipher contains the encryption key for cryptographic operations.
	Cipher struct {
		CryptoKey string `mapstructure:"cryptoKey"` // The cryptographic key used for encryption.
	} `mapstructure:"cipher"`

	// Middleware contains the list of API keys used for middleware authentication.
	Middleware struct {
		Keys []string `mapstructure:"keys"` // List of API keys for middleware.
	} `mapstructure:"middleware"`

	// Store contains the database and cache configuration.
	Store struct {
		// Database contains the properties for connecting to a database.
		Database struct {
			Host     string `mapstructure:"host"`     // Database host address.
			Port     string `mapstructure:"port"`     // Database port.
			Username string `mapstructure:"username"` // Database username.
			Password string `mapstructure:"password"` // Database password.
			Name     string `mapstructure:"name"`     // Database name.
			Prefix   string `mapstructure:"prefix"`   // Prefix used in database tables.
		} `mapstructure:"database"`

		// Cache contains the configuration for caching.
		Cache struct {
			Heap struct {
				Enabled     bool `mapstructure:"enabled"`      // Indicates whether heap cache is enabled.
				MaxCapacity int  `mapstructure:"max_capacity"` // Maximum capacity of the heap cache.
				Expiry      int  `mapstructure:"expiry"`       // Cache expiry time in seconds.
			} `mapstructure:"heap"`
		} `mapstructure:"cache"`
	} `mapstructure:"store"`

	// Api contains the API configuration for different services.
	Api api `mapstructure:"api"` // API configuration with various endpoints.
}

// logger struct defines the logging configuration for the application, including log level,
// encoding method, and log file path.
type logger struct {
	Level    int // Level defines the severity of logs (e.g., ERROR, INFO, DEBUG).
	Encoding struct {
		Method int  // Encoding method for logs (e.g., JSON, plaintext).
		Caller bool // Indicates if the caller info (file, line number) should be included in logs.
	}
	Path string // Path where log files will be stored.
}

// api struct contains the configuration for various API endpoints for different actions
// related to accounts, securities, stocks, and mutual funds.
type api struct {
	// Account-related API configurations.
	AccountCreate     apiData `mapstructure:"accountCreate"`     // Account creation API.
	AccountAll        apiData `mapstructure:"accountAll"`        // Get all accounts API.
	AccountGet        apiData `mapstructure:"accountGet"`        // Get account details API.
	AccountUpdate     apiData `mapstructure:"accountUpdate"`     // Update account API.
	AccountActivate   apiData `mapstructure:"accountActivate"`   // Activate account API.
	AccountInactivate apiData `mapstructure:"accountInactivate"` // Inactivate account API.

	// Security-related API configurations.
	SecurityCreate apiData `mapstructure:"securityCreate"` // Create security API.
	SecurityUpdate apiData `mapstructure:"securityUpdate"` // Update security API.
	SecurityGet    apiData `mapstructure:"securityGet"`    // Get security details API.
	SecurityAll    apiData `mapstructure:"securityAll"`    // Get all securities API.
	SecuritySearch apiData `mapstructure:"securitySearch"` // Search securities API.

	// Stock-related API configurations.
	StockBuy              apiData `mapstructure:"stockBuy"`              // Buy stock API.
	StockSell             apiData `mapstructure:"stockSell"`             // Sell stock API.
	StockSplit            apiData `mapstructure:"stockSplit"`            // Stock split API.
	StockDividendAdd      apiData `mapstructure:"stockDividendAdd"`      // Add stock dividend API.
	StockSummary          apiData `mapstructure:"stockSummary"`          // Get stock summary API.
	StockInventories      apiData `mapstructure:"stockInventories"`      // Get stock inventories API.
	StockInventoryLedgers apiData `mapstructure:"stockInventiryLedgers"` // Get stock inventory ledgers API.

	// Mutual fund-related API configurations.
	MutualFundBuy         apiData `mapstructure:"mutualFundBuy"`         // Buy mutual fund API.
	MutualFundSell        apiData `mapstructure:"mutualFundSell"`        // Sell mutual fund API.
	MutualFundSummary     apiData `mapstructure:"mutualFundSummary"`     // Get mutual fund summary API.
	MutualFundInventory   apiData `mapstructure:"mutualFundInventory"`   // Get mutual fund inventory API.
	MutualFundTransaction apiData `mapstructure:"mutualFundTransaction"` // Mutual fund transaction API.
}

// apiData struct defines the configuration for a single API endpoint, including whether
// it is enabled and its route and HTTP method.
type apiData struct {
	Enabled bool   `mapstructure:"enabled"` // Indicates if the API endpoint is enabled.
	Route   string `mapstructure:"route"`   // The route (URL path) for the API endpoint.
	Method  string `mapstructure:"method"`  // The HTTP method (e.g., GET, POST) for the API endpoint.
}

package config

type app struct {
	Application struct {
		Name string `mapstructure:"name"`
		Port string `mapstructure:"port"`
	} `mapstructure:"application"`
	Logger logger `mapstructure:"logger"`
	Cipher struct {
		CryptoKey string `mapstructure:"cryptoKey"`
	} `mapstructure:"cipher"`
	Middleware struct {
		Keys []string `mapstructure:"keys"`
	} `mapstructure:"middleware"`
	Store struct {
		Database struct {
			Host     string `mapstructure:"host"`
			Port     string `mapstructure:"port"`
			Username string `mapstructure:"username"`
			Password string `mapstructure:"password"`
			Name     string `mapstructure:"name"`
			Prefix   string `mapstructure:"prefix"`
		} `mapstructure:"database"`
		Cache struct {
			Heap struct {
				Enabled     bool `mapstructure:"enabled"`
				MaxCapacity int  `mapstructure:"max_capacity"`
				Expiry      int  `mapstructure:"expiry"`
			} `mapstructure:"heap"`
		} `mapstructure:"cache"`
	} `mapstructure:"store"`
	Api api `mapstructure:"api"`
}

type logger struct {
	Level    int
	Encoding struct {
		Method int
		Caller bool
	}
	Path    string
	ErrPath string
}

type api struct {
	AccountCreate     apiData `mapstructure:"accountCreate"`
	AccountAll        apiData `mapstructure:"accountAll"`
	AccountGet        apiData `mapstructure:"accountGet"`
	AccountUpdate     apiData `mapstructure:"accountUpdate"`
	AccountActivate   apiData `mapstructure:"accountActivate"`
	AccountInactivate apiData `mapstructure:"accountInactivate"`

	SecurityCreate apiData `mapstructure:"securityCreate"`
	SecurityUpdate apiData `mapstructure:"securityUpdate"`
	SecurityGet    apiData `mapstructure:"securityGet"`
	SecurityAll    apiData `mapstructure:"securityAll"`
	SecuritySearch apiData `mapstructure:"securitySearch"`

	StockBuy                   apiData `mapstructure:"stockBuy"`
	StockSell                  apiData `mapstructure:"stockSell"`
	StockDividendAdd           apiData `mapstructure:"stockDividendAdd"`
	StockSummary               apiData `mapstructure:"stockSummary"`
	StockInventory             apiData `mapstructure:"stockInventory"`
	StockInventoryTransactions apiData `mapstructure:"stockInventoryTransaction"`

	MutualFundBuy         apiData `mapstructure:"mutualFundBuy"`
	MutualFundSell        apiData `mapstructure:"mutualFundSell"`
	MutualFundSummary     apiData `mapstructure:"mutualFundSummary"`
	MutualFundInventory   apiData `mapstructure:"mutualFundInventory"`
	MutualFundTransaction apiData `mapstructure:"mutualFundTransaction"`
}

type apiData struct {
	Enabled bool   `mapstructure:"enabled"`
	Route   string `mapstructure:"route"`
	Method  string `mapstructure:"method"`
}

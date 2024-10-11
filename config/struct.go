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
	AccountNew    apiData `mapstructure:"accountNew"`
	AccountUpdate apiData `mapstructure:"accountUpdate"`
}

type apiData struct {
	Enabled bool   `mapstructure:"enabled"`
	Route   string `mapstructure:"route"`
	Method  string `mapstructure:"method"`
}

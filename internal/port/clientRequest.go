package port

type AccountCreateClientRequest interface {
	Validate() error
	GetUserId() int
	GetName() string
}

type AccountAllClientRequest interface {
	Validate() error
	GetUserId() int
}

type AccountGetClientRequest interface {
	Validate() error
	GetAccountId() int
	GetUserId() int
}

type AccountUpdateClientRequest interface {
	Validate() error
	GetUserId() int
	GetAccountId() int
	GetName() string
}

type AccountActivateClientRequest interface {
	Validate() error
	GetAccountId() int
	GetUserId() int
}

type AccountInactivateClientRequest interface {
	Validate() error
	GetAccountId() int
	GetUserId() int
}

type SecurityCreateClientRequest interface {
	Validate() error
	GetName() string
	GetType() string
	GetSymbol() string
	GetExchange() string
}

type SecurityUpdateClientRequest interface {
	Validate() error
	GetSecuriryId() int
	GetName() string
	GetType() string
	GetSymbol() string
	GetExchange() string
}

type SecurityGetClientRequest interface {
	Validate() error
	GetSecuriryId() int
}

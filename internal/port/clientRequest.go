package port

type AccountNewClientRequest interface {
	Validate() error
	GetUserId() int
	GetName() string
}

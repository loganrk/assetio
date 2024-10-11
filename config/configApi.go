package config

type Api interface {
	GetAccountNewEnabled() bool
	GetAccountNewProperties() (string, string)

	GetAccountUpdateEnabled() bool
	GetAccountUpdateProperties() (string, string)
}

func (a api) GetAccountNewEnabled() bool {

	return a.AccountNew.Enabled
}

func (a api) GetAccountNewProperties() (string, string) {
	apiData := a.AccountNew

	return apiData.Method, apiData.Route
}

func (a api) GetAccountUpdateEnabled() bool {

	return a.AccountUpdate.Enabled
}

func (a api) GetAccountUpdateProperties() (string, string) {
	apiData := a.AccountUpdate

	return apiData.Method, apiData.Route
}

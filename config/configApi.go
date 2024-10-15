package config

type Api interface {
	GetAccountCreateEnabled() bool
	GetAccountCreateProperties() (string, string)

	GetAccountAllEnabled() bool
	GetAccountAllProperties() (string, string)

	GetAccountGetEnabled() bool
	GetAccountGetProperties() (string, string)

	GetAccountUpdateEnabled() bool
	GetAccountUpdateProperties() (string, string)

	GetAccountActivateEnabled() bool
	GetAccountActivateProperties() (string, string)

	GetAccountInactivateEnabled() bool
	GetAccountInactivateProperties() (string, string)

	GetSecurityCreateEnabled() bool
	GetSecurityCreateProperties() (string, string)

	GetSecurityUpdateEnabled() bool
	GetSecurityUpdateProperties() (string, string)

	GetSecurityAllEnabled() bool
	GetSecurityAllProperties() (string, string)

	GetSecurityGetEnabled() bool
	GetSecurityGetProperties() (string, string)

	GetSecuritySearchEnabled() bool
	GetSecuritySearchProperties() (string, string)
}

func (a api) GetAccountCreateEnabled() bool {

	return a.AccountCreate.Enabled
}

func (a api) GetAccountCreateProperties() (string, string) {
	apiData := a.AccountCreate

	return apiData.Method, apiData.Route
}

func (a api) GetAccountUpdateEnabled() bool {

	return a.AccountUpdate.Enabled
}

func (a api) GetAccountUpdateProperties() (string, string) {
	apiData := a.AccountUpdate

	return apiData.Method, apiData.Route
}

func (a api) GetAccountAllEnabled() bool {

	return a.AccountAll.Enabled
}

func (a api) GetAccountAllProperties() (string, string) {
	apiData := a.AccountAll

	return apiData.Method, apiData.Route
}

func (a api) GetAccountGetEnabled() bool {

	return a.AccountGet.Enabled
}

func (a api) GetAccountGetProperties() (string, string) {
	apiData := a.AccountGet

	return apiData.Method, apiData.Route
}

func (a api) GetAccountActivateEnabled() bool {

	return a.AccountActivate.Enabled
}

func (a api) GetAccountActivateProperties() (string, string) {
	apiData := a.AccountActivate

	return apiData.Method, apiData.Route
}

func (a api) GetAccountInactivateEnabled() bool {

	return a.AccountInactivate.Enabled
}

func (a api) GetAccountInactivateProperties() (string, string) {
	apiData := a.AccountInactivate

	return apiData.Method, apiData.Route
}

func (a api) GetSecurityCreateEnabled() bool {

	return a.SecurityCreate.Enabled
}

func (a api) GetSecurityCreateProperties() (string, string) {
	apiData := a.SecurityCreate

	return apiData.Method, apiData.Route
}

func (a api) GetSecurityUpdateEnabled() bool {

	return a.SecurityUpdate.Enabled
}

func (a api) GetSecurityUpdateProperties() (string, string) {
	apiData := a.SecurityUpdate

	return apiData.Method, apiData.Route
}

func (a api) GetSecurityGetEnabled() bool {

	return a.SecurityGet.Enabled
}

func (a api) GetSecurityGetProperties() (string, string) {
	apiData := a.SecurityGet

	return apiData.Method, apiData.Route
}

func (a api) GetSecurityAllEnabled() bool {

	return a.SecurityAll.Enabled
}

func (a api) GetSecurityAllProperties() (string, string) {
	apiData := a.SecurityAll

	return apiData.Method, apiData.Route
}

func (a api) GetSecuritySearchEnabled() bool {

	return a.SecuritySearch.Enabled
}

func (a api) GetSecuritySearchProperties() (string, string) {
	apiData := a.SecuritySearch

	return apiData.Method, apiData.Route
}

package config

type Api interface {
	// Account API Methods
	// Returns whether the account creation feature is enabled
	GetAccountCreateEnabled() bool

	// Returns the HTTP method and route for creating an account
	GetAccountCreateProperties() (string, string)

	// Returns whether the account listing feature is enabled
	GetAccountAllEnabled() bool

	// Returns the HTTP method and route for fetching all accounts
	GetAccountAllProperties() (string, string)

	// Returns whether the account retrieval feature is enabled
	GetAccountGetEnabled() bool

	// Returns the HTTP method and route for fetching a specific account
	GetAccountGetProperties() (string, string)

	// Returns whether the account update feature is enabled
	GetAccountUpdateEnabled() bool

	// Returns the HTTP method and route for updating an account
	GetAccountUpdateProperties() (string, string)

	// Returns whether the account activation feature is enabled
	GetAccountActivateEnabled() bool

	// Returns the HTTP method and route for activating an account
	GetAccountActivateProperties() (string, string)

	// Returns whether the account inactivation feature is enabled
	GetAccountInactivateEnabled() bool

	// Returns the HTTP method and route for inactivating an account
	GetAccountInactivateProperties() (string, string)

	// Security API Methods
	// Returns whether the security creation feature is enabled
	GetSecurityCreateEnabled() bool

	// Returns the HTTP method and route for creating a security
	GetSecurityCreateProperties() (string, string)

	// Returns whether the security update feature is enabled
	GetSecurityUpdateEnabled() bool

	// Returns the HTTP method and route for updating a security
	GetSecurityUpdateProperties() (string, string)

	// Returns whether the security retrieval feature is enabled
	GetSecurityGetEnabled() bool

	// Returns the HTTP method and route for fetching a specific security
	GetSecurityGetProperties() (string, string)

	// Returns whether the security listing feature is enabled
	GetSecurityAllEnabled() bool

	// Returns the HTTP method and route for fetching all securities
	GetSecurityAllProperties() (string, string)

	// Returns whether the security search feature is enabled
	GetSecuritySearchEnabled() bool

	// Returns the HTTP method and route for searching for securities
	GetSecuritySearchProperties() (string, string)

	// Stock API Methods
	// Returns whether the stock buy feature is enabled
	GetStockBuyEnabled() bool

	// Returns the HTTP method and route for buying stocks
	GetStockBuyProperties() (string, string)

	// Returns whether the stock sell feature is enabled
	GetStockSellEnabled() bool

	// Returns the HTTP method and route for selling stocks
	GetStockSellProperties() (string, string)

	// Returns whether the stock dividend addition feature is enabled
	GetStockDividendAddEnabled() bool

	// Returns the HTTP method and route for adding dividends to stocks
	GetStockDividendAddProperties() (string, string)

	// Returns whether the stock dividend list feature is enabled
	GetStockDividendsEnabled() bool

	// Returns the HTTP method and route for dividends list
	GetStockDividendsProperties() (string, string)

	// Returns whether the stock split feature is enabled
	GetStockSplitEnabled() bool

	// Returns the HTTP method and route for splitting stocks
	GetStockSplitProperties() (string, string)

	// Returns whether the bonus stock feature is enabled
	GetStockBonusEnabled() bool

	// Returns the HTTP method and route for add bonus for the stocks
	GetStockBonusProperties() (string, string)

	// Returns whether the stock merge feature is enabled
	GetStockMergeEnabled() bool

	// Returns the HTTP method and route for merging stocks
	GetStockMergeProperties() (string, string)

	// Returns whether the stock demerge feature is enabled
	GetStockDemergeEnabled() bool

	// Returns the HTTP method and route for demerging stocks
	GetStockDemergeProperties() (string, string)

	// Returns whether the stock summary feature is enabled
	GetStockSummarylEnabled() bool

	// Returns the HTTP method and route for fetching stock summaries
	GetStockSummaryProperties() (string, string)

	// Returns whether the stock inventories feature is enabled
	GetStockInventorieslEnabled() bool

	// Returns the HTTP method and route for fetching stock inventories
	GetStockInventoriesProperties() (string, string)

	// Returns whether the stock inventory ledgers feature is enabled
	GetStockInventoryLedgerslEnabled() bool

	// Returns the HTTP method and route for fetching stock inventory ledgers
	GetStockInventoryLedgersProperties() (string, string)
}

// GetAccountCreateEnabled checks if account creation is enabled and returns a boolean.
func (a api) GetAccountCreateEnabled() bool {
	return a.AccountCreate.Enabled
}

// GetAccountCreateProperties returns the HTTP method and route for the account creation API.
func (a api) GetAccountCreateProperties() (string, string) {
	apiData := a.AccountCreate
	return apiData.Method, apiData.Route
}

// GetAccountUpdateEnabled checks if account update is enabled and returns a boolean.
func (a api) GetAccountUpdateEnabled() bool {
	return a.AccountUpdate.Enabled
}

// GetAccountUpdateProperties returns the HTTP method and route for the account update API.
func (a api) GetAccountUpdateProperties() (string, string) {
	apiData := a.AccountUpdate
	return apiData.Method, apiData.Route
}

// GetAccountAllEnabled checks if fetching all accounts is enabled and returns a boolean.
func (a api) GetAccountAllEnabled() bool {
	return a.AccountAll.Enabled
}

// GetAccountAllProperties returns the HTTP method and route for the account listing API.
func (a api) GetAccountAllProperties() (string, string) {
	apiData := a.AccountAll
	return apiData.Method, apiData.Route
}

// GetAccountGetEnabled checks if fetching a specific account is enabled and returns a boolean.
func (a api) GetAccountGetEnabled() bool {
	return a.AccountGet.Enabled
}

// GetAccountGetProperties returns the HTTP method and route for fetching a specific account.
func (a api) GetAccountGetProperties() (string, string) {
	apiData := a.AccountGet
	return apiData.Method, apiData.Route
}

// GetAccountActivateEnabled checks if activating an account is enabled and returns a boolean.
func (a api) GetAccountActivateEnabled() bool {
	return a.AccountActivate.Enabled
}

// GetAccountActivateProperties returns the HTTP method and route for activating an account.
func (a api) GetAccountActivateProperties() (string, string) {
	apiData := a.AccountActivate
	return apiData.Method, apiData.Route
}

// GetAccountInactivateEnabled checks if inactivating an account is enabled and returns a boolean.
func (a api) GetAccountInactivateEnabled() bool {
	return a.AccountInactivate.Enabled
}

// GetAccountInactivateProperties returns the HTTP method and route for inactivating an account.
func (a api) GetAccountInactivateProperties() (string, string) {
	apiData := a.AccountInactivate
	return apiData.Method, apiData.Route
}

// GetSecurityCreateEnabled checks if security creation is enabled and returns a boolean.
func (a api) GetSecurityCreateEnabled() bool {
	return a.SecurityCreate.Enabled
}

// GetSecurityCreateProperties returns the HTTP method and route for creating a security.
func (a api) GetSecurityCreateProperties() (string, string) {
	apiData := a.SecurityCreate
	return apiData.Method, apiData.Route
}

// GetSecurityUpdateEnabled checks if security update is enabled and returns a boolean.
func (a api) GetSecurityUpdateEnabled() bool {
	return a.SecurityUpdate.Enabled
}

// GetSecurityUpdateProperties returns the HTTP method and route for updating a security.
func (a api) GetSecurityUpdateProperties() (string, string) {
	apiData := a.SecurityUpdate
	return apiData.Method, apiData.Route
}

// GetSecurityGetEnabled checks if fetching a specific security is enabled and returns a boolean.
func (a api) GetSecurityGetEnabled() bool {
	return a.SecurityGet.Enabled
}

// GetSecurityGetProperties returns the HTTP method and route for fetching a specific security.
func (a api) GetSecurityGetProperties() (string, string) {
	apiData := a.SecurityGet
	return apiData.Method, apiData.Route
}

// GetSecurityAllEnabled checks if fetching all securities is enabled and returns a boolean.
func (a api) GetSecurityAllEnabled() bool {
	return a.SecurityAll.Enabled
}

// GetSecurityAllProperties returns the HTTP method and route for fetching all securities.
func (a api) GetSecurityAllProperties() (string, string) {
	apiData := a.SecurityAll
	return apiData.Method, apiData.Route
}

// GetSecuritySearchEnabled checks if security search is enabled and returns a boolean.
func (a api) GetSecuritySearchEnabled() bool {
	return a.SecuritySearch.Enabled
}

// GetSecuritySearchProperties returns the HTTP method and route for searching securities.
func (a api) GetSecuritySearchProperties() (string, string) {
	apiData := a.SecuritySearch
	return apiData.Method, apiData.Route
}

// GetStockBuyEnabled checks if stock buying is enabled and returns a boolean.
func (a api) GetStockBuyEnabled() bool {
	return a.StockBuy.Enabled
}

// GetStockBuyProperties returns the HTTP method and route for buying stocks.
func (a api) GetStockBuyProperties() (string, string) {
	apiData := a.StockBuy
	return apiData.Method, apiData.Route
}

// GetStockSellEnabled checks if stock selling is enabled and returns a boolean.
func (a api) GetStockSellEnabled() bool {
	return a.StockSell.Enabled
}

// GetStockSellProperties returns the HTTP method and route for selling stocks.
func (a api) GetStockSellProperties() (string, string) {
	apiData := a.StockSell
	return apiData.Method, apiData.Route
}

// GetStockDividendAddEnabled checks if adding dividends to stocks is enabled and returns a boolean.
func (a api) GetStockDividendAddEnabled() bool {
	return a.StockDividendAdd.Enabled
}

// GetStockDividendAddProperties returns the HTTP method and route for adding stock dividends.
func (a api) GetStockDividendAddProperties() (string, string) {
	apiData := a.StockDividendAdd
	return apiData.Method, apiData.Route
}

// GetStockDividendsEnabled checks if list of dividends to stocks is enabled and returns a boolean.
func (a api) GetStockDividendsEnabled() bool {
	return a.StockDividends.Enabled
}

// GetStockDividendsProperties returns the HTTP method and route for list of stock dividends.
func (a api) GetStockDividendsProperties() (string, string) {
	apiData := a.StockDividends
	return apiData.Method, apiData.Route
}

// GetStockSplitEnabled checks if stock splitting is enabled and returns a boolean.
func (a api) GetStockSplitEnabled() bool {
	return a.StockSplit.Enabled
}

// GetStockSplitProperties returns the HTTP method and route for splitting stocks.
func (a api) GetStockSplitProperties() (string, string) {
	apiData := a.StockSplit
	return apiData.Method, apiData.Route
}

// GetStockSplitEnabled checks if stock bonus is enabled and returns a boolean.
func (a api) GetStockBonusEnabled() bool {
	return a.StockBonus.Enabled
}

// GetStockSplitProperties returns the HTTP method and route for bonus for the stocks.
func (a api) GetStockBonusProperties() (string, string) {
	apiData := a.StockBonus
	return apiData.Method, apiData.Route
}

// GetStockMergeEnabled checks if stock merging is enabled and returns a boolean.
func (a api) GetStockMergeEnabled() bool {
	return a.StockMerge.Enabled
}

// GetStockMergeProperties returns the HTTP method and route for merging stocks.
func (a api) GetStockMergeProperties() (string, string) {
	apiData := a.StockMerge
	return apiData.Method, apiData.Route
}

// GetStockDemergeEnabled checks if stock demerging is enabled and returns a boolean.
func (a api) GetStockDemergeEnabled() bool {
	return a.StockDemerge.Enabled
}

// GetStockDemergeProperties returns the HTTP method and route for demerging stocks.
func (a api) GetStockDemergeProperties() (string, string) {
	apiData := a.StockDemerge
	return apiData.Method, apiData.Route
}

// GetStockSummarylEnabled checks if fetching stock summaries is enabled and returns a boolean.
func (a api) GetStockSummarylEnabled() bool {
	return a.StockSummary.Enabled
}

// GetStockSummaryProperties returns the HTTP method and route for fetching stock summaries.
func (a api) GetStockSummaryProperties() (string, string) {
	apiData := a.StockSummary
	return apiData.Method, apiData.Route
}

// GetStockInventorieslEnabled checks if fetching stock inventories is enabled and returns a boolean.
func (a api) GetStockInventorieslEnabled() bool {
	return a.StockInventories.Enabled
}

// GetStockInventoriesProperties returns the HTTP method and route for fetching stock inventories.
func (a api) GetStockInventoriesProperties() (string, string) {
	apiData := a.StockInventories
	return apiData.Method, apiData.Route
}

// GetStockInventoryLedgerslEnabled checks if fetching stock inventory ledgers is enabled and returns a boolean.
func (a api) GetStockInventoryLedgerslEnabled() bool {
	return a.StockInventoryLedgers.Enabled
}

// GetStockInventoryLedgersProperties returns the HTTP method and route for fetching stock inventory ledgers.
func (a api) GetStockInventoryLedgersProperties() (string, string) {
	apiData := a.StockInventoryLedgers
	return apiData.Method, apiData.Route
}

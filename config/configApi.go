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

	GetStockBuyEnabled() bool
	GetStockBuyProperties() (string, string)

	GetStockSellEnabled() bool
	GetStockSellProperties() (string, string)

	GetStockDividendAddEnabled() bool
	GetStockDividendAddProperties() (string, string)

	GetStockSummarylEnabled() bool
	GetStockSummaryProperties() (string, string)

	GetStockInventorylEnabled() bool
	GetStockInventoryProperties() (string, string)

	GetStockInventoryTransactionslEnabled() bool
	GetStockInventoryTransactionsProperties() (string, string)

	GetMutualFundBuyEnabled() bool
	GetMutualFundBuyProperties() (string, string)

	GetMutualFundSellEnabled() bool
	GetMutualFundSellProperties() (string, string)

	GetMutualFundSummarylEnabled() bool
	GetMutualFundSummaryProperties() (string, string)

	GetMutualFundInventorylEnabled() bool
	GetMutualFundInventoryProperties() (string, string)

	GetMutualFundTransactionlEnabled() bool
	GetMutualFundTransactionProperties() (string, string)
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

func (a api) GetStockBuyEnabled() bool {

	return a.StockBuy.Enabled
}

func (a api) GetStockBuyProperties() (string, string) {
	apiData := a.StockBuy

	return apiData.Method, apiData.Route
}

func (a api) GetStockSellEnabled() bool {

	return a.StockSell.Enabled
}

func (a api) GetStockSellProperties() (string, string) {
	apiData := a.StockSell

	return apiData.Method, apiData.Route
}

func (a api) GetStockDividendAddEnabled() bool {

	return a.StockDividendAdd.Enabled
}

func (a api) GetStockDividendAddProperties() (string, string) {
	apiData := a.StockDividendAdd

	return apiData.Method, apiData.Route
}

func (a api) GetStockSummarylEnabled() bool {

	return a.StockSummary.Enabled
}

func (a api) GetStockSummaryProperties() (string, string) {
	apiData := a.StockSummary

	return apiData.Method, apiData.Route
}

func (a api) GetStockInventorylEnabled() bool {

	return a.StockInventory.Enabled
}

func (a api) GetStockInventoryProperties() (string, string) {
	apiData := a.StockInventory

	return apiData.Method, apiData.Route
}

func (a api) GetStockInventoryTransactionslEnabled() bool {

	return a.StockInventoryTransactions.Enabled
}

func (a api) GetStockInventoryTransactionsProperties() (string, string) {
	apiData := a.StockInventoryTransactions

	return apiData.Method, apiData.Route
}

func (a api) GetMutualFundBuyEnabled() bool {

	return a.MutualFundBuy.Enabled
}

func (a api) GetMutualFundBuyProperties() (string, string) {
	apiData := a.MutualFundBuy

	return apiData.Method, apiData.Route
}

func (a api) GetMutualFundSellEnabled() bool {

	return a.MutualFundSell.Enabled
}

func (a api) GetMutualFundSellProperties() (string, string) {
	apiData := a.MutualFundSell

	return apiData.Method, apiData.Route
}

func (a api) GetMutualFundSummarylEnabled() bool {

	return a.MutualFundSummary.Enabled
}

func (a api) GetMutualFundSummaryProperties() (string, string) {
	apiData := a.MutualFundSummary

	return apiData.Method, apiData.Route
}

func (a api) GetMutualFundInventorylEnabled() bool {

	return a.MutualFundInventory.Enabled
}

func (a api) GetMutualFundInventoryProperties() (string, string) {
	apiData := a.MutualFundInventory

	return apiData.Method, apiData.Route
}

func (a api) GetMutualFundTransactionlEnabled() bool {

	return a.MutualFundTransaction.Enabled
}

func (a api) GetMutualFundTransactionProperties() (string, string) {
	apiData := a.MutualFundTransaction

	return apiData.Method, apiData.Route
}

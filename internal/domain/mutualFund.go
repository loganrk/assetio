package domain

type ClientMutualFundBuyRequest struct {
	UserId       int     `json:"uid" schema:"uid"`
	AccountId    int     `json:"account_id" schema:"account_id"`
	MutualFundId int     `json:"mutual_fund_id" schema:"mutual_fund_id"`
	InventoryId  int     `json:"inventory_id" schema:"inventory_id"`
	Quantity     float64 `json:"quantity" schema:"quantity"`
	AveragePrice float64 `json:"average_price" schema:"average_price"`
	FeeAmount    float64 `json:"fee_amount" schema:"fee_amount"`
}

type ClientMutualFundSellRequest struct {
	UserId       int     `json:"uid" schema:"uid"`
	AccountId    int     `json:"account_id" schema:"account_id"`
	MutualFundId int     `json:"mutual_fund_id" schema:"mutual_fund_id"`
	InventoryId  int     `json:"inventory_id" schema:"inventory_id"`
	Quantity     float64 `json:"quantity" schema:"quantity"`
	AveragePrice float64 `json:"average_price" schema:"average_price"`
	FeeAmount    float64 `json:"fee_amount" schema:"fee_amount"`
}

type ClientMutualFundSellResponse struct {
	Message string `json:"message" schema:"message"`
}

type ClientMutualFundAddRequest struct {
	UserId       int     `json:"uid" schema:"uid"`
	AccountId    int     `json:"account_id" schema:"account_id"`
	MutualFundId int     `json:"mutual_fund_id" schema:"mutual_fund_id"`
	InventoryId  int     `json:"inventory_id" schema:"inventory_id"`
	Quantity     float64 `json:"quantity" schema:"quantity"`
	AveragePrice float64 `json:"average_price" schema:"average_price"`
	FeeAmount    float64 `json:"fee_amount" schema:"fee_amount"`
}

type ClientMutualFundSummaryRequest struct {
	UserId    int `json:"uid" schema:"uid"`
	AccountId int `json:"account_id" schema:"account_id"`
}

type ClientMutualFundSummaryResponse struct {
	MutualFundId       int     `gorm:"mutual_fund_id" schema:"mutual_fund_id"`
	MutualFundSymbol   string  `gorm:"mutual_fund_symbol" schema:"mutual_fund_symbol"`
	MutualFundExchange string  `gorm:"mutual_fund_exchange" schema:"mutual_fund_exchange"`
	MutualFundName     string  `gorm:"mutual_fund_name" schema:"mutual_fund_name"`
	Quantity           int     `gorm:"quantity" schema:"quantity"`
	Amount             float64 `gorm:"amount" schema:"amount"`
}

type ClientMutualFundInventoryRequest struct {
	UserId       int `json:"uid" schema:"uid"`
	AccountId    int `json:"account_id" schema:"account_id"`
	MutualFundId int `json:"mutual_fund_id" schema:"mutual_fund_id"`
}

type ClientMutualFundInventoryResponse struct {
	InventoryId int     `json:"inventory_id" schema:"inventory_id"`
	Quantity    int     `json:"quantity" schema:"quantity"`
	Amount      float64 `json:"amount" schema:"amount"`
	Date        string  `json:"date" schema:"date"`
}

type ClientMutualFundInventoryLedgersRequest struct {
	UserId      int `json:"uid" schema:"uid"`
	AccountId   int `json:"account_id" schema:"account_id"`
	InventoryId int `json:"inventory_id" schema:"inventory_id"`
}

type ClientMutualFundInventoryLedgersResponse struct {
	TransactionId   int     `json:"transaction_id" schema:"transaction_id"`
	TransactionType string  `json:"transaction_ype" schema:"transaction_ype"`
	Quantity        int     `json:"quantity" schema:"quantity"`
	Amount          float64 `json:"amount" schema:"amount"`
	Fee             float64 `json:"fee" schema:"fee"`
	Date            string  `json:"date" schema:"date"`
}

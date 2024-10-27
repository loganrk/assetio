package domain

type ClientStockBuyRequest struct {
	UserId            int     `json:"uid" schema:"uid"`
	AccountId         int     `json:"account_id" schema:"account_id"`
	StockId           int     `json:"stock_id" schema:"stock_id"`
	Quantity          float64 `json:"quantity" schema:"quantity"`
	AmountPerQuantity float64 `json:"amount_per_quantity" schema:"amount_per_quantity"`
	FeeAmount         float64 `json:"fee_amount" schema:"fee_amount"`
}

type ClientStockBuyResponse struct {
	Message string `json:"message" schema:"message"`
}

type ClientStockSellRequest struct {
	UserId            int     `json:"uid" schema:"uid"`
	AccountId         int     `json:"account_id" schema:"account_id"`
	StockId           int     `json:"stock_id" schema:"stock_id"`
	InventoryId       int     `json:"inventory_id" schema:"inventory_id"`
	Quantity          float64 `json:"quantity" schema:"quantity"`
	AmountPerQuantity float64 `json:"amount_per_quantity" schema:"amount_per_quantity"`
	FeeAmount         float64 `json:"fee_amount" schema:"fee_amount"`
}
type ClientStockSellResponse struct {
	Message string `json:"message" schema:"message"`
}

type ClientStockDividendAddRequest struct {
	UserId            int     `json:"uid" schema:"uid"`
	AccountId         int     `json:"account_id" schema:"account_id"`
	StockId           int     `json:"stock_id" schema:"stock_id"`
	InventoryId       int     `json:"inventory_id" schema:"inventory_id"`
	Quantity          float64 `json:"quantity" schema:"quantity"`
	AmountPerQuantity float64 `json:"amount_per_quantity" schema:"amount_per_quantity"`
}

type ClientStockDividendResponse struct {
	Message string `json:"message" schema:"message"`
}

type ClientStockSummaryRequest struct {
	UserId    int `json:"uid" schema:"uid"`
	AccountId int `json:"account_id" schema:"account_id"`
}

type ClientStockSummaryResponse struct {
	StockId       int     `gorm:"stock_id" schema:"ustock_idid"`
	StockSymbol   string  `gorm:"stock_symbol" schema:"stock_symbol"`
	StockExchange string  `gorm:"stock_exchange" schema:"stock_exchange"`
	StockName     string  `gorm:"stock_name" schema:"stock_name"`
	Quantity      int     `gorm:"quantity" schema:"quantity"`
	Amount        float64 `gorm:"amount" schema:"amount"`
}

type ClientStockInventoryRequest struct {
	UserId    int `json:"uid" schema:"uid"`
	AccountId int `json:"account_id" schema:"account_id"`
	StockId   int `json:"stock_id" schema:"stock_id"`
}

type ClientStockInventoryResponse struct {
	InventoryId int     `json:"inventory_id" schema:"inventory_id"`
	Quantity    int     `json:"quantity" schema:"quantity"`
	Amount      float64 `json:"amount" schema:"amount"`
	Date        string  `json:"date" schema:"date"`
}

type ClientStockInventoryTransactionsRequest struct {
	UserId      int `json:"uid" schema:"uid"`
	AccountId   int `json:"account_id" schema:"account_id"`
	InventoryId int `json:"inventory_id" schema:"inventory_id"`
}

type ClientStockInventoryTransactionsResponse struct {
	TransactionId   int     `json:"transaction_id" schema:"transaction_id"`
	TransactionType string  `json:"transaction_ype" schema:"transaction_ype"`
	Quantity        int     `json:"quantity" schema:"quantity"`
	Amount          float64 `json:"amount" schema:"amount"`
	Fee             float64 `json:"fee" schema:"fee"`
	Date            string  `json:"date" schema:"date"`
}

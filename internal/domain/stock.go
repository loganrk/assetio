package domain

type ClientStockBuyRequest struct {
	UserId       int     `json:"uid" schema:"uid"`
	AccountId    int     `json:"account_id" schema:"account_id"`
	StockId      int     `json:"stock_id" schema:"stock_id"`
	InventoryId  int     `json:"inventory_id" schema:"inventory_id"`
	Date         string  `json:"date" schema:"date"`
	Quantity     float64 `json:"quantity" schema:"quantity"`
	AveragePrice float64 `json:"average_price" schema:"average_price"`
	FeeAmount    float64 `json:"fee_amount" schema:"fee_amount"`
}

type ClientStockBuyResponse struct {
	Message string `json:"message" schema:"message"`
}

type ClientStockSellRequest struct {
	UserId       int     `json:"uid" schema:"uid"`
	AccountId    int     `json:"account_id" schema:"account_id"`
	StockId      int     `json:"stock_id" schema:"stock_id"`
	InventoryId  int     `json:"inventory_id" schema:"inventory_id"`
	Date         string  `json:"date" schema:"date"`
	Quantity     float64 `json:"quantity" schema:"quantity"`
	AveragePrice float64 `json:"average_price" schema:"average_price"`
	FeeAmount    float64 `json:"fee_amount" schema:"fee_amount"`
}
type ClientStockSellResponse struct {
	Message string `json:"message" schema:"message"`
}

type ClientStockSplitRequest struct {
	UserId      int     `json:"uid" schema:"uid"`
	AccountId   int     `json:"account_id" schema:"account_id"`
	StockId     int     `json:"stock_id" schema:"stock_id"`
	InventoryId int     `json:"inventory_id" schema:"inventory_id"`
	Quantity    float64 `json:"quantity" schema:"quantity"`
	FeeAmount   float64 `json:"fee_amount" schema:"fee_amount"`
}

type ClientStockSplitResponse struct {
	Message string `json:"message" schema:"message"`
}

type ClientStockDividendAddRequest struct {
	UserId            int     `json:"uid" schema:"uid"`
	AccountId         int     `json:"account_id" schema:"account_id"`
	StockId           int     `json:"stock_id" schema:"stock_id"`
	InventoryId       int     `json:"inventory_id" schema:"inventory_id"`
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
	StockId             int     `json:"stock_id" schema:"stock_id"`
	StockSymbol         string  `json:"stock_symbol" schema:"stock_symbol"`
	StockExchange       string  `json:"stock_exchange" schema:"stock_exchange"`
	StockName           string  `json:"stock_name" schema:"stock_name"`
	Quantity            int     `json:"quantity" schema:"quantity"`
	Amount              float64 `json:"amount" schema:"amount"`
	MarketPrice         float64 `json:"market_price" schema:"market_price"`
	MarketChange        float64 `json:"market_change" schema:"market_change"`
	MarketChangePercent float64 `json:"market_change_percent" schema:"market_change_percent"`
}

type ClientStockInventoriesRequest struct {
	UserId    int `json:"uid" schema:"uid"`
	AccountId int `json:"account_id" schema:"account_id"`
	StockId   int `json:"stock_id" schema:"stock_id"`
}

type ClientStockInventoriesResponse struct {
	InventoryId         int     `json:"inventory_id" schema:"inventory_id"`
	Quantity            int     `json:"quantity" schema:"quantity"`
	Amount              float64 `json:"amount" schema:"amount"`
	Date                string  `json:"date" schema:"date"`
	MarketPrice         float64 `json:"market_price" schema:"market_price"`
	MarketChange        float64 `json:"market_change" schema:"market_change"`
	MarketChangePercent float64 `json:"market_change_percent" schema:"market_change_percent"`
}

type ClientStockInventoryLedgersRequest struct {
	UserId      int `json:"uid" schema:"uid"`
	AccountId   int `json:"account_id" schema:"account_id"`
	InventoryId int `json:"inventory_id" schema:"inventory_id"`
}

type ClientStockInventoryLedgersResponse struct {
	TransactionId   int     `json:"transaction_id" schema:"transaction_id"`
	TransactionType string  `json:"transaction_ype" schema:"transaction_ype"`
	Quantity        int     `json:"quantity" schema:"quantity"`
	Amount          float64 `json:"amount" schema:"amount"`
	Date            string  `json:"date" schema:"date"`
}

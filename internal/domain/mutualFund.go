package domain

type ClientMutualFundBuyRequest struct {
	UserId            int     `json:"uid" schema:"uid"`
	AccountId         int     `json:"account_id" schema:"account_id"`
	MutualFundId      int     `json:"stock_id" schema:"stock_id"`
	InventoryId       int     `json:"inventory_id" schema:"inventory_id"`
	Quantity          float64 `json:"quantity" schema:"quantity"`
	AmountPerQuantity float64 `json:"amount_per_quantity" schema:"amount_per_quantity"`
	FeeAmount         float64 `json:"fee_amount" schema:"fee_amount"`
}

type ClientMutualFundSellRequest struct {
	UserId            int     `json:"uid" schema:"uid"`
	AccountId         int     `json:"account_id" schema:"account_id"`
	MutualFundId      int     `json:"stock_id" schema:"stock_id"`
	InventoryId       int     `json:"inventory_id" schema:"inventory_id"`
	Quantity          float64 `json:"quantity" schema:"quantity"`
	AmountPerQuantity float64 `json:"amount_per_quantity" schema:"amount_per_quantity"`
	FeeAmount         float64 `json:"fee_amount" schema:"fee_amount"`
}

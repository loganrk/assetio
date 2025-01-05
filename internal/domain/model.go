package domain

import "time"

type TransactionType string

const (
	BUY               TransactionType = "BUY"
	SELL              TransactionType = "SELL"
	DIVIDEND          TransactionType = "DIVIDEND"
	SPLIT             TransactionType = "SPLIT"
	BONUS             TransactionType = "BONUS"
	MERGER            TransactionType = "MERGER"
	DEMERGER          TransactionType = "DEMERGER"
	DEMERGER_TRANSFER TransactionType = "DEMERGER_TRANSFER"
)

type Accounts struct {
	Id        int       `gorm:"primarykey;size:16"`
	UserId    int       `gorm:"column:user_id;size:16"`
	Status    int       `gorm:"column:status;size:11"`
	Name      string    `gorm:"column:name;size:255"`
	CreatedAt time.Time `gorm:"autoCreateTime,column:created_at"`
	UpdatedAt time.Time `gorm:"autoUpdateTime,column:updated_at"`
}

type Inventories struct {
	Id                int       `gorm:"primarykey;size:16"`
	AccountId         int       `gorm:"column:account_id;size:16"`
	SecurityId        int       `gorm:"column:security_id;size:16"`
	AvailableQuantity float64   `gorm:"type:decimal(12,4);column:available_quantity"`
	AveragePrice      float64   `gorm:"type:decimal(12,4);column:average_price"`
	TotalValue        float64   `gorm:"type:decimal(12,4);column:total_value"`
	Date              time.Time `gorm:"column:date"`
	State             int       `gorm:"column:state;size:11;"`
	CreatedAt         time.Time `gorm:"autoCreateTime,column:created_at"`
	UpdatedAt         time.Time `gorm:"autoUpdateTime,column:updated_at"`
}

type InventoryLedger struct {
	Id            int             `gorm:"primarykey;size:16"`
	InventoryId   int             `gorm:"column:inventory_id;size:16"`
	TransactionId int             `gorm:"column:transaction_id;size:16"`
	Type          TransactionType `gorm:"type:enum('BUY', 'SELL', 'DIVIDEND', 'SPLIT', 'BONUS' , 'DEMERGER', 'DEMERGER_TRANSFER');column:type;size:16"`
	Quantity      float64         `gorm:"type:decimal(12,4);column:quantity"`
	AveragePrice  float64         `gorm:"type:decimal(12,4);column:average_price"`
	TotalValue    float64         `gorm:"type:decimal(12,4);column:total_value"`
	Fee           float64         `gorm:"type:decimal(12,4);column:fee"`
	Date          time.Time       `gorm:"column:date"`
	CreatedAt     time.Time       `gorm:"autoCreateTime,column:created_at"`
	UpdatedAt     time.Time       `gorm:"autoUpdateTime,column:updated_at"`
}

type Transactions struct {
	Id           int             `gorm:"primarykey;size:16"`
	AccountId    int             `gorm:"column:account_id;size:16"`
	SecurityId   int             `gorm:"column:security_id;size:16"`
	Type         TransactionType `gorm:"type:enum('BUY', 'SELL', 'DIVIDEND', 'SPLIT','BONUS','DEMERGER', 'DEMERGER_TRANSFER');column:type;size:16"`
	Quantity     float64         `gorm:"type:decimal(12,4);column:quantity"`
	AveragePrice float64         `gorm:"type:decimal(12,4);column:average_price"`
	TotalValue   float64         `gorm:"type:decimal(12,4);column:total_value"`
	Fee          float64         `gorm:"type:decimal(12,4);column:fee"`
	State        int             `gorm:"column:state;size:11;"`
	Date         time.Time       `gorm:"column:date"`
	CreatedAt    time.Time       `gorm:"autoCreateTime,column:created_at"`
	UpdatedAt    time.Time       `gorm:"autoUpdateTime,column:updated_at"`
}

type Securities struct {
	Id        int       `gorm:"primarykey;size:16"`
	Type      int       `gorm:"index:idx_type_exchange_symbol,unique;column:type;size:16"`
	Exchange  int       `gorm:"index:idx_type_exchange_symbol,unique;column:exchange;size:16"`
	Symbol    string    `gorm:"index:idx_type_exchange_symbol,unique;column:symbol;size:255"`
	Name      string    `gorm:"column:name;size:255"`
	CreatedAt time.Time `gorm:"autoCreateTime,column:created_at"`
	UpdatedAt time.Time `gorm:"autoUpdateTime,column:updated_at"`
}

type InventorySummary struct {
	Id                int     `gorm:"column:id"`
	AccountId         int     `gorm:"column:account_id"`
	SecurityId        int     `gorm:"column:security_id"`
	SecurityExchange  string  `gorm:"column:security_exchange"`
	SecuritySymbol    string  `gorm:"column:security_symbol"`
	SecurityName      string  `gorm:"column:security_name"`
	AvailableQuantity float64 `gorm:"column:available_quantity"`
	TotalValue        float64 `gorm:"column:total_value"`
}
type InventoryDetails struct {
	Id                int       `gorm:"column:id"`
	AvailableQuantity float64   `gorm:"column:available_quantity"`
	Price             float64   `gorm:"column:price"`
	TotalValue        float64   `gorm:"column:total_value"`
	Date              time.Time `gorm:"column:date"`
}

type InventoryLedgers struct {
	Id         int             `gorm:"column:id"`
	Type       TransactionType `gorm:"column:type"`
	Quantity   float64         `gorm:"column:quantity"`
	Price      float64         `gorm:"column:average_price"`
	TotalValue float64         `gorm:"column:total_value"`
	Date       time.Time       `gorm:"column:date"`
}

type DividendTransaction struct {
	Quantity   float64   `gorm:"column:quantity"`
	Price      float64   `gorm:"column:average_price"`
	TotalValue float64   `gorm:"column:total_value"`
	Date       time.Time `gorm:"column:date"`
}

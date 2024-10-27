package domain

import "time"

type TransactionType string

const (
	Buy      TransactionType = "buy"
	Sell     TransactionType = "sell"
	Dividend TransactionType = "dividend"
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
	AvailableQuantiry float64   `gorm:"type:decimal(12,4);column:available_quantity"`
	Price             float64   `gorm:"type:decimal(12,4);column:price"`
	State             int       `gorm:"column:state;size:11;"`
	CreatedAt         time.Time `gorm:"autoCreateTime,column:created_at"`
	UpdatedAt         time.Time `gorm:"autoUpdateTime,column:updated_at"`
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

type Transactions struct {
	Id          int             `gorm:"primarykey;size:16"`
	AccountId   int             `gorm:"column:account_id;size:16"`
	InventoryId int             `gorm:"column:inventory_id;size:16"`
	Type        TransactionType `gorm:"type:enum('buy', 'sell', 'dividend');column:type;size:16"`
	Quantity    float64         `gorm:"type:decimal(12,4);column:quantity"`
	Price       float64         `gorm:"type:decimal(12,4);column:price"`
	Fee         float64         `gorm:"type:decimal(12,4);column:fee"`
	Date        time.Time       `gorm:"column:created_at"`
	CreatedAt   time.Time       `gorm:"autoCreateTime,column:created_at"`
	UpdatedAt   time.Time       `gorm:"autoUpdateTime,column:updated_at"`
}

type InventorySummary struct {
	Id               int     `gorm:"column:id"`
	AccountId        int     `gorm:"column:account_id"`
	SecurityId       int     `gorm:"column:security_id"`
	SecurityExchange string  `gorm:"column:security_exchange"`
	SecuritySymbol   string  `gorm:"column:security_symbol"`
	SecurityName     string  `gorm:"column:security_name"`
	Quantity         float64 `gorm:"column:quantity"`
	Amount           float64 `gorm:"column:amount"`
}
type InventoryDetails struct {
	Id                int       `gorm:"column:id"`
	AvailableQuantiry float64   `gorm:"column:available_quantity"`
	Price             float64   `gorm:"column:price"`
	Date              time.Time `gorm:"column:date"`
}

type InventoryTransactions struct {
	Id       int             `gorm:"column:id"`
	Type     TransactionType `gorm:"column:type"`
	Quantity float64         `gorm:"column:quantity"`
	Price    float64         `gorm:"column:price"`
	Fee      float64         `gorm:"column:fee"`
	Date     time.Time       `gorm:"column:date"`
}

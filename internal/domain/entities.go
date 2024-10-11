package domain

import (
	"time"
)

type TransactionType string

const (
	Buy      TransactionType = "buy"
	Sell     TransactionType = "sell"
	Dividend TransactionType = "dividend"
)

type Account struct {
	Id        int       `gorm:"primarykey;size:16"`
	UserId    int       `gorm:"column:user_id;size:16"`
	Status    int       `gorm:"column:status;size:11"`
	Name      string    `gorm:"column:name;size:255"`
	CreatedAt time.Time `gorm:"autoCreateTime,column:created_at"`
	UpdatedAt time.Time `gorm:"autoUpdateTime,column:updated_at"`
}

type Inventory struct {
	Id         int       `gorm:"primarykey;size:16"`
	AccountId  int       `gorm:"column:account_id;size:16"`
	SecruityId int       `gorm:"column:secruity_id;size:16"`
	State      int       `gorm:"column:state;size:11;"`
	CreatedAt  time.Time `gorm:"autoCreateTime,column:created_at"`
	UpdatedAt  time.Time `gorm:"autoUpdateTime,column:updated_at"`
}

type Security struct {
	Id        int       `gorm:"primarykey;size:16"`
	Type      int       `gorm:"column:type;size:16"`
	Exchange  string    `gorm:"column:exchange;size:50"`
	Symbol    string    `gorm:"column:symbol;size:255"`
	Name      string    `gorm:"column:name;size:255"`
	CreatedAt time.Time `gorm:"autoCreateTime,column:created_at"`
	UpdatedAt time.Time `gorm:"autoUpdateTime,column:updated_at"`
}

type Transaction struct {
	Id          int             `gorm:"primarykey;size:16"`
	AccountId   int             `gorm:"column:account_id;size:16"`
	InventoryId int             `gorm:"column:inventory_id;size:16"`
	Type        TransactionType `gorm:"type:enum('buy', 'sell', 'dividend');column:type;size:16"`
	Quantity    int             `gorm:"column:quantity;size:16"`
	Price       float64         `gorm:"type:type:decimal(10,4);column:price"`
	Fee         float64         `gorm:"type:type:decimal(10,4);column:fee"`
	Date        time.Time       `gorm:"column:created_at"`
	CreatedAt   time.Time       `gorm:"autoCreateTime,column:created_at"`
	UpdatedAt   time.Time       `gorm:"autoUpdateTime,column:updated_at"`
}

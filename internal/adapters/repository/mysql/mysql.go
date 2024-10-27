package mysql

import (
	"assetio/internal/domain"
	"assetio/internal/port"
	"context"

	gormMysql "gorm.io/driver/mysql"

	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

type mysql struct {
	dialer *gorm.DB
	prefix string
}

func New(hostname, port, username, password, name string, prefix string) (port.RepositoryStore, error) {
	dsn := username + ":" + password + "@tcp(" + hostname + ":" + port + ")/" + name + "?charset=utf8mb4&parseTime=True&loc=Local"
	dialer, err := gorm.Open(gormMysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Error),
		NamingStrategy: schema.NamingStrategy{
			TablePrefix: prefix,
		},
	})

	return &mysql{
		dialer: dialer,
		prefix: prefix,
	}, err

}

func (m *mysql) AutoMigrate() {
	m.dialer.AutoMigrate(&domain.Accounts{}, &domain.Inventories{}, &domain.Securities{}, &domain.Transactions{})
}

func (m *mysql) CreateAccount(ctx context.Context, accountData domain.Accounts) (int, error) {
	result := m.dialer.WithContext(ctx).Model(&domain.Accounts{}).Create(&accountData)
	return accountData.Id, result.Error

}

func (m *mysql) GetAccount(ctx context.Context, accountId, userId int) (domain.Accounts, error) {
	var accountData domain.Accounts
	result := m.dialer.WithContext(ctx).Model(&domain.Accounts{}).Select("id", "name", "status").Where("id = ? and user_id = ? ", accountId, userId).First(&accountData)

	if result.Error == gorm.ErrRecordNotFound {
		result.Error = nil
	}
	return accountData, result.Error
}

func (m *mysql) GetAccounts(ctx context.Context, userId int) ([]domain.Accounts, error) {
	var accountsData []domain.Accounts
	result := m.dialer.WithContext(ctx).Model(&domain.Accounts{}).Select("id", "name", "status").Where("user_id = ? ", userId).Find(&accountsData)

	if result.Error == gorm.ErrRecordNotFound {
		result.Error = nil
	}
	return accountsData, result.Error
}

func (m *mysql) UpdateAccount(ctx context.Context, accountId, userId int, accountData domain.Accounts) error {
	result := m.dialer.WithContext(ctx).Model(&domain.Accounts{}).Where("id = ? and user_id = ? ", accountId, userId).Updates(&accountData)
	return result.Error

}

func (m *mysql) CreateSecuriry(ctx context.Context, securityData domain.Securities) (int, error) {
	result := m.dialer.WithContext(ctx).Model(&domain.Securities{}).Create(&securityData)
	return securityData.Id, result.Error

}

func (m *mysql) GetSecuriry(ctx context.Context, types, exchange int, symbol string) (domain.Securities, error) {
	var securityData domain.Securities

	result := m.dialer.WithContext(ctx).Model(&domain.Securities{}).Select("id").Where("type = ? and exchange = ? and symbol =? ", types, exchange, symbol).First(&securityData)
	if result.Error == gorm.ErrRecordNotFound {
		result.Error = nil
	}

	return securityData, result.Error
}

func (m *mysql) GetSecuriryById(ctx context.Context, securityId int) (domain.Securities, error) {
	var securityData domain.Securities

	result := m.dialer.WithContext(ctx).Model(&domain.Securities{}).Select("id", "type", "exchange", "symbol", "name").Where("id = ?", securityId).First(&securityData)
	if result.Error == gorm.ErrRecordNotFound {
		result.Error = nil
	}

	return securityData, result.Error
}

func (m *mysql) UpdateSecuriry(ctx context.Context, securityId int, securityData domain.Securities) error {
	result := m.dialer.WithContext(ctx).Model(&domain.Securities{}).Where("id = ?", securityId).Updates(&securityData)
	return result.Error

}

func (m *mysql) GetSecurities(ctx context.Context, types, exchange int) ([]domain.Securities, error) {
	var securitiesData []domain.Securities

	result := m.dialer.WithContext(ctx).Model(&domain.Securities{}).Select("id", "type", "exchange", "symbol", "name").Where("type = ? and exchange = ?", types, exchange).Find(&securitiesData)
	if result.Error == gorm.ErrRecordNotFound {
		result.Error = nil
	}

	return securitiesData, result.Error
}

func (m *mysql) SearchSecurities(ctx context.Context, types, exchange int, search string) ([]domain.Securities, error) {
	var securitiesData []domain.Securities

	result := m.dialer.WithContext(ctx).Model(&domain.Securities{}).Select("id", "type", "exchange", "symbol", "name").Where("type = ? and exchange = ? and (name LIKE ? or symbol LIKE ?)", types, exchange, "%"+search+"%", "%"+search+"%").Find(&securitiesData)
	if result.Error == gorm.ErrRecordNotFound {
		result.Error = nil
	}

	return securitiesData, result.Error
}

func (m *mysql) InsertTransaction(ctx context.Context, transactionData domain.Transactions) (domain.Transactions, error) {
	result := m.dialer.WithContext(ctx).Model(&domain.Transactions{}).Create(&transactionData)
	return transactionData, result.Error
}
func (m *mysql) InsertInventory(ctx context.Context, inventoryData domain.Inventories) (domain.Inventories, error) {
	result := m.dialer.WithContext(ctx).Model(&domain.Inventories{}).Create(&inventoryData)
	return inventoryData, result.Error
}

func (m *mysql) GetInventoryById(ctx context.Context, inventoryId int) (domain.Inventories, error) {
	var inventoryData domain.Inventories

	result := m.dialer.WithContext(ctx).Model(&domain.Inventories{}).Select("id", "account_id", "security_id", "available_quantity").Where("id = ? ", inventoryId).Find(&inventoryData)
	if result.Error == gorm.ErrRecordNotFound {
		result.Error = nil
	}

	return inventoryData, result.Error
}
func (m *mysql) UpdateAvailableQuanityToInventoryById(ctx context.Context, inventoryId int, quantity float64) error {
	result := m.dialer.WithContext(ctx).
		Model(&domain.Inventories{}).
		Where("id = ?", inventoryId).
		Updates(map[string]interface{}{
			"available_quantity": quantity,
		})

	return result.Error

}
func (m *mysql) GetActiveInventoriesByAccountIdAndSecurityId(ctx context.Context, accountId, securityId int) ([]domain.Inventories, error) {
	var InventoriesData []domain.Inventories

	result := m.dialer.WithContext(ctx).Model(&domain.Inventories{}).Select("id", "available_quantity").
		Where("account_id = ? and security_id = ? and available_quantity > 0", accountId, securityId).
		Find(&InventoriesData)

	if result.Error == gorm.ErrRecordNotFound {
		result.Error = nil
	}

	return InventoriesData, result.Error
}

func (m *mysql) SelectInvertriesSummaryByAccountIdAndSecurityType(ctx context.Context, accountId, securityType int) ([]domain.InventorySummary, error) {
	var inventoryData []domain.InventorySummary

	result := m.dialer.WithContext(ctx).
		Model(&domain.Inventories{}).
		Select(m.prefix+"inventories.id", m.prefix+"inventories.account_id", m.prefix+"inventories.security_id", m.prefix+"securities.name as security_name", m.prefix+"securities.exchange as security_exchange", m.prefix+"securities.symbol as security_symbol", "SUM("+m.prefix+"inventories.available_quantity) as quantity", "SUM("+m.prefix+"inventories.available_quantity * "+m.prefix+"inventories.price) as amount").
		Joins("JOIN "+m.prefix+"securities ON (security_id = "+m.prefix+"securities.id and type = ? )", securityType).
		Where("account_id = ? and available_quantity > 0 ", accountId).
		Group("security_id").
		Find(&inventoryData)

	if result.Error == gorm.ErrRecordNotFound {
		result.Error = nil
	}

	return inventoryData, result.Error
}

func (m *mysql) SelectInvertriesByAccountIdAndStockId(ctx context.Context, accountId, securityId int) ([]domain.InventoryDetails, error) {
	var inventoryData []domain.InventoryDetails

	result := m.dialer.WithContext(ctx).
		Model(&domain.Inventories{}).
		Select("id", "available_quantity", "price", "created_at as date").
		Where("account_id = ? and security_id = ? and available_quantity > 0 ", accountId, securityId).
		Find(&inventoryData)

	if result.Error == gorm.ErrRecordNotFound {
		result.Error = nil
	}

	return inventoryData, result.Error
}

func (m *mysql) SelectInvertriesTransactionByIdAndAccountId(ctx context.Context, accountId, inventoryId int) ([]domain.InventoryTransactions, error) {
	var inventoryData []domain.InventoryTransactions

	result := m.dialer.WithContext(ctx).
		Model(&domain.Transactions{}).
		Select("id", "type", "quantity", "price", "fee", "created_at as date").
		Where("account_id = ? and inventory_id = ?", accountId, inventoryId).
		Find(&inventoryData)

	if result.Error == gorm.ErrRecordNotFound {
		result.Error = nil
	}

	return inventoryData, result.Error
}

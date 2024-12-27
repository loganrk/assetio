package mysql

import (
	"assetio/internal/domain"
	"assetio/internal/port"
	"context"
	"time"

	gormMysql "gorm.io/driver/mysql"

	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

type mysql struct {
	dialer *gorm.DB
	prefix string
}

// New creates a new MySQL database connection using GORM with custom configurations.
// It takes database credentials and a prefix for table naming conventions, returning a
// RepositoryStore interface for accessing database methods or an error if connection fails.
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

// AutoMigrate automatically migrates all defined models, creating or updating tables
// to match the structs in the domain package. Used for schema versioning.
func (m *mysql) AutoMigrate() {
	m.dialer.AutoMigrate(&domain.Accounts{}, &domain.Securities{}, &domain.Inventories{}, &domain.InventoryLedger{}, &domain.Transactions{})
}

// InsertAccountData adds a new account entry to the Accounts table.
// Returns the created account data along with any error encountered during insertion.
func (m *mysql) InsertAccountData(ctx context.Context, accountData domain.Accounts) (domain.Accounts, error) {
	// Create a new record in the Accounts table with the provided account data
	result := m.dialer.WithContext(ctx).Model(&domain.Accounts{}).Create(&accountData)
	return accountData, result.Error
}

// GetAccountDataByIdAndUserId fetches account details based on account ID and user ID.
// Only selects specific fields: id, name, and status. Returns the account data if found, or nil if no matching record exists.
func (m *mysql) GetAccountDataByIdAndUserId(ctx context.Context, accountId, userId int) (domain.Accounts, error) {
	var accountData domain.Accounts

	// Query the Accounts table for a record that matches the specified account ID and user ID
	result := m.dialer.WithContext(ctx).Model(&domain.Accounts{}).
		Select("id", "name", "status").
		Where("id = ? and user_id = ?", accountId, userId).
		First(&accountData)

	// Set result.Error to nil if no record is found to avoid a "record not found" error
	if result.Error == gorm.ErrRecordNotFound {
		result.Error = nil
	}
	return accountData, result.Error
}

// GetAccountsData retrieves all accounts associated with the specified user ID.
// Returns a slice of account records or an empty slice if none are found.
func (m *mysql) GetAccountsData(ctx context.Context, userId int) ([]domain.Accounts, error) {
	var accountsData []domain.Accounts

	// Query the Accounts table for records that match the specified user ID
	result := m.dialer.WithContext(ctx).Model(&domain.Accounts{}).
		Select("id", "name", "status").
		Where("user_id = ?", userId).
		Find(&accountsData)

	// Set result.Error to nil if no records are found, avoiding a "record not found" error
	if result.Error == gorm.ErrRecordNotFound {
		result.Error = nil
	}
	return accountsData, result.Error
}

// UpdateAccountData updates account information for a specific account and user ID.
// Only updates the account record if both account ID and user ID match.
func (m *mysql) UpdateAccountData(ctx context.Context, accountId, userId int, accountData domain.Accounts) error {
	// Update the account record where the ID and user ID match the specified values
	result := m.dialer.WithContext(ctx).Model(&domain.Accounts{}).
		Where("id = ? and user_id = ?", accountId, userId).
		Updates(&accountData)
	return result.Error
}

// InsertSecurityData adds a new security entry to the Securities table.
// Returns the created security data along with any error encountered during insertion.
func (m *mysql) InsertSecurityData(ctx context.Context, securityData domain.Securities) (domain.Securities, error) {
	// Create a new record in the Securities table with the provided security data
	result := m.dialer.WithContext(ctx).Model(&domain.Securities{}).Create(&securityData)
	return securityData, result.Error
}

// GetSecurityDataById retrieves security information based on the security ID.
// It selects specific fields: id, type, exchange, symbol, and name.
// Returns the security data if found, or nil if no matching record exists.
func (m *mysql) GetSecurityDataById(ctx context.Context, securityId int) (domain.Securities, error) {
	var securityData domain.Securities

	// Query the Securities table for a record matching the specified security ID
	result := m.dialer.WithContext(ctx).
		Model(&domain.Securities{}).
		Select("id", "type", "exchange", "symbol", "name").
		Where("id = ?", securityId).
		First(&securityData)

	// Set result.Error to nil if no record is found to avoid a "record not found" error
	if result.Error == gorm.ErrRecordNotFound {
		result.Error = nil
	}
	return securityData, result.Error
}

// GetSecurityDataByTypeAndExchangeAndSymbol fetches security information based on the given type, exchange, and symbol.
// Returns the security data if found, or nil if no matching record exists.
func (m *mysql) GetSecurityDataByTypeAndExchangeAndSymbol(ctx context.Context, types, exchange int, symbol string) (domain.Securities, error) {
	var securityData domain.Securities

	// Query the Securities table to find the security that matches the provided type, exchange, and symbol
	result := m.dialer.WithContext(ctx).Model(&domain.Securities{}).
		Select("id").
		Where("type = ? and exchange = ? and symbol = ?", types, exchange, symbol).
		First(&securityData)

	// If no record is found, set result.Error to nil to avoid returning a "record not found" error
	if result.Error == gorm.ErrRecordNotFound {
		result.Error = nil
	}
	return securityData, result.Error
}

// UpdateSecurityData modifies security information for a specified security ID with the provided security data.
// Returns any error encountered during the update.
func (m *mysql) UpdateSecurityData(ctx context.Context, securityId int, securityData domain.Securities) error {
	// Update the security record where the ID matches the specified securityId
	result := m.dialer.WithContext(ctx).Model(&domain.Securities{}).
		Where("id = ?", securityId).
		Updates(&securityData)
	return result.Error
}

// GetSecuritiesDataByType retrieves all securities data that match the given type and exchange.
// Returns a slice of Securities if records are found, or an empty slice if no records match.
func (m *mysql) GetSecuritiesDataByType(ctx context.Context, types int) ([]domain.Securities, error) {
	var securitiesData []domain.Securities

	// Query the Securities table for records that match the given type and exchange
	result := m.dialer.WithContext(ctx).Model(&domain.Securities{}).
		Select("id", "type", "exchange", "symbol", "name").
		Where("type = ? ", types).
		Find(&securitiesData)

	// Set result.Error to nil if no record is found, preventing "record not found" error
	if result.Error == gorm.ErrRecordNotFound {
		result.Error = nil
	}
	return securitiesData, result.Error
}

// SearchSecuritiesDataByTypeAndExchange performs a search for securities by type, exchange, name, or symbol.
// The search term is matched partially with both the name and symbol fields using SQL LIKE.
// Returns a slice of matching Securities records or an empty slice if none found.
func (m *mysql) SearchSecuritiesDataByTypeAndExchange(ctx context.Context, types, exchange int, search string) ([]domain.Securities, error) {
	var securitiesData []domain.Securities

	// Query the Securities table to find records that match the type, exchange, and partially match the search term in name or symbol
	result := m.dialer.WithContext(ctx).Model(&domain.Securities{}).
		Select("id", "type", "exchange", "symbol", "name").
		Where("type = ? and exchange = ? and (name LIKE ? or symbol LIKE ?)", types, exchange, "%"+search+"%", "%"+search+"%").
		Find(&securitiesData)

	// Set result.Error to nil if no record is found
	if result.Error == gorm.ErrRecordNotFound {
		result.Error = nil
	}
	return securitiesData, result.Error
}

// InsertInventoryLedger adds a new entry to the InventoryLedger table with the provided inventory ledger data.
// Returns the inserted inventory ledger data along with any error encountered.
func (m *mysql) InsertInventoryLedger(ctx context.Context, inventoryLedgerData domain.InventoryLedger) (domain.InventoryLedger, error) {
	// Insert the new ledger entry into the InventoryLedger table
	result := m.dialer.WithContext(ctx).Model(&domain.InventoryLedger{}).Create(&inventoryLedgerData)
	return inventoryLedgerData, result.Error
}

// UpdateInventoryDetailsById updates an inventory record by its ID with the provided inventory data.
// Returns any error encountered during the update.
func (m *mysql) UpdateInventoryDetailsById(ctx context.Context, inventoryId int, availableQuantity, averagePrice, totalValue float64) error {
	// Update the inventory record where the ID matches the given inventoryId
	result := m.dialer.WithContext(ctx).Model(&domain.Inventories{}).Where("id = ?", inventoryId).Updates(map[string]interface{}{
		"available_quantity": availableQuantity,
		"average_price":      averagePrice,
		"total_value":        totalValue,
	})
	return result.Error
}

// InsertTransaction adds a new transaction entry to the Transactions table using the provided transaction data.
// Returns the inserted transaction data along with any error encountered.
func (m *mysql) InsertTransaction(ctx context.Context, transactionData domain.Transactions) (domain.Transactions, error) {
	// Insert the new transaction entry into the Transactions table
	result := m.dialer.WithContext(ctx).Model(&domain.Transactions{}).Create(&transactionData)
	return transactionData, result.Error
}

// UpdateInventoryLedgerTransactionIdById updates the transaction ID in an inventory ledger record by its ledger ID.
// Returns any error encountered during the update.
func (m *mysql) UpdateInventoryLedgerTransactionIdById(ctx context.Context, ledgerId, transactionId int) error {
	// Update the transaction_id field in the InventoryLedger record with the specified ledgerId
	result := m.dialer.WithContext(ctx).Model(&domain.InventoryLedger{}).Where("id = ? ", ledgerId).Updates(map[string]interface{}{
		"transaction_id": transactionId,
	})
	return result.Error
}

// UpdateInventoryLedgerTransactionIdByIds updates the transaction ID for multiple inventory ledger records based on a list of ledger IDs.
// It sets the transaction_id field for all records where the id is in the provided ledgerIds slice.
func (m *mysql) UpdateInventoryLedgerTransactionIdByIds(ctx context.Context, ledgerIds []int, transactionId int) error {
	// Update the transaction_id field for records with IDs in the ledgerIds slice
	result := m.dialer.WithContext(ctx).Model(&domain.InventoryLedger{}).
		Where("id IN ?", ledgerIds).
		Updates(map[string]interface{}{
			"transaction_id": transactionId,
		})
	return result.Error
}

// UpdateAvailableQuantityToInventoryById sets the available quantity in an inventory record by its inventory ID.
// Returns any error encountered during the update.
func (m *mysql) UpdateAvailableQuantityToInventoryById(ctx context.Context, inventoryId int, quantity float64) error {
	// Update the available_quantity field in the Inventories record with the specified inventoryId
	result := m.dialer.WithContext(ctx).Model(&domain.Inventories{}).Where("id = ?", inventoryId).Updates(map[string]interface{}{
		"available_quantity": quantity,
	})
	return result.Error
}

// GetInvertriesSummaryByAccountIdAndSecurityType retrieves a summary of inventories for a specific account ID and security type.
// The summary includes security details (name, exchange, symbol) and aggregate values for available quantity and total value.
func (m *mysql) GetInvertriesSummaryByAccountIdAndSecurityType(ctx context.Context, accountId, securityType int) ([]domain.InventorySummary, error) {
	var inventoryData []domain.InventorySummary

	// Query to join inventories with securities to get a summary, filtered by account ID and security type
	result := m.dialer.WithContext(ctx).
		Model(&domain.Inventories{}).
		Select(m.prefix+"inventories.id", m.prefix+"inventories.account_id", m.prefix+"inventories.security_id", m.prefix+"securities.name as security_name", m.prefix+"securities.exchange as security_exchange", m.prefix+"securities.symbol as security_symbol", "SUM("+m.prefix+"inventories.available_quantity) as available_quantity", "SUM("+m.prefix+"inventories.total_value) as total_value").
		Joins("JOIN "+m.prefix+"securities ON (security_id = "+m.prefix+"securities.id and type = ? )", securityType).
		Where("account_id = ? and available_quantity > 0 ", accountId).
		Group("security_id"). // Group by security_id to get summary data per security
		Find(&inventoryData)

	// If no record found, set error to nil for empty results
	if result.Error == gorm.ErrRecordNotFound {
		result.Error = nil
	}

	return inventoryData, result.Error
}

// GetInvertriesByAccountIdAndSecurityId retrieves detailed inventory records for a specific account and security.
// The details include ID, available quantity, and total value, ordered by the creation date in descending order.
func (m *mysql) GetInvertriesByAccountIdAndSecurityId(ctx context.Context, accountId, securityId int) ([]domain.InventoryDetails, error) {
	var inventoryData []domain.InventoryDetails

	// Query to get inventory details based on account and security IDs, ordered by the latest creation date
	result := m.dialer.WithContext(ctx).
		Model(&domain.Inventories{}).
		Select("id", "available_quantity", "total_value", "date").
		Where("account_id = ? and security_id = ? and available_quantity > 0 ", accountId, securityId).
		Order("date desc"). // Retrieve the latest data first
		Find(&inventoryData)

	// If no record found, set error to nil for empty results
	if result.Error == gorm.ErrRecordNotFound {
		result.Error = nil
	}

	return inventoryData, result.Error
}

// InsertTransactionData adds a new transaction entry to the Transactions table.
// Returns the created transaction data along with any error encountered during insertion.
func (m *mysql) InsertTransactionData(ctx context.Context, transactionData domain.Transactions) (domain.Transactions, error) {
	// Create a new record in the Transactions table with the provided transaction data
	result := m.dialer.WithContext(ctx).Model(&domain.Transactions{}).Create(&transactionData)
	return transactionData, result.Error
}

// InsertInventoryData adds a new inventory entry to the Inventories table.
// Returns the created inventory data along with any error encountered during insertion.
func (m *mysql) InsertInventoryData(ctx context.Context, inventoryData domain.Inventories) (domain.Inventories, error) {
	// Create a new record in the Inventories table with the provided inventory data
	result := m.dialer.WithContext(ctx).Model(&domain.Inventories{}).Create(&inventoryData)
	return inventoryData, result.Error
}

// GetInventoryDataById retrieves an inventory record by its ID.
// It fetches specific fields: id, account_id, security_id, available_quantity, and total_value.
// Returns the inventory data if found, or nil if no matching record exists.
func (m *mysql) GetInventoryDataById(ctx context.Context, inventoryId int) (domain.Inventories, error) {
	var inventoryData domain.Inventories

	// Query the Inventories table for a record matching the specified inventory ID
	result := m.dialer.WithContext(ctx).Model(&domain.Inventories{}).
		Select("*").
		Where("id = ?", inventoryId).
		Find(&inventoryData)

	// Set result.Error to nil if no record is found to avoid a "record not found" error
	if result.Error == gorm.ErrRecordNotFound {
		result.Error = nil
	}
	return inventoryData, result.Error
}

// UpdateAvailableQuantityToInventoryById updates the available quantity for a specific inventory record by its ID.
// Takes the inventory ID and the new quantity as parameters.
// Returns an error if the update operation fails.
func (m *mysql) UpdateAvailableQuanityToInventoryById(ctx context.Context, inventoryId int, quantity float64) error {
	// Update the available_quantity field for the record with the specified inventory ID
	result := m.dialer.WithContext(ctx).
		Model(&domain.Inventories{}).
		Where("id = ?", inventoryId).
		Updates(map[string]interface{}{
			"available_quantity": quantity,
		})

	return result.Error
}

// GetActiveInventoriesByAccountIdAndSecurityId retrieves active inventory records with available quantities greater than zero
// for a specific account and security. Returns a list of inventory records with ID and available quantity fields.
func (m *mysql) GetActiveInventoriesByAccountIdAndSecurityId(ctx context.Context, accountId, securityId int) ([]domain.Inventories, error) {
	var InventoriesData []domain.Inventories

	// Query to find active inventories based on account and security IDs with positive available quantity
	result := m.dialer.WithContext(ctx).Model(&domain.Inventories{}).Select("id", "available_quantity", "date").
		Where("account_id = ? and security_id = ? and available_quantity > 0", accountId, securityId).
		Order("id"). // Fetch old data first by ordering by ID
		Find(&InventoriesData)

	// If no record found, set error to nil to avoid returning an error for empty results
	if result.Error == gorm.ErrRecordNotFound {
		result.Error = nil
	}

	return InventoriesData, result.Error
}

// GetInventoryLedgersByInventoryIdAndAccountId retrieves ledger entries associated with a specific inventory ID, ordered by date.
// Each ledger entry includes ID, type, quantity, average price, total value, and date fields.
func (m *mysql) GetInventoryLedgersByInventoryIdAndAccountId(ctx context.Context, accountId, inventoryId int) ([]domain.InventoryLedgers, error) {
	var inventoryLedgerData []domain.InventoryLedgers

	// Query to find inventory ledger data based on inventory ID, ordered by date in descending order
	result := m.dialer.WithContext(ctx).
		Model(&domain.InventoryLedger{}).
		Select("id", "type", "quantity", "average_price", "total_value", "date").
		Where("account_id =? and inventory_id = ?", accountId, inventoryId).
		Order("date desc"). // Fetch the latest data first
		Find(&inventoryLedgerData)

	// If no record found, set error to nil for empty results
	if result.Error == gorm.ErrRecordNotFound {
		result.Error = nil
	}

	return inventoryLedgerData, result.Error
}

func (m *mysql) GetDividendTransactionsByAccountIdAndSecurityId(ctx context.Context, accountId, securityId int) ([]domain.DividendTransaction, error) {
	var transactionsData []domain.DividendTransaction

	// Query to find inventory ledger data based on inventory ID, ordered by date in descending order
	result := m.dialer.WithContext(ctx).
		Model(&domain.Transactions{}).
		Select("quantity", "average_price", "total_value", "date").
		Where("account_id =? and security_id = ? and type =?", accountId, securityId, domain.DIVIDEND).
		Order("date desc"). // Fetch the latest data first
		Find(&transactionsData)

		// If no record found, set error to nil for empty results
	if result.Error == gorm.ErrRecordNotFound {
		result.Error = nil
	}
	return transactionsData, result.Error

}

func (m *mysql) GetInventoryAvailableQuanitityBySecurityIdAndDate(ctx context.Context, accountId, securityId int, date time.Time) (float64, error) {
	var totalQuantity float64

	// Query to calculate the total quantity

	result := m.dialer.WithContext(ctx).
		Model(&domain.Transactions{}).
		Select(`SUM(
        CASE 
            WHEN type = ? AND date < ? THEN quantity   
            WHEN type = ?  AND date <= ? THEN -quantity  
            ELSE 0                            
        END
    ) as total_quantity`, domain.BUY, date, domain.SELL, date).
		Where("account_id = ? AND security_id = ?", accountId, securityId).
		Scan(&totalQuantity)

		// If no record found, set error to nil for empty results
	if result.Error == gorm.ErrRecordNotFound {
		result.Error = nil
	}

	return totalQuantity, result.Error

}

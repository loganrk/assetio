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
}

func New(hostname, port, username, password, name string, prefix string) (port.RepositoryMySQL, error) {
	dsn := username + ":" + password + "@tcp(" + hostname + ":" + port + ")/" + name + "?charset=utf8mb4&parseTime=True&loc=Local"
	dialer, err := gorm.Open(gormMysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Error),
		NamingStrategy: schema.NamingStrategy{
			TablePrefix: prefix,
		},
	})

	return &mysql{
		dialer: dialer,
	}, err

}

func (m *mysql) AutoMigrate() {
	m.dialer.AutoMigrate(&domain.Account{}, &domain.Inventory{}, &domain.Security{}, &domain.Transaction{})
}

func (m *mysql) CreateAccount(ctx context.Context, accountData domain.Account) (int, error) {
	result := m.dialer.WithContext(ctx).Model(&domain.Account{}).Create(&accountData)
	return accountData.Id, result.Error

}

func (m *mysql) GetAccount(ctx context.Context, accountId, userId int) (domain.Account, error) {
	var accountData domain.Account
	result := m.dialer.WithContext(ctx).Model(&domain.Account{}).Select("id", "name", "status").Where("id = ? and user_id = ? ", accountId, userId).First(&accountData)

	if result.Error == gorm.ErrRecordNotFound {
		result.Error = nil
	}
	return accountData, result.Error
}

func (m *mysql) GetAccounts(ctx context.Context, userId int) ([]domain.Account, error) {
	var accountsData []domain.Account
	result := m.dialer.WithContext(ctx).Model(&domain.Account{}).Select("id", "name", "status").Where("user_id = ? ", userId).Find(&accountsData)

	if result.Error == gorm.ErrRecordNotFound {
		result.Error = nil
	}
	return accountsData, result.Error
}

func (m *mysql) UpdateAccount(ctx context.Context, accountId, userId int, accountData domain.Account) error {
	result := m.dialer.WithContext(ctx).Model(&domain.Account{}).Where("id = ? and user_id = ? ", accountId, userId).Updates(&accountData)
	return result.Error

}

func (m *mysql) CreateSecuriry(ctx context.Context, securityData domain.Security) (int, error) {
	result := m.dialer.WithContext(ctx).Model(&domain.Security{}).Create(&securityData)
	return securityData.Id, result.Error

}

func (m *mysql) GetSecuriry(ctx context.Context, types, exchange int, symbol string) (domain.Security, error) {
	var securityData domain.Security

	result := m.dialer.WithContext(ctx).Model(&domain.Security{}).Select("id").Where("type = ? and exchange = ? and symbol =? ", types, exchange, symbol).First(&securityData)
	if result.Error == gorm.ErrRecordNotFound {
		result.Error = nil
	}

	return securityData, result.Error
}

func (m *mysql) GetSecuriryById(ctx context.Context, secruityId int) (domain.Security, error) {
	var securityData domain.Security

	result := m.dialer.WithContext(ctx).Model(&domain.Security{}).Select("id", "type", "exchange", "symbol", "name").Where("id = ?", secruityId).First(&securityData)
	if result.Error == gorm.ErrRecordNotFound {
		result.Error = nil
	}

	return securityData, result.Error
}

func (m *mysql) UpdateSecuriry(ctx context.Context, secruityId int, securityData domain.Security) error {
	result := m.dialer.WithContext(ctx).Model(&domain.Security{}).Where("id = ?", secruityId).Updates(&securityData)
	return result.Error

}

func (m *mysql) GetSecurities(ctx context.Context, types, exchange int) ([]domain.Security, error) {
	var securitiesData []domain.Security

	result := m.dialer.WithContext(ctx).Model(&domain.Security{}).Select("id").Where("type = ? and exchange = ?", types, exchange).Find(&securitiesData)
	if result.Error == gorm.ErrRecordNotFound {
		result.Error = nil
	}

	return securitiesData, result.Error
}

func (m *mysql) SearchSecurities(ctx context.Context, types, exchange int, search string) ([]domain.Security, error) {
	var securitiesData []domain.Security

	result := m.dialer.WithContext(ctx).Model(&domain.Security{}).Select("id").Where("type = ? and exchange = ? and (name =? or symbol ?)", types, exchange, "%"+search+"%", "%"+search+"%").Find(&securitiesData)
	if result.Error == gorm.ErrRecordNotFound {
		result.Error = nil
	}

	return securitiesData, result.Error
}

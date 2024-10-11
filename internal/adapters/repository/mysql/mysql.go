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

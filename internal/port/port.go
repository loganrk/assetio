package port

import (
	"assetio/internal/domain"
	"context"
	"net/http"
	"time"
)

type Handler interface {
	AccountCreate(w http.ResponseWriter, r *http.Request)
	AccountAll(w http.ResponseWriter, r *http.Request)
	AccountGet(w http.ResponseWriter, r *http.Request)
	AccountUpdate(w http.ResponseWriter, r *http.Request)
	AccountActivate(w http.ResponseWriter, r *http.Request)
	AccountInactivate(w http.ResponseWriter, r *http.Request)

	SecurityCreate(w http.ResponseWriter, r *http.Request)
	SecurityUpdate(w http.ResponseWriter, r *http.Request)
	SecurityAll(w http.ResponseWriter, r *http.Request)
	SecurityGet(w http.ResponseWriter, r *http.Request)
	SecuritySearch(w http.ResponseWriter, r *http.Request)
}

type RepositoryMySQL interface {
	AutoMigrate()
	CreateAccount(ctx context.Context, accountData domain.Account) (int, error)
	GetAccounts(ctx context.Context, userId int) ([]domain.Account, error)
	GetAccount(ctx context.Context, accountId, userId int) (domain.Account, error)
	UpdateAccount(ctx context.Context, accountId, userId int, accountData domain.Account) error

	CreateSecuriry(ctx context.Context, securityData domain.Security) (int, error)
	GetSecuriry(ctx context.Context, types, exchange int, symbol string) (domain.Security, error)
	GetSecuriryById(ctx context.Context, secruityId int) (domain.Security, error)
	UpdateSecuriry(ctx context.Context, secruityId int, securityData domain.Security) error
}

type Router interface {
	RegisterRoute(method, path string, handlerFunc http.HandlerFunc)
	StartServer(port string) error
	UseBefore(middlewares ...http.HandlerFunc)
	NewGroup(groupName string) RouterGroup
}

type RouterGroup interface {
	RegisterRoute(method, path string, handlerFunc http.HandlerFunc)
	UseBefore(middlewares ...http.HandlerFunc)
}

type Cipher interface {
	Encrypt(text string) (string, error)
	Decrypt(cryptoText string) (string, error)
	GetKey() string
}

type Token interface {
	GetAccessTokenData(encryptedToken string) (int, time.Time, error)
}

type Auth interface {
	ValidateApiKey() http.HandlerFunc
	ValidateAccessToken() http.HandlerFunc
}

type Logger interface {
	Debug(ctx context.Context, messages ...any)
	Info(ctx context.Context, messages ...any)
	Warn(ctx context.Context, messages ...any)
	Error(ctx context.Context, messages ...any)
	Fatal(ctx context.Context, messages ...any)
	Debugf(ctx context.Context, template string, args ...any)
	Infof(ctx context.Context, template string, args ...any)
	Warnf(ctx context.Context, template string, args ...any)
	Errorf(ctx context.Context, template string, args ...any)
	Fatalf(ctx context.Context, template string, args ...any)
	Debugw(ctx context.Context, msg string, keysAndValues ...any)
	Infow(ctx context.Context, msg string, keysAndValues ...any)
	Warnw(ctx context.Context, msg string, keysAndValues ...any)
	Errorw(ctx context.Context, msg string, keysAndValues ...any)
	Fatalw(ctx context.Context, msg string, keysAndValues ...any)
	Sync(ctx context.Context) error
}

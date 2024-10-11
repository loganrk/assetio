package port

import (
	"assetio/internal/domain"
	"context"
	"net/http"
	"time"
)

type Handler interface {
	AccountAdd(w http.ResponseWriter, r *http.Request)
	AccountUpdate(w http.ResponseWriter, r *http.Request)
}

type RepositoryMySQL interface {
	AutoMigrate()
	CreateAccount(ctx context.Context, accountData domain.Account) (int, error)
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

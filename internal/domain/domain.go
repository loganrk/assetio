package domain

import (
	"context"
)

type List struct {
	Account AccountSvr
}

type AccountSvr interface {
	CreateAccount(ctx context.Context, userId int, name string) (int, error)
}

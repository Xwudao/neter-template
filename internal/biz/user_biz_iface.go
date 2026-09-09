package biz

import (
	"context"

	"github.com/Xwudao/neter-template/internal/data/sqlc"
	"github.com/Xwudao/neter-template/internal/domain/params"
)

// UserBizIface is the interface consumed by route handlers.
// Defining the interface here (alongside its implementation) keeps the biz
// package as the single source of truth; wire.Bind maps *UserBiz → UserBizIface.
type UserBizIface interface {
	Login(ctx context.Context, p *params.UserLoginParams) (*sqlc.User, string, error)
	Delete(ctx context.Context, id int64) error
	Get(ctx context.Context, id int64) (*sqlc.User, error)
	GetBy(ctx context.Context, p *params.GetUserByParams) (*sqlc.User, error)
	Create(ctx context.Context, p *params.CreateUserParams) (*sqlc.User, error)
	GetAll(ctx context.Context) ([]*sqlc.User, error)
}

// Compile-time assertion: *UserBiz must satisfy UserBizIface.
var _ UserBizIface = (*UserBiz)(nil)

package biz

import (
	"context"
	"errors"

	"go.uber.org/zap"

	"github.com/Xwudao/neter-template/internal/data/ent"
	"github.com/Xwudao/neter-template/internal/domain/params"
	"github.com/Xwudao/neter-template/internal/system"
	"github.com/Xwudao/neter-template/pkg/utils/bcrypt"
	"github.com/Xwudao/neter-template/pkg/utils/jwt"
)

type UserRepository interface {
	GetAll(ctx context.Context) ([]*ent.User, error)
	DeleteByID(ctx context.Context, id int64) error
	GetByID(ctx context.Context, id int64) (*ent.User, error)
	GetBy(ctx context.Context, p *params.GetUserByParams) (*ent.User, error)
	Create(ctx context.Context, p *params.CreateUserParams) (*ent.User, error)
}

type UserBiz struct {
	log    *zap.SugaredLogger
	appCtx *system.AppContext
	ur     UserRepository

	jwt *jwt.Client
}

func NewUserBiz(log *zap.SugaredLogger, ur UserRepository, jwt *jwt.Client, appCtx *system.AppContext) *UserBiz {
	return &UserBiz{
		log:    log.Named("user-biz"),
		appCtx: appCtx,
		ur:     ur, jwt: jwt,
	}
}

func (h *UserBiz) Index() string {
	panic("TODO implement")
}

func (h *UserBiz) Delete(ctx context.Context, id int64) error {
	return h.ur.DeleteByID(ctx, id)
}

func (h *UserBiz) Get(ctx context.Context, id int64) (*ent.User, error) {
	return h.ur.GetByID(ctx, id)
}

func (h *UserBiz) GetBy(ctx context.Context, p *params.GetUserByParams) (*ent.User, error) {
	return h.ur.GetBy(ctx, p)
}

// Login 登录
func (h *UserBiz) Login(ctx context.Context, p *params.UserLoginParams) (*ent.User, string, error) {
	us, err := h.ur.GetBy(ctx, &params.GetUserByParams{Username: p.Username})
	if err != nil {
		return nil, "", err
	}

	ok := bcrypt.ValidatePassword(us.Password, p.Password)
	if !ok {
		return nil, "", errors.New("密码错误")
	}

	token, err := h.jwt.Generate(us.ID)
	if err != nil {
		return nil, "", err
	}

	return us, token, nil
}

func (h *UserBiz) Create(ctx context.Context, p *params.CreateUserParams) (*ent.User, error) {

	generatePassword, err := bcrypt.GeneratePassword(p.Password)
	if err != nil {
		return nil, err
	}
	p.Password = generatePassword

	return h.ur.Create(ctx, p)
}

func (h *UserBiz) GetAll(ctx context.Context) ([]*ent.User, error) {
	return h.ur.GetAll(ctx)
}

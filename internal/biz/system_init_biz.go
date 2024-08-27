package biz

import (
	"go.uber.org/zap"

	"github.com/Xwudao/neter-template/internal/data/ent"
	"github.com/Xwudao/neter-template/internal/data/ent/user"
	"github.com/Xwudao/neter-template/internal/domain/params"
	"github.com/Xwudao/neter-template/internal/system"
	"github.com/Xwudao/neter-template/pkg/utils"
	"github.com/Xwudao/neter-template/pkg/utils/bcrypt"
)

type SystemInitBiz struct {
	log    *zap.SugaredLogger
	appCtx *system.AppContext
	ur     UserRepository
}

func NewSystemInitBiz(log *zap.SugaredLogger, us UserRepository, appCtx *system.AppContext) *SystemInitBiz {
	return &SystemInitBiz{
		log:    log.Named("system-init-biz"),
		appCtx: appCtx,
		ur:     us,
	}
}

func (h *SystemInitBiz) AddAdminUser() error {
	var randomPass = utils.GenerateRandomString(12)

	password, err := bcrypt.GeneratePassword(randomPass)
	if err != nil {
		return err
	}

	_, err = h.ur.GetByID(h.appCtx.Ctx, 1)
	if err != nil && ent.IsNotFound(err) {
		_, err = h.ur.Create(h.appCtx.Ctx, &params.CreateUserParams{
			Username: "admin",
			Password: password,
			Role:     user.RoleAdmin,
		})
		if err != nil {
			return err
		}

		h.log.Infof("Admin user created, username: admin, password: %s", randomPass)
		return nil
	}

	if err != nil {
		return err
	}

	return nil
}

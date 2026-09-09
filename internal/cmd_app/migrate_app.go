package cmd_app

import (
	"context"
	"fmt"
	"net/url"

	gomigrate "github.com/golang-migrate/migrate/v4"
	"github.com/knadh/koanf/v2"
	"go.uber.org/zap"

	"github.com/Xwudao/neter-template/internal/domain/payloads"
	"github.com/Xwudao/neter-template/internal/system"

	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

type MigrateApp struct {
	ctx  context.Context
	conf *koanf.Koanf
	log  *zap.SugaredLogger
}

func NewMigrateApp(app *system.AppContext, conf *koanf.Koanf, log *zap.SugaredLogger) *MigrateApp {
	return &MigrateApp{ctx: app.Ctx, conf: conf, log: log}
}

// Run intentionally does not infer schema changes. With sqlc, migrations are
// authored and reviewed as SQL in db/migrations, then sqlc generate validates
// queries against that schema.
func (a *MigrateApp) Run(_ string) {
	a.log.Info("add a paired .up.sql/.down.sql migration under db/migrations; sqlc does not generate migrations")
}

func (a *MigrateApp) dsn() (string, error) {
	var db payloads.DBConfig
	if err := a.conf.Unmarshal("db", &db); err != nil {
		return "", err
	}
	return fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=disable",
		url.QueryEscape(db.Username), url.QueryEscape(db.Password), db.Host, db.Port, url.PathEscape(db.Database)), nil
}

func (a *MigrateApp) migrate() (*gomigrate.Migrate, error) {
	dsn, err := a.dsn()
	if err != nil {
		return nil, err
	}
	return gomigrate.New("file://db/migrations", dsn)
}

func (a *MigrateApp) Up(upAll bool) {
	m, err := a.migrate()
	if err != nil {
		a.log.Errorf("create migrator: %v", err)
		return
	}
	defer func() { _, _ = m.Close() }()
	if upAll {
		err = m.Up()
	} else {
		err = m.Steps(1)
	}
	if err != nil && err != gomigrate.ErrNoChange {
		a.log.Errorf("migration up: %v", err)
		return
	}
	a.log.Info("migration up success")
}

func (a *MigrateApp) Down(downAll bool) {
	m, err := a.migrate()
	if err != nil {
		a.log.Errorf("create migrator: %v", err)
		return
	}
	defer func() { _, _ = m.Close() }()
	if downAll {
		err = m.Down()
	} else {
		err = m.Steps(-1)
	}
	if err != nil && err != gomigrate.ErrNoChange {
		a.log.Errorf("migration down: %v", err)
		return
	}
	a.log.Info("migration down success")
}

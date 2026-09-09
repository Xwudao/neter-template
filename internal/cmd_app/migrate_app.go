package cmd_app

import (
	"context"
	"errors"
	"fmt"
	"net/url"
	"os"
	"strconv"

	gomigrate "github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/source/iofs"
	"github.com/knadh/koanf/v2"
	"go.uber.org/zap"

	"github.com/Xwudao/neter-template/db/migrations"
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

// Run is kept for `./app migrate` without a subcommand.
func (a *MigrateApp) Run(_ string) {
	a.log.Info("add a paired .up.sql/.down.sql migration under db/migrations; sqlc does not generate migrations")
}

// AutoMigrate applies every pending migration before the application serves
// traffic. golang-migrate takes a Postgres advisory lock, so concurrent
// replicas serialize: one migrates, the rest observe ErrNoChange.
func (a *MigrateApp) AutoMigrate() error {
	if !a.autoMigrateEnabled() {
		a.log.Info("auto migrate disabled; skipping database migrations")
		return nil
	}

	m, err := a.migrate()
	if err != nil {
		return fmt.Errorf("create migrator: %w", err)
	}
	defer func() { _, _ = m.Close() }()

	before, _, _ := m.Version()
	switch err := m.Up(); {
	case err == nil:
		after, _, _ := m.Version()
		a.log.Infof("database migrated from version %d to %d", before, after)
	case errors.Is(err, gomigrate.ErrNoChange):
		a.log.Infof("database schema is up to date (version %d)", before)
	default:
		return fmt.Errorf("apply migrations: %w", err)
	}
	return nil
}

// autoMigrateEnabled resolves NETER_AUTO_MIGRATE first, then db.autoMigrate.
func (a *MigrateApp) autoMigrateEnabled() bool {
	if v, ok := os.LookupEnv("NETER_AUTO_MIGRATE"); ok {
		if b, err := strconv.ParseBool(v); err == nil {
			return b
		}
	}

	var db payloads.DBConfig
	if err := a.conf.Unmarshal("db", &db); err != nil {
		return false
	}
	return db.AutoMigrate
}

// dsn resolves NETER_MIGRATE_DSN, then db.migrateDsn, then the db block.
func (a *MigrateApp) dsn() (string, error) {
	if dsn := os.Getenv("NETER_MIGRATE_DSN"); dsn != "" {
		return dsn, nil
	}

	var db payloads.DBConfig
	if err := a.conf.Unmarshal("db", &db); err != nil {
		return "", err
	}
	if db.MigrateDsn != "" {
		return db.MigrateDsn, nil
	}
	return fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=disable",
		url.QueryEscape(db.Username), url.QueryEscape(db.Password), db.Host, db.Port, url.PathEscape(db.Database)), nil
}

// migrate builds a migrator from the embedded migrations. Setting db.migratePath
// switches to a directory on disk, which is useful while developing.
func (a *MigrateApp) migrate() (*gomigrate.Migrate, error) {
	dsn, err := a.dsn()
	if err != nil {
		return nil, err
	}

	var db payloads.DBConfig
	if err := a.conf.Unmarshal("db", &db); err == nil && db.MigratePath != "" {
		return gomigrate.New("file://"+db.MigratePath, dsn)
	}

	src, err := iofs.New(migrations.FS, ".")
	if err != nil {
		return nil, err
	}
	return gomigrate.NewWithSourceInstance("iofs", src, dsn)
}

func (a *MigrateApp) Up(upAll bool) error {
	m, err := a.migrate()
	if err != nil {
		return fmt.Errorf("create migrator: %w", err)
	}
	defer func() { _, _ = m.Close() }()

	if upAll {
		err = m.Up()
	} else {
		err = m.Steps(1)
	}
	if err != nil && !errors.Is(err, gomigrate.ErrNoChange) {
		return fmt.Errorf("migration up: %w", err)
	}
	a.log.Info("migration up success")
	return nil
}

func (a *MigrateApp) Down(downAll bool) error {
	m, err := a.migrate()
	if err != nil {
		return fmt.Errorf("create migrator: %w", err)
	}
	defer func() { _, _ = m.Close() }()

	if downAll {
		err = m.Down()
	} else {
		err = m.Steps(-1)
	}
	if err != nil && !errors.Is(err, gomigrate.ErrNoChange) {
		return fmt.Errorf("migration down: %w", err)
	}
	a.log.Info("migration down success")
	return nil
}

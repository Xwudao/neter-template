package cmd_app

import (
	"context"
	"fmt"

	"ariga.io/atlas/sql/sqltool"
	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql/schema"
	"github.com/Xwudao/neter-template/internal/data/ent/migrate"
	"github.com/Xwudao/neter-template/internal/system"
	gomigrate "github.com/golang-migrate/migrate/v4"
	"github.com/knadh/koanf/v2"
	"go.uber.org/zap"

	_ "github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

type MigrateApp struct {
	ctx  context.Context
	conf *koanf.Koanf
	log  *zap.SugaredLogger
}

func NewMigrateApp(
	app *system.AppContext,
	conf *koanf.Koanf,
	log *zap.SugaredLogger,
) *MigrateApp {
	return &MigrateApp{
		ctx:  app.Ctx,
		conf: conf,
		log:  log,
	}
}

func (a *MigrateApp) Run(migrateName string) {
	var (
		log  = a.log
		conf = a.conf
	)

	dsn := conf.String("db.devDsn")

	dir, err := sqltool.NewGolangMigrateDir("migrations")
	if err != nil {
		log.Fatalf("failed creating atlas migration directory: %v", err)
	}
	// Write migration diff.
	opts := []schema.MigrateOption{
		schema.WithDir(dir),                         // provide migration directory
		schema.WithMigrationMode(schema.ModeReplay), // provide migration mode
		schema.WithDialect(dialect.MySQL),           // Ent dialect to use
		schema.WithFormatter(sqltool.GolangMigrateFormatter),
	}
	if migrateName == "" {
		err = migrate.Diff(a.ctx, dsn, opts...)
	} else {
		err = migrate.NamedDiff(a.ctx, dsn, migrateName, opts...)
	}
	if err != nil {
		log.Fatalf("failed generating migration file: %v", err)
	}
}

// migrate up

func (a *MigrateApp) Up(upAll bool) {
	var (
		conf = a.conf
		log  = a.log
	)
	dl := conf.String("db.dialect")
	host := conf.String("db.host")
	port := conf.String("db.port")
	username := conf.String("db.username")
	password := conf.String("db.password")
	database := conf.String("db.database")

	dsn := fmt.Sprintf("%s://%s:%s@tcp(%s:%s)/%s?multiStatements=true", dl, username, password, host, port, database)

	m, err := gomigrate.New(fmt.Sprintf("file://%s", "migrations"), dsn)
	if err != nil {
		log.Fatalf("failed creating migrate: %v", err)
	}

	if upAll {
		err = m.Up()
	} else {
		err = m.Steps(1)
	}
	if err != nil {
		log.Errorf("failed creating migrate: %v", err)
		return
	}

	log.Infof("migration up success")
}

// migrate down

func (a *MigrateApp) Down(downAll bool) {
	var (
		conf = a.conf
		log  = a.log
	)
	dl := conf.String("db.dialect")
	host := conf.String("db.host")
	port := conf.String("db.port")
	username := conf.String("db.username")
	password := conf.String("db.password")
	database := conf.String("db.database")

	dsn := fmt.Sprintf("%s://%s:%s@tcp(%s:%s)/%s?multiStatements=true", dl, username, password, host, port, database)

	m, err := gomigrate.New(fmt.Sprintf("file://%s", "migrations"), dsn)
	if err != nil {
		log.Fatalf("failed creating migrate: %v", err)
	}

	if downAll {
		err = m.Down()
	} else {
		err = m.Steps(-1)
	}
	if err != nil {
		log.Errorf("failed creating migrate: %v", err)
		return
	}

	log.Infof("migration down success")
}

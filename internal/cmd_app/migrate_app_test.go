package cmd_app

import (
	"testing"

	"github.com/knadh/koanf/v2"
	"go.uber.org/zap"

	"github.com/Xwudao/neter-template/internal/system"
)

func testMigrateApp(t *testing.T, values map[string]any) *MigrateApp {
	t.Helper()
	k := koanf.New(".")
	for key, val := range values {
		if err := k.Set(key, val); err != nil {
			t.Fatalf("set %s: %v", key, err)
		}
	}
	return NewMigrateApp(system.NewTestAppContext(), k, zap.NewNop().Sugar())
}

func TestAutoMigrateEnabled(t *testing.T) {
	app := testMigrateApp(t, map[string]any{"db.autoMigrate": true})

	t.Setenv("NETER_AUTO_MIGRATE", "false")
	if app.autoMigrateEnabled() {
		t.Fatal("NETER_AUTO_MIGRATE=false must override db.autoMigrate=true")
	}

	t.Setenv("NETER_AUTO_MIGRATE", "true")
	if !app.autoMigrateEnabled() {
		t.Fatal("NETER_AUTO_MIGRATE=true must enable auto migration")
	}

	// An empty env value is ignored, so the config file wins.
	t.Setenv("NETER_AUTO_MIGRATE", "")
	if !app.autoMigrateEnabled() {
		t.Fatal("empty NETER_AUTO_MIGRATE must fall back to db.autoMigrate")
	}

	off := testMigrateApp(t, map[string]any{"db.autoMigrate": false})
	if off.autoMigrateEnabled() {
		t.Fatal("db.autoMigrate=false must disable auto migration")
	}
}

func TestMigrateDSNPrecedence(t *testing.T) {
	app := testMigrateApp(t, map[string]any{
		"db.host":     "db.local",
		"db.port":     5432,
		"db.username": "app",
		"db.password": "p@ss",
		"db.database": "shop",
	})

	t.Setenv("NETER_MIGRATE_DSN", "")
	dsn, err := app.dsn()
	if err != nil {
		t.Fatalf("dsn: %v", err)
	}
	if dsn != "postgres://app:p%40ss@db.local:5432/shop?sslmode=disable" {
		t.Fatalf("unexpected dsn: %s", dsn)
	}

	t.Setenv("NETER_MIGRATE_DSN", "postgres://migrator:secret@db.local:5432/shop?sslmode=disable")
	dsn, err = app.dsn()
	if err != nil {
		t.Fatalf("dsn: %v", err)
	}
	if dsn != "postgres://migrator:secret@db.local:5432/shop?sslmode=disable" {
		t.Fatalf("NETER_MIGRATE_DSN must win, got: %s", dsn)
	}
}

func TestMigrateDsnFromConfig(t *testing.T) {
	app := testMigrateApp(t, map[string]any{
		"db.migrateDsn": "postgres://migrator@db.local/shop",
		"db.host":       "ignored",
	})
	t.Setenv("NETER_MIGRATE_DSN", "")

	dsn, err := app.dsn()
	if err != nil {
		t.Fatalf("dsn: %v", err)
	}
	if dsn != "postgres://migrator@db.local/shop" {
		t.Fatalf("db.migrateDsn must be used, got: %s", dsn)
	}
}

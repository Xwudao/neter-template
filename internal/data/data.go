package data

import (
	"context"
	"fmt"
	"net/url"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/knadh/koanf/v2"
	"go.uber.org/zap"

	"github.com/Xwudao/neter-template/internal/data/sqlc"
	"github.com/Xwudao/neter-template/internal/domain/payloads"
)

// Data owns the PostgreSQL connection pool and sqlc query set. Repositories
// depend on this small facade rather than opening database connections directly.
type Data struct {
	Pool    *pgxpool.Pool
	Queries *sqlc.Queries
}

func NewData(dbConf *payloads.DBConfig, conf *koanf.Koanf, log *zap.SugaredLogger) (*Data, func(), error) {
	dsn := fmt.Sprintf(
		"postgres://%s:%s@%s:%d/%s?sslmode=disable",
		url.QueryEscape(dbConf.Username),
		url.QueryEscape(dbConf.Password),
		dbConf.Host,
		dbConf.Port,
		url.PathEscape(dbConf.Database),
	)

	poolConfig, err := pgxpool.ParseConfig(dsn)
	if err != nil {
		return nil, nil, err
	}
	// SQL tracing is restricted to debug mode to avoid noisy production logs.
	if conf.String("app.mode") == "debug" {
		poolConfig.ConnConfig.Tracer = sqlTracer{log: log.Named("sql")}
	}

	pool, err := pgxpool.NewWithConfig(context.Background(), poolConfig)
	if err != nil {
		return nil, nil, err
	}
	if err := pool.Ping(context.Background()); err != nil {
		pool.Close()
		return nil, nil, err
	}

	data := &Data{Pool: pool, Queries: sqlc.New(pool)}
	return data, pool.Close, nil
}

package data

import (
	"context"
	"fmt"
	"net/url"

	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/Xwudao/neter-template/internal/data/sqlc"
	"github.com/Xwudao/neter-template/internal/domain/payloads"
)

// Data owns the PostgreSQL connection pool and sqlc query set. Repositories
// depend on this small facade rather than opening database connections directly.
type Data struct {
	Pool    *pgxpool.Pool
	Queries *sqlc.Queries
}

func NewData(dbConf *payloads.DBConfig) (*Data, func(), error) {
	dsn := fmt.Sprintf(
		"postgres://%s:%s@%s:%d/%s?sslmode=disable",
		url.QueryEscape(dbConf.Username),
		url.QueryEscape(dbConf.Password),
		dbConf.Host,
		dbConf.Port,
		url.PathEscape(dbConf.Database),
	)

	pool, err := pgxpool.New(context.Background(), dsn)
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

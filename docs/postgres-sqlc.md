# PostgreSQL + pgx + sqlc

This project uses SQL as the source of truth:

- `db/migrations/`: versioned PostgreSQL schema changes. Each change needs a reviewed `up` and `down` file. They are embedded into the binary (`db/migrations/embed.go`), so a deployed build migrates itself.
- `db/query/`: named SQL queries. These are the only files edited for ordinary reads/writes.
- `sqlc.yaml`: sqlc generation contract (PostgreSQL + pgx/v5, generated result pointers, JSON tags).
- `db/migrations/000001_init.up.sql`: `users.role` is a real PostgreSQL enum (`CREATE TYPE user_role AS ENUM ('user','admin')`), so sqlc generates the closed Go type `sqlc.UserRole` with `UserRoleUser`/`UserRoleAdmin`; pgx scans it directly.
- `internal/data/sqlc/`: committed generated code; do not edit it.
- `internal/data/`: pgxpool lifecycle, repository implementations, and transactions. `Data` is injected by Wire and its cleanup closes the pool.
- `internal/biz/`: repository interfaces and application behavior. It currently uses sqlc result structs just as the original template used Ent structs, avoiding an unnecessarily broad API refactor during this experiment.

## Workflow

1. Create a paired migration: `nr migrate new add_orders` (writes `db/migrations/000NNN_add_orders.{up,down}.sql`).
2. Apply it to the configured database: `nr migrate up` (or `go run ./cmd/app migrate up --all`).
3. Add or update named statements in `db/query/*.sql`.
4. Run `make sqlc`, then commit generated `internal/data/sqlc/` files with the SQL changes.
5. Run `make test`.

### `nr migrate` commands

| Command | Purpose |
| --- | --- |
| `nr migrate new <name>` | Create the next `000NNN_<name>.up.sql` / `.down.sql` pair. |
| `nr migrate up [--steps N]` | Apply all pending migrations (or at most N). |
| `nr migrate down [--all] [--steps N]` | Roll back one migration, N migrations, or all of them. |
| `nr migrate status` | Show the current version and dirty flag. |
| `nr migrate force <version>` | Recover a dirty database by pinning the schema version. |

The DSN is resolved from `--dsn`, then `NETER_DSN` / `DATABASE_URL`, then the `db` block in `config.yml`. Pass `--path` to use a migration directory other than `db/migrations`.

`migrate` no longer generates a schema diff: that was Ent/Atlas behavior. SQL migrations are authored explicitly so production DDL is reviewable. `UpdateOrder` demonstrates the repository boundary for a pgx transaction (`Queries.WithTx`).

## Auto-migrate on startup

The migrations are embedded in the binary and applied before the app serves traffic (see `MainApp.Run` → `MigrateApp.AutoMigrate`). Enable it in `config.yml`:

```yaml
db:
  autoMigrate: true
```

or override at deploy time with `NETER_AUTO_MIGRATE=true|false`. When disabled, run migrations as a separate step: `./app migrate up --all` (or `nr migrate up`, or a Kubernetes `Job`).

Behaviour and safeguards:

- **Concurrency safe.** golang-migrate takes a Postgres advisory lock, so several replicas starting at once serialize: one migrates, the rest see `ErrNoChange`.
- **Fail fast.** A failed migration makes `Run` return an error and the process exits non-zero, so an orchestrator keeps the instance out of rotation. A dirty schema keeps failing until you run `nr migrate force <version>` after fixing the SQL.
- **Dedicated DDL user (optional).** Set `db.migrateDsn` or `NETER_MIGRATE_DSN` to run migrations with a different role than the application.
- **Backward-compatible DDL.** During a rolling deploy old and new code run side by side; use expand/contract (add nullable column → backfill → enforce) instead of destructive changes.
- **Local files (optional).** `db.migratePath` reads migrations from disk instead of the embedded copy while developing.

The DSN for migrations resolves as `NETER_MIGRATE_DSN` → `db.migrateDsn` → the `db` block.

## Generating a resource

On a new sqlc project, `nr gen biz -n order --with-crud --model Order` scaffolds the biz interface and repository. Fill in the matching `db/query/orders.sql` queries (`ListOrders`, `GetOrder`, `CreateOrder`, `UpdateOrder`, `DeleteOrder`), run `make sqlc`, then replace the generated `errOrderNotImplemented` bodies with `data.Queries` calls. Legacy Ent projects keep using `--ent-name`.

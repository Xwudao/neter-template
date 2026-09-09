package biz

// nr gen biz 可选参数：
//   --no-repo        不生成 data Repository
//   --with-crud      在 Biz/Repository 中生成基础 CRUD 方法
//   --model NAME     sqlc 模型名，例如 Order（新项目使用，配合 --with-crud）
//   --plural NAME    sqlc 列表方法复数，默认 <Model>s，例如 Orders
//   --ent-name NAME  旧 Ent 项目的实体名（legacy 项目使用）
//   --with-params    生成 internal/domain/params 请求参数文件
//   --with-contracts 生成不依赖 HTTP 的 Command/Query 业务合约
//   --with-iface     生成 Biz 接口，并在 mocks/mock_gen.go 中增加 mockgen 指令
//
// 本项目使用 PostgreSQL + pgx + sqlc（SQL 为唯一事实来源）。
// --with-crud 会按 sqlc 约定生成接口与仓储骨架：先在 db/query/<table>.sql
// 补齐 List<Plural>/Get<Model>/Create<Model>/Update<Model>/Delete<Model>，
// 运行 make sqlc，再用 data.Queries 调用替换 err<Model>NotImplemented 占位。
//go:generate nr gen biz -n system_init --no-repo
//go:generate nr gen biz -n site_help --no-repo
//go:generate nr gen biz -n html_help --no-repo
//go:generate nr gen biz -n seo_biz --no-repo

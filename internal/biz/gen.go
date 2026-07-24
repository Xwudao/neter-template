package biz

// nr gen -t biz 可选参数：
//   --no-repo        不生成 data Repository
//   --with-crud      在 Biz/Repository 中生成基础 CRUD 方法（需同时指定 --ent-name）
//   --ent-name NAME  CRUD 对应的 Ent 实体名，例如 User
//   --with-params    生成 internal/domain/params 请求参数文件
//   --with-contracts 生成不依赖 HTTP 的 Command/Query 业务合约
//   --with-iface     生成 Biz 接口，并在 mocks/mock_gen.go 中增加 mockgen 指令
//go:generate nr gen -t biz -n user --with-crud --ent-name User
//go:generate nr gen -t biz -n system_init --no-repo
//go:generate nr gen -t biz -n site_help --no-repo
//go:generate nr gen -t biz -n html_help --no-repo
//go:generate nr gen -t biz -n seo_biz --no-repo
//go:generate nr gen -t biz -n site_config --with-crud --ent-name SiteConfig
//go:generate nr gen -t biz -n data_list --with-crud --ent-name DataList

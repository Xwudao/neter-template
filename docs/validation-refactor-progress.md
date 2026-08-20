# 代码式参数校验重构进度（validation-refactor-progress）

> 将 Go + Gin 项目中基于 `binding:"..."` tag 的 go-playground/validator 校验，迁移为项目内 `internal/validate` 包提供的代码式、无 tag、无反射的 `Validate() error`，并移除对 validator 的直接依赖。

## 一、范围

- 迁移对象：`internal/domain/params` 下全部 DTO（5 个文件，26 个 `binding` tag）。
- 新增：`internal/validate` 包（规则 + 表驱动测试）。
- 接入：`internal/core/rtn.go` 的 `bindAndRespond` 绑定后处理，插入 `Validate()` 步骤。
- 清理：删除全部 `binding` tag、`GetMessages` 旧映射、`validator/v10` 直接依赖。
- 不动：`internal/data/ent/**`（生成代码）、Ent schema 持久化校验、Biz/Domain 层外部状态规则。

### Tag 分布（迁移前盘点）

| 文件 | tag 数 | 类型 |
| --- | --- | --- |
| `cmn_params.go` | 3 | `required`（int64、[]int64、[]int） |
| `data_list_params.go` | 10 | `required` ×8、`min=1` ×1、`min=1,max=100` ×1 |
| `user_params.go` | 4 | `required`（string ×4） |
| `upload_params.go` | 3 | `required`（string ×2、`*multipart.FileHeader` ×1） |
| `site_config_params.go` | 6 | `required`（string ×6） |

> 实际用到的 tag 只有 `required` / `min` / `max`，未出现 `omitempty`、`oneof`、`email`、`ip`、`datetime` 等；`validate` 包仍按提示词最小规则集补齐了 `Optional`/`When`/`OneOf`/`Email`/`MinLen`/`MaxLen`/`MinItems`/`MaxItems`/`NotZero` 等通用规则，供后续扩展。

## 二、目标架构

### 2.1 `internal/validate` 包

- 纯函数规则 `type Rule[T any] func(T) error`：不接收 context、不访问 DB/仓储/外部服务。
- 无反射、无 tag、无全局注册器。
- `Validate(fields ...fieldResult) error` 聚合每个字段的第一条失败规则（按字段声明顺序，消息以 `; ` 连接）。
- `Field(name, value, rules...)` 对单字段按声明顺序短路。
- `Errors` 实现 `error` 接口，`First()`/`FirstError()` 供 HTTP 入口只暴露第一条消息。
- 规则默认错误文案统一为 `参数错误`（与旧兜底 `ErrParams` 一致）；DTO 层通过 `Message(msg, rule)` 替换为自定义文案。

### 2.2 框架接入（`internal/core/rtn.go`）

绑定后处理顺序固定为：

```
绑定成功 → Validate()（失败则返回第一条消息） → Optimize()（默认值/钳制/offset） → handler
```

- 新增接口 `RequestValidator { Validate() error }`，`bindAndRespond` 在 `optimizeRequest` 之前调用 `validateRequest`。
- 绑定解析失败（JSON 语法/类型错误、空 body、strconv 错误）仍走 `defaultBindErrorMapper` 安全映射，统一返回 `参数错误`，不泄漏原始解析错误。
- 本项目无直接 `ShouldBindQuery`/`ShouldBindJSON` 的旁路绑定端点（全部经 `core.JSONE/RequestE/NoInputE`），无需额外显式接入点。

## 三、语义决策清单

1. **消息保留**：原自定义消息映射中存在的消息全部原样保留（含历史怪文案，见第 2 条）。没有自定义消息的规则（仅 `UploadToS3Params`，死代码）使用兜底 `参数错误`。
2. **保留怪文案**：`UpdateSiteConfigParams` 的 `Name.required` 旧文案为 `"ID不能为空"`（疑似复制粘贴历史遗留）。按"消息必须原样保留"决策，迁移后保持 `Name` 字段失败返回 `ID不能为空`，未顺手"修复"。
3. **空值语义**（逐字段核对）：
   - `ListDataByKindParams.Page/Size` 旧 tag 为 `min=1`/`min=1,max=100`，**无 `omitempty`**，因此空值（0）必须失败 → 直接用 `Min(1)`/`Max(100)`，不加 `When` 跳过。已用 DTO 测试断言 `Page=0,Size=0 → "Page最小值为1; Size最小值为1"`。
   - 全部 `required` 字段（string 用 `Required()` 不 trim、int64/指针用 `NotZero`、切片用 `MinItems(1)`）在空值时失败，与原行为一致。
4. **死映射删除**（不补校验、不改行为）：
   - `DeleteIDParams` 映射 key 为 `"Name.required"`（字段实为 `ID`），属失效 key。因 `ID` 的 `required` tag 真实存在、消息 `ID必填` 语义明确，按"消息保留"决策用于 `ID`（修复失效 key，而非删除校验）。
   - `ListDataByKindParams` 的 `"Kind.required"` 映射：`Kind` 字段**没有** binding tag（可选），该映射是死代码，已删除，**未**给 `Kind` 增加原本不存在的必填校验。
5. **聚合与首条消息**：`Validate()` 返回聚合错误（DTO 测试断言完整文本）；HTTP 入口 `FirstError` 只取第一条消息，维持既有 `msg` 行为。
6. **`Optimize` 不迁入校验**：默认值（`ItemOrder`=1）、offset 计算（`ListDataByKindParams`）保留在 DTO 原 `Optimize()`；保持"校验失败 → 不进 Optimize"的先后关系（路由级测试覆盖）。
7. **min/max 的 rune 语义**：本项目 `min/max` 均用于纯数值字段（int），不涉及 go-playground 字符串按字节计 vs validate 按 rune 计的差异；`MinLen/MaxLen` 按 rune（`utf8.RuneCountInString`）实现并有专门测试，供后续字符串长度规则使用。
8. **未迁移的外部状态规则**：无。本项目 tag 校验全部是字段自身合法性；数据库存在性/权限等校验原本就在 Biz 层，未改动。

## 四、测试结果

新增测试：

| 位置 | 内容 |
| --- | --- |
| `internal/validate/validate_test.go` | 12 组表驱动测试：Required、NotZero、Min/Max、MinLen/MaxLen（rune）、Email、OneOf、MinItems/MaxItems、Optional、When、Message、Field 短路、Validate 聚合、FirstError |
| `internal/domain/params/params_test.go` | 19 个 DTO 用例：聚合错误文本与原自定义消息逐字一致；首条消息；`ItemOrderParams` 校验/Optimize 解耦；`UploadToS3Params` 文件指针 nil 失败；分页 `Optimize` offset 语义 |
| `internal/core/rtn_test.go` | 路由级测试：绑定成功 → Validate 先于 Optimize；缺参返回原包装结构 + 中文首条消息；JSON 语法错误安全映射 `参数错误` |

执行结果（沙箱缓存 `/tmp/neter-gocache`、`/tmp/neter-gopath`）：

```
go test ./internal/validate/ ./internal/core/ ./internal/domain/params/ ./internal/routes/valid/  → 全绿
go vet  上述 4 个变更包                                                            → 通过
go build 上述 4 个变更包                                                            → 通过
git diff --check                                                                 → 通过
```

## 五、既有阻断（与本重构无关，未顺手修改）

全量 `go test ./...` / `go vet ./...` 存在以下**迁移前就存在**的失败（已用 `git stash` 基线对比确认前后一致）：

1. `assets/holder.go` 的 `//go:embed all:dist` 因 `assets/dist` 为空目录报 `pattern all:dist: no matching files found`，导致 `assets`、`cmd/app`、`internal/cmd`、`internal/routes` 无法编译。**前置条件：先构建前端产出 `dist`（`pnpm build` 后复制到 `assets/dist`）。**
2. `internal/routes/v1/site_config_routes.go:50,58` 调用了 `core.NoInput` 却传入返回 `error` 的函数（应为 `core.NoInputE`），类型不匹配编译失败。
3. `internal/libx/s3_client_test.go` 因测试环境未配置静态凭据失败（`EmptyStaticCreds`）。

## 六、完成标准核对

| 标准 | 状态 |
| --- | --- |
| `rg 'binding:"'`（排除 `internal/data/ent/**`）无结果 | ✅ |
| `rg 'GetMessages\|ValidatorMessages' internal/` 无结果 | ✅ |
| `go.mod` 中 `validator/v10` 为 `// indirect`（gin 传递依赖） | ✅ |
| `defaultBindErrorMapper`/`GetErrorMsg` 不再引用 validator 类型；解析失败仍安全映射中文消息 | ✅ |
| 变更包 `go test`/`go vet`/`go build` 通过；`git diff --check` 通过 | ✅ |
| 本进度文档已更新 | ✅ |

## 七、后续事项

- 若需解除 `internal/routes` 等包的编译阻断，按"既有阻断"第 1、2 条处理（前端构建 + `NoInputE` 修正），属独立任务。
- 新 DTO 校验规则直接写入对应 `Validate()` 方法；需要新规则时在 `internal/validate` 添加纯函数 + 表驱动测试。

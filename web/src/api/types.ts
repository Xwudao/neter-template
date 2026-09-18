/**
 * Hand-written API contracts for the Go backend.
 *
 * These definitions are maintained manually; there is no generator keeping
 * them in sync. Update them alongside the matching Go sources:
 *
 * - routes: `internal/routes/v1/*_routes.go`
 * - DTOs:   `internal/domain/params/*.go`
 * - models: `internal/data/sqlc/models.go`
 * - envelope: `internal/core/rtn.go` (`WrappedResp`)
 */

/** Unified response envelope returned by every handler. */
export interface ApiResponse<T> {
  code: number
  msg: string
  data: T
}

/** `sqlc.UserRole`. */
export type UserRole = 'user' | 'admin'

/** `sqlc.User`. */
export interface User {
  id: number
  username: string
  password: string
  role: UserRole
}

/** `sqlc.DataList`. */
export interface DataList {
  id: number
  label: string
  kind: string
  key: string
  value: string
  item_order: number
  /** RFC 3339 timestamp serialized from Go `time.Time`. */
  create_time: string
  update_time: string
}

/** `sqlc.SiteConfig`. */
export interface SiteConfig {
  id: number
  name: string
  config: string
  create_time: string
  update_time: string
}

/** Paginated response of `GET /admin/v1/data_list/list`. */
export interface DataListPage {
  list: DataList[]
  total: number
}

/** Response of `GET /v1/site_config/all` and `GET /admin/v1/site_config/all`. */
export type SiteConfigMap = Record<string, string>

/** Body of `POST /v1/user/login` (`params.UserLoginParams`). */
export interface UserLoginParams {
  username: string
  password: string
}

/** Data of `POST /v1/user/login` (`v1.UserLoginResponse`). */
export interface UserLoginResult {
  user: User | null
  token: string
}

/** Query of `GET /admin/v1/data_list/list` (`params.ListDataByKindParams`). */
export interface ListDataByKindParams {
  kind: string
  page: number
  size: number
}

/** Query of `GET /admin/v1/data_list/sort_data` (`params.GetDataListSortDataParams`). */
export interface GetDataListSortDataParams {
  kind: string
}

/** Body of `POST /admin/v1/data_list/create` (`params.CreateDataListParams`). */
export interface CreateDataListParams {
  label: string
  key: string
  kind: string
  value: string
  item_order: number
}

/** Body of `POST /admin/v1/data_list/update` (`params.UpdateDataListParams`). */
export interface UpdateDataListParams {
  id: number
  key: string
  value: string
  item_order?: number | null
}

/** Body of `POST /admin/v1/data_list/update_order` (`params.ItemOrderParams`). */
export interface ItemOrderParams {
  ids: number[]
  orders: number[]
}

/** Body of `POST /admin/v1/data_list/delete` (`params.DeleteIDParams`). */
export interface DeleteIDParams {
  id: number
}

/** Body of `POST /admin/v1/site_config/update` (`params.UpdateSiteConfigParams`). */
export interface UpdateSiteConfigParams {
  name: string
  config: string
}

/** Body of `POST /admin/v1/site_config/write_file` (`params.WriteFileParams`). */
export interface WriteFileParams {
  filename: string
  data: string
}

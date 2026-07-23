package core

// Page is the transport-neutral result of a paginated business query.
//
// Biz and repository code can return it directly; route handlers may pass it
// to NoInput/JSON without first rebuilding a gin.H value.
type Page[T any] struct {
	List  []T   `json:"list"`
	Total int64 `json:"total"`
}

func NewPage[T any](list []T, total int64) Page[T] {
	return Page[T]{List: list, Total: total}
}

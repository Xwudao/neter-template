package valid

import (
	"errors"
)

var (
	ErrParams = errors.New("参数错误")
)

// GetErrorMsg 获取错误信息（绑定解析失败的安全映射）。
// 字段校验失败由 DTO 的 Validate() 在绑定成功后处理，不再经过本函数；
// 这里只保留对解析错误（JSON 语法/类型错误、空 body、strconv 错误）的统一兜底，
// 不把原始解析错误暴露给客户端。
func GetErrorMsg(request any, err error) error {
	return ErrParams
}

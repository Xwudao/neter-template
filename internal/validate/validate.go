// Package validate 提供代码式、无 tag、无反射的请求参数校验规则。
//
// 设计约束：
//   - 规则是纯函数（Rule[T]），不接收 context，不访问数据库/仓储/外部服务；
//   - 不引入 reflect、DSL 或全局注册器；
//   - Validate 聚合每个字段的第一条失败规则，HTTP 入口通过 FirstError 只暴露第一条消息。
package validate

import (
	"cmp"
	"errors"
	"net/mail"
	"slices"
	"strings"
	"unicode/utf8"
)

// Rule 单字段校验规则：value 不合法时返回非 nil error。
type Rule[T any] func(T) error

// FieldError 单个字段的校验失败结果。
type FieldError struct {
	Field   string
	Message string
}

// Errors 聚合多个字段错误，实现 error 接口。
type Errors []FieldError

// Error 返回聚合错误文本（所有字段消息以 "; " 连接）。
func (e Errors) Error() string {
	if len(e) == 0 {
		return ""
	}
	msgs := make([]string, 0, len(e))
	for _, fe := range e {
		msgs = append(msgs, fe.Message)
	}
	return strings.Join(msgs, "; ")
}

// First 返回第一条错误消息；空 Errors 返回 nil。
// 供通用绑定入口只向响应暴露第一条消息，维持既有 msg 行为。
func (e Errors) First() error {
	if len(e) == 0 {
		return nil
	}
	return errors.New(e[0].Message)
}

// FirstError 返回聚合错误中的第一条消息；非聚合错误原样返回。
func FirstError(err error) error {
	var es Errors
	if errors.As(err, &es) && len(es) > 0 {
		return es.First()
	}
	return err
}

// fieldResult Field 的单字段结果，由 Validate 聚合。
type fieldResult struct {
	name string
	err  error
}

// Validate 聚合每个字段的第一条失败规则；全部通过时返回 nil。
func Validate(fields ...fieldResult) error {
	var es Errors
	for _, f := range fields {
		if f.err == nil {
			continue
		}
		es = append(es, FieldError{Field: f.name, Message: f.err.Error()})
	}
	if len(es) == 0 {
		return nil
	}
	return es
}

// Field 按声明顺序对单字段短路：返回第一条失败规则的结果。
func Field[T any](name string, value T, rules ...Rule[T]) fieldResult {
	for _, rule := range rules {
		if err := rule(value); err != nil {
			return fieldResult{name: name, err: err}
		}
	}
	return fieldResult{name: name}
}

// Message 保留规则逻辑、替换错误文案。
func Message[T any](msg string, rule Rule[T]) Rule[T] {
	return func(v T) error {
		if err := rule(v); err != nil {
			return errors.New(msg)
		}
		return nil
	}
}

// Required 拒绝空字符串；不 trim，保持 Gin `required` 的零值语义。
func Required() Rule[string] {
	return func(v string) error {
		if v == "" {
			return errors.New("参数错误")
		}
		return nil
	}
}

// NotZero 拒绝零值；需显式类型参数，如 NotZero[int64]()。
func NotZero[T comparable]() Rule[T] {
	var zero T
	return func(v T) error {
		if v == zero {
			return errors.New("参数错误")
		}
		return nil
	}
}

// Min 数值下界（含边界）。
func Min[T cmp.Ordered](bound T) Rule[T] {
	return func(v T) error {
		if v < bound {
			return errors.New("参数错误")
		}
		return nil
	}
}

// Max 数值上界（含边界）。
func Max[T cmp.Ordered](bound T) Rule[T] {
	return func(v T) error {
		if v > bound {
			return errors.New("参数错误")
		}
		return nil
	}
}

// MinLen 字符串最小长度（按 rune 数，非字节数）。
func MinLen(n int) Rule[string] {
	return func(v string) error {
		if utf8.RuneCountInString(v) < n {
			return errors.New("参数错误")
		}
		return nil
	}
}

// MaxLen 字符串最大长度（按 rune 数，非字节数）。
func MaxLen(n int) Rule[string] {
	return func(v string) error {
		if utf8.RuneCountInString(v) > n {
			return errors.New("参数错误")
		}
		return nil
	}
}

// Email RFC 5322 单地址校验。
func Email() Rule[string] {
	return func(v string) error {
		if _, err := mail.ParseAddress(v); err != nil {
			return errors.New("参数错误")
		}
		return nil
	}
}

// OneOf 枚举校验；零值不在集合中即失败（配合 When/Optional 模拟 omitempty）。
func OneOf[T comparable](options ...T) Rule[T] {
	return func(v T) error {
		if slices.Contains(options, v) {
			return nil
		}
		return errors.New("参数错误")
	}
}

// MinItems 切片元素数下界。
func MinItems[T any](n int) Rule[[]T] {
	return func(v []T) error {
		if len(v) < n {
			return errors.New("参数错误")
		}
		return nil
	}
}

// MaxItems 切片元素数上界。
func MaxItems[T any](n int) Rule[[]T] {
	return func(v []T) error {
		if len(v) > n {
			return errors.New("参数错误")
		}
		return nil
	}
}

// Optional nil 指针跳过，显式提供的值仍校验（等价 omitempty 的指针字段）。
func Optional[T any](rule Rule[T]) Rule[*T] {
	return func(v *T) error {
		if v == nil {
			return nil
		}
		return rule(*v)
	}
}

// When 条件为真才校验（等价 omitempty 的非指针字段）。
func When[T any](cond bool, rule Rule[T]) Rule[T] {
	return func(v T) error {
		if !cond {
			return nil
		}
		return rule(v)
	}
}

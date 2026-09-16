package params

import (
	"errors"
	"mime/multipart"
	"testing"

	"github.com/Xwudao/go-validate"
)

// TestDTOValidateMessages 断言 Validate() 的字段名和消息按声明顺序聚合。
func TestDTOValidateMessages(t *testing.T) {
	tests := []struct {
		name string
		err  error
		want string // 为空表示期望 nil
	}{
		{
			name: "CreateDataListParams all empty",
			err:  (&CreateDataListParams{}).Validate(),
			want: "label: 标签必填; key: Key必填; kind: 分类必填; value: 内容必填",
		},
		{
			name: "CreateDataListParams valid",
			err:  (&CreateDataListParams{Label: "l", Key: "k", Kind: "x", Value: "v"}).Validate(),
			want: "",
		},
		{
			name: "UpdateDataListParams all empty",
			err:  (&UpdateDataListParams{}).Validate(),
			want: "id: ID必填; key: Key必填; value: 内容必填",
		},
		{
			name: "ListDataByKindParams page and size zero",
			err:  (&ListDataByKindParams{}).Validate(),
			want: "page: Page最小值为1; size: Size最小值为1",
		},
		{
			name: "ListDataByKindParams size too large",
			err:  (&ListDataByKindParams{Page: 1, Size: 101}).Validate(),
			want: "size: Size最大值为100",
		},
		{
			name: "ListDataByKindParams valid",
			err:  (&ListDataByKindParams{Kind: "friend_link", Page: 2, Size: 10}).Validate(),
			want: "",
		},
		{
			name: "GetDataListSortDataParams kind empty",
			err:  (&GetDataListSortDataParams{}).Validate(),
			want: "kind: Kind必填",
		},
		{
			name: "DeleteIDParams id zero",
			err:  (&DeleteIDParams{}).Validate(),
			want: "id: ID必填",
		},
		{
			name: "DeleteIDParams id valid",
			err:  (&DeleteIDParams{ID: 5}).Validate(),
			want: "",
		},
		{
			name: "ItemOrderParams ids empty",
			err:  (&ItemOrderParams{}).Validate(),
			want: "ids: ID必填; orders: 排序必填",
		},
		{
			name: "ItemOrderParams valid",
			err:  (&ItemOrderParams{IDs: []int64{1}, Orders: []int{0}}).Validate(),
			want: "",
		},
		{
			name: "CreateUserParams all empty",
			err:  (&CreateUserParams{}).Validate(),
			want: "username: 用户名不能为空; password: 密码不能为空",
		},
		{
			name: "UserLoginParams all empty",
			err:  (&UserLoginParams{}).Validate(),
			want: "username: 用户名不能为空; password: 密码不能为空",
		},
		{
			name: "CreateSiteConfigParams all empty",
			err:  (&CreateSiteConfigParams{}).Validate(),
			want: "name: 名称不能为空; config: 配置不能为空",
		},
		{
			// 旧自定义消息映射中 Name.required 的文案为 "ID不能为空"（历史文案），按原样保留。
			name: "UpdateSiteConfigParams all empty",
			err:  (&UpdateSiteConfigParams{}).Validate(),
			want: "name: ID不能为空; config: 配置不能为空",
		},
		{
			name: "WriteFileParams all empty",
			err:  (&WriteFileParams{}).Validate(),
			want: "filename: 文件名不能为空; data: 数据不能为空",
		},
		{
			name: "UploadToS3Params all empty (default messages)",
			err:  (&UploadToS3Params{}).Validate(),
			want: "prefix: 不能为空; object: 不能为空; file: 不能为空",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.want == "" {
				if tt.err != nil {
					t.Fatalf("expected nil error, got %v", tt.err)
				}
				return
			}
			if tt.err == nil {
				t.Fatal("expected error, got nil")
			}
			got := tt.err.Error()
			if got != tt.want {
				t.Fatalf("Validate() = %q, want %q", got, tt.want)
			}
		})
	}
}

// TestValidateFirstMessage 断言 HTTP 入口可从聚合错误中获取首条消息。
func TestValidateFirstMessage(t *testing.T) {
	err := (&CreateDataListParams{}).Validate()
	if err == nil {
		t.Fatal("expected error")
	}
	var validationErrors validate.Errors
	if !errors.As(err, &validationErrors) {
		t.Fatalf("expected validate.Errors, got %T", err)
	}
	if first := validationErrors.First(); first == nil || first.Error() != "标签必填" {
		t.Fatalf("first message = %v, want 标签必填", first)
	}
}

// TestItemOrderParamsOptimizeNotAffectedByValidate 断言：校验与 Optimize 解耦，
// 合法的 IDs/Orders 通过校验后仍可进入 Optimize（长度不一致由 Optimize 拦截）。
func TestItemOrderParamsOptimize(t *testing.T) {
	p := &ItemOrderParams{IDs: []int64{1, 2}, Orders: []int{0}}
	if err := p.Validate(); err != nil {
		t.Fatalf("Validate should pass: %v", err)
	}
	if err := p.Optimize(); err == nil {
		t.Fatal("Optimize should reject mismatched lengths")
	}
}

// TestUploadToS3ParamsFileNil 断言文件指针为零值时校验失败。
func TestUploadToS3ParamsFileNil(t *testing.T) {
	p := &UploadToS3Params{Prefix: "p", Object: "o"}
	if err := p.Validate(); err == nil {
		t.Fatal("nil file should fail validation")
	}
	p.File = &multipart.FileHeader{}
	if err := p.Validate(); err != nil {
		t.Fatalf("non-nil file should pass: %v", err)
	}
}

// TestListDataByKindParamsOffsetAfterOptimize 断言分页 DTO 的 Optimize 默认值语义不变。
func TestListDataByKindParamsOffsetAfterOptimize(t *testing.T) {
	p := &ListDataByKindParams{Kind: "nav", Page: 2, Size: 10}
	if err := p.Validate(); err != nil {
		t.Fatalf("Validate should pass: %v", err)
	}
	p.Optimize()
	if p.Offset != 10 {
		t.Fatalf("Offset = %d, want 10", p.Offset)
	}
}

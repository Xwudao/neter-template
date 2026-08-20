package params

import (
	"mime/multipart"
	"strings"
	"testing"
)

// TestDTOValidateMessages 断言 Validate() 聚合错误文本与旧自定义消息映射保持一致。
// 消息按字段声明顺序聚合（每条用 "; " 分隔），与旧 validator 字段序遍历一致。
func TestDTOValidateMessages(t *testing.T) {
	tests := []struct {
		name string
		err  error
		want string // 为空表示期望 nil
	}{
		{
			name: "CreateDataListParams all empty",
			err:  (&CreateDataListParams{}).Validate(),
			want: "标签必填; Key必填; 分类必填; 内容必填",
		},
		{
			name: "CreateDataListParams valid",
			err:  (&CreateDataListParams{Label: "l", Key: "k", Kind: "x", Value: "v"}).Validate(),
			want: "",
		},
		{
			name: "UpdateDataListParams all empty",
			err:  (&UpdateDataListParams{}).Validate(),
			want: "ID必填; Key必填; 内容必填",
		},
		{
			name: "ListDataByKindParams page and size zero",
			err:  (&ListDataByKindParams{}).Validate(),
			want: "Page最小值为1; Size最小值为1",
		},
		{
			name: "ListDataByKindParams size too large",
			err:  (&ListDataByKindParams{Page: 1, Size: 101}).Validate(),
			want: "Size最大值为100",
		},
		{
			name: "ListDataByKindParams valid",
			err:  (&ListDataByKindParams{Kind: "friend_link", Page: 2, Size: 10}).Validate(),
			want: "",
		},
		{
			name: "GetDataListSortDataParams kind empty",
			err:  (&GetDataListSortDataParams{}).Validate(),
			want: "Kind必填",
		},
		{
			name: "DeleteIDParams id zero",
			err:  (&DeleteIDParams{}).Validate(),
			want: "ID必填",
		},
		{
			name: "DeleteIDParams id valid",
			err:  (&DeleteIDParams{ID: 5}).Validate(),
			want: "",
		},
		{
			name: "ItemOrderParams ids empty",
			err:  (&ItemOrderParams{}).Validate(),
			want: "ID必填; 排序必填",
		},
		{
			name: "ItemOrderParams valid",
			err:  (&ItemOrderParams{IDs: []int64{1}, Orders: []int{0}}).Validate(),
			want: "",
		},
		{
			name: "CreateUserParams all empty",
			err:  (&CreateUserParams{}).Validate(),
			want: "用户名不能为空; 密码不能为空",
		},
		{
			name: "UserLoginParams all empty",
			err:  (&UserLoginParams{}).Validate(),
			want: "用户名不能为空; 密码不能为空",
		},
		{
			name: "CreateSiteConfigParams all empty",
			err:  (&CreateSiteConfigParams{}).Validate(),
			want: "名称不能为空; 配置不能为空",
		},
		{
			// 旧自定义消息映射中 Name.required 的文案为 "ID不能为空"（历史文案），按原样保留。
			name: "UpdateSiteConfigParams all empty",
			err:  (&UpdateSiteConfigParams{}).Validate(),
			want: "ID不能为空; 配置不能为空",
		},
		{
			name: "WriteFileParams all empty",
			err:  (&WriteFileParams{}).Validate(),
			want: "文件名不能为空; 数据不能为空",
		},
		{
			name: "UploadToS3Params all empty (no custom messages -> 参数错误)",
			err:  (&UploadToS3Params{}).Validate(),
			want: "参数错误; 参数错误; 参数错误",
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

// TestValidateFirstMessage 断言 HTTP 入口取第一条消息的行为（维持既有 msg）。
func TestValidateFirstMessage(t *testing.T) {
	err := (&CreateDataListParams{}).Validate()
	if err == nil {
		t.Fatal("expected error")
	}
	first := strings.Split(err.Error(), "; ")[0]
	if first != "标签必填" {
		t.Fatalf("first message = %q, want 标签必填", first)
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

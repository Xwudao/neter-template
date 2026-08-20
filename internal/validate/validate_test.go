package validate

import (
	"errors"
	"strings"
	"testing"
)

func TestRequired(t *testing.T) {
	tests := []struct {
		name string
		val  string
		want bool
	}{
		{"empty fails", "", true},
		{"whitespace passes (no trim)", "  ", false},
		{"non-empty passes", "abc", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := Required()(tt.val)
			if (err != nil) != tt.want {
				t.Fatalf("Required(%q) err=%v, wantErr=%v", tt.val, err, tt.want)
			}
		})
	}
}

func TestNotZero(t *testing.T) {
	tests := []struct {
		name string
		val  any
		fn   func(any) error
		want bool
	}{
		{"zero int64 fails", int64(0), func(v any) error { return NotZero[int64]()(v.(int64)) }, true},
		{"non-zero int64 passes", int64(7), func(v any) error { return NotZero[int64]()(v.(int64)) }, false},
		{"nil pointer fails", (*int)(nil), func(v any) error { return NotZero[*int]()(v.(*int)) }, true},
		{"non-nil pointer passes", intPtr(1), func(v any) error { return NotZero[*int]()(v.(*int)) }, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.fn(tt.val)
			if (err != nil) != tt.want {
				t.Fatalf("%s err=%v, wantErr=%v", tt.name, err, tt.want)
			}
		})
	}
}

func intPtr(v int) *int { return &v }

func TestMinMax(t *testing.T) {
	tests := []struct {
		name string
		val  int
		rule Rule[int]
		want bool
	}{
		{"min pass boundary", 1, Min(1), false},
		{"min fail below", 0, Min(1), true},
		{"max pass boundary", 100, Max(100), false},
		{"max fail above", 101, Max(100), true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.rule(tt.val)
			if (err != nil) != tt.want {
				t.Fatalf("%s err=%v, wantErr=%v", tt.name, err, tt.want)
			}
		})
	}
}

func TestMinLenMaxLenRunesNotBytes(t *testing.T) {
	// "中文标题" 4 runes / 12 bytes：按 rune 长度 4 应通过 MaxLen(5)、按字节 12 应失败。
	if err := MaxLen(5)("中文标题"); err != nil {
		t.Fatalf("MaxLen(5) on 4-rune string should pass by rune count: %v", err)
	}
	if err := MinLen(4)("中文标题"); err != nil {
		t.Fatalf("MinLen(4) on 4-rune string should pass: %v", err)
	}
	if err := MinLen(5)("中文标题"); err != nil {
		// 4 runes < 5，应失败
	} else {
		t.Fatal("MinLen(5) on 4-rune string should fail")
	}
	if err := MaxLen(3)("中文标题"); err == nil {
		t.Fatal("MaxLen(3) on 4-rune string should fail")
	}
}

func TestEmail(t *testing.T) {
	tests := []struct {
		name string
		val  string
		want bool
	}{
		{"valid", "user@example.com", false},
		{"missing at", "not-an-email", true},
		{"empty", "", true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := Email()(tt.val)
			if (err != nil) != tt.want {
				t.Fatalf("Email(%q) err=%v, wantErr=%v", tt.val, err, tt.want)
			}
		})
	}
}

func TestOneOf(t *testing.T) {
	tests := []struct {
		name string
		val  string
		want bool
	}{
		{"in set passes", "asc", false},
		{"not in set fails", "desc2", true},
		{"empty fails (zero value)", "", true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := OneOf("asc", "desc")(tt.val)
			if (err != nil) != tt.want {
				t.Fatalf("OneOf(%q) err=%v, wantErr=%v", tt.val, err, tt.want)
			}
		})
	}
}

func TestMinItemsMaxItems(t *testing.T) {
	if err := MinItems[int64](1)(nil); err == nil {
		t.Fatal("MinItems(1) on nil slice should fail")
	}
	if err := MinItems[int64](1)([]int64{1}); err != nil {
		t.Fatalf("MinItems(1) on len-1 slice should pass: %v", err)
	}
	if err := MaxItems[int](2)([]int{1, 2, 3}); err == nil {
		t.Fatal("MaxItems(2) on len-3 slice should fail")
	}
}

func TestOptional(t *testing.T) {
	if err := Optional(Min(1))(nil); err != nil {
		t.Fatalf("Optional on nil should skip: %v", err)
	}
	zero := 0
	if err := Optional(Min(1))(&zero); err == nil {
		t.Fatal("Optional should still validate provided value")
	}
	one := 1
	if err := Optional(Min(1))(&one); err != nil {
		t.Fatalf("Optional should pass valid provided value: %v", err)
	}
}

func TestWhen(t *testing.T) {
	if err := When(false, Min(1))(0); err != nil {
		t.Fatalf("When(false) should skip: %v", err)
	}
	if err := When(true, Min(1))(0); err == nil {
		t.Fatal("When(true) should validate")
	}
}

func TestMessageReplacesText(t *testing.T) {
	err := Message("名称不能为空", Required())("")
	if err == nil || err.Error() != "名称不能为空" {
		t.Fatalf("Message should replace text, got %v", err)
	}
	if err := Message("名称不能为空", Required())("ok"); err != nil {
		t.Fatalf("valid value should pass: %v", err)
	}
}

func TestFieldShortCircuitsOnFirstRule(t *testing.T) {
	res := Field("name", "",
		Message("不能为空", Required()),
		Message("不能超长", MaxLen(10)),
	)
	if res.err == nil || res.err.Error() != "不能为空" {
		t.Fatalf("expected first rule message, got %v", res.err)
	}

	res = Field("name", "ok",
		Message("不能为空", Required()),
		Message("不能超长", MaxLen(1)),
	)
	if res.err == nil || res.err.Error() != "不能超长" {
		t.Fatalf("expected second rule message, got %v", res.err)
	}
}

func TestValidateAggregatesPerFieldFirstError(t *testing.T) {
	err := Validate(
		Field("a", "",
			Message("A必填", Required()),
			Message("A超长", MaxLen(1)),
		),
		Field("b", 0,
			Message("B最小为1", Min(1)),
		),
		Field("c", "ok"),
	)
	if err == nil {
		t.Fatal("expected aggregate error")
	}
	got := err.Error()
	for _, want := range []string{"A必填", "B最小为1"} {
		if !strings.Contains(got, want) {
			t.Fatalf("aggregate %q missing %q", got, want)
		}
	}
	if strings.Contains(got, "A超长") {
		t.Fatalf("aggregate should keep first rule per field, got %q", got)
	}
	if strings.Contains(got, "c") {
		t.Fatalf("valid field should not appear, got %q", got)
	}

	if err := Validate(Field("a", "ok")); err != nil {
		t.Fatalf("all-pass should be nil, got %v", err)
	}
}

func TestFirstError(t *testing.T) {
	es := Errors{{Field: "a", Message: "A必填"}, {Field: "b", Message: "B必填"}}
	if err := FirstError(es); err == nil || err.Error() != "A必填" {
		t.Fatalf("FirstError should return first message, got %v", err)
	}
	plain := errors.New("boom")
	if FirstError(plain) != plain {
		t.Fatal("FirstError should pass through non-aggregate errors")
	}
	if err := FirstError(nil); err != nil {
		t.Fatalf("FirstError(nil) should be nil, got %v", err)
	}
}

func TestErrorsFirstEmpty(t *testing.T) {
	if err := Errors(nil).First(); err != nil {
		t.Fatalf("empty Errors.First() should be nil, got %v", err)
	}
}

package core

import (
	"errors"
	"testing"

	"github.com/Xwudao/neter-template/internal/domain/errs"
)

func TestNewRtnWithErrUsesBusinessMessage(t *testing.T) {
	status := NewRtnWithErr(errs.NewConflict("名称已存在", errors.New("unique index")))

	if status.Code != CodeError || status.Message != "名称已存在" {
		t.Fatalf("unexpected status: %#v", status)
	}
}

package varx

import (
	"reflect"
	"testing"
)

func TestTrimArray(t *testing.T) {
	type args[T comparable] struct {
		arr []T
	}
	type testCase[T comparable] struct {
		name string
		args args[T]
		want []T
	}
	tests := []testCase[string]{
		{
			name: "test1",
			args: args[string]{
				arr: []string{"a", "b", "c", ""},
			},
			want: []string{"a", "b", "c"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := TrimArray(tt.args.arr, nil); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("TrimArray() = %v, want %v", got, tt.want)
			}
		})
	}
}

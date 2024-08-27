package models

import (
	"reflect"
	"testing"

	"github.com/Xwudao/neter-template/internal/data/ent"
)

func TestUnmarshalDataList(t *testing.T) {
	type args struct {
		arr  []*ent.DataList
		kind string
	}
	type testCase[T any] struct {
		name string
		args args
		want []T
	}
	tests := []testCase[DataLink]{
		{
			name: "case1",
			args: args{
				arr: []*ent.DataList{
					{
						Value: `{"name":"test","link":"http://test.com","open_blank":true,"enable":true}`,
						Label: `友链`,
						Kind:  `friend_link`,
						Key:   `some_key`,
					},
				},
				kind: "friend_link",
			},
			want: []DataLink{
				{
					Name:      "test",
					Link:      "http://test.com",
					OpenBlank: true,
					Enable:    true,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := UnmarshalDataList[DataLink](tt.args.arr, tt.args.kind); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UnmarshalDataList() = %v, want %v", got, tt.want)
			}
		})
	}
}

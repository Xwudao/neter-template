package models

import (
	json "github.com/json-iterator/go"

	"go-kitboxpro/internal/data/ent"
)

//type DataListBase struct {
//	Label string `json:"label"`
//	Kind  string `json:"kind"`
//	Value string `json:"value"`
//}

//type DataListModel interface {
//	GetKey() string
//	GetData() (*DataListBase, error)
//}
//func (d *DataLink) GetKey() string {
//	return utils.Md5Str(d.Link)
//}
//
//func (d *DataLink) GetData() (*DataListBase, error) {
//	value, err := sonic.MarshalString(d)
//	if err != nil {
//		return nil, err
//	}
//	return &DataListBase{
//		Label: d.Name,
//		Kind:  "friend_link",
//		Value: value,
//	}, nil
//}

type DataLink struct {
	Name      string `json:"name"`
	Link      string `json:"link"`
	OpenBlank bool   `json:"open_blank"`
	Enable    bool   `json:"enable"`
}

func UnmarshalDataList[T any](arr []*ent.DataList, kind string) []T {
	var res []T
	for _, v := range arr {
		if v.Kind != kind {
			continue
		}
		var item T
		if err := json.UnmarshalFromString(v.Value, &item); err != nil {
			continue
		}
		res = append(res, item)
	}

	return res
}

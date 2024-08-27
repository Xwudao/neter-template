package varx

import (
	"encoding/json"

	"github.com/Xwudao/neter-template/internal/domain/models"
)

func MustMarshal(v any) string {
	marshal, err := json.Marshal(v)
	if err != nil {
		return ""
	}

	return string(marshal)
}

func MustMarshalDefault(v any) string {
	if v, ok := v.(models.ConfigDefault); ok {
		v.GetDefault()
	}

	return MustMarshal(v)
}

package database

import (
	"github.com/json-iterator/go"
)

func ToJSON(src interface{}) []byte {
	if src == nil {
		return nil
	}

	jval, _ := jsoniter.Marshal(src)
	return jval
}

func FromJSON(src interface{}) interface{} {
	if src == nil {
		return nil
	}
	if bb, ok := src.([]byte); ok {
		var out interface{}
		_ = jsoniter.Unmarshal(bb, &out)
		return out
	}
	return nil
}

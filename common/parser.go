package common

import jsoniter "github.com/json-iterator/go"

func FromMapToByte(src interface{}) []byte {
	if src == nil {
		return nil
	}
	value, err := jsoniter.Marshal(src)
	if err != nil {
		return nil
	}
	return value
}

func FromByteToMap(src []byte) interface{} {
	if src == nil {
		return nil
	}

	var out interface{}
	err := jsoniter.Unmarshal(src, &out)
	if err != nil {
		return nil
	}

	return out
}

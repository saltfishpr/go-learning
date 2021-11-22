// @description: json序列化与反序列化
// @file: json.go
// @date: 2021/11/21

package utils

import jsoniter "github.com/json-iterator/go"

var json = jsoniter.ConfigCompatibleWithStandardLibrary

func JsonMarshal(data interface{}) ([]byte, error) {
	return json.Marshal(&data)
}

func JsonUnmarshal(data []byte, v interface{}) error {
	return json.Unmarshal(data, v)
}

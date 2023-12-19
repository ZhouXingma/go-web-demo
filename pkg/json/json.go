package json

import (
	"encoding/json"
)

// 将一个对象转换为string字符串
// 参数：
// data： 要转换为字符串的对象
// 返回：
// string: 字符串，json字符串
// error: 异常
func ToJsonString(data interface{}) (string, error) {
	s, e := json.Marshal(data)
	if e != nil {
		return "", e
	}
	return string(s), nil
}

// 将一个字json符串转换为对象
// 参数：
// jsonStr：字符串，json字符串
// v：要转换的对象
// 返回：
// error: 异常
func Parse2Object(jsonStr string, v any) error {
	e := json.Unmarshal([]byte(jsonStr), v)
	return e
}

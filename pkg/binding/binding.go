package binding

import (
	"github.com/go-playground/validator/v10"
	"reflect"
)

// 获取binding中tag为msg的信息
// 参数：
// error：异常信息，主要用于获取异常的字段
// obj: 要获取信息的结构
// 返回：
// string: msg信息
func GetBindingErrorMsg(err error, obj any) string {
	// 使用的时候，需要传obj的指针
	getObj := reflect.TypeOf(obj)
	if errs, ok := err.(validator.ValidationErrors); ok {
		// 断言成功
		for _, e := range errs {
			// 循环每一个错误信息
			// 根据报错字段名，获取结构体的具体字段
			f, exits := getObj.Elem().FieldByName(e.Field())
			if exits {
				msg := f.Tag.Get("msg")
				return msg
			}
		}
	}
	return err.Error()
}

// 获取某个结构的tag信息
// 参数：
// obj: 某个对象，需要传obj的指针
// fieldName： 字段名称
// tag: 标签名称
// 返回：
// string: msg信息
func GetBindingTagInfo(obj any, fieldName string, tag string) string {
	getObj := reflect.TypeOf(obj)
	f, exits := getObj.Elem().FieldByName(fieldName)
	if exits {
		tagValue := f.Tag.Get(tag)
		return tagValue
	}
	return ""
}

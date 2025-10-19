package utils

import (
	"encoding/json"
	"reflect"
)

// ToJSON 将任意对象转换为JSON字符串，便于打印日志
func ToJSON(v interface{}) string {
	if v == nil {
		return "null"
	}

	// 如果已经是字符串，直接返回
	if str, ok := v.(string); ok {
		return str
	}

	// 如果是指针，获取指向的值
	val := reflect.ValueOf(v)
	if val.Kind() == reflect.Ptr {
		if val.IsNil() {
			return "null"
		}
		v = val.Elem().Interface()
	}

	// 转换为JSON
	data, err := json.Marshal(v)
	if err != nil {
		return "{}"
	}

	return string(data)
}

// ToPrettyJSON 将任意对象转换为格式化的JSON字符串，便于调试
func ToPrettyJSON(v interface{}) string {
	if v == nil {
		return "null"
	}

	// 如果是指针，获取指向的值
	val := reflect.ValueOf(v)
	if val.Kind() == reflect.Ptr {
		if val.IsNil() {
			return "null"
		}
		v = val.Elem().Interface()
	}

	// 转换为格式化的JSON
	data, err := json.MarshalIndent(v, "", "  ")
	if err != nil {
		return "{}"
	}

	return string(data)
}

// ToJSONBytes 将任意对象转换为JSON字节数组
func ToJSONBytes(v interface{}) []byte {
	if v == nil {
		return []byte("null")
	}

	// 如果是指针，获取指向的值
	val := reflect.ValueOf(v)
	if val.Kind() == reflect.Ptr {
		if val.IsNil() {
			return []byte("null")
		}
		v = val.Elem().Interface()
	}

	// 转换为JSON
	data, err := json.Marshal(v)
	if err != nil {
		return []byte("{}")
	}

	return data
}

// FromJSON 将JSON字符串转换为指定类型的对象
func FromJSON(data string, v interface{}) error {
	return json.Unmarshal([]byte(data), v)
}

// FromJSONBytes 将JSON字节数组转换为指定类型的对象
func FromJSONBytes(data []byte, v interface{}) error {
	return json.Unmarshal(data, v)
}

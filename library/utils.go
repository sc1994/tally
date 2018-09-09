package library

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"reflect"
)

// Md5String 获取字符串的MD5
func Md5String(str string) string {
	h := md5.New()
	h.Write([]byte(str))
	sum := h.Sum(nil)
	return hex.EncodeToString(sum)
}

// IsEmpty 判断一个数据是否为空，支持int, float, string, slice, array, map的判断
func IsEmpty(value interface{}) bool {
	if value == nil {
		return true
	}
	switch reflect.TypeOf(value).Kind() {
	case reflect.String, reflect.Slice, reflect.Array, reflect.Map:
		if reflect.ValueOf(value).Len() == 0 {
			return true
		}
		return false
	}
	return false
}

// GetString 通过interface{}获取字符串
func GetString(val interface{}) string {
	return fmt.Sprintf("%v", val)
}

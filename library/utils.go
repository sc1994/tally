package library

import (
	"bytes"
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"io/ioutil"
	"net/http"
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

// HTTPRequest HttpRequest
func HTTPRequest(url string, jsonStr interface{}, token string) string {
	var req *http.Request
	var err error
	if jsonStr != nil {
		// var jsonStr = []byte(`{"pwd":"123123","name":"test"}`)
		str := []byte(fmt.Sprintf("%v", jsonStr))
		req, err = http.NewRequest("POST", url, bytes.NewBuffer(str))
		req.Header.Set("Content-Type", "application/json")
	} else {
		req, err = http.NewRequest("GET", url, nil)
	}

	req.Header.Set("token", token)
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	return string(body[:])
}

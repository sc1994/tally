package test

import (
	"encoding/json"
	"tally/library"
	"tally/models"
)

var baseURL = "http://localhost/"
var token = "21bd59b757e5"

func httpRequest(router string, jsonStr interface{}) models.BaseResponse {
	url := baseURL + router
	body := library.HTTPRequest(url, jsonStr, token)
	result := models.BaseResponse{}
	json.Unmarshal([]byte(body), &result)
	return result
}

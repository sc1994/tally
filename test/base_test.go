package test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"tally/models"
)

var baseURL = "http://localhost/"
var token = "21bd5911a84679cd19480a616f19a941|1ee169d11f1e1aa85cd0585d515eef8c"

func httpRequest(router string, jsonStr interface{}) models.BaseResponse {
	url := baseURL + router
	var req *http.Request
	var err error
	if jsonStr != nil {
		// var jsonStr = []byte(`{"pwd":"123123","name":"test"}`)
		str := []byte(fmt.Sprintf("%v", jsonStr))
		req, err = http.NewRequest("POST", url, bytes.NewBuffer(str))
	} else {
		req, err = http.NewRequest("GET", url, nil)
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("token", token)
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	result := models.BaseResponse{}
	json.Unmarshal(body, &result)
	return result
}

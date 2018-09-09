package test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"tally/models"
	"testing"
)

var baseURL = "http://localhost/"
var token = "44870d91af48e2601bb1afeee70edf5f|30eb8a1dc2566aca3dfef68ec180e628"

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

func Test_Login_1(t *testing.T) {
	result := httpRequest("signin", `{"name":"test","pwd":"123123"}`)
	if result.Code != 0 {
		j, _ := json.Marshal(result)
		t.Fatal("Code==0", string(j))
	}
	if result.Data == nil {
		j, _ := json.Marshal(result)
		t.Fatal("Data==nil", string(j))
	}
	token = fmt.Sprintf("%v", result.Data)
}

func Test_CheckName_1(t *testing.T) {
	result := httpRequest("signupcheck\\test", nil)
	if result.Code != 0 {
		j, _ := json.Marshal(result)
		t.Fatal("Code==0", string(j))
	}
}

func Test_Logout_1(t *testing.T) {
	result := httpRequest("removetoken", nil)
	if result.Code != 0 {
		j, _ := json.Marshal(result)
		t.Fatal("Code==0", string(j))
	}
}

func Test_LogUp_1(t *testing.T) {
	result := httpRequest("signup", `{"name":"test6","pwd":"123123"}`)
	if result.Code != 0 {
		j, _ := json.Marshal(result)
		t.Fatal("Code==0", string(j))
	}
}

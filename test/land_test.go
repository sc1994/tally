package test

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"tally/models"
	"testing"
)

var baseURL = "http://localhost/"

func Test_Login_1(t *testing.T) {
	url := baseURL + "signin"
	var jsonStr = []byte(`{"pwd":"123123","name":"test"}`)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("token", "token_test")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	result := models.BaseResponse{}
	json.Unmarshal(body, &result)
	j, _ := json.Marshal(result)
	if result.Code != 0 {
		t.Fatal("Code==0", string(j))
	}
	if result.Data == nil {
		t.Fatal("Data==nil", string(j))
	}
}

package test

import (
	"encoding/json"
	"fmt"
	"testing"
)

func Test_Login_1(t *testing.T) {
	result := httpRequest("land/signin", `{"name":"test","pwd":"123123"}`)
	if result.Msg != "success" {
		j, _ := json.Marshal(result)
		t.Fatal("result.Msg != success", string(j))
	}
	if result.Data == nil {
		j, _ := json.Marshal(result)
		t.Fatal("Data==nil", string(j))
	}
	token = fmt.Sprintf("%v", result.Data)
}

func Test_CheckName_1(t *testing.T) {
	result := httpRequest("signupcheck/test", nil)
	if result.Msg != "success" {
		j, _ := json.Marshal(result)
		t.Fatal("result.Msg != success", string(j))
	}
}

func Test_Logout_1(t *testing.T) {
	result := httpRequest("land/signout", nil)
	if result.Msg != "success" {
		j, _ := json.Marshal(result)
		t.Fatal("result.Msg != success", string(j))
	}
}

func Test_LogUp_1(t *testing.T) {
	result := httpRequest("/land/signup", `{"name":"test","pwd":"123123"}`)
	if result.Msg != "success" {
		j, _ := json.Marshal(result)
		t.Fatal("result.Msg != success", string(j))
	}
}

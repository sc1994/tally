package test

import (
	"encoding/json"
	"testing"
)

func Test_AddTally_1(t *testing.T) {
	jsonStr := `{"money":100,"type":"吃饭","mode":"支出","channel":"支付宝","remark":"吃饭支出了100元，通过支付宝支付。","ttime":"2018-09-10T12:02:23.046Z"}`
	result := httpRequest("inserttally", jsonStr)
	if result.Msg != "success" {
		j, _ := json.Marshal(result)
		t.Fatal("result.Msg != success", string(j))
	}
}

func Test_GetTally_2(t *testing.T) {
	jsonStr := `{"uids":["5b965bf4e3790f47a434e0dd"],"btime":"2018-09-01T12:02:23.046Z","etime":"2018-09-30T12:02:23.046Z","bMoney":0,"eMoney":999999,"types":[],"modes":[],"channels":[],"pageIndex":1,"pageSize":12}`
	result := httpRequest("gettallybyuser", jsonStr)
	if result.Msg != "success" {
		j, _ := json.Marshal(result)
		t.Fatal("result.Msg != success", string(j))
	}
	t.Log(result)
}

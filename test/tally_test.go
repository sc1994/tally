package test

import (
	"encoding/json"
	"tally/data"
	"testing"

	"gopkg.in/mgo.v2/bson"
)

func Test_AddTally_1(t *testing.T) {
	jsonStr := `{"money":100,"type":"吃饭","mode":"支出","channel":"支付宝","remark":"吃饭支出了100元，通过支付宝支付。","ttime":"2018-09-10T12:02:23.046Z"}`
	result := httpRequest("tally/add", jsonStr)
	if result.Msg != "success" {
		j, _ := json.Marshal(result)
		t.Fatal("result.Msg != success", string(j))
	}
}

func Test_GetTally_2(t *testing.T) {
	jsonStr := `{"uids":["5b965bf4e3790f47a434e0dd"],"btime":"2018-09-01T12:02:23.046Z","etime":"2018-09-30T12:02:23.046Z","bMoney":0,"eMoney":999999,"types":[],"modes":[],"channels":[],"pageIndex":1,"pageSize":12}`
	result := httpRequest("tally/get", jsonStr)
	if result.Msg != "success" {
		j, _ := json.Marshal(result)
		t.Fatal("result.Msg != success", string(j))
	}
	t.Log(result)
}

func Test_Sum_1(t *testing.T) {
	var result []map[string]interface{}
	/*
		bson.M{
			"$match": bson.M{"money": bson.M{"$gte": 1}}},
		bson.M{
			"$group": bson.M{"_id": nil, "totals": bson.M{"$sum": "$money"}}}
	*/
	data.Pipe("tally",
		[]bson.M{
			{"$match": bson.M{"money": bson.M{"$gte": 1}}},
			{"$group": bson.M{"_id": nil, "totals": bson.M{"$sum": "$money"}}},
		}, &result)
	t.Error(result)
}

type TestTest struct {
	id    string  `bson:"_id"`
	Total float64 `bson:"totals"`
}

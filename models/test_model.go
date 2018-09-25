package models

import (
	"tally/data"
)

const testTable = "test"

// Tally 账单
type TestStruct struct {
	BaseModel
	Remark string `json:"remark" bson:"remark"`
}

func AddTestStruct() {
	// data.Insert(testTable, TestStruct{
	// 	ID:         bson.NewObjectId(),
	// 	UpdateTime: time.Now(),
	// })
}

func GetTestStruct() (result []TestStruct) {
	data.Find(testTable, nil, &result)
	return result
}

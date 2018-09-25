package models

import (
	"reflect"
	"time"

	"gopkg.in/mgo.v2"

	"gopkg.in/mgo.v2/bson"
)

// BaseModel 实体基
type BaseModel struct {
	ID         bson.ObjectId `json:"id" bson:"_id"`      // ID
	CreateTime time.Time     `json:"ctime" bson:"ctime"` // CreateTime 创建时间
	UpdateTime time.Time     `json:"utime" bson:"utime"` // UpdateTime 更新时间
}

// BaseResponse 响应基类
type BaseResponse struct {
	Code   int         `json:"code"`
	Data   interface{} `json:"data"`
	Msg    string      `json:"msg"`
	Status string
}

// Baser 基类接口
type Baser interface {
	Get(search map[string]interface{}) (result []interface{})
	Set(update map[string]interface{}, selector map[string]interface{}) *mgo.ChangeInfo
	Add()
	Delete(selector map[string]interface{})
}

func baseUpdateField(update *map[string]interface{}) {
	if val, ok := (*update)["$set"]; ok {
		if reflect.Map == reflect.TypeOf(val).Kind() {
			m := val.(bson.M)
			if _, o := m["utime"]; !o {
				m["utime"] = time.Now()
			}
		}
	} else {
		(*update)["$set"] = bson.M{"utime": time.Now()}
	}
}

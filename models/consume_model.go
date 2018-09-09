package models

import (
	"gopkg.in/mgo.v2/bson"
)

// Consume 消费类型
type Consume struct {
	BaseModel
	UserID  bson.ObjectId `json:"uid" bson:"uid"`         // 关联用户id
	Content string        `json:"content" bson:"content"` // 内容
	Count   int64         `json:"count" bson:"count"`     // 使用次数
	Default []string      `json:"default" bson:"default"` // 默认消费的模式
}

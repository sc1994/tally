package models

import (
	"gopkg.in/mgo.v2/bson"
)

// Channel 消费渠道
type Channel struct {
	BaseModel
	UserID  bson.ObjectId `json:"uid" bson:"uid"`         // 关联用户id
	Content string        `json:"content" bson:"content"` // 内容
	Count   int64         `json:"count" bson:"count"`     // 使用次数
	Default []string      `json:"default" bson:"default"` // 默认渠道
}

// ChannelRequest 请求
type ChannelRequest struct {
	Channel
}

// ChannelResponse 响应
type ChannelResponse struct {
	Channel
}

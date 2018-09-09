package models

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

// Tally 账单
type Tally struct {
	ID         bson.ObjectId `json:"id" bson:"_id,omitempty"`
	CreateTime time.Time     `json:"ctime" bson:"ctime"`
	UpdateTime time.Time     `json:"utime" bson:"utime"`
	UserID     bson.ObjectId `json:"userID" bson:"uid"` //todo userID ==> uid
	Money      float32       `json:"money" bson:"money"`
	Type       string        `json:"type" bson:"type"`       //消费类型
	Mode       string        `json:"mode" bson:"mode"`       // 消费模式(预支,收入等)
	Channel    string        `json:"channel" bson:"channel"` // 消费渠道
	Remark     string        `json:"remark" bson:"remark"`
}

// TallyRequest 请求
type TallyRequest struct {
	Tally
	PageIndex int  `json:"pageIndex"`
	PageSize  int  `json:"pageSize"`
	OnlyMe    bool `json:"onlyMe"`
}

// TallyResponse 响应
type TallyResponse struct {
	Tally
	User
	TID        bson.ObjectId `json:"tid"`
	CreateTime time.Time     `json:"ctime"`
}

const tallyTable string = "tally"

func AddTally(t ...*Tally) {

}

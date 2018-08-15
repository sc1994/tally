package model

import (
	"tally/data"
	"time"

	"gopkg.in/mgo.v2/bson"
)

// Tally 账单模型
type Tally struct {
	ID         bson.ObjectId `json:"id" bson:"_id,omitempty"`
	UserID     bson.ObjectId `json:"userID" bson:"uid"`
	Money      float32       `json:"money" bson:"money"`
	Type       string        `json:"type" bson:"type"`
	Mode       string        `json:"mode" bson:"mode"`
	Channel    string        `json:"channel" bson:"channel"`
	CreateTime time.Time     `json:"ctime" bson:"ctime"`
	Remark     string        `json:"remark" bson:"remark"`
}

// TallyRequest 账单模型的请求模型
type TallyRequest struct {
	Token      string    `json:"token"`
	Money      float32   `json:"money"`
	Type       string    `json:"type"`
	Mode       string    `json:"mode"`
	Channel    string    `json:"channel"`
	Remark     string    `json:"remark"`
	CreateTime time.Time `json:"ctime"`
	PageIndex  int       `json:"pageIndex"`
	PageSize   int       `json:"pageSize"`
}

const tallyDB string = "tally"
const tallyTable string = "tally"

// InsertTally 新增消费信息
func (t *Tally) InsertTally() bool {
	return data.Insert(tallyDB, tallyTable, t)
}

// FindTally 获取消费信息
func (t *Tally) FindTally(pageIndex int, pageSize int, search interface{}, result interface{}) bool {
	return data.Page(tallyDB, tallyTable, pageIndex, pageSize, "-ctime", search, result)
}

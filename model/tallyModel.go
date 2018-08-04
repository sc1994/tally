package model

import (
	"tally/data"
	"time"

	"gopkg.in/mgo.v2/bson"
)

type Tally struct {
	Id         bson.ObjectId `json:"id" bson:"_id,omitempty"`
	UserId     bson.ObjectId `json:"userId" bson:"uid"`
	Money      float32       `json:"money" bson:"money"`
	Type       string        `json:"type" bson:"type"`
	Mode       string        `json:"mode" bson:"mode"`
	Channel    string        `json:"channel" bson:"channel"`
	CreateTime time.Time     `json:"ctime" bson:"ctime"`
	Remark     string        `json:"remark" bson:"remark"`
}

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

// GetTally 获取消费信息
func (t *Tally) GetTally(pageIndex int, pageSize int, search interface{}, result interface{}) bool {
	return data.Page(tallyDB, tallyTable, pageIndex, pageSize, "-ctime", search, result)
}

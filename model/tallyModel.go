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
	ID         bson.ObjectId `json:"id"`
	Token      string        `json:"token"`
	Money      float32       `json:"money"`
	Type       string        `json:"type"`
	Mode       string        `json:"mode"`
	Channel    string        `json:"channel"`
	Remark     string        `json:"remark"`
	CreateTime time.Time     `json:"ctime"`
	PageIndex  int           `json:"pageIndex"`
	PageSize   int           `json:"pageSize"`
}

// TallyResponse 响应实体
type TallyResponse struct {
	Tally
	User
	TID        bson.ObjectId `json:"tid"`
	CreateTime time.Time     `json:"ctime"`
}

const tallyDB string = "tally"
const tallyTable string = "tally"

// InsertTally 新增消费信息
func (t *Tally) InsertTally() bool {
	return data.Insert(tallyDB, tallyTable, t)
}

// FindTallyPage 获取消费分页信息 等待重构,查询语法不应该出现在controller
func (t *Tally) FindTallyPage(pageIndex int, pageSize int, search interface{}, result interface{}) int {
	return data.Page(tallyDB, tallyTable, pageIndex, pageSize, "-ctime", search, result)
}

// FindTallyByMonth 获取当前月消费信息
func FindTallyByMonth(userIDs []bson.ObjectId, result interface{}) {
	now := time.Now()
	currentYear, currentMonth, _ := now.Date()
	currentLocation := now.Location()
	firstOfMonth := time.Date(currentYear, currentMonth, 1, 0, 0, 0, 0, currentLocation)
	search := bson.M{"uid": bson.M{"$in": userIDs}, "ctime": bson.M{"$gte": firstOfMonth}}
	data.Find(tallyDB, tallyTable, search, result)
}

// UpdateTallyByID 依据id更新数据
func (t *Tally) UpdateTallyByID() bool {
	selector := bson.M{"_id": t.ID}
	update := bson.M{"$set": bson.M{
		"money":   t.Money,
		"type":    t.Type,
		"mode":    t.Mode,
		"channel": t.Channel,
		"remark":  t.Remark,
	},
	}
	return data.Update(tallyDB, tallyTable, selector, update)
}

package model

import (
	"tally/data"
	"time"

	"gopkg.in/mgo.v2/bson"
)

// Message 消息实体
type Message struct {
	ID         bson.ObjectId `json:"id" bson:"_id"`              // 主键
	FromID     bson.ObjectId `json:"fid" bson:"fid"`             // 来自谁
	ToID       bson.ObjectId `json:"tid" bson:"tid"`             // 发送给谁
	Content    string        `json:"content" bson:"content"`     // 消息内容
	Status     int           `json:"status" bson:"status"`       // 状态 1.未读 2.已读 3.同意 4.拒绝
	NeedTouch  int           `json:"needTouch" bson:"needTouch"` // 需要点击阅读 0.false 1.true
	CreateTime time.Time     `json:"ctime" bson:"ctime"`         // 创建时间
	UpdateTime time.Time     `json:"utime" bson:"utime"`         // 更新时间(代表消息的状态修改时间)
}

// MessageRequest 消息的请求实体
type MessageRequest struct {
	ID        bson.ObjectId `json:"id"`        // 主键
	FromID    bson.ObjectId `json:"fid"`       // 来自谁
	ToID      bson.ObjectId `json:"tid"`       // 发送给谁
	Content   string        `json:"content" `  // 消息内容
	Type      int           `json:"type"`      // 消息类型 1.邀请记账
	NeedTouch int           `json:"needTouch"` // 需要点击阅读 0.false 1.true
}

const messageDB string = "tally"
const messageTable string = "message"

// InsertMessage 新增一条消息
func (m *Message) InsertMessage() bool {
	m.ID = bson.NewObjectId()
	m.CreateTime = time.Now()
	m.UpdateTime = time.Now()
	m.Status = 1
	return data.Insert(messageDB, messageTable, m)
}

// UpdateMessageStatus 更新消息状态
func UpdateMessageStatus(id bson.ObjectId, status int) bool {
	update := bson.M{"status": status}
	selector := bson.M{"_id": id}
	return data.Update(messageDB, messageTable, selector, update)
}

// FindMessageByID 依据id获悉消息详细
func FindMessageByID(id bson.ObjectId) (Message, bool) {
	search := bson.M{"_id": id}
	var r []Message
	data.Find(messageDB, messageTable, search, &r)
	if len(r) < 0 {
		return Message{}, false
	}
	return r[0], true
}

// FindMessageByMe 获取我的全部消息
func FindMessageByMe(toID bson.ObjectId, result interface{}) {
	search := bson.M{"tid": toID}
	data.Find(messageDB, messageTable, search, result)
}

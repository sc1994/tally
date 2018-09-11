package models

import (
	"tally/data"
	"time"

	linq "github.com/ahmetb/go-linq"
	"gopkg.in/mgo.v2/bson"
)

const messageTable string = "message"

// Message 消息
type Message struct {
	ID         bson.ObjectId `json:"id" bson:"_id"`              // ID
	CreateTime time.Time     `json:"ctime" bson:"ctime"`         // CreateTime 创建时间
	UpdateTime time.Time     `json:"utime" bson:"utime"`         // UpdateTime 更新时间
	FromID     bson.ObjectId `json:"fid" bson:"fid"`             // 来自谁
	ToID       bson.ObjectId `json:"tid" bson:"tid"`             // 发送给谁
	Content    string        `json:"content" bson:"content"`     // 消息内容
	Type       int           `json:"type" bson:"type"`           // 消息类型 1.小伙伴邀请 2.邀请回执 3.账单消息
	Status     int           `json:"status" bson:"status"`       // 状态 1.未读 2.已读 3.同意 4.拒绝
	NeedTouch  bool          `json:"needTouch" bson:"needTouch"` // 需要点击阅读
}

// MessageRequest 请求
type MessageRequest struct {
	ID        bson.ObjectId `json:"id"`             // 主键
	FromID    bson.ObjectId `json:"fid"`            // 来自谁
	ToID      bson.ObjectId `json:"tid" bson:"tid"` // 发送给谁
	Content   string        `json:"content" `       // 消息内容
	Type      int           `json:"type"`           // 消息类型 1.小伙伴邀请 2.邀请回执 3.账单消息
	NeedTouch bool          `json:"needTouch"`      // 需要点击阅读
	PageIndex int           `json:"pageIndex"`
	PageSize  int           `json:"pageSize"`
}

// MessageResponse 响应
type MessageResponse struct {
	*Message
	FromNick string `json:"fnick" bson:"fnick"` // 昵称
	FromImg  string `json:"fimg" bson:"fimg"`   // 头像
	ToNick   string `json:"tnick" bson:"tnick"` // 昵称
	ToImg    string `json:"timg" bson:"timg"`   // 头像
}

// AddMessage AddMessage
func AddMessage(docs ...*Message) []bson.ObjectId {
	d := make([]interface{}, len(docs))
	for i, v := range docs {
		v.ID = bson.NewObjectId()
		v.CreateTime = time.Now()
		v.UpdateTime = time.Now()
		d[i] = *v
	}
	data.Insert(messageTable, d)
	var result []bson.ObjectId
	linq.From(docs).Select(func(x interface{}) interface{} {
		return x.(*Message).ID
	}).ToSlice(&result)
	return result
}

// Get Get
func (c *MessageRequest) Get(search map[string]interface{}) (result []*MessageResponse) {
	var flag []*MessageRequest
	data.Page(messageTable, search, c.PageIndex, c.PageSize, &flag, "-ctime")
	uids := make([]bson.ObjectId, 0)
	linq.From(flag).SelectMany(func(x interface{}) linq.Query {
		return linq.From([]bson.ObjectId{x.(*Message).ToID, x.(*Message).FromID})
	}).Distinct().ToSlice(&uids)
	users := new(UserRequest).Get(bson.M{"_id": bson.M{"$in": uids}}) // 相关用户
	linq.From(flag).Select(func(x interface{}) interface{} {
		tu := linq.From(users).FirstWith(func(f interface{}) bool { // 接受用户
			return f.(*User).ID == x.(*Message).ToID
		})
		fu := linq.From(users).FirstWith(func(f interface{}) bool { // 发送用户
			return f.(*User).ID == x.(*Message).FromID
		})
		if tu == nil || fu == nil { // 任意为空返回 nil ,后续有剔除操作
			return nil
		}
		return &MessageResponse{ // 返回最终数据格式
			Message:  x.(*Message),
			FromNick: fu.(*User).NickName,
			FromImg:  fu.(*User).HeadImg,
			ToNick:   tu.(*User).NickName,
			ToImg:    tu.(*User).HeadImg,
		}
	}).Where(func(x interface{}) bool {
		return x != nil
	}).ToSlice(&result)
	return
}

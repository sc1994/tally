package models

import (
	"tally/data"
	"time"

	linq "github.com/ahmetb/go-linq"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

const channelTable string = "channel"

// Channel 消费渠道
type Channel struct {
	ID         bson.ObjectId `json:"id" bson:"_id"`          // ID
	CreateTime time.Time     `json:"ctime" bson:"ctime"`     // CreateTime 创建时间
	UpdateTime time.Time     `json:"utime" bson:"utime"`     // UpdateTime 更新时间
	UserID     bson.ObjectId `json:"uid" bson:"uid"`         // 关联用户id
	Content    string        `json:"content" bson:"content"` // 内容
	Count      int64         `json:"count" bson:"count"`     // 使用次数
	Default    []string      `json:"default" bson:"default"` // 默认渠道
}

// ChannelRequest 请求
type ChannelRequest struct {
	Channel
}

// ChannelResponse 响应
type ChannelResponse struct {
	Channel
}

// Get Get
func (c *ChannelRequest) Get(search map[string]interface{}) (result []*Channel) {
	data.Find(channelTable, search, &result)
	return
}

// Set Set
func (c *ChannelRequest) Set(update map[string]interface{}, selector map[string]interface{}) *mgo.ChangeInfo {
	return data.Update(channelTable, update, selector)
}

// AddChannel Add
func AddChannel(docs ...*Channel) []bson.ObjectId {
	d := make([]interface{}, len(docs))
	for i, v := range docs {
		v.ID = bson.NewObjectId()
		v.CreateTime = time.Now()
		v.UpdateTime = time.Now()
		d[i] = *v
	}

	data.InsertBatch(channelTable, d)
	var result []bson.ObjectId
	linq.From(docs).Select(func(x interface{}) interface{} {
		return x.(*Channel).ID
	}).ToSlice(&result)
	return result
}

// Delete Delete
func (c *ChannelRequest) Delete(selector map[string]interface{}) {
	data.Remove(channelTable, selector)
}

// CopyChannel Copy
func CopyChannel() {
	data.CopyTable(channelTable)
}

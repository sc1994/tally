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
	BaseModel  `bson:",inline"`
	UserID     bson.ObjectId `json:"uid" bson:"uid"`             // 关联用户id
	Content    string        `json:"content" bson:"content"`     // 内容
	Count      int64         `json:"count" bson:"count"`         // 使用次数
	RefundDate []string      `json:"rdate" bson:"rdate"`         // 还款日期
	IsAdvance  bool          `json:"isAdvance" bson:"isAdvance"` // 是否预支渠道
	Default    []string      `json:"default" bson:"default"`     // 默认渠道
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
	baseUpdateField(&update)
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

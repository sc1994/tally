package models

import (
	"tally/data"
	"time"

	linq "github.com/ahmetb/go-linq"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

const consumeTable = "consume"

// Consume 消费类型
type Consume struct {
	ID         bson.ObjectId `json:"id" bson:"_id"`          // ID
	CreateTime time.Time     `json:"ctime" bson:"ctime"`     // CreateTime 创建时间
	UpdateTime time.Time     `json:"utime" bson:"utime"`     // UpdateTime 更新时间
	UserID     bson.ObjectId `json:"uid" bson:"uid"`         // 关联用户id
	Content    string        `json:"content" bson:"content"` // 内容
	Count      int64         `json:"count" bson:"count"`     // 使用次数
	Default    []string      `json:"default" bson:"default"` // 默认消费的模式
}

// ConsumeRequest 请求
type ConsumeRequest struct {
	Consume
}

// ConsumeResponse 响应
type ConsumeResponse struct {
	Consume
}

// Get Get
func (c *ConsumeRequest) Get(search map[string]interface{}) (result []*Consume) {
	data.Find(consumeTable, search, &result)
	return
}

// Set Set
func (c *ConsumeRequest) Set(update map[string]interface{}, selector map[string]interface{}) *mgo.ChangeInfo {
	return data.Update(consumeTable, update, selector)
}

// AddConsume Add
func AddConsume(docs ...*Consume) []bson.ObjectId {
	d := make([]interface{}, len(docs))
	for i, v := range docs {
		v.ID = bson.NewObjectId()
		v.CreateTime = time.Now()
		v.UpdateTime = time.Now()
		d[i] = *v
	}
	data.InsertBatch(consumeTable, d)
	var result []bson.ObjectId
	linq.From(docs).Select(func(x interface{}) interface{} {
		return x.(*Consume).ID
	}).ToSlice(&result)
	return result
}

// Delete Delete
func (c *ConsumeRequest) Delete(selector map[string]interface{}) {
	data.Remove(consumeTable, selector)
}

// CopyConsume Copy
func CopyConsume() {
	data.CopyTable(consumeTable)
}

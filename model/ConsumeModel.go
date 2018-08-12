package model

import (
	"sort"
	"tally/common"
	"tally/data"
	"time"

	"gopkg.in/mgo.v2/bson"
)

const consumeDB string = "tally"
const consumeTable string = "consume"

// Consume 消费类型
type Consume struct {
	Id         bson.ObjectId `json:"id" bson:"_id"`          // 主键
	UserId     bson.ObjectId `json:"uid" bson:"uid"`         // 关联用户id
	Content    string        `json:"content" bson:"content"` // 内容
	Count      int64         `json:"count" bson:"count"`     // 使用次数
	Default    []string      `json:"default" bson:"default"` // 默认消费的模式
	CreateTime time.Time     `json:"ctime" bson:"ctime"`     // CreateTime 创建时间
}

// InsertConsume 新增一条消费类型
func (c *Consume) InsertConsume() bool {
	c.CreateTime = time.Now()
	c.Id = bson.NewObjectId()
	return data.Insert(consumeDB, consumeTable, c)
}

// RemoveConsumeById 依据主键移除类型
func (c *Consume) RemoveConsumeById(id bson.ObjectId) bool {
	search := bson.M{"_id": id}
	return data.Delete(consumeDB, consumeTable, search)
}

// FindConsumeByUserId 依据用户Id查找消费类型
func FindConsumeByUserId(userId bson.ObjectId) (result []Consume) {
	search := bson.M{"uid": userId}
	data.Find(consumeDB, consumeTable, search, &result)
	if result == nil {
		result = []Consume{}
	}
	sort.Sort(ConsumebyCount(result))
	return result
}

// ExistConsume 是否存在相同的类型
func ExistConsume(userId bson.ObjectId, content string) bool {
	search := bson.M{"uid": userId, "content": content}
	var result []Consume
	data.Find(consumeDB, consumeTable, search, &result)
	if len(result) > 0 {
		return true
	}
	return false
}

// UpdateConsume 更新内容和默认模式信息
func (c *Consume) UpdateConsume() bool {
	selector := bson.M{"_id": c.Id}
	update := bson.M{"$set": bson.M{"content": c.Content, "default": c.Default}}
	return data.Update(consumeDB, consumeTable, selector, update)
}

// IncConsumeCount 类型累加1
func IncConsumeCount(userId bson.ObjectId, content string) bool {
	selector := bson.M{"content": content, "uid": userId}
	update := bson.M{"$inc": bson.M{"count": 1}}
	return data.Update(consumeDB, consumeTable, selector, update)
}

// InitConsume 初始化用户的消费类型
func InitConsume(userId bson.ObjectId) int {
	consumes := []Consume{
		Consume{
			UserId:  userId,
			Content: "吃饭",
			Count:   0,
			Default: []string{common.TallyMode[1], common.TallyMode[2]},
		},
		Consume{
			UserId:  userId,
			Content: "房租",
			Count:   0,
			Default: []string{common.TallyMode[1], common.TallyMode[2]},
		},
		Consume{
			UserId:  userId,
			Content: "工资",
			Count:   0,
			Default: []string{common.TallyMode[0]},
		},
	}
	result := 0
	for _, v := range consumes {
		b := v.InsertConsume()
		if b {
			result++
		}
	}
	return result
}

type ConsumebyCount []Consume

func (c ConsumebyCount) Len() int           { return len(c) }
func (c ConsumebyCount) Swap(i, j int)      { c[i], c[j] = c[j], c[i] }
func (c ConsumebyCount) Less(i, j int) bool { return c[i].Count > c[j].Count }

package model

import (
	"tally/common"
	"tally/data"
	"time"

	"gopkg.in/mgo.v2/bson"
)

const channelDB string = "tally"
const channelTable string = "channel"

// Channel 消费渠道
type Channel struct {
	Id         bson.ObjectId `json:"id" bson:"_id"`          // 主键
	UserId     bson.ObjectId `json:"uid" bson:"uid"`         // 关联用户id
	Content    string        `json:"content" bson:"content"` // 内容
	Count      int64         `json:"count" bson:"count"`     // 使用次数
	Default    []string      `json:"default" bson:"default"` // 默认渠道
	CreateTime time.Time     `json:"ctime" bson:"ctime"`     // CreateTime 创建时间
}

// InsertChannel 新增一条消费渠道
func (c *Channel) InsertChannel() bool {
	c.CreateTime = time.Now()
	c.Id = bson.NewObjectId()
	return data.Insert(channelDB, channelTable, c)
}

// RemoveChannelById 依据主键移除类型
func (c *Channel) RemoveChannelById(id bson.ObjectId) bool {
	search := bson.M{"_id": id}
	return data.Delete(channelDB, channelTable, search)
}

// FindChannelByUserId 依据用户Id查找信息
func FindChannelByUserId(userId bson.ObjectId) (result []Channel) {
	search := bson.M{"uid": userId}
	data.Find(channelDB, channelTable, search, &result)
	if result == nil {
		result = []Channel{}
	}
	return result
}

// ExistChannel 是否存在相同的类型
func ExistChannel(userId bson.ObjectId, content string) bool {
	search := bson.M{"uid": userId, "content": content}
	var result []*Channel
	data.Find(channelDB, channelTable, search, result)
	if len(result) > 0 {
		return true
	}
	return false
}

// UpdateChannel 更新内容和默认模式信息
func (c *Channel) UpdateChannel() bool {
	selector := bson.M{"_id": c.Id}
	update := bson.M{"content": c.Content, "default": c.Default}
	return data.Update(channelDB, channelTable, selector, update)
}

// IncChannelCount 渠道使用次数累加一
func IncChannelCount(id bson.ObjectId) bool {
	selector := bson.M{"_id": id}
	update := bson.M{"$inc": bson.M{"count": 1}}
	return data.Update(channelDB, channelTable, selector, update)
}

// InitChannel 初始化用户的渠道数据
func InitChannel(userId bson.ObjectId) int {
	channels := []Channel{
		Channel{
			UserId:  userId,
			Content: "支付宝",
			Count:   0,
			Default: []string{common.TallyMode[0], common.TallyMode[1]},
		},
		Channel{
			UserId:  userId,
			Content: "微信",
			Count:   0,
			Default: []string{common.TallyMode[0], common.TallyMode[1]},
		},
		Channel{
			UserId:  userId,
			Content: "银行卡",
			Count:   0,
			Default: []string{common.TallyMode[0], common.TallyMode[1]},
		},
		Channel{
			UserId:  userId,
			Content: "信用卡",
			Count:   0,
			Default: []string{common.TallyMode[2]},
		},
		Channel{
			UserId:  userId,
			Content: "现金",
			Count:   0,
			Default: []string{common.TallyMode[0], common.TallyMode[1]},
		},
		Channel{
			UserId:  userId,
			Content: "花呗",
			Count:   0,
			Default: []string{common.TallyMode[2]},
		},
		Channel{
			UserId:  userId,
			Content: "白条",
			Count:   0,
			Default: []string{common.TallyMode[2]},
		},
	}
	result := 0
	for _, v := range channels {
		b := v.InsertChannel()
		if b {
			result++
		}
	}
	return result
}

type ChannelbyCount []Channel

func (c ChannelbyCount) Len() int           { return len(c) }
func (c ChannelbyCount) Swap(i, j int)      { c[i], c[j] = c[j], c[i] }
func (c ChannelbyCount) Less(i, j int) bool { return c[i].Count > c[j].Count }

package models

import (
	"tally/data"
	"time"

	linq "github.com/ahmetb/go-linq"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

const tallyTable = "tally"

// Tally 账单
type Tally struct {
	ID         bson.ObjectId `json:"id" bson:"_id"`      // ID
	CreateTime time.Time     `json:"ctime" bson:"ctime"` // CreateTime 创建时间
	UpdateTime time.Time     `json:"utime" bson:"utime"` // UpdateTime 更新时间
	TallyTime  time.Time     `json:"ttime" bson:"ttime"` // 消费时间 (新加字段)
	UserID     bson.ObjectId `json:"uid" bson:"uid"`
	Money      float32       `json:"money" bson:"money"`
	Type       string        `json:"type" bson:"type"`       //消费类型
	Mode       string        `json:"mode" bson:"mode"`       // 消费模式(预支,收入等)
	Channel    string        `json:"channel" bson:"channel"` // 消费渠道
	Remark     string        `json:"remark" bson:"remark"`
}

// TallyRequest 请求
type TallyRequest struct {
	*Tally
	PageIndex  int             `json:"pageIndex"`
	PageSize   int             `json:"pageSize"`
	OnlyMe     bool            `json:"onlyMe"` // only ==> UIDS
	UserIDs    []bson.ObjectId `json:"uids"`
	BeginTime  time.Time       `json:"btime"`
	EndTime    time.Time       `json:"etime"`
	BeginMoney float32         `json:"bMoney"`
	EndMoney   float32         `json:"eMoney"`
	Types      []string        `json:"types"`
	Modes      []string        `json:"modes"`
	Channels   []string        `json:"channels"`
}

// TallyResponse 响应
type TallyResponse struct {
	*Tally
	*User
	TID        bson.ObjectId `json:"tid"`
	CreateTime time.Time     `json:"ctime"`
	CanEdit    bool          `json:"canEdit"`
}

// Add AddTally
func (c *TallyRequest) Add() bson.ObjectId {
	c.Tally.ID = bson.NewObjectId()
	c.Tally.CreateTime = time.Now()
	c.Tally.UpdateTime = time.Now()
	data.Insert(tallyTable, c.Tally)
	return c.Tally.ID
}

// Get get
func (c *TallyRequest) Get(seach map[string]interface{}) (result []*Tally) {
	data.Find(tallyTable, seach, &result)
	return
}

// Page 分页查询
func (c *TallyRequest) Page(search map[string]interface{}) (result []*TallyResponse) {
	var flag []*Tally
	data.Page(tallyTable, search, c.PageIndex, c.PageSize, &flag, "-ttime", "-ctime", "-utime")
	var uids []bson.ObjectId
	linq.From(flag).Select(func(x interface{}) interface{} {
		return x.(*Tally).UserID
	}).Distinct().ToSlice(&uids)
	users := new(UserRequest).Get(bson.M{"_id": bson.M{"$in": uids}})
	linq.From(flag).Select(func(x interface{}) interface{} {
		cuser := linq.From(users).FirstWith(func(f interface{}) bool {
			return f.(*User).ID == x.(*Tally).UserID
		})
		r := new(TallyResponse)
		r.Tally = (x.(*Tally))
		r.User = (cuser.(*User))
		r.TID = x.(*Tally).ID
		r.CreateTime = x.(*Tally).CreateTime
		return r
	}).ToSlice(&result)
	return
}

// Set Set
func (c *TallyRequest) Set(update map[string]interface{}, selector map[string]interface{}) *mgo.ChangeInfo {
	baseUpdateField(&update)
	return data.Update(tallyTable, update, selector)
}

// Delete Delete
func (c *TallyRequest) Delete(selector map[string]interface{}) {
	data.Remove(tallyTable, selector)
}

// Pipe Pipe
func (c *TallyRequest) Pipe(match map[string]interface{}, group map[string]interface{}) (result []map[string]interface{}) {
	data.Pipe(tallyTable, []bson.M{match, group}, &result)
	return
}

// CopyTally Copy
func CopyTally() {
	data.CopyTable(tallyTable)
}

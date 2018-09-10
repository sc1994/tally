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
	ID         bson.ObjectId `json:"id" bson:"_id,omitempty"`
	CreateTime time.Time     `json:"ctime" bson:"ctime"`
	UpdateTime time.Time     `json:"utime" bson:"utime"`
	TallyTime  time.Time     `json:"ttime" bson:"ttime"` // 消费时间 (新加字段)
	UserID     bson.ObjectId `json:"userID" bson:"uid"`  //todo userID ==> uid
	Money      float32       `json:"money" bson:"money"`
	Type       string        `json:"type" bson:"type"`       //消费类型
	Mode       string        `json:"mode" bson:"mode"`       // 消费模式(预支,收入等)
	Channel    string        `json:"channel" bson:"channel"` // 消费渠道
	Remark     string        `json:"remark" bson:"remark"`
}

// TallyRequest 请求
type TallyRequest struct {
	Tally
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
	Tally
	User
	TID        bson.ObjectId `json:"tid"`
	CreateTime time.Time     `json:"ctime"`
	CanEdit    bool          `json:"canEdit"`
}

// AddTally AddTally
func AddTally(docs ...*Tally) []bson.ObjectId {
	d := make([]interface{}, len(docs))
	for i, v := range docs {
		v.ID = bson.NewObjectId()
		v.CreateTime = time.Now()
		v.UpdateTime = time.Now()
		d[i] = *v
	}
	data.Insert(tallyTable, d)
	var result []bson.ObjectId
	linq.From(docs).Select(func(x interface{}) interface{} {
		return x.(*Tally).ID
	}).ToSlice(&result)
	return result
}

// PageTally 分页查询
func PageTally(search map[string]interface{}, index int, size int) (result []*TallyResponse) {
	var flag []*Tally
	data.Page(tallyTable, search, index, size, &flag, "-ttime", "-ctime", "-utime")
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
		r.Tally = *(x.(*Tally))
		r.User = *(cuser.(*User))
		r.TID = x.(*Tally).ID
		r.CreateTime = x.(*Tally).CreateTime
		return r
	}).ToSlice(&result)
	return
}

// Set Set
func (c *TallyRequest) Set(update map[string]interface{}, selector map[string]interface{}) *mgo.ChangeInfo {
	return data.Update(tallyTable, update, selector)
}

// Delete Delete
func (c *TallyRequest) Delete(selector map[string]interface{}) {
	data.Remove(tallyTable, selector)
}

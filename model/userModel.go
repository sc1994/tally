package model

import (
	"tally/data"
	"time"

	"encoding/json"

	"gopkg.in/mgo.v2/bson"
)

// User 用户信息实体
type User struct {
	Id         bson.ObjectId `json:"id" bson:"_id,omitempty"`    // Id
	Password   string        `json:"_" bson:"pwd"`               // Password
	Name       string        `json:"name" bson:"name"`           // Name
	NickName   string        `json:"nick" bson:"nick"`           // NickName 昵称
	CreateTime time.Time     `json:"ctime" bson:"ctime"`         // CreateTime 创建时间
	Consumes   []ManyType    `json:"consumes" bson:"consumes"`   // Consumes 消费类型
	Channels   []ManyType    `json:"channels" bson:"channels"`   // Channels 消费渠道
	Modes      []ManyType    `json:"modes" bson:"modes"`         // Modes 消费模式-->支出/收入/预支
	HeadImg    string        `json:"headImg" bson:"himg"`        // HeadImg 头像
	Budget     float32       `json:"budget" bson:"budget"`       // 月预算
	FixDate    float32       `json:"fixDate" bson:"fixDate"`     // 定期
	WechatPay  float32       `json:"wechatPay" bson:"wechatPay"` // 微信
	Alipay     float32       `json:"aliPay" bson:"aliPay"`       // 支付宝
	BackCard   float32       `json:"backCard" bson:"backCard"`   // 银行卡
}

// ManyType 泛指各种类型
type ManyType struct {
	Content string   `json:"content" bson:"content"` // 内容
	Count   int64    `json:"count" bson:"count"`     // 使用次数
	Default []string `json:"default" bson:"default"` // 默认关联内容
}

type ByCount []ManyType

func (c ByCount) Len() int           { return len(c) }
func (c ByCount) Swap(i, j int)      { c[i], c[j] = c[j], c[i] }
func (c ByCount) Less(i, j int) bool { return c[i].Count > c[j].Count }

// UserRequest 用户信息请求参数
type UserRequest struct {
	Token    string `json:"token"`
	Name     string `json:"name"`
	Password string `json:"password"`
	Content  string `json:"content"`
	Remember bool   `json:"remember"`
}

const userDB string = "tally"
const userTable string = "user"

// FindOne 获取一条用户数据
func (u *User) FindOne(search interface{}) {
	r := new(User).Find(search)
	if len(r) > 0 {
		*u = *(r[0])
	}
}

// Find 获取满足条件的用户数据
func (u *User) Find(search interface{}) (result []*User) {
	data.Find(userDB, userTable, search, &result)
	return
}

// Insert 新增一条用户信息
func (u *User) Insert() bool {
	return data.Insert(userDB, userTable, u)
}

// Update 更新一条用户信息
func (u *User) Update(token string, selector interface{}, update interface{}) bool {
	b := data.Update(userDB, userTable, selector, update)
	u.FindOne(selector)
	if data.RedisDel(token) {
		j, _ := json.Marshal(u)
		data.RedisSet(token, j, 7*24*60)
	}
	return b
}

// Init 依据token获取用户信息
func (u *User) Init(token string) bool {
	s, b := data.RedisGet(token)
	if b {
		e := json.Unmarshal([]byte(s), u)
		if e != nil {
			return false
		}
		u.Password = "*"
		return true
	}
	return false
}

// RemoveToken 移除token
func (u *User) RemoveToken(token string) {
	data.RedisDel(token)
}

package models

import (
	"tally/data"
	"tally/library"
	"time"

	"github.com/astaxie/beego"

	"gopkg.in/mgo.v2/bson"
)

// User 用户信息实体
type User struct {
	BaseModel
	Password   string          `json:"pwd" bson:"pwd"`               // Password
	Name       string          `json:"name" bson:"name"`             // Name
	NickName   string          `json:"nick" bson:"nick"`             // NickName 昵称
	Intro      string          `json:"intro" bson:"intro"`           // 简介
	HeadImg    string          `json:"headImg" bson:"himg"`          // HeadImg 头像
	Budget     float32         `json:"budget" bson:"budget"`         // 月预算
	FixDate    float32         `json:"fixDate" bson:"fixDate"`       // 定期
	WechatPay  float32         `json:"wechatPay" bson:"wechatPay"`   // 微信
	Alipay     float32         `json:"aliPay" bson:"aliPay"`         // 支付宝
	BackCard   float32         `json:"backCard" bson:"backCard"`     // 银行卡
	CreditCard float32         `json:"creditCard" bson:"creditCard"` // 信用卡
	Cash       float32         `json:"cash" bson:"cash"`             // 现金
	AntCheck   float32         `json:"antCheck" bson:"antCheck"`     // 花呗
	WhiteBar   float32         `json:"whiteBar" bson:"whiteBar"`     // 白条
	Partners   []bson.ObjectId `json:"partners" bson:"partners"`     // 小伙伴
}

// UserRequest 用户信息请求参数
type UserRequest struct {
	User
	IDS      []bson.ObjectId
	Content  string `json:"content"`
	Remember bool   `json:"remember"`
}

// UserResponse 用户信息的相应模型
type UserResponse struct {
	*User
	Consumes        []Consume `json:"consumes"`
	Channels        []Channel `json:"channels"`
	HaveBeenUsed    float32   `json:"haveBeenUsed"`
	HaveBeenAdvance float32   `json:"haveBeenAdvance"`
	Partners        []*User   `json:"partners"`
}

var table = "user"

// Get 获取用户数据
func (u *UserRequest) Get(search map[string]interface{}) (result []*User) {
	data.Find(table, search, result)
	return
}

// GetResponse 获取响应数据
func (u *UserRequest) GetResponse(search map[string]interface{}) (result UserResponse) {
	users := u.Get(search)
	if library.IsEmpty(users) {
		beego.Info("不存在的用户数据")
		return
	}
	user := users[0]
	user.Password = "*"
	result.User = user
	searchP := bson.M{"_id": bson.M{"$in": user.Partners}}
	result.Partners = u.Get(searchP)
	return
}

// Set 用户数据设置
func (u *UserRequest) Set(update map[string]interface{}, selector interface{}) {
	if library.IsEmpty(update["utime"]) {
		update["utime"] = time.Now()
	}
	data.Update(table, update, selector)
}

// Add 添加用户
func (u *UserRequest) Add() {
	u.User.CreateTime = time.Now()
	u.User.UpdateTime = time.Now()

	data.Insert(table, u.User)
}

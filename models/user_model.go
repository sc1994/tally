package models

import (
	"encoding/json"
	"tally/data"
	"tally/library"
	"time"

	"github.com/astaxie/beego"

	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// User 用户信息实体
type User struct {
	ID         bson.ObjectId   `json:"id" bson:"_id"`                // ID
	CreateTime time.Time       `json:"ctime" bson:"ctime"`           // CreateTime 创建时间
	UpdateTime time.Time       `json:"utime" bson:"utime"`           // UpdateTime 更新时间
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
	Consumes        []*Consume `json:"consumes"`
	Channels        []*Channel `json:"channels"`
	HaveBeenUsed    float32    `json:"haveBeenUsed"`
	HaveBeenAdvance float32    `json:"haveBeenAdvance"`
	Partners        []*User    `json:"partners"`
}

const userTable string = "user"

// Get Get
func (u *UserRequest) Get(search map[string]interface{}) (result []*User) {
	data.Find(userTable, search, &result)
	for _, i := range result {
		i.Password = "*"
	}
	return
}

// GetResponse 获取响应数据
func (u *UserRequest) GetResponse(search map[string]interface{}) (result UserResponse, b bool) {
	b = false
	users := u.Get(search)
	if library.IsEmpty(users) {
		beego.Info("不存在的用户数据")
		return
	}
	user := users[0]
	result.User = user
	searchP := bson.M{"_id": bson.M{"$in": user.Partners}}
	result.Partners = u.Get(searchP)
	result.Consumes = new(ConsumeRequest).Get(bson.M{"uid": result.User.ID})
	result.Channels = new(ChannelRequest).Get(bson.M{"uid": result.User.ID})
	b = true
	return
}

// Set Set
func (u *UserRequest) Set(update map[string]interface{}, selector map[string]interface{}) *mgo.ChangeInfo {
	if library.IsEmpty(update["utime"]) {
		update["utime"] = time.Now()
	}
	return data.Update(userTable, update, selector)
}

// Add Add
func (u *UserRequest) Add() bson.ObjectId {
	u.User.CreateTime = time.Now()
	u.User.UpdateTime = time.Now()
	u.User.ID = bson.NewObjectId()
	u.User.NickName = u.Name
	u.User.HeadImg = "/static/images/head-default.png"
	u.User.Budget = 1000
	u.User.Intro = "这家伙很懒,啥都没说"
	data.Insert(userTable, u.User)
	return u.User.ID
}

// Delete Delete
func (u *UserRequest) Delete(selector map[string]interface{}) {
	panic("no realize")
}

// GetUserToken 获取用户token
func GetUserToken(id bson.ObjectId, name string, pwd string) string {
	flag := id.Hex() + name + pwd
	head := string([]byte(library.Md5String(flag)[:8]))
	food := string([]byte(library.Md5String(library.GetString(time.Now().Unix()))[:8]))
	return head + food
}

// RefreshUserRedis 刷新用户缓存
func RefreshUserRedis(user UserResponse) string {
	token := GetUserToken(user.ID, user.Name, user.Password)
	flag := string([]byte(token)[:8])
	keys := library.GetRedisKeys(flag + "*")
	for _, k := range keys {
		library.DelRedis(k)
	}
	j, _ := json.Marshal(user)
	library.SetRedis(token, j, 7*24*60)
	return token
}

// ChangeUserMoney 变更用户金额
func (u *User) ChangeUserMoney(uid bson.ObjectId, mode string, channel string, money float32) bool {
	var updateField string
	var updateValue float32
	switch channel {
	case "支付宝":
		if mode == "收入" {
			u.Alipay += money
			updateValue = money
		} else if mode == "支出" {
			u.Alipay -= money
			updateValue = -money
		}
		updateField = "aliPay"
	case "微信":
		if mode == "收入" {
			u.WechatPay += money
			updateValue = money
		} else if mode == "支出" {
			u.WechatPay -= money
			updateValue = -money
		}
		updateField = "wechatPay"
	case "银行卡":
		if mode == "收入" {
			u.BackCard += money
			updateValue = money
		} else if mode == "支出" {
			u.BackCard -= money
			updateValue = -money
		}
		updateField = "backCard"
	case "现金":
		if mode == "收入" {
			u.Cash += money
			updateValue = money
		} else if mode == "支出" {
			u.Cash -= money
			updateValue = -money
		}
		updateField = "cash"
	case "信用卡":
		u.CreditCard -= money // 信用卡 只有支出,还款在job中进行
		updateValue = -money
		updateField = "creditCard"
	case "花呗":
		u.AntCheck -= money
		updateValue = -money
		updateField = "antCheck"
	case "白条":
		u.WhiteBar -= money
		updateValue = -money
		updateField = "whiteBar"
	}
	update := bson.M{"$inc": bson.M{updateField: updateValue}}
	selector := bson.M{"_id": uid}
	return new(UserRequest).Set(update, selector) != nil
}

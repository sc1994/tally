package model

import (
	"tally/data"
	"time"

	"encoding/json"

	"gopkg.in/mgo.v2/bson"
)

const userDB string = "tally"
const userTable string = "user"

// User 用户信息实体
type User struct {
	Id         bson.ObjectId `json:"id" bson:"_id,omitempty"`      // Id
	Password   string        `json:"pwd" bson:"pwd"`               // Password
	Name       string        `json:"name" bson:"name"`             // Name
	NickName   string        `json:"nick" bson:"nick"`             // NickName 昵称
	CreateTime time.Time     `json:"ctime" bson:"ctime"`           // CreateTime 创建时间
	HeadImg    string        `json:"headImg" bson:"himg"`          // HeadImg 头像
	Budget     float32       `json:"budget" bson:"budget"`         // 月预算
	FixDate    float32       `json:"fixDate" bson:"fixDate"`       // 定期
	WechatPay  float32       `json:"wechatPay" bson:"wechatPay"`   // 微信
	Alipay     float32       `json:"aliPay" bson:"aliPay"`         // 支付宝
	BackCard   float32       `json:"backCard" bson:"backCard"`     // 银行卡
	CreditCard float32       `json:"creditCard" bson:"creditCard"` // 信用卡
	Cash       float32       `json:"cash" bson:"cash"`             // 现金
	AntCheck   float32       `json:"antCheck" bson:"antCheck"`     // 花呗
	WhiteBar   float32       `json:"whiteBar" bson:"whiteBar"`     // 白条
}

// UserRequest 用户信息请求参数
type UserRequest struct {
	Token    string `json:"token"`
	Name     string `json:"name"`
	Password string `json:"password"`
	Content  string `json:"content"`
	Remember bool   `json:"remember"`
}

type UserResponse struct {
	User
	Consumes []Consume `json:"consumes"`
	Channels []Channel `json:"channels"`
}

// FindOne 获取一条用户数据
func (u *User) FindOneUser(name string, pwd string) bool {
	search := bson.M{"name": name}
	if len(pwd) > 0 {
		search["pwd"] = pwd
	}
	r := FindUser(search)
	if len(r) > 0 {
		*u = *(r[0])
		return true
	}
	return false
}

// Find 获取满足条件的用户数据
func FindUser(search interface{}) (result []*User) {
	data.Find(userDB, userTable, search, &result)
	return
}

// Insert 新增一条用户信息
func (u *User) InsertUser() bool {
	u.CreateTime = time.Now()
	u.Id = bson.NewObjectId()
	return data.Insert(userDB, userTable, u)
}

// Update 更新一条用户信息
func UpdateUser(name string, update interface{}) bool {
	b := data.Update(userDB, userTable, bson.M{"name": name}, update)
	return b
}

// GetUserByToken 依据token获取用户全部信息
func (u *UserResponse) GetUserByToken(token string) bool {
	s, b := data.GetRedis(token)
	if b {
		e := json.Unmarshal([]byte(s), u)
		if e != nil {
			return false
		}
		return true
	}
	return false
}

// RemoveUserToken 移除token
func RemoveUserToken(token string) {
	data.DelRedis(token)
}

// 初始化用户
func InitUser(name string, pwd string) (bool, User) {
	u := User{
		Name:     name,
		Password: pwd,
		NickName: "用户: " + name,
		HeadImg:  "/static/images/head-default.png",
	}
	b := u.InsertUser()
	return b, u
}

// RefreshUserRedis 刷新用户缓存
func RefreshUserRedis(token string) (bool, string) {
	u := User{}
	ur := UserResponse{}
	if !ur.GetUserByToken(token) {
		return false, "token不存在"
	}
	if !data.DelRedis(token) {
		return false, "移除token失败"
	}
	if !u.FindOneUser(u.Name, "") {
		return false, "用户信息不存在"
	}
	ur = UserResponse{
		User:     u,
		Consumes: FindConsumeByUserId(u.Id),
		Channels: FindChannelByUserId(u.Id),
	} // 获取完整用户信息
	j, _ := json.Marshal(ur)         // 序列化用户数据
	data.SetRedis(token, j, 7*24*60) // 设置到缓存
	return true, ""
}

func (u *User) ChangeUserMoney(mode string, channel string, money float32) bool {
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
		updateField = "alipay"
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
	return UpdateUser(u.Name, update)
}

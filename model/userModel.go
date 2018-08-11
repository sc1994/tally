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
	Id         bson.ObjectId `json:"id" bson:"_id,omitempty"`    // Id
	Password   string        `json:"pwd" bson:"pwd"`             // Password
	Name       string        `json:"name" bson:"name"`           // Name
	NickName   string        `json:"nick" bson:"nick"`           // NickName 昵称
	CreateTime time.Time     `json:"ctime" bson:"ctime"`         // CreateTime 创建时间
	HeadImg    string        `json:"headImg" bson:"himg"`        // HeadImg 头像
	Budget     float32       `json:"budget" bson:"budget"`       // 月预算
	FixDate    float32       `json:"fixDate" bson:"fixDate"`     // 定期
	WechatPay  float32       `json:"wechatPay" bson:"wechatPay"` // 微信
	Alipay     float32       `json:"aliPay" bson:"aliPay"`       // 支付宝
	BackCard   float32       `json:"backCard" bson:"backCard"`   // 银行卡
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
func (u *User) FindOneUser(search interface{}) {
	r := new(User).FindUser(search)
	if len(r) > 0 {
		*u = *(r[0])
	}
}

// Find 获取满足条件的用户数据
func (u *User) FindUser(search interface{}) (result []*User) {
	data.Find(userDB, userTable, search, &result)
	return
}

// Insert 新增一条用户信息
func (u *User) InsertUser() bool {
	return data.Insert(userDB, userTable, u)
}

// Update 更新一条用户信息
func (u *User) UpdateUser(token string, selector interface{}, update interface{}) bool {
	b := data.Update(userDB, userTable, selector, update)
	u.FindOneUser(selector)
	if data.DelRedis(token) {
		j, _ := json.Marshal(u)
		data.SetRedis(token, j, 7*24*60)
	}
	return b
}

// GetUserByToken 依据token获取用户信息
func (u *User) GetUserByToken(token string) bool {
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

// RemoveToken 移除token
func (u *User) RemoveToken(token string) {
	data.DelRedis(token)
}

func (u *User) InitUser(name string, pwd string) bool {
	if len(name) <= 0 || len(pwd) <= 0 {
		return false
	}
	*u = User{
		Name:     name,
		Password: pwd,
		NickName: "用户: " + name,
		HeadImg:  "",
	}
	return true
}

/*
Consumes: []model.ManyType{
			model.ManyType{
				Content: "吃饭",
				Count:   0,
				Default: []string{"支出", "预支"},
			},
			model.ManyType{
				Content: "房租",
				Count:   0,
				Default: []string{"支出", "预支"},
			},
			model.ManyType{
				Content: "工资",
				Count:   0,
				Default: []string{"收入"},
			},
		},
		Channels: []model.ManyType{
			model.ManyType{
				Content: "银行卡",
				Count:   0,
				Default: []string{"收入", "支出"},
			},
			model.ManyType{
				Content: "信用卡",
				Count:   0,
				Default: []string{"预支"},
			},
			model.ManyType{
				Content: "支付宝",
				Count:   0,
				Default: []string{"收入", "支出"},
			},
			model.ManyType{
				Content: "花呗",
				Count:   0,
				Default: []string{"预支"},
			},
			model.ManyType{
				Content: "微信",
				Count:   0,
				Default: []string{"收入", "支出"},
			},
			model.ManyType{
				Content: "现金",
				Count:   0,
				Default: []string{"收入", "支出"},
			},
		},
*/

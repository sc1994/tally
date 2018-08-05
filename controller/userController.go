package controller

import (
	"encoding/json"
	"sort"
	"tally/common"
	"tally/data"
	"tally/model"
	"time"

	"github.com/gin-gonic/gin"
	"gopkg.in/mgo.v2/bson"
)

// GetUser 使用Token获取用户信息
func GetUser(service string, r *gin.Engine) {
	r.GET(service, func(c *gin.Context) {
		user := model.User{}
		b := user.Init(c.Param("token"))
		if b {
			sort.Sort(model.ByCount(user.Consumes))
		}
		c.JSON(200, gin.H{
			"result": b,
			"user":   user,
		})
	})
}

// RemoveToken 移除Token，用于注销用户登录
func RemoveToken(service string, r *gin.Engine) {
	r.GET(service, func(c *gin.Context) {
		new(model.User).RemoveToken(c.Param("token"))
	})
}

// GetExist 获取用户名是否存在
func GetExist(service string, r *gin.Engine) {
	r.GET(service, func(c *gin.Context) {
		_, _, t := getOneUser(c.Param("name"), "", 0)
		c.JSON(200, gin.H{
			"exist": t == 2,
		})
	})
}

// InsertUser 用户注册
func InsertUser(service string, r *gin.Engine) {
	r.POST(service, func(c *gin.Context) {
		request := new(model.UserRequest)
		common.BindExtend(c, request)
		_, _, t := getOneUser(request.Name, request.Password, 0)
		var b bool
		u := model.User{
			Name:       request.Name,
			Password:   request.Password,
			CreateTime: time.Now(),
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
			Modes: []model.ManyType{
				model.ManyType{
					Content: "收入",
					Count:   0,
				}, model.ManyType{
					Content: "支出",
					Count:   0,
				}, model.ManyType{
					Content: "预支",
					Count:   0,
				},
			},
		}
		if t == 2 || t == 3 {
			b = false
		} else {
			b = u.Insert()
		}
		c.JSON(200, gin.H{
			"result": b,
		})
	})
}

// GetOneUser 获取用户信息用于登陆
func GetOneUser(service string, r *gin.Engine) {
	r.POST(service, func(c *gin.Context) {
		request := new(model.UserRequest)
		common.BindExtend(c, request)
		var hour int32
		if request.Remember {
			hour = 7 * 24
		} else {
			hour = 2
		}
		u, m, t := getOneUser(request.Name, request.Password, hour)
		if t == 1 || t == 2 {
			c.JSON(200, gin.H{
				"msg": m,
			})
		} else {
			c.JSON(200, gin.H{
				"token": m,
				"user":  u,
			})
		}
	})
}

// InstertConsume 新增消费类型
func InstertConsume(service string, r *gin.Engine) {
	r.POST(service, func(c *gin.Context) {
		request := new(model.UserRequest)
		common.BindExtend(c, request)
		e, m, u := existConsume(request.Content, request.Token)
		var result bool
		var msg string
		if e {
			result = false
			msg = m
		} else {
			newConsume := model.ManyType{
				Content: request.Content,
				Count:   0,
				Default: []string{"收入", "支出", "预支"},
			}
			r := u.Update(
				request.Token,
				bson.M{"_id": u.Id},
				bson.M{"$push": bson.M{"consumes": newConsume}})
			result = r
			msg = ""
		}
		c.JSON(200, gin.H{
			"result": result,
			"msg":    msg,
		})
	})
}

// RemoveConsume 移除消费类型
func RemoveConsume(service string, r *gin.Engine) {
	r.POST(service, func(c *gin.Context) {
		request := new(model.UserRequest)
		common.BindExtend(c, request)
		r, m, u := existConsume(request.Content, request.Token)
		if !r {
			c.JSON(200, gin.H{
				"result": r,
				"msg":    m,
			})
			return
		}
		var consume model.ManyType
		for _, v := range u.Consumes {
			if v.Content == request.Content {
				consume = v
				break
			}
		}
		r = u.Update(
			request.Token,
			bson.M{"_id": u.Id},
			bson.M{"$pull": bson.M{"consumes": consume}})
		c.JSON(200, gin.H{
			"result": r,
			"msg":    "",
		})
	})

}

// IncConsumeCount 消费类型使用次数加1
func IncConsumeCount(token string, content string) bool {
	user := model.User{}
	b := user.Init(token)
	consume := model.ManyType{}
	for _, v := range user.Consumes {
		if v.Content == content {
			consume = v
			break
		}
	}
	if !b || consume.Content != content {
		b = user.Update(
			token,
			bson.M{"_id": user.Id},
			bson.M{"$push": bson.M{"consumes": model.ManyType{Content: content, Count: 1}}})
		return b
	}
	b = user.Update(token,
		bson.M{"_id": user.Id, "consumes": bson.M{"$elemMatch": bson.M{"content": consume.Content}}},
		bson.M{"$set": bson.M{"consumes.$.count": consume.Count + 1}})
	return b
}

// getOneUser 获取一条用户信息
func getOneUser(name string, password string, hour int32) (model.User, string, int16) {
	user := model.User{}
	user.FindOne(bson.M{"name": name})
	if user.Name != name {
		return model.User{}, "用户名称不存在", 1
	}
	if user.Password != password {
		return user, "密码不正确", 2
	}
	user.Password = "*"
	var k bson.ObjectId
	if hour > 0 {
		k = bson.NewObjectId()
		user.Id = k
		j, _ := json.Marshal(user)
		data.RedisSet(k.Hex(), j, hour*60)
	}
	return user, k.Hex(), 3
}

// existConsume 是否存在的消费类型
func existConsume(content string, token string) (bool, string, model.User) {
	user := model.User{}
	user.Init(token)
	if len(user.Name) <= 0 {
		return true, "token已失效", user
	}
	found := -1
	for i, v := range user.Consumes {
		if v.Content == content {
			found = i
			break
		}
	}
	if found > -1 {
		return true, "已存在的消费类型", user
	}
	return false, "不存在的消费类型", user
}

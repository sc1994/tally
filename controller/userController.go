package controller

import (
	"encoding/json"
	"tally/common"
	"tally/data"
	"tally/model"

	"github.com/gin-gonic/gin"
	"gopkg.in/mgo.v2/bson"
)

// GetUser 使用Token获取用户信息
func GetUser(c *gin.Context) {
	user := model.User{}
	b := user.GetUserByToken(c.Param("token"))
	if b {
		consumes := model.FindConsumeByUserId(user.Id)
		channels := model.FindChannelByUserId(user.Id)
		response := model.UserResponse{
			User:     user,
			Consumes: consumes,
			Channels: channels,
		}
		c.JSON(200, gin.H{
			"result": b,
			"user":   response,
		})
	} else {
		c.JSON(200, gin.H{
			"result": b,
		})
	}

}

// RemoveToken 移除Token，用于注销用户登录
func RemoveToken(c *gin.Context) {
	new(model.User).RemoveToken(c.Param("token"))
}

// ExistUser 获取用户名是否存在
func ExistUser(c *gin.Context) {
	_, _, t := getOneUser(c.Param("name"), "", 0)
	c.JSON(200, gin.H{
		"exist": t == 2,
	})
}

// InsertUser 用户注册
func InsertUser(c *gin.Context) {
	request := new(model.UserRequest)
	common.BindExtend(c, request)
	u, _, t := getOneUser(request.Name, request.Password, 0)
	var b bool

	if t == 2 || t == 3 {
		b = false
	} else {
		b = u.InitUser(request.Name, request.Password)
		if b {
			b = u.InsertUser()
		}
	}
	c.JSON(200, gin.H{
		"result": b,
	})
}

// FindOneUser 获取用户信息用于登陆
func FindOneUser(c *gin.Context) {
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
}

// getOneUser 获取一条用户信息
func getOneUser(name string, password string, hour int32) (model.User, string, int16) {
	user := model.User{}
	user.FindOneUser(bson.M{"name": name})
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
		j, _ := json.Marshal(user)
		data.SetRedis(k.Hex(), j, hour*60)
	}
	return user, k.Hex(), 3
}

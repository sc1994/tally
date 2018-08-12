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
	user := model.UserResponse{}
	b := user.GetUserByToken(c.Param("token"))
	if b {
		c.JSON(200, gin.H{
			"result": b,
			"user":   user,
		})
	} else {
		c.JSON(200, gin.H{
			"result": b,
		})
	}

}

// RemoveToken 移除Token，用于注销用户登录
func RemoveToken(c *gin.Context) {
	model.RemoveUserToken(c.Param("token"))
}

// ExistUser 获取用户名是否存在
func ExistUser(c *gin.Context) {
	e := new(model.User).FindOneUser(c.Param("name"), "")
	c.JSON(200, gin.H{
		"exist": e,
	})
}

// InsertUser 用户注册
func InsertUser(c *gin.Context) {
	request := new(model.UserRequest)
	common.BindExtend(c, request)
	e := new(model.User).FindOneUser(request.Name, "")
	if e {
		c.JSON(200, gin.H{
			"result": false,
		})
	} else {
		b, u := model.InitUser(request.Name, request.Password) // 添加用户
		if b {
			model.InitChannel(u.Id) // 添加消费渠道
			model.InitConsume(u.Id) // 添加消费类型
		}
		c.JSON(200, gin.H{
			"result": b,
		})
	}
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
	u := model.User{}
	e := u.FindOneUser(request.Name, request.Password)
	if e {
		ur := model.UserResponse{
			User:     u,
			Consumes: model.FindConsumeByUserId(u.Id),
			Channels: model.FindChannelByUserId(u.Id),
		} // 获取完整用户信息
		k := bson.NewObjectId()            // new token
		j, _ := json.Marshal(ur)           // 序列化用户数据
		data.SetRedis(k.Hex(), j, hour*60) // 设置到缓存
		c.JSON(200, gin.H{
			"result": e,
			"token":  k,
		})
	} else {
		c.JSON(200, gin.H{
			"result": e,
		})
	}
}

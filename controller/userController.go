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
			model.InitChannel(u.ID) // 添加消费渠道
			model.InitConsume(u.ID) // 添加消费类型
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
		ur := u.GetUserResponse()
		j, _ := json.Marshal(ur)               // 序列化用户数据
		data.SetRedis(ur.ID.Hex(), j, hour*60) // 设置到缓存
		c.JSON(200, gin.H{
			"result": e,
			"token":  ur.ID.Hex(),
		})
	} else {
		c.JSON(200, gin.H{
			"result": e,
		})
	}
}

// SetUserBaseInfo 设置用户基本信息
func SetUserBaseInfo(c *gin.Context) {
	request := new(model.UserRequest)
	common.BindExtend(c, request)
	updater := bson.M{
		"$set": bson.M{
			"nick":      request.NickName,
			"budget":    request.Budget,
			"fixDate":   request.FixDate,
			"wechatPay": request.WechatPay,
			"aliPay":    request.Alipay,
			"backCard":  request.BackCard,
			"cash":      request.Cash,
		},
	}
	b := model.UpdateUserByID(request.ID, updater)
	if b {
		model.RefreshUserRedis(request.Token)
	}
	c.JSON(200, gin.H{
		"result": b,
	})
}

// FindUsersByName 依据昵称搜索用户
func FindUsersByName(c *gin.Context) {
	search := bson.M{"name": bson.M{"$regex": c.Param("name")}}
	result := model.FindUser(search)
	c.JSON(200, gin.H{
		"result": result,
	})
}

// SetUserHeadImage 设置头像
func SetUserHeadImage(c *gin.Context) {
	var request model.UserRequest
	common.BindExtend(c, &request)
	update := bson.M{"$set": bson.M{"himg": request.HeadImg}}
	r := model.UpdateUserByID(request.ID, update)
	if r {
		go model.RefreshUserRedis(request.ID.Hex())
	}
	c.JSON(200, gin.H{
		"result": r,
	})
}

package controller

import (
	"tally/common"
	"tally/model"

	"gopkg.in/mgo.v2/bson"

	"github.com/gin-gonic/gin"
)

// InsertTally 新增一条记录
func InsertTally(c *gin.Context) {
	request := new(model.TallyRequest)
	common.BindExtend(c, request)
	u := model.UserResponse{}
	b := u.GetUserByToken(request.Token)
	if !b {
		c.JSON(200, gin.H{
			"result": b,
		})
		return
	}
	t := model.Tally{
		UserId:     u.Id,
		Money:      request.Money,
		Type:       request.Type,
		Mode:       request.Mode,
		Channel:    request.Channel,
		CreateTime: request.CreateTime,
		Remark:     request.Remark,
	}
	b = t.InsertTally()
	if b {
		// b = IncConsumeCount(request.Token, request.Type)
	}
	c.JSON(200, gin.H{
		"result": b,
	})
}

// GetTallyByUser 获取消费记录
func GetTallyByUser(c *gin.Context) {
	request := new(model.TallyRequest)
	common.BindExtend(c, request)
	u := model.UserResponse{}
	b := u.GetUserByToken(request.Token)
	if !b {
		c.JSON(200, gin.H{
			"result": b,
			"msg":    "用户信息不存在",
		})
		return
	}
	t := model.Tally{}
	var result []model.Tally
	b = t.FindTally(
		request.PageIndex,
		request.PageSize,
		bson.M{"uid": u.Id},
		&result)
	if !b {
		c.JSON(200, gin.H{
			"result": b,
			"msg":    "查询异常",
		})
		return
	}
	c.JSON(200, gin.H{
		"result": b,
		"msg":    "查询成功",
		"body":   result,
	})
	return
}

// GetTest 函数封装方式
func GetTest(c *gin.Context) {
	c.JSON(200, gin.H{
		"msg": "查询成功",
	})
}

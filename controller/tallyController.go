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
	// 获取用户信息
	b := u.GetUserByToken(request.Token)
	if !b {
		c.JSON(200, gin.H{
			"result": b,
		})
		return
	}
	// 检查消费类型是否存在
	e := model.ExistConsume(u.Id, request.Type)
	if !e {
		// 新增消费类型
		consume := model.Consume{
			Content: request.Type,
			Default: []string{common.TallyMode[0], common.TallyMode[1], common.TallyMode[2]},
			Count:   1,
			UserId:  u.Id,
		}
		consume.InsertConsume()
	} else {
		model.IncConsumeCount(u.Id, request.Type)
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
		if request.Channel == "" {

		}
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

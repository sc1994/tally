package controller

import (
	"tally/common"
	"tally/model"

	"gopkg.in/mgo.v2/bson"

	linq "github.com/ahmetb/go-linq"
	"github.com/gin-gonic/gin"
)

// GetMessage 获取消息
func GetMessage(c *gin.Context) {
	uid := c.Param("uid")
	var r []model.Message
	model.FindMessageByMe(bson.ObjectIdHex(uid), &r)
	linq.From(r).OrderByDescending(
		func(x interface{}) interface{} {
			return x.(model.Message).CreateTime.String()
		},
	).ToSlice(&r)
	c.JSON(200, gin.H{
		"result": &r,
	})
}

// SendMessage 添加消息
func SendMessage(c *gin.Context) {
	var request model.MessageRequest
	common.BindExtend(c, &request)
	m := model.Message{
		FromID:    request.FromID,
		FromNick:  request.FromNick,
		FromImg:   request.FromImg,
		ToID:      request.ToID,
		ToNick:    request.ToNick,
		ToImg:     request.ToImg,
		Content:   request.Content,
		NeedTouch: request.NeedTouch,
		Type:      request.Type,
	}
	b := m.InsertMessage()
	c.JSON(200, gin.H{
		"result": b,
	})
}

// AgreeMessage 同意
func AgreeMessage(c *gin.Context) {
	var request model.MessageRequest
	common.BindExtend(c, &request)
	b := model.UpdateMessageStatus(request.ID, 3)
	if b {
		switch request.Type {
		case 1:
			// todo 添加绑定关系
		}
	}
	c.JSON(200, gin.H{
		"result": b,
	})
}

// RejectMessage 拒绝
func RejectMessage(c *gin.Context) {
	var request model.MessageRequest
	common.BindExtend(c, &request)
	m, b := model.FindMessageByID(request.ID)
	if !b {
		c.JSON(200, gin.H{
			"result": b,
			"msg":    "消息不存在",
		})
		return
	}
	b = model.UpdateMessageStatus(request.ID, 4)
	if b {
		switch request.Type {
		case 1:
			// 回执消息
			receipt := model.Message{
				FromID:  m.ToID,
				ToID:    m.FromID,
				Content: "伙伴邀请被拒绝", // todo
			}
			receipt.InsertMessage()
		}
	}
	c.JSON(200, gin.H{
		"result": b,
	})
}

// ReadMessage 阅读消息
func ReadMessage(c *gin.Context) {

}

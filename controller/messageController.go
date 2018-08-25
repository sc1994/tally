package controller

import (
	"strconv"
	"tally/common"
	"tally/model"

	"gopkg.in/mgo.v2/bson"

	linq "github.com/ahmetb/go-linq"
	"github.com/gin-gonic/gin"
)

// GetMessages 获取消息
func GetMessages(c *gin.Context) {
	uid := c.Param("uid")
	size, _ := strconv.Atoi(c.Param("size"))
	index, _ := strconv.Atoi(c.Param("index"))
	var r []model.Message
	count := model.FindMessageByMe(bson.ObjectIdHex(uid), index, size, &r)
	linq.From(r).OrderByDescending(
		func(x interface{}) interface{} {
			return x.(model.Message).CreateTime.String()
		},
	).ToSlice(&r)
	c.JSON(200, gin.H{
		"result": &r,
		"count":  count,
	})
}

// GetMessageUnreadCount 获取未读的消息数量
func GetMessageUnreadCount(c *gin.Context) {
	uid := c.Param("uid")
	i := model.FindMessageCountByStatus(bson.ObjectIdHex(uid), 1)
	c.JSON(200, gin.H{
		"result": i,
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

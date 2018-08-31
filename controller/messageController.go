package controller

import (
	"strconv"
	"tally/common"
	"tally/model"

	"gopkg.in/mgo.v2/bson"

	"github.com/gin-gonic/gin"
)

// GetMessages 获取消息
func GetMessages(c *gin.Context) {
	uid := c.Param("uid")
	size, _ := strconv.Atoi(c.Param("size"))
	index, _ := strconv.Atoi(c.Param("index"))
	count, r := model.FindMessageByMe(bson.ObjectIdHex(uid), index, size)
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
		ToID:      request.ToID,
		Content:   request.Content,
		NeedTouch: request.NeedTouch,
		Type:      request.Type,
	}
	b := m.InsertMessage()
	c.JSON(200, gin.H{
		"result": b,
		"msg":    "发送成功",
	})
}

// AgreeMessage 同意
func AgreeMessage(c *gin.Context) {
	var request model.MessageRequest
	common.BindExtend(c, &request)
	b := model.UpdateMessageStatus(request.ID, 3)
	tu, _ := model.GetUserByID(request.ToID)
	fu, _ := model.GetUserByID(request.FromID)
	if b {
		switch request.Type {
		case 1: // 添加互相之间的绑定关系
			go func() {
				tu.AddUserPartner(request.FromID)
				model.RefreshUserRedis(request.FromID.Hex())
				m := model.Message{
					FromID:    bson.ObjectIdHex(common.AdminID),
					ToID:      request.FromID,
					Content:   tu.NickName + "同意了你的邀请",
					Type:      2,
					Status:    1,
					NeedTouch: false,
				}
				m.InsertMessage()
			}()
			go func() {
				fu.AddUserPartner(request.ToID)
				model.RefreshUserRedis(request.ToID.Hex())
			}()
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
				FromID:  bson.ObjectIdHex(common.AdminID),
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

// ReadMessageAll 阅读消息
func ReadMessageAll(c *gin.Context) {

}

// ReadMessageAll 阅读消息
// func ReadMessageAll(c *gin.Context) {
// 	// c.Param("id")

// }

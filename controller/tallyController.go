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
	t := model.Tally{
		UserID:     u.ID,
		Money:      request.Money,
		Type:       request.Type,
		Mode:       request.Mode,
		Channel:    request.Channel,
		CreateTime: request.CreateTime,
		Remark:     request.Remark,
	}
	b = t.InsertTally()
	if b {
		// 检查消费类型是否存在
		e := model.ExistConsume(u.ID, request.Type)
		if !e {
			// 新增消费类型
			consume := model.Consume{
				Content: request.Type,
				Default: []string{common.TallyMode[0], common.TallyMode[1], common.TallyMode[2]},
				Count:   1,
				UserID:  u.ID,
			}
			consume.InsertConsume()
		} else {
			model.IncConsumeCount(u.ID, request.Type)
		}
		b = u.ChangeUserMoney(request.Mode, request.Channel, request.Money)
		model.RefreshUserRedis(request.Token)
		go func() {
			for _, val := range u.Partners {
				m := model.Message{
					FromID:    u.ID,
					ToID:      val.ID,
					Content:   val.NickName + "添加了一笔50元的消费",
					Type:      3,
					NeedTouch: false,
				}
				m.InsertMessage()
			}
		}()
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
	t.FindTallyPage(
		request.PageIndex,
		request.PageSize,
		bson.M{"uid": u.ID},
		&result)
	c.JSON(200, gin.H{
		"result": true,
		"msg":    "查询成功",
		"body":   result,
	})
	return
}

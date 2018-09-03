package controller

import (
	"tally/common"
	"tally/model"

	"github.com/ahmetb/go-linq"
	"github.com/leekchan/accounting"

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
		go func() {
			b = u.ChangeUserMoney(request.Mode, request.Channel, request.Money)
			model.RefreshUserRedis(request.Token)
			ac := accounting.Accounting{Symbol: "", Precision: 2}
			for _, val := range u.Partners {
				m := model.Message{
					FromID:    u.ID,
					ToID:      val.ID,
					Content:   "添加了一笔" + ac.FormatMoney(request.Money) + "元的" + request.Mode,
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
	uids := make([]bson.ObjectId, 0)
	uids = append(uids, u.ID)
	if !request.OnlyMe {
		for _, val := range u.Partners {
			uids = append(uids, val.ID)
		}
	}
	us := model.GetUsersByIDs(uids)
	var result []model.Tally
	t.FindTallyPage(
		request.PageIndex,
		request.PageSize,
		bson.M{"uid": bson.M{"$in": uids}},
		&result)

	response := make([]model.TallyResponse, 0)
	for _, val := range result {
		first := linq.From(us).WhereT(func(x model.User) bool {
			return x.ID == val.UserID
		}).First().(model.User)
		r := model.TallyResponse{
			Tally:      val,
			User:       first,
			TID:        val.ID,
			CreateTime: val.CreateTime,
		}
		response = append(response, r)
	}
	c.JSON(200, gin.H{
		"result": true,
		"msg":    "查询成功",
		"body":   response,
	})
	return
}

// UpdateTallyByID 更新一条消费记录
func UpdateTallyByID(c *gin.Context) {
	var request model.TallyRequest
	common.BindExtend(c, &request)
	t := &model.Tally{
		ID:      request.ID,
		Money:   request.Money,
		Type:    request.Type,
		Mode:    request.Mode,
		Channel: request.Channel,
		Remark:  request.Remark,
	}
	b := t.UpdateTallyByID()
	go model.RefreshUserRedis(request.Token)
	c.JSON(200, gin.H{
		"result": b,
	})
}

// DeleteTallyByID 删除一条消费记录
func DeleteTallyByID(c *gin.Context) {
	id := c.Param("id")
	token := c.Param("token")
	b := model.DeleteTallyByID(id)
	go model.RefreshUserRedis(token)
	c.JSON(200, gin.H{
		"result": b,
	})
}

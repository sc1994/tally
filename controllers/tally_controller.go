package controllers

import (
	"tally/library"
	"tally/models"

	"github.com/leekchan/accounting"
	"gopkg.in/mgo.v2/bson"
)

// TallyController 消费
type TallyController struct {
	BaseController
}

// Add Add
func (c *TallyController) Add() {
	var request models.TallyRequest
	c.RequestObject(&request)
	request.Tally.UserID = CurrentUser.ID
	id := request.Add()
	code := 0
	token := ""
	if len(id.Hex()) < 1 {
		code = 1
	} else {
		// go CurrentUser.User.ChangeUserMoney(CurrentUser.ID, request.Mode, request.Channel, request.Money)
		consumes := new(models.ConsumeRequest).Get(bson.M{"uid": CurrentUser.ID, "content": request.Type})
		if len(consumes) > 0 {
			c := consumes[0]
			go new(models.ConsumeRequest).Set(bson.M{"$inc": bson.M{"count": 1}}, bson.M{"_id": c.ID})
		} else {
			models.AddConsume(&models.Consume{
				Content: request.Type,
				UserID:  CurrentUser.ID,
				Count:   1,
				Default: []string{library.TallyMode[0], library.TallyMode[1], library.TallyMode[2]},
			})
		}
		token = models.RefreshUserRedis(CurrentUser)
		go func() {
			ac := accounting.Accounting{Symbol: "", Precision: 2}
			for _, val := range CurrentUser.Partners {
				m := &models.MessageRequest{
					Message: &models.Message{
						FromID:    CurrentUser.ID,
						ToID:      val.ID,
						Content:   "添加了一笔" + ac.FormatMoney(request.Money) + "元的" + request.Mode,
						Type:      3,
						NeedTouch: false,
						Status:    1,
					},
				}
				m.Add()
			}
		}()
	}
	c.ResponseJSON(models.BaseResponse{
		Code: code,
		Data: map[string]string{"token": token},
		Msg:  "success",
	})
}

// Get Get
func (c *TallyController) Get() {
	var request models.TallyRequest
	c.RequestObject(&request)
	search := getSearch(request)
	result := request.Page(search)
	for _, v := range result {
		v.CanEdit = CurrentUser.ID == v.Tally.UserID
	}
	c.ResponseJSON(models.BaseResponse{
		Code: 0,
		Data: result,
		Msg:  "success",
	})
}

// Set Set
func (c *TallyController) Set() {
	var request models.TallyRequest
	c.RequestObject(&request)
	if library.IsEmpty(request.ID) {
		c.ResponseJSON(models.BaseResponse{
			Code: 1,
			Msg:  "id is null",
		})
	}
	selector := bson.M{"_id": request.ID}
	update := bson.M{"$set": bson.M{
		"money":   request.Money,
		"type":    request.Type,
		"mode":    request.Mode,
		"channel": request.Channel,
		"remark":  request.Remark,
		"ttime":   request.TallyTime,
	},
	}
	result := request.Set(update, selector)
	c.ResponseJSON(models.BaseResponse{
		Code: 0,
		Data: result,
		Msg:  "success",
	})
}

// Delete Delete
func (c *TallyController) Delete() {
	id := c.Ctx.Input.Param(":id")
	if library.IsEmpty(id) {
		c.ResponseJSON(models.BaseResponse{
			Code: 1,
			Msg:  "id is null",
		})
	}
	selector := bson.M{
		"_id": bson.ObjectIdHex(id),
		"uid": CurrentUser.ID,
	}
	new(models.TallyRequest).Delete(selector)
	token := models.RefreshUserRedis(CurrentUser)
	c.ResponseJSON(models.BaseResponse{
		Code: 0,
		Data: map[string]string{"token": token},
		Msg:  "success",
	})
}

// Total 消费统计
func (c *TallyController) Total() {
	var request models.TallyRequest
	c.RequestObject(&request)
	search := bson.M{"$match": getSearch(request)}
	group := bson.M{"$group": bson.M{"_id": nil, "total": bson.M{"$sum": "$money"}}}
	result := request.Pipe(search, group)
	if len(result) > 0 {
		c.ResponseJSON(models.BaseResponse{
			Code: 0,
			Data: result[0]["total"],
			Msg:  "success",
		})
	}
	c.ResponseJSON(models.BaseResponse{
		Code: 0,
		Data: 0,
		Msg:  "success",
	})
}

// GetTypes 获取消费类型
func (c *TallyController) GetTypes() {
	var request models.TallyRequest
	c.RequestObject(&request)
	types := request.Pipe(
		bson.M{
			"$match": bson.M{
				"uid":   bson.M{"$in": request.UserIDs},
				"mode":  bson.M{"$in": request.Modes},
				"ttime": bson.M{"$gte": request.BeginTime, "$lte": request.EndTime},
			},
		},
		bson.M{
			"$group": bson.M{
				"_id":   "$type",
				"count": bson.M{"$sum": 1},
				"utime": bson.M{"$max": "$utime"},
				"ctime": bson.M{"$max": "$ctime"},
			},
		})
	c.ResponseJSON(models.BaseResponse{
		Code: 0,
		Data: types,
		Msg:  "success",
	})
}

// GetAdvance 预支信息
func (c *TallyController) GetAdvance() {
	var request models.TallyRequest
	c.RequestObject(&request)
	result := request.Pipe(
		bson.M{
			"$match": bson.M{
				"uid":     CurrentUser.ID,
				"ttime":   bson.M{"$gte": request.BeginTime, "$lte": request.EndTime},
				"channel": bson.M{"$in": []string{"信用卡", "花呗", "白条"}},
			},
		},
		bson.M{
			"$group": bson.M{
				"_id":   "$channel",
				"money": bson.M{"$sum": "$money"},
			},
		})
	c.ResponseJSON(models.BaseResponse{
		Code: 0,
		Data: result,
		Msg:  "success",
	})
}

func getSearch(request models.TallyRequest) map[string]interface{} {
	search := bson.M{
		"uid":   bson.M{"$in": request.UserIDs},
		"ttime": bson.M{"$gte": request.BeginTime, "$lte": request.EndTime},
		"money": bson.M{"$gte": request.BeginMoney, "$lte": request.EndMoney},
		"type":  bson.M{"$in": request.Types},
	}
	if len(request.Types) > 0 {
		search["type"] = bson.M{"$in": request.Types}
	}
	if len(request.Modes) > 0 {
		search["mode"] = bson.M{"$in": request.Modes}
	}
	if len(request.Channels) > 0 {
		search["channel"] = bson.M{"$in": request.Channels}
	}
	return search
}

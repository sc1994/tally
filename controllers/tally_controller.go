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
	if len(id.Hex()) < 1 {
		code = 1
	} else {
		CurrentUser.User.ChangeUserMoney(CurrentUser.ID, request.Mode, request.Channel, request.Money)
		consumes := new(models.ConsumeRequest).Get(bson.M{"uid": CurrentUser.ID, "content": request.Type})
		if len(consumes) > 0 {
			c := consumes[0]
			new(models.ConsumeRequest).Set(bson.M{"$inc": bson.M{"count": 1}}, bson.M{"_id": c.ID})
		} else {
			models.AddConsume(&models.Consume{
				Content: request.Type,
				UserID:  CurrentUser.ID,
				Count:   1,
				Default: []string{library.TallyMode[0], library.TallyMode[1], library.TallyMode[2]},
			})
		}
		models.RefreshUserRedis(CurrentUser)
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
					},
				}
				m.Add()
			}
		}()
	}
	c.ResponseJSON(models.BaseResponse{
		Code: code,
		Data: id,
		Msg:  "success",
	})
}

// Get Get
func (c *TallyController) Get() {
	var request models.TallyRequest
	c.RequestObject(&request)
	search := bson.M{
		"uid":   bson.M{"$in": request.UserIDs},
		"ttime": bson.M{"$gte": request.BeginTime, "$lte": request.EndTime},
		"money": bson.M{"$gte": request.BeginMoney, "$lte": request.EndMoney},
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
	selector := bson.M{"_id": id}
	new(models.TallyRequest).Delete(selector)
	c.ResponseJSON(models.BaseResponse{
		Code: 0,
		Msg:  "success",
	})
}

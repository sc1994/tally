package controllers

import (
	"tally/models"

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
	ids := models.AddTally(&request.Tally)
	code := 0
	if len(ids) < 1 {
		code = 1
	}
	c.ResponseJSON(models.BaseResponse{
		Code: code,
		Data: ids,
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
	result := models.PageTally(search, request.PageIndex, request.PageSize)
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
	id := c.Ctx.Input.Param("id")
	selector := bson.M{"_id": id}
	new(models.TallyRequest).Delete(selector)
	c.ResponseJSON(models.BaseResponse{
		Code: 0,
		Msg:  "success",
	})
}

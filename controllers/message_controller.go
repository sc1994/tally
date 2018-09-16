package controllers

import (
	"tally/models"

	"gopkg.in/mgo.v2/bson"
)

// MessageController 消息
type MessageController struct {
	BaseController
}

// Get Get
func (c *MessageController) Get() {
	var request models.MessageRequest
	c.RequestObject(&request)
	search := bson.M{
		"tid": CurrentUser.ID,
	}
	if request.Type != 0 {
		search["status"] = request.Status
	}
	if request.Type != 0 {
		search["type"] = request.Type
	}
	result := request.Page(search)
	c.ResponseJSON(models.BaseResponse{
		Code: 0,
		Data: result,
		Msg:  "success",
	})
}

// Add add
func (c *MessageController) Add() {
	var request models.MessageRequest
	c.RequestObject(&request)
	if request.Type == 1 {
		request.Message.NeedTouch = true // 不知为何bool值不能传入 todo
	}
	id := request.Add()
	code := 0
	if len(id.Hex()) < 1 {
		code = 1
	}
	c.ResponseJSON(models.BaseResponse{
		Code: code,
		Data: id,
		Msg:  "success",
	})
}

// GetCount GetCount
func (c *MessageController) GetCount() {
	var request models.MessageRequest
	c.RequestObject(&request)
	search := bson.M{
		"tid":    CurrentUser.ID,
		"status": request.Status,
	}
	result := request.GetCount(search)
	c.ResponseJSON(models.BaseResponse{
		Code: 0,
		Data: result,
		Msg:  "success",
	})
}

// Set Set
func (c *MessageController) Set() {
	var request models.MessageRequest
	c.RequestObject(&request)
	if len(request.IDs) < 1 {
		c.ResponseJSON(models.BaseResponse{
			Code: 1,
			Msg:  "ids is null",
		})
	}
	selector := bson.M{
		"_id": bson.M{"$in": request.IDs},
		"tid": CurrentUser.ID,
	}
	update := bson.M{
		"$set": bson.M{"status": request.Status},
	}
	result := request.Set(update, selector)
	c.ResponseJSON(models.BaseResponse{
		Code: 0,
		Data: result,
		Msg:  "success",
	})
}

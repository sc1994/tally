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
		"tid":    request.ToID,
		"type":   request.Type,
		"status": request.Status,
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
	c.RequestObject(request)
	search := bson.M{
		"tid":    request.ToID,
		"type":   request.Type,
		"status": request.Status,
	}
	result := request.GetCount(search)
	c.ResponseJSON(models.BaseResponse{
		Code: 0,
		Data: result,
		Msg:  "success",
	})
}

// func (c *MessageController)

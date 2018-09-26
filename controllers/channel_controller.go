package controllers

import (
	"tally/models"

	"gopkg.in/mgo.v2/bson"
)

// ChannelController 登陆登出
type ChannelController struct {
	BaseController
}

// Set Set
func (c *ChannelController) Set() {
	var request models.ChannelRequest
	c.RequestObject(&request)
	r := request.Set(
		bson.M{
			"$set": bson.M{
				"content":   request.Content,
				"default":   request.Default,
				"isAdvance": request.IsAdvance,
				"rdate":     request.RefundDate,
			},
		},
		bson.M{
			"_id": request.ID,
			"uid": CurrentUser.ID,
		})
	c.ResponseJSON(models.BaseResponse{
		Code: 0,
		Msg:  "success",
		Data: r,
	})
}

// Add Add
func (c ChannelController) Add() {
	var request models.ChannelRequest
	c.RequestObject(&request)
	request.Channel.UserID = CurrentUser.ID
	r := models.AddChannel(&request.Channel)
	c.ResponseJSON(models.BaseResponse{
		Code: 0,
		Msg:  "success",
		Data: r,
	})
}

package controllers

import (
	"tally/models"

	"gopkg.in/mgo.v2/bson"
)

// UserController 用户
type UserController struct {
	BaseController
}

// Get Get
func (u *UserController) Get() {
	u.ResponseJSON(models.BaseResponse{
		Code: 0,
		Data: CurrentUser,
		Msg:  "success",
	})
}

// Search 用户搜索
func (u *UserController) Search() {
	name := u.Ctx.Input.Param(":name")
	if len(name) < 1 {
		u.ResponseJSON(models.BaseResponse{
			Code: 1,
			Msg:  "name is null",
		})
	}
	result := new(models.UserRequest).Get(bson.M{"name": bson.M{"$regex": name}})
	if len(result) < 1 {
		u.ResponseJSON(models.BaseResponse{
			Code: 1,
			Msg:  "查无结果",
		})
	}
	u.ResponseJSON(models.BaseResponse{
		Code: 0,
		Data: result,
		Msg:  "success",
	})
}

// Set set
func (u *UserController) Set() {
	var request models.UserRequest
	u.RequestObject(&request)
	updater := bson.M{
		"$set": bson.M{
			"himg":   request.HeadImg,
			"nick":   request.NickName,
			"budget": request.Budget,
			// "fixDate":   request.FixDate,
			// "wechatPay": request.WechatPay,
			// "aliPay":    request.Alipay,
			// "backCard":  request.BackCard,
			// "cash":      request.Cash,
			// "utime": request.UpdateTime,
		},
	}
	request.Set(updater, bson.M{"_id": CurrentUser.ID})
	token := models.RefreshUserRedis(CurrentUser)
	u.ResponseJSON(models.BaseResponse{
		Code: 0,
		Data: map[string]string{"token": token},
		Msg:  "success",
	})
}

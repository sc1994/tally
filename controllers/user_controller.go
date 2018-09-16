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

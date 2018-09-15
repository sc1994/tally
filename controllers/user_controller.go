package controllers

import (
	"tally/models"
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

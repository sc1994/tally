package controllers

import (
	"tally/models"

	"gopkg.in/mgo.v2/bson"
)

// LandController 登陆登出
type LandController struct {
	BaseController
}

// Login 登陆
func (c *LandController) Login() {
	request := models.UserRequest{}
	c.RequestObject(&request)
	search := bson.M{"name": request.Name, "pwd": request.Password}
	r := request.GetResponse(search)
	c.ResponseJSON(models.BaseResponse{
		Code: 0,
		Data: r,
		Msg:  "success",
	})
}

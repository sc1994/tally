package controllers

import (
	"tally/models"

	"gopkg.in/mgo.v2/bson"
)

// DatabaseController 登陆登出
type DatabaseController struct {
	BaseController
}

// Copy 数据备份
func (c *DatabaseController) Copy() {
	auth := c.Ctx.Input.Param("auth")
	if auth != "64860BF9-C180-4B61-8D98-689E1EC6E14A" {
		c.ResponseJSON(models.BaseResponse{
			Code: 1,
			Msg:  "auth info error",
		})
	}
	users := new(models.UserRequest).Get(bson.M{"": ""})
}

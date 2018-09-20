package controllers

import "tally/models"

// DatabaseController 数据备份
type DatabaseController struct {
	BaseController
}

// Copy Copy
func (c *DatabaseController) Copy() {
	auth := c.Ctx.Input.Param(":auth")
	if auth != "5b9621bff35d959bf1b9608a" {
		c.ResponseJSON(models.BaseResponse{
			Code: 1,
			Msg:  "auth is error",
		})
	}
	models.CopyUser()
	models.CopyChannel()
	models.CopyConsume()
	models.CopyMessage()
	models.CopyTally()
}

package controllers

import (
	"tally/library"
	"tally/models"

	"gopkg.in/mgo.v2/bson"
)

// LandController 登陆登出
type LandController struct {
	BaseController
}

// Signin 登陆
func (c *LandController) Signin() {
	request := models.UserRequest{}
	c.RequestObject(&request)
	search := bson.M{"name": request.Name, "pwd": request.Password}
	r, b := request.GetResponse(search)
	if b {
		t := models.RefreshUserRedis(r)
		c.ResponseJSON(models.BaseResponse{
			Code: 0,
			Data: map[string]string{"token": t},
			Msg:  "success",
		})
	} else {
		c.ResponseJSON(models.BaseResponse{
			Code: 1,
			Data: nil,
			Msg:  "用户名或者密码错误",
		})
	}
}

// CheckName 用户名验证
func (c *LandController) CheckName() {
	name := c.Ctx.Input.Param(":name")
	if library.IsEmpty(name) {
		c.ResponseJSON(models.BaseResponse{
			Code: 1,
			Msg:  "name is null",
		})
	}
	request := models.UserRequest{}
	r := request.Get(bson.M{"name": name})
	if len(r) > 0 {
		c.ResponseJSON(models.BaseResponse{
			Code: 1,
			Msg:  "用户名已存在",
			Data: true,
		})
	} else {
		c.ResponseJSON(models.BaseResponse{
			Code: 0,
			Msg:  "success",
			Data: false,
		})
	}

}

// Logout 登出
func (c *LandController) Logout() {
	library.DelRedis(c.Ctx.Input.Header(":token"))
	c.ResponseJSON(models.BaseResponse{
		Code: 0,
		Msg:  "success",
		Data: true,
	})
}

// Signup 注册
func (c *LandController) Signup() {
	request := new(models.UserRequest)
	c.RequestObject(request)
	// 验证用户名是否存在
	if len(request.Get(bson.M{"name": request.Name})) > 0 {
		c.ResponseJSON(models.BaseResponse{
			Code: 1,
			Msg:  "已存在的用户名",
			Data: request.Name,
		})
	}
	uid := request.Add()
	if len(uid.Hex()) < 1 {
		c.ResponseJSON(models.BaseResponse{
			Code: 1,
			Msg:  "注册失败",
		})
	}

	models.AddChannel(
		&models.Channel{
			UserID:  uid,
			Content: "余额",
			Default: []string{library.TallyMode[0], library.TallyMode[1]},
		},
		&models.Channel{
			UserID:  uid,
			Content: "花呗",
			Default: []string{library.TallyMode[2]},
		})
	models.AddConsume(
		&models.Consume{
			UserID:  uid,
			Content: "吃饭",
			Default: []string{library.TallyMode[1], library.TallyMode[2]},
		},
		&models.Consume{
			UserID:  uid,
			Content: "房租",
			Default: []string{library.TallyMode[1], library.TallyMode[2]},
		},
		&models.Consume{
			UserID:  uid,
			Content: "工资",
			Default: []string{library.TallyMode[0]},
		})
	c.ResponseJSON(models.BaseResponse{
		Code: 0,
		Data: true,
		Msg:  "success",
	})
}

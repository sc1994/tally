package controllers

import (
	"encoding/json"
	"tally/library"
	"tally/models"
	"time"

	"github.com/astaxie/beego"
)

// Token TODO 存储用户信息
var Token string

// BaseController 控制器基类
type BaseController struct {
	beego.Controller
}

// Prepare 请求开始
func (c *BaseController) Prepare() {
	Token = c.Ctx.Input.Header("token")
	beego.Info("TOKEN --:--> " + Token) // todo 权限验证
}

// ResponseJSON 响应数据
func (c *BaseController) ResponseJSON(b models.BaseResponse) {
	c.Data["json"] = map[string]interface{}{
		"code":  b.Code,
		"msg":   b.Msg,
		"data":  b.Data,
		"stime": time.Now(),
	}
	c.ServeJSON()
	if library.IsEmpty(b.Status) {
		b.Status = "200"
	}
	c.Abort(b.Status)
}

// RequestObject POST请求实体
func (c *BaseController) RequestObject(result interface{}) {
	err := json.Unmarshal(c.Ctx.Input.RequestBody, &result)
	beego.Info(string(c.Ctx.Input.RequestBody))
	if err != nil {
		c.ResponseJSON(models.BaseResponse{
			Code:   1,
			Data:   err,
			Msg:    "请求参数序列化异常",
			Status: "500",
		})
	}
}

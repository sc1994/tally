package controllers

import (
	"encoding/json"
	"strings"
	"tally/library"
	"tally/models"
	"time"

	"github.com/astaxie/beego"
)

// CurrentUser 登陆的用户信息
var CurrentUser models.UserResponse

// BaseController 控制器基类
type BaseController struct {
	beego.Controller
}

// Prepare 请求拦截
func (c *BaseController) Prepare() {
	url := c.Ctx.Input.URL()
	if strings.Contains(url, "/signin") ||
		strings.Contains(url, "/signupcheck/") ||
		strings.Contains(url, "/signup") {
		// 屏蔽掉不需要验证的请求
		beego.Info("跳过登陆验证")
	} else {
		token := c.Ctx.Input.Header("token")
		if library.IsEmpty(token) {
			c.ResponseJSON(models.BaseResponse{
				Code: 1,
				Msg:  "token为空",
			})
		}
		j, r := library.GetRedis(token)
		if !r {
			c.ResponseJSON(models.BaseResponse{
				Code: 1,
				Msg:  "登陆信息失效",
			})
		}
		e := json.Unmarshal([]byte(j), &CurrentUser)
		if e != nil {
			c.ResponseJSON(models.BaseResponse{
				Code: 1,
				Msg:  "数据异常",
			})
		}
	}
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
	if err != nil {
		c.ResponseJSON(models.BaseResponse{
			Code:   1,
			Data:   err,
			Msg:    "请求参数序列化异常",
			Status: "500",
		})
	}
}

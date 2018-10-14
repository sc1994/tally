package controllers

import (
	"fmt"
	"tally/library"
	"tally/models"
)

// WechatController 用户
type WechatController struct {
	BaseController
}

var appid = "wx1c1903a0843cf2c7"
var secret = "2987228fedbbf33b8f865262c6562a0c"

// GetUserInfo GetUserInfo
func (w *WechatController) GetUserInfo() {
	code := w.Ctx.Input.Param(":code")
	var url = fmt.Sprint("https://api.weixin.qq.com/sns/jscode2session?appid=&secret=&js_code=${res.code}&grant_type=authorization_code", appid, secret, code)
	result := library.HTTPRequest(url, nil, "")
	w.ResponseJSON(models.BaseResponse{
		Code: 0,
		Data: result,
		Msg:  "success",
	})
}

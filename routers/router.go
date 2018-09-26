package routers

import (
	"tally/controllers"
	"time"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/plugins/cors"
)

func init() {
	beego.InsertFilter("*", beego.BeforeRouter, cors.Allow(&cors.Options{
		AllowAllOrigins: true,
		AllowOrigins:    []string{"*"},
		AllowMethods:    []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:    []string{"Origin", "UserToken", "Authorization", "Access-Control-Allow-Origin", "Access-Control-Allow-Headers", "Content-Type", "token"},
		ExposeHeaders:   []string{"Content-Length", "Access-Control-Allow-Origin", "Access-Control-Allow-Headers", "Content-Type"},
		MaxAge:          2 * time.Minute,
	}))

	// 部分接口需要简化命名
	//登陆=====================================================================================================
	land := &controllers.LandController{}
	beego.Router("/land/signin", land, "post:Signin")         // 登陆
	beego.Router("/signupcheck/:name", land, "get:CheckName") // 验证用户名是否存在
	beego.Router("/land/signout/:token", land, "get:Logout")  // 登出
	beego.Router("/land/signup", land, "post:Signup")
	//用户=====================================================================================================
	user := &controllers.UserController{}
	beego.Router("/user/get", user, "get:Get")             // 获取用户数据
	beego.Router("/user/search/:name", user, "get:Search") // 用户搜索
	beego.Router("/user/set", user, "post:Set")            // 设置
	//账单=====================================================================================================
	tally := &controllers.TallyController{}
	beego.Router("/tally/add", tally, "post:Add")               // 添加账单
	beego.Router("/tally/get", tally, "post:Get")               // 获取基本账单
	beego.Router("/tally/set", tally, "post:Set")               // 更新账单
	beego.Router("/tally/delete/:id", tally, "get:Delete")      // 删除账单
	beego.Router("/tally/total", tally, "post:Total")           // 总金额统计
	beego.Router("/tally/gettypes", tally, "post:GetTypes")     // 账单类型获取
	beego.Router("/tally/getadvance", tally, "post:GetAdvance") // 获取预支数据
	//渠道=====================================================================================================
	channel := &controllers.ChannelController{}
	beego.Router("/channel/set", channel, "post:Set") // 设置渠道
	beego.Router("/channel/add", channel, "post:Add") // 新增渠道
	//消息=====================================================================================================
	msg := &controllers.MessageController{}
	beego.Router("/message/add", msg, "post:Add")                  // 发送消息
	beego.Router("/message/get", msg, "post:Get")                  // 获取消息
	beego.Router("/message/getcount/:status", msg, "get:GetCount") // 获取消息量
	beego.Router("/message/set", msg, "post:Set")                  // 更新消息状态
	//数据备份=================================================================================================
	beego.Router("/database/copy/:auth", &controllers.DatabaseController{}, "get:Copy") // 数据备份
}

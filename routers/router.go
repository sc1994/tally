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
	beego.Router("/land/signin", &controllers.LandController{}, "post:Signin")         // 登陆
	beego.Router("/signupcheck/:name", &controllers.LandController{}, "get:CheckName") // 验证用户名是否存在
	beego.Router("/removetoken", &controllers.LandController{}, "get:Logout")          // 登出
	beego.Router("/land/signup", &controllers.LandController{}, "post:Signup")
	//用户=====================================================================================================
	beego.Router("/user/get", &controllers.UserController{}, "get:Get") // 获取用户数据
	//账单=====================================================================================================
	beego.Router("/tally/add", &controllers.TallyController{}, "post:Add")              // 添加账单
	beego.Router("/tally/get", &controllers.TallyController{}, "post:Get")              // 获取基本账单
	beego.Router("/updatetallybyid", &controllers.TallyController{}, "post:Set")        // 更新账单
	beego.Router("/deletetallybyid/:id", &controllers.TallyController{}, "post:Delete") // 删除账单
	//消息=====================================================================================================
	beego.Router("/sendmessage", &controllers.MessageController{}, "post:Add")           // 发送消息
	beego.Router("/getmessage", &controllers.MessageController{}, "post:Get")            // 获取消息
	beego.Router("/message/getcount", &controllers.MessageController{}, "post:GetCount") // 获取消息量
	beego.Router("/updatemessage", &controllers.MessageController{}, "post:Set")         // 更新消息状态
}

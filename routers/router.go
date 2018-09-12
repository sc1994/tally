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
		AllowHeaders:    []string{"Origin", "UserToken", "Authorization", "Access-Control-Allow-Origin", "Access-Control-Allow-Headers", "Content-Type"},
		ExposeHeaders:   []string{"Content-Length", "Access-Control-Allow-Origin", "Access-Control-Allow-Headers", "Content-Type"},
		MaxAge:          2 * time.Minute,
	}))

	// 部分接口需要简化命名
	//登陆=====================================================================================================
	beego.Router("/signin", &controllers.LandController{}, "post:Login")               // 登陆
	beego.Router("/signupcheck/:name", &controllers.LandController{}, "get:CheckName") // 验证用户名是否存在
	beego.Router("/removetoken", &controllers.LandController{}, "get:Logout")          // 登出
	beego.Router("/signup", &controllers.LandController{}, "post:Logup")               // 注册
	//账单=====================================================================================================
	beego.Router("/inserttally", &controllers.TallyController{}, "post:Add")            // 添加账单
	beego.Router("/gettallybyuser", &controllers.TallyController{}, "post:Get")         // 获取基本账单
	beego.Router("/updatetallybyid", &controllers.TallyController{}, "post:Set")        // 更新账单
	beego.Router("/deletetallybyid/:id", &controllers.TallyController{}, "post:Delete") // 删除账单
	//消息=====================================================================================================

	// r.POST("/sendmessage", controller.SendMessage)                         // 发送消息
	// r.GET("/getmessage/:uid/:index/:size", controller.GetMessages)         // 获取全部消息
	// r.GET("/getmessageunreadcount/:uid", controller.GetMessageUnreadCount) // 获取未读的消息量
	// r.POST("/agreemessage", controller.AgreeMessage)                       // 同意对方的消息
	// r.GET("/readmessageall/:uid", controller.ReadMessageAll)               // 无需阅读的消息全部已读
}

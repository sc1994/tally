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
	//登陆相关=====================================================================================================
	beego.Router("/signin", &controllers.LandController{}, "post:Login")               // 登陆
	beego.Router("/signupcheck/:name", &controllers.LandController{}, "get:CheckName") // 验证用户名是否存在
	beego.Router("/removetoken", &controllers.LandController{}, "get:Logout")          // 登出
	beego.Router("/signup", &controllers.LandController{}, "post:Logup")               // 注册
	//账单相关=====================================================================================================
	beego.Router("/inserttally", &controllers.TallyController{}, "post:Add")    // 添加账单
	beego.Router("/gettallybyuser", &controllers.TallyController{}, "post:Get") // 获取基本账单
}

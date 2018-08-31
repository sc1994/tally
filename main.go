package main

import (
	"tally/controller"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	// // http.HandleFunc("/echo", echo) todo websocker
	// // http.HandleFunc("/", home)
	// log.Fatal(http.ListenAndServe("localhost:8080", nil))

	// api ------------------------------------------------------------------
	r := gin.Default()
	config := cors.DefaultConfig()
	// config.AllowAllOrigins = true
	config.AllowCredentials = true
	config.AllowOrigins = []string{"http://localhost:8080", "http://suncheng.xyz:7777"}
	r.Use(cors.New(config))
	r.Static("static", "static")
	// 用户相关的操作接口======================================================
	r.POST("/signup", controller.InsertUser)                    //  注册
	r.GET("/signupcheck/:name", controller.ExistUser)           //  验证用户名是否存在
	r.GET("/getuser/:token", controller.GetUser)                //  获取用户信息
	r.POST("/signin", controller.FindOneUser)                   //  登陆验证
	r.GET("/removetoken/:token", controller.RemoveToken)        //  移除Token
	r.POST("/setuserbaseinfo", controller.SetUserBaseInfo)      // 设置用户基本信息
	r.GET("/findusersbyname/:name", controller.FindUsersByName) // 查询用户
	r.POST("/setuserheadimage", controller.SetUserHeadImage)    // 设置用户头像
	// 账单相关接口=============================================================
	r.POST("/inserttally", controller.InsertTally)       //   添加消费记录
	r.POST("/gettallybyuser", controller.GetTallyByUser) //  获取当前用户的消费记录
	// 消息相关=================================================================
	r.POST("/sendmessage", controller.SendMessage)                         // 发送消息
	r.GET("/getmessage/:uid/:index/:size", controller.GetMessages)         // 获取全部消息
	r.GET("/getmessageunreadcount/:uid", controller.GetMessageUnreadCount) // 获取未读的消息量
	r.POST("/agreemessage", controller.AgreeMessage)                       // 同意对方的消息

	// 功能或者包的试验---------------------------------------------------------------
	r.GET("/test1", controller.Test1) // ling selectMary

	r.Run(":80")
}

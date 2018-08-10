package main

import (
	"tally/controller"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	r.Use(cors.New(config))
	r.Static("static", "static")
	// 用户相关的操作接口======================================================
	r.POST("/signup", controller.InsertUser)             //  注册
	r.GET("/signupcheck/:name", controller.GetExist)     //  验证用户名是否存在
	r.GET("/getuser/:token", controller.GetUser)         //  获取用户信息
	r.POST("/signin", controller.GetOneUser)             //  登陆验证
	r.POST("/addconsume", controller.InstertConsume)     //  新增消费类型
	r.POST("/removeconsume", controller.RemoveConsume)   //  移除消费类型
	r.GET("/removetoken/:token", controller.RemoveToken) //  移除Token
	// 账单相关接口=============================================================
	r.POST("/inserttally", controller.InsertTally)       //   添加消费记录
	r.POST("/gettallybyuser", controller.GetTallyByUser) //  获取当前用户的消费记录

	r.Run(":80")
}

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
	controller.InsertUser("/signup", r)              // POST 注册
	controller.GetExist("/signupcheck/:name", r)     // GET 验证用户名是否存在
	controller.GetUser("/getuser/:token", r)         // GET 获取用户信息
	controller.GetOneUser("/signin", r)              // POST 登陆验证
	controller.InstertConsume("/addconsume", r)      // POST 新增消费类型
	controller.RemoveConsume("/removeconsume", r)    // POST 移除消费类型
	controller.RemoveToken("/removetoken/:token", r) // GET 移除Token
	// 账单相关接口=============================================================
	controller.InsertTally("/inserttally", r)       // POST  添加消费记录
	controller.GetTallyByUser("/gettallybyuser", r) // POST 获取当前用户的消费记录

	r.Run(":80")
}

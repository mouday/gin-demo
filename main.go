package main

import (
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/thinkerou/favicon"
)

func userMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.Set("userId", "100")
	}
}

func main() {
	// 创建一个服务
	server := gin.Default()

	// set favicon middleware
	server.Use(favicon.New("./favicon.ico"))

	// 全局使用中间件
	server.Use(userMiddleware())

	// 加载模板文件
	server.LoadHTMLGlob("templates/*")

	// 静态资源
	server.Static("/static", "./static")

	// 处理请求
	server.GET("/", func(ctx *gin.Context) {
		// 返回html
		ctx.HTML(
			http.StatusOK,
			"index.html",
			gin.H{
				"name": "Golang",
			})
	})

	// 接收路由参数
	// http://127.0.0.1:8000/user/100
	server.GET("/user/:userId", func(ctx *gin.Context) {

		userId := ctx.Param("userId")

		ctx.JSON(200, gin.H{
			"userId": userId,
		})
	})

	// post请求
	server.POST("/post", func(ctx *gin.Context) {
		// 返回json数据
		ctx.JSON(200, gin.H{"msg": "Hello"})
	})

	// 接收json数据
	// POST http://127.0.0.1:8000/json
	server.POST("/json", func(ctx *gin.Context) {

		// 解析json数据
		rawData, _ := ctx.GetRawData()
		var data gin.H
		json.Unmarshal(rawData, &data)

		// 返回json数据
		ctx.JSON(http.StatusOK, data)
	})

	// 表单页面
	// GET http://127.0.0.1:8000/post-form
	server.GET("/post-form", func(ctx *gin.Context) {
		ctx.HTML(
			http.StatusOK,
			"post-form.html",
			nil,
		)
	})

	// 接收表单数据
	// POST http://127.0.0.1:8000/post-form
	server.POST("/post-form", func(ctx *gin.Context) {
		name := ctx.PostForm("name")
		age := ctx.PostForm("age")

		// 返回json数据
		ctx.JSON(http.StatusOK, gin.H{
			"name": name,
			"age":  age,
		})
	})

	// 使用中间件
	server.GET("/user", userMiddleware(), func(ctx *gin.Context) {

		userId := ctx.GetString("userId")

		ctx.JSON(http.StatusOK, gin.H{
			"userId": userId,
		})
	})

	// 启动服务
	server.Run(":8000")
}

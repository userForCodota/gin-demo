package main

import (
	"github.com/gin-gonic/gin"
	"github.com/thinkerou/favicon"
)

func main() {
	// 创建一个服务
	ginServer := gin.Default()

	// 图标
	ginServer.Use(favicon.New("./favicon.ico"))

	// 注册一个路由
	ginServer.GET("/hello", func(context *gin.Context) {
		context.JSON(200, gin.H{
			"message": "Hello, world!",
		})
	})

	// 注册一个路由:POST
	ginServer.POST("/hello", func(context *gin.Context) {
		context.JSON(200, gin.H{
			"message": "Hello, world!",
		})
	})

	// 注册一个路由:PUT
	ginServer.PUT("/hello", func(context *gin.Context) {
		context.JSON(200, gin.H{
			"message": "Hello, world!",
		})
	})

	// 注册一个路由:DELETE
	ginServer.DELETE("/hello", func(context *gin.Context) {
		context.JSON(200, gin.H{
			"message": "Hello, world!",
		})
	})

	//	服务器监听端口
	ginServer.Run(":8082")

}

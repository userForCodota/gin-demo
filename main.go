package main

import (
	"github.com/gin-gonic/gin"
	"github.com/thinkerou/favicon"
	"net/http"
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

	// 响应一个页面给前端
	ginServer.LoadHTMLGlob("templates/*") // 加载全部模板文件
	// 加载静态文件,第一个参数是前端访问的路径,第二个参数是后端的路径，
	//前端访问/static/xxx,后端就会去./static/xxx找文件
	ginServer.Static("/static", "./static")

	ginServer.GET("/index", func(context *gin.Context) {
		context.HTML(http.StatusOK, "index.html", gin.H{
			"msg": "这是go后台传递来的数据",
		})
	})

	//	服务器监听端口
	ginServer.Run(":8082")

}

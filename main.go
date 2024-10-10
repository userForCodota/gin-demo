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

	// 接收前端传递的参数,场景一:URL参数,如：http://localhost:8082/hello2?name=张三
	ginServer.GET("/hello2", func(context *gin.Context) {
		name := context.Query("name")
		context.JSON(200, gin.H{
			"message": "Hello, " + name,
		})
	})

	// 接收前端传递的参数,场景二:表单参数
	ginServer.POST("/hello3", func(context *gin.Context) {
		name := context.PostForm("name")
		context.JSON(200, gin.H{
			"message": "Hello, " + name,
		})
	})

	// 接收前端传递的参数,场景三:JSON参数
	ginServer.POST("/hello4", func(context *gin.Context) {
		var jsonParam map[string]interface{}
		context.BindJSON(&jsonParam)
		context.JSON(200, gin.H{
			"message": "Hello, " + jsonParam["name"].(string),
		})
	})

	// 接收前端传递的参数,场景四:路径参数
	ginServer.GET("/hello5/:name", func(context *gin.Context) {
		name := context.Param("name")
		context.JSON(200, gin.H{
			"message": "Hello, " + name,
		})
	})

	//	注册一个路由,301,重定向至百度
	ginServer.GET("/baidu", func(context *gin.Context) {
		context.Redirect(http.StatusMovedPermanently, "https://www.baidu.com")
	})

	//	注册一个路由,404,跳自定义的404页面
	ginServer.NoRoute(func(context *gin.Context) {
		context.HTML(http.StatusNotFound, "404.html", gin.H{
			"msg": "404,页面找不到",
		})
	})

	//	服务器监听端口
	ginServer.Run(":8082")

}

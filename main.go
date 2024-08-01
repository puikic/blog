package main

import (
	"blog/handler"
	"blog/util"

	"github.com/gin-gonic/gin"
)

func Home(ctx *gin.Context) {
	ctx.String(200, "你好，世界！")
}
func Init() {
	util.InitLog("log")
}
func main() {
	Init()
	router := gin.Default()
	router.Static("/js", "views/js") //在url是访问目录/js相当于访问文件系统中的views/js目录
	router.GET("/", Home)            //使用这些.html文件时就不需要加路径了
	router.LoadHTMLFiles("views/login.html")
	router.GET("/login", func(ctx *gin.Context) {
		ctx.HTML(200, "login.html", nil)
	})
	router.POST("/login/submit", handler.Login)
	//restful风格,参数放在url路径里
	router.GET("/blog/list/:uid", handler.BlogList)
	router.GET("blog/:bid", handler.BlogDetail)
	router.POST("blog/update", middleware.Auth(), handler.BlogUpdate)

	router.Run("127.0.0.1:5678")
}

// go run ./main.go

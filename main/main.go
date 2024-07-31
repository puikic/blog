package main

import "github.com/gin-gonic/gin"

func Home(ctx *gin.Context) {
	ctx.String(200, "你好，世界！")
}

func Html(ctx *gin.Context) {
	ctx.HTML(200, "home.html", nil)
}
func main() {
	router := gin.Default()
	router.GET("/", Home)
	router.LoadHTMLFiles("views/home.html")
	router.GET("/login", Html)
	router.Run("127.0.0.1:5678")
}

// go run ./main/main.go

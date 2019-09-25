package main

import (
	"github.com/gin-gonic/gin"
	"io"
	"os"
)

func main() {
	f, _ := os.Create("gin.log")
	gin.DefaultWriter = io.MultiWriter(f)
	gin.DefaultErrorWriter = io.MultiWriter(f)

	r := gin.New()
	//Logger的作用就是在控制台打印请求的一系列信息，例如控制台白色字体的信息
	//Recover的作用就是：手动触发一个panic，如果没有Recovery,则服务器崩溃
	r.Use(gin.Logger(), gin.Recovery())
	r.GET("/test", func(c *gin.Context) {
		name := c.DefaultQuery("name", "default_name")
		panic("test panic!")
		c.String(200, name)
	})
	r.Run(":8080")
}

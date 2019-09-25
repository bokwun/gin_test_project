package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()  //创建gin的一个默认实例
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message":"pong",
		})  //H是一个map类型
	})
	r.Run()  //默认8080端口
}

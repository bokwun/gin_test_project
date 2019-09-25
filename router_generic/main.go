package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()
	//所有/user/任何  都执行这个回调函数
	r.GET("/user/*action", func(c *gin.Context) {
		c.String(200, "hello_world")
	})
	r.Run()
}

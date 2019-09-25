package main

import (
	"github.com/gin-gonic/gin"
)
//绑定由form确定，验证由binding确定
type Person struct {
	Age int `form:"age" binding:"required,gt=10"`
	Name string `form:"name" binding:"required"`
	Address string `form:"address" binding:"required"`
}

func main() {
	r := gin.Default()
	r.GET("/testing", func(c *gin.Context) {
		var person Person
		err := c.ShouldBind(&person)
		if err != nil {
			c.String(200, "%v", err)
			c.Abort()
			return
		}
		c.String(200, "%v", person)

	})
	r.Run(":8080")
}

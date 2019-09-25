package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()
	r.GET("/test", func(c *gin.Context) {
		firstname := c.Query("first_name")
		middlename := c.Query("middle_name")
		lastname := c.DefaultQuery("last_name", "last_default_name")
		c.String(http.StatusOK, "%s, %s, %s",firstname, middlename, lastname)
	})
	r.Run()
}

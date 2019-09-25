package main

import (
	"bytes"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
)

func main() {
	r := gin.Default()
	r.POST("/test", func(c *gin.Context) {
		//bodyByts是切片类型
		bodyByts, err := ioutil.ReadAll(c.Request.Body)
		if err != nil {
			//不为空就是有错误，就要打印错误信息
			c.String(http.StatusBadRequest, err.Error())
			c.Abort()
		}
		//c.request.body里面的东西被读取出来了，存在bodyByts切片里
		//需要把他读取回去
		c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(bodyByts))

		firstname := c.PostForm("firstname")
		lastname := c.DefaultPostForm("lastname", "last_default_name")
		c.String(http.StatusOK, "%s, %s, %s",firstname, lastname, string(bodyByts))
	})
	r.Run(":8080")
}

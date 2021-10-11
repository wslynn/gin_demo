package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default() // 返回默认的路由引擎
	r.GET("/hello", sayHello)
	r.GET("/book", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "get",
		})
	})
	r.POST("/book", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "post",
		})
	})

	r.Run(":9090")
}

func sayHello(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Hello golang!",
	})
}

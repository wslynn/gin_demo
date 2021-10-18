package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()
	r.GET("/index", func(c *gin.Context) {
		// 方法1: 使用map
		data := map[string]interface{}{
			"name":    "小王子",
			"message": "hello world",
			"age":     18,
		}
		c.JSON(http.StatusOK, data)
	})
	r.Run(":9090")
}

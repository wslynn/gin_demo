package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()
	r.GET("/index", func(c *gin.Context) {
		// ---- 获取浏览器URL携带的Query string参数
		//name := c.Query("name")
		//name := c.DefaultQuery("name", "啥都没填的默认值")
		name, ok := c.GetQuery("name")
		if !ok {
			name = "啥都没填的默认值"
		}
		// ---- 获取浏览器URL携带的form参数

		c.JSON(http.StatusOK, gin.H{
			"name": name,
		})
	})
	r.Run(":9090")
}

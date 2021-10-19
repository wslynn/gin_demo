package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()
	r.GET("/index", func(c *gin.Context) {
		// 方法1: 使用map
		//data := map[string]interface{}{
		//	"name":    "小王子",
		//	"message": "hello world",
		//	"age":     18,
		//}
		// 方法2: 使用gin.H
		//data := gin.H{
		//	"name":    "小王子",
		//	"message": "hello world",
		//	"age":     18,
		//}
		// 方法3: 使用结构体, 变量名必须大写才能导出, 利用tag变为小写
		type msg struct {
			Name    string `json:"name"`
			Message string
			Age     int
		}
		data := msg{"小王子", "hello world", 18}
		c.JSON(http.StatusOK, data)
	})
	r.Run(":9090")
}

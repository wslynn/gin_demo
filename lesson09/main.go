package main

import (
	"github.com/gin-gonic/gin"
	"html/template"
	"net/http"
)

func main() {
	r := gin.Default()
	// 设置静态文件路径
	r.Static("/xxx", "./static")
	// 设置自定义的js函数
	r.SetFuncMap(template.FuncMap{
		"safe": func(str string) template.HTML {
			return template.HTML(str)
		},
	})
	r.LoadHTMLFiles("templates/index.html")

	r.GET("/index", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", "<a href='https://www.baidu.com'>百度</a>")
	})
	r.Run(":9090")
}

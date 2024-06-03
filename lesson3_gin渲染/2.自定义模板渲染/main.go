package main

import (
	"html/template"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.SetFuncMap(template.FuncMap{
		"safe": func(str string) template.HTML {
			return template.HTML(str)
		},
	})
	r.LoadHTMLFiles("./index.tmpl")
	r.GET("/index", func(c *gin.Context) {
		c.HTML(
			http.StatusOK,
			"index.tmpl",
			"<a href='https://www.baidu.com'>百度一下</a>",
		)
	})
	r.Run(":8080")
}

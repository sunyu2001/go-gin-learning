package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	// r.LoadHTMLGlob("templates/**/*") // 匹配多个
	r.LoadHTMLFiles("templates/posts/index.html", "templates/users/index.html")
	r.GET("/posts/index", func(c *gin.Context) {
		c.HTML(http.StatusOK, "posts/index.html", gin.H{
			"title": "post/index",
		})
	})
	r.GET("/user/index", func(c *gin.Context) {
		c.HTML(http.StatusOK, "users/index.html", gin.H{
			"title": "user/index",
		})
	})
	r.Run()
}

package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	// 创建一个默认路由引擎
	r := gin.Default()
	r.GET("/hello", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "hello world",
		})
	})
	r.Run() // 默认端口0.0.0.0:8080
}

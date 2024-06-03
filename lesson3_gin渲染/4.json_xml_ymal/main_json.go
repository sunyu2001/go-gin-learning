package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// gin.H是map[string] interface{}的缩写
	r.GET("/someJson", func(c *gin.Context) { // 自己拼接
		c.JSON(http.StatusOK, gin.H{
			"message": "Hello World!",
		})
	})
	r.GET("/moreJSON", func(c *gin.Context) { // 使用结构体
		var msg struct {
			Name    string `json:"user"`
			Message string
			Age     int
		}
		msg.Name = "小王子"
		msg.Message = "Hello World!"
		msg.Age = 18
		c.JSON(http.StatusOK, msg)
	})
	r.Run(":8080")
}

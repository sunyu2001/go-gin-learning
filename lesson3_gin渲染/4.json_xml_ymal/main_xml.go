package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// gin.H是map[string] interface{}的缩写
	r.GET("/someXML", func(c *gin.Context) { // 自己拼接
		c.XML(http.StatusOK, gin.H{
			"message": "Hello World!",
		})
	})
	r.GET("/moreXML", func(c *gin.Context) { // 使用结构体
		type MessageRecord struct {
			Name    string
			Message string
			Age     int
		}
		var msg MessageRecord
		msg.Name = "小王子"
		msg.Message = "Hello World!"
		msg.Age = 18
		c.XML(http.StatusOK, msg)
	})
	r.Run(":8080")
}

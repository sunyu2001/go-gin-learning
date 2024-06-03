package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/testdata/protoexample"
)

func main() {
	r := gin.Default()

	// gin.H是map[string] interface{}的缩写
	r.GET("/someYAML", func(c *gin.Context) { // 自己拼接
		c.YAML(http.StatusOK, gin.H{
			"message": "Hello World!",
		})
	})
	r.GET("/moreYAML", func(c *gin.Context) { // 使用结构体
		type MessageRecord struct {
			Name    string
			Message string
			Age     int
		}
		var msg MessageRecord
		msg.Name = "小王子"
		msg.Message = "Hello World!"
		msg.Age = 18
		c.YAML(http.StatusOK, msg)
	})
	r.GET("/someProtoBuf", func(c *gin.Context) {
		reps := []int64{int64(1), int64(2)}
		label := "test"
		data := &protoexample.Test{
			Label: &label,
			Reps:  reps,
		}
		c.ProtoBuf(http.StatusOK, data)
	})
	r.Run(":8080")
}

package main

import (
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func StatCost() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		c.Set("name", "sunyu")
		c.Next()
		cost := time.Since(start)
		log.Println(cost)
	}
}

func main() {
	r := gin.New()

	// 加载模板
	r.Use(StatCost())

	r.GET("/test", func(c *gin.Context) {
		name := c.MustGet("name").(string) // 从上下文取值
		log.Println(name)
		c.JSON(http.StatusOK, gin.H{
			"message": "Hello world!",
		})
	})
	r.Run(":8080")
}

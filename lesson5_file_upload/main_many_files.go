package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// 找到模板路径
	wd, err := os.Getwd()
	if err != nil {
		fmt.Printf("get wd failed, err: %v\n", wd)
		return
	}

	// 加载模板
	fmt.Print(wd)
	r.LoadHTMLFiles(wd + "/index.tmpl")
	r.GET("/index", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.tmpl", nil)
	})

	r.POST("/upload", func(c *gin.Context) {
		form, _ := c.MultipartForm()
		files := form.File["file"]
		for index, file := range files {
			log.Println(file.Filename)
			dst := fmt.Sprintf("E:/Go/code/go-gin-learning/%s_%d", file.Filename, index)
			c.SaveUploadedFile(file, dst)
		}

		c.JSON(http.StatusOK, gin.H{
			"message": fmt.Sprintf("'%d' files uploaded!", len(files)),
		})
	})
	r.Run(":8080")
}

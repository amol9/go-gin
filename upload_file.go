package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.POST("/upload", func(c *gin.Context) {
		file, _ := c.FormFile("file")
		c.SaveUploadedFile(file, "uploaded")

		c.String(http.StatusOK, "%s uploaded", file.Filename)
	})

	r.Run()
}

package main

import (
	"log"
	"mime/multipart"
	"net/http"

	"github.com/gin-gonic/gin"
)

func saveUploadedFile(c *gin.Context, file *multipart.FileHeader) {
	err := c.SaveUploadedFile(file, "/home/amol2/Documents/code/learn/gin/uploaded/"+file.Filename)
	if err != nil {
		log.Println("file save error:", err)
	}
}

func main() {
	r := gin.Default()

	r.POST("/single", func(c *gin.Context) {
		file, _ := c.FormFile("file")
		log.Println("uploaded filename: " + file.Filename)

		saveUploadedFile(c, file)

		c.String(http.StatusOK, "uploaded: %s", file.Filename)
	})

	r.POST("multiple", func(c *gin.Context) {
		form, err := c.MultipartForm()
		if err != nil {
			log.Fatal(err)
		}
		files := form.File["upload[]"]
		s := ""
		for _, file := range files {
			saveUploadedFile(c, file)
			s += " " + file.Filename
		}
		c.String(http.StatusOK, "uploaded: %s", s)
	})

	r.Run()
}

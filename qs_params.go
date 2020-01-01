package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()

	r.GET("/username", func(c *gin.Context) {
		firstname := c.DefaultQuery("firstname", "John")
		lastname := c.Query("lastname")

		c.String(200, "Hello %s %s", firstname, lastname)
	})

	r.Run()
}

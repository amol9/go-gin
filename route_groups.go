package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

/*func makeHandler(output string): (func(*gin.Context)) = {
	return func(c *gin.Context) {
		c.String(http.StatusOK, output)
	}
}*/

func main() {
	r := gin.Default()

	v1 := r.Group("/v1")
	{
		v1.POST("/login", func(c *gin.Context) { c.String(http.StatusOK, "login v1\n") })
		v1.POST("/submit", func(c *gin.Context) { c.String(http.StatusOK, "submit v1\n") })
		v1.POST("/register", func(c *gin.Context) { c.String(http.StatusOK, "register v1\n") })
	}

	v2 := r.Group("/v2")
	{
		v2.POST("/login", func(c *gin.Context) { c.String(http.StatusOK, "login v2\n") })
		v2.POST("/submit", func(c *gin.Context) { c.String(http.StatusOK, "submit v2\n") })
		v2.POST("/register", func(c *gin.Context) { c.String(http.StatusOK, "register v2\n") })
	}

	r.Run()
}

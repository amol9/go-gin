package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/user/:name", func(c *gin.Context) {
		name := c.Param("name")
		c.String(200, "<h1>Hello, %s</h1>", name)
	})

	/*r.GET("/user/:name", func(c *gin.Context) {
		name := c.Param("name")
		html := fmt.Sprintf("<h1>Hello, %s</h1>", name)
		c.HTML(200, html)
	})*/

	r.GET("/user/:name/*action", func(c *gin.Context) {
		name := c.Param("name")
		action := c.Param("action")
		out := name + " is " + action
		c.String(200, out)
	})

	r.Run()
}

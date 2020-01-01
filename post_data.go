package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.POST("/userdata", func(c *gin.Context) {
		name := c.PostForm("name")
		email := c.PostForm("email")
		message := c.DefaultPostForm("message", "Hello")

		c.JSON(200, gin.H{
			"name":    name,
			"email":   email,
			"message": message,
		})
	})

	/*form := `
		<form action="/userdata" method="POST">
			Name: <input type="text" name="name"><br>
			Email: <input type="text" name="email"><br>
			Message: <input type="textarea" name="message"><br>
			<input type="submit" value="Submit">
		</form>
	`*/

	r.LoadHTMLGlob("userform.html")

	r.GET("/userform", func(c *gin.Context) {
		c.HTML(200, "userform.html", nil)
	})

	r.Run()
}

// curl -d "name=Amol%20Umrale&email=a@b.com" "http://localhost:8080/userdata"

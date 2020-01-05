package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Login struct {
	User     string `json:"user" xml:"user" form:"user" binding:"required"`
	Password string `json:"password" xml:"password" form:"password" binding:"required"`
}

func auth(username string, password string) bool {
	if username == "test" && password == "123" {
		return true
	} else {
		return false
	}
}

func authUser(login Login, c *gin.Context) {
	log.Println("login obj: %v", login)
	if auth(login.User, login.Password) {
		c.String(http.StatusOK, "authorized")
	} else {
		c.String(http.StatusUnauthorized, "unauthorized")
	}
}

func main() {
	r := gin.Default()

	r.POST("/loginJson", func(c *gin.Context) {
		var json Login
		if err := c.ShouldBindJSON(&json); err != nil {
			log.Println(err)
			c.String(http.StatusInternalServerError, "%v", err)
			return
		}
		authUser(json, c)
	})

	// NOT WORKING, fixed (there should be no space between : and " in the  struct member tags)
	r.POST("/loginXML", func(c *gin.Context) {
		var xml Login
		if err := c.ShouldBindXML(&xml); err != nil {
			log.Println(err)
			c.String(http.StatusInternalServerError, "%v", err)
			return
		}
		log.Println("x %v", xml)
		authUser(xml, c)
	})

	r.POST("/loginForm", func(c *gin.Context) {
		var form Login
		if err := c.ShouldBind(&form); err != nil {
			log.Println(err)
			c.String(http.StatusInternalServerError, "%v", err)
			return
		}
		authUser(form, c)
	})

	r.Run()
}

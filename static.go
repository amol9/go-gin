package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.Static("/assets", "./assets")
	r.StaticFS("/moreassets", gin.Dir("/home/amol2/Pictures/", true))
	r.StaticFile("/favicon.ico", "./files/fav.ico")

	r.Run()
}

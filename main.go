package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.StaticFile("/", "static/index.html")
	router.StaticFile("/vue.js", "static/vue.js")
	router.StaticFile("/style.css", "static/style.css")
	router.StaticFile("/fontawesome-all.css", "static/css/fontawesome-all.css")

	router.Run(":8900")
}

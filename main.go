package main

import (
	"github.com/gin-gonic/gin"
)

var router *gin.Engine

func main() {
	// create Default router
	router = gin.Default()
	// Load templates/
	router.LoadHTMLGlob("templates/*")

	router.GET("/", showIndexPage)
	router.GET("/article/view/:article_id", getArticle)

	router.Run()
}

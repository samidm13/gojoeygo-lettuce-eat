package main

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func setupRouter() *gin.Engine {
	router := gin.Default()
	router.LoadHTMLGlob("templates/*")
	router.GET("/", showIndexPage)
	return router
}

func main() {
	port := os.Getenv("PORT")

	if port == "" {
		port = "8080"
	}
	router := setupRouter()
	router.Run(":" + port)
}

func showIndexPage(c *gin.Context) {
	c.HTML(
		http.StatusOK,
		"index.html",
		gin.H{
			"title": "Home Page",
		},
	)
}

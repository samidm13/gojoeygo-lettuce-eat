package main

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

const (
	host   = "localhost"
	port   = 5432
	dbname = "lettuce_eat"
)

func setupRouter() *gin.Engine {
	router := gin.Default()
	router.LoadHTMLGlob("templates/*")
	router.GET("/", showIndexPage)
	router.GET("/signup", showSignUpPage)
	router.POST("/signup", signUp)
	router.GET("/login", showLogInPage)
	router.POST("/login", logIn)
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

package main

import "github.com/gin-gonic/gin"

func setupRouter() *gin.Engine {

	router := gin.Default()

	router.LoadHTMLGlob("templates/*")

	router.GET("/", showIndexPage)

  router.GET("/signup", showSignUpPage)
  
	router.POST("/signup", signUp)

	return router
}

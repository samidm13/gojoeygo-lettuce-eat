package main

import "github.com/gin-gonic/gin"

func setupRouter() *gin.Engine {
	router := gin.Default()
	router.LoadHTMLGlob("templates/*")
	router.Use(setUserStatus())
	router.GET("/", ensureLoggedIn(), showIndexPage)
	router.GET("/signup", ensureNotLoggedIn(), showSignUpPage)
	router.POST("/signup", ensureNotLoggedIn(), signUp)
	router.GET("/login", ensureNotLoggedIn(), showLogInPage)
	router.POST("/login", ensureNotLoggedIn(), logIn)
	router.GET("/logout", ensureLoggedIn(), logOut)
	return router
}

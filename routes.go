package main

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

func setupRouter() *gin.Engine {

	router := gin.Default()

	store := cookie.NewStore([]byte("secret"))

	router.Use(sessions.Sessions("mysession", store))

	router.LoadHTMLGlob("templates/*")

	router.Static("public", "./public")

	router.Use(setUserStatus())

	router.GET("/", ensureLoggedIn(), showIndexPage)

	router.GET("/signlog", ensureNotLoggedIn(), showSignLogPage)

	router.POST("/signup", ensureNotLoggedIn(), signUp)

	router.POST("/login", ensureNotLoggedIn(), logIn)

	router.GET("/logout", ensureLoggedIn(), logOut)

	router.GET("/restaurants", ensureLoggedIn(), showRestaurants)

	router.POST("/orders", ensureLoggedIn(), orderSetUp)

	router.GET("/orders", ensureLoggedIn(), showOrderPage)

	router.GET("/menu", ensureLoggedIn())

	router.POST("/basket", ensureLoggedIn(), createBasket)

	return router
}

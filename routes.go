package main

import "github.com/gin-gonic/gin"

func setupRouter() *gin.Engine {
	router := gin.Default()
	router.LoadHTMLGlob("templates/*")
<<<<<<< HEAD

	router.GET("/", showIndexPage)

  router.GET("/signup", showSignUpPage)

	router.POST("/signup", signUp)

	router.GET("/restaurants", showRestaurants)

	router.POST("/orders", orderSetUp)

	router.GET("/orders", showOrderPage)


=======
	router.Use(setUserStatus())
	router.GET("/", ensureLoggedIn(), showIndexPage)
	router.GET("/signup", ensureNotLoggedIn(), showSignUpPage)
	router.POST("/signup", ensureNotLoggedIn(), signUp)
	router.GET("/login", ensureNotLoggedIn(), showLogInPage)
	router.POST("/login", ensureNotLoggedIn(), logIn)
	router.GET("/logout", ensureLoggedIn(), logOut)
>>>>>>> 7968f279252dfa3df3cc1da07471be6d7e8ab139
	return router
}

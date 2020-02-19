package main

import "github.com/gin-gonic/gin"

func setUserStatus() gin.HandlerFunc {
	return func(c *gin.Context) {
		if cookieValue, err := c.Cookie("name"); err == nil || cookieValue != "" {
			c.Set("is_logged_in", true)
		} else {
			c.Set("is_logged_in", false)
		}
	}
}
func ensureLoggedIn() gin.HandlerFunc {
	return func(c *gin.Context) {
		loggedInInterface, _ := c.Get("is_logged_in")
		loggedIn := loggedInInterface.(bool)
		if !loggedIn {
			c.Redirect(
				303,
				"/login",
			)
		}
	}
}
func ensureNotLoggedIn() gin.HandlerFunc {
	return func(c *gin.Context) {
		loggedInInterface, _ := c.Get("is_logged_in")
		loggedIn := loggedInInterface.(bool)
		if loggedIn {
			c.Redirect(
				303,
				"/",
			)
		}
	}
}

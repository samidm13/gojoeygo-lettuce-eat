package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"strings"
	"strconv"
	"math/rand"
)

func showIndexPage(c *gin.Context) {
	c.HTML(
		http.StatusOK,
		"index.html",
		gin.H{
			"title": "Home Page",
		},
	)
}

func showSignUpPage(c *gin.Context) {
  c.HTML(
    http.StatusOK,
    "signup.html",
    gin.H{
      "title": "Sign Up",
    },
  )
}

func signUp(c *gin.Context) {
  firstname := c.PostForm("first_name")
  lastname := c.PostForm("last_name")
  mail := c.PostForm("email")
  pass := c.PostForm("password")

  registerNewUser(firstname, lastname, mail, pass)

	token := strconv.FormatInt(rand.Int63(), 16)
	c.SetCookie("token", token, 3600, "", "", false, true)
	c.Set("is_logged_in", true)

  c.Redirect(
    303,
    "/",
  )
}

func showLogInPage(c *gin.Context) {
	c.HTML(
		http.StatusOK,
		"login.html",
		gin.H{
			"title": "Log In",
		},
	)
}

func logIn(c *gin.Context) {
	remail := strings.TrimSpace(c.PostForm("email"))
	rpassword := strings.TrimSpace(c.PostForm("password"))

	varify := userLogIn(remail, rpassword)

	if varify == true {

		token := strconv.FormatInt(rand.Int63(), 16)
		c.SetCookie("token", token, 3600, "", "", false, true)
		c.Set("is_logged_in", true)

	c.Redirect(
		303,
		"/",
	)
} else {
	c.Redirect(
		303,
		"/signup",
	)
}

}

func logOut(c *gin.Context) {
	// Clear the cookie
	c.SetCookie("token", "", -1, "", "", false, true)
	c.Redirect(
		303,
		"/login",
	)
}

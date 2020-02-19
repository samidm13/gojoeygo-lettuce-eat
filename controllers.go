package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
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

  c.Redirect(
    303,
    "/",
  )
}

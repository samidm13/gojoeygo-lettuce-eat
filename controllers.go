package main

import (
	"net/http"
	"math/rand"
	"strconv"
	"time"

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

func showRestaurants(c *gin.Context) {
  restaurants := getAllRestaurants()
	rand.Seed(time.Now().UnixNano())
  token := rand.Intn(100000)
  c.HTML(
    http.StatusOK,
    "restaurants.html",
    gin.H{
      "title": "Restaurant page",
      "payload": restaurants,
      "token": token,
    },
  )
}

func orderSetUp(c *gin.Context) {
  restID := c.PostForm("RestID")
  token := c.PostForm("token")
	address := c.PostForm("address")
	// userID := c.PostForm("user_id")

	rest_ID, _ := strconv.Atoi(restID)
	tok, _ := strconv.Atoi(token)
	// usID, _ := strconv.Atoi(userID)

	orderPlacing(rest_ID, tok, address)

	c.Redirect(
		303,
		"/orders",
	)
}

func showOrderPage(c *gin.Context) {
	tok := getOrders()
	c.HTML(
		http.StatusOK,
		"orders.html",
		gin.H{
			"title": "Order",
			"payload": tok,
		},
	)
}

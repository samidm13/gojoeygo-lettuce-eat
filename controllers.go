package main

import (
	"net/http"
	"math/rand"
	"strconv"
	"time"
	"strings"

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
	rtoken := c.PostForm("token")
	token, _ := strconv.Atoi(rtoken)

	user_id := userSignUp(firstname, lastname, mail, pass)
	cookieValue := strconv.Itoa(user_id)
	c.SetCookie("name", cookieValue, 3600, "", "", false, true)
	c.Set("is_logged_in", true)

	if token == 0 {
		c.Redirect(
			303,
			"/restaurants",
	)
 } else {
	 _, error := getToken(token)

	 if error != nil {
		 c.Redirect(
			 303,
			 "/",
	 	)
	 } else {
	 c.Redirect(
		 303,
		 "/login",
 	)
}
 }
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

	user_id, password := userLogIn(remail, rpassword)

	cookieValue := strconv.Itoa(user_id)

	if CheckPasswordHash(rpassword, password) {
		c.SetCookie("name", cookieValue, 3600, "", "", false, true)
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
	c.SetCookie("name", "", -1, "", "", false, true)
	c.Redirect(
		303,
		"/login",
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

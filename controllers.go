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

	user_id := userSignUp(firstname, lastname, mail, pass)
	cookieValue := strconv.Itoa(user_id)
	c.SetCookie("name", cookieValue, 3600, "", "", false, true)
	c.Set("is_logged_in", true)

	if rtoken == "" {
		c.Redirect(
			303,
			"/restaurants",
	)
	return
	}

token, _ := strconv.Atoi(rtoken)

validTokens := getToken(token)

	 if len(validTokens) == 0 {
		 c.Redirect(
			 303,
			 "/restaurants",
	 	)
		return
	 }

	 if validTokens[0].Expiration.Before(time.Now()) {
		 c.Redirect(
			 303,
			 "/restaurants",
	 	)
		return
	 }

	 c.Redirect(
		 303,
		 "/menu",
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

	user_id, password := userLogIn(remail, rpassword)

	cookieValue := strconv.Itoa(user_id)

	if CheckPasswordHash(rpassword, password) {
		c.SetCookie("name", cookieValue, 3600, "", "", false, true)
		c.Set("is_logged_in", true)
		c.Redirect(
			303,
			"/restaurants",
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
	// now := time.Now().Format(time.RFC3339)
	// end := time.Now()+
  c.HTML(
    http.StatusOK,
    "restaurants.html",
    gin.H{
      "title": "Restaurant page",
      "payload": restaurants,
      "token": token,
			// "now": now,
    },
  )
}




func orderSetUp(c *gin.Context) {
  restID := c.PostForm("RestID")
  token := c.PostForm("token")
	address := c.PostForm("address")
	userID, _ := c.Cookie("name")
	time := c.PostForm("appt")

	rest_ID, _ := strconv.Atoi(restID)
	tok, _ := strconv.Atoi(token)
	usID, _ := strconv.Atoi(userID)

	orderPlacing(rest_ID, tok, address, usID, time)

	c.Redirect(
		303,
		"/orders",
	)
}

func showOrderPage(c *gin.Context) {
	userID, _ := c.Cookie("name")
	usID, _ := strconv.Atoi(userID)
	tok := getOrders(usID)
	c.HTML(
		http.StatusOK,
		"orders.html",
		gin.H{
			"title": "Order",
			"payload": tok,
		},
	)
}

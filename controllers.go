package main

import (
	"math/rand"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/sessions"
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
func showSignLogPage(c *gin.Context) {
	c.HTML(
		http.StatusOK,
		"signlog.html",
		gin.H{
			"title": "Sign Up / Log In",
		},
	)
}
func signUp(c *gin.Context) {
	firstname := c.PostForm("first_name")
	lastname := c.PostForm("last_name")
	mail := c.PostForm("email")
	pass := c.PostForm("password")
	rtoken := strings.TrimSpace(c.PostForm("signtoken"))

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
		session := sessions.Default(c)
		session.AddFlash("Invalid Token")
		flash := session.Flashes()
		session.Save()
		c.HTML(
			http.StatusOK,
			"signlog.html",
			gin.H{
				"title": "Sign Up/ Log In",
				"flashes": flash,
		})

		return
	}

	if validTokens[0].Expiration.Before(time.Now()) {

		session := sessions.Default(c)
		session.AddFlash("Expired Token")
		flash := session.Flashes()
		session.Save()
		c.HTML(
			http.StatusOK,
			"signlog.html",
			gin.H{
				"title": "Sign Up/ Log In",
				"flashes": flash,
		})

		return
	}

	menuList := displayMenu(token)
	c.HTML(
		http.StatusOK,
		"menu.html",
		gin.H{
			"title":   "Menu",
			"payload": menuList,
		},
	)
}

// func showMenuPage(c *gin.Context) {
// 	token := c.PostForm("token")
// 	fmt.Println(token)
// 	tokenint, _ := strconv.Atoi(token)
// 	fmt.Println(tokenint)
// 	validTokens := displayMenu(tokenint)
// 	fmt.Println("!")
// 	fmt.Println(validTokens)

// 	c.HTML(
// 		http.StatusOK,
// 		"menu.html",
// 		gin.H{
// 			"title":   "Menu",
// 			"payload": validTokens,
// 		},
// 	)

// }

func logIn(c *gin.Context) {
	remail := strings.TrimSpace(c.PostForm("username"))
	rpassword := strings.TrimSpace(c.PostForm("pass"))
	rtoken := strings.TrimSpace(c.PostForm("logtoken"))

	user_id, password := userLogIn(remail, rpassword)

	cookieValue := strconv.Itoa(user_id)

	if CheckPasswordHash(rpassword, password) {
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
			session := sessions.Default(c)
			session.AddFlash("Invalid Token")
			flash := session.Flashes()
			session.Save()
			c.HTML(
				http.StatusOK,
				"signlog.html",
				gin.H{
					"title": "Sign Up/ Log In",
					"flashes": flash,
			})
				return
			 }

			if validTokens[0].Expiration.Before(time.Now()) {
				session := sessions.Default(c)
			 session.AddFlash("Expired Token")
			 flash := session.Flashes()
			 session.Save()
			 c.HTML(
				 http.StatusOK,
				 "signlog.html",
				 gin.H{
					 "title": "Sign Up/ Log In",
					 "flashes": flash,
			 })
				return
			 }

			 menuList := displayMenu(token)
		 	c.HTML(
		 		http.StatusOK,
		 		"menu.html",
		 		gin.H{
		 			"title":   "Menu",
		 			"payload": menuList,
		 		},
		 	)

			} else {
					session := sessions.Default(c)
					session.AddFlash("The email or password is incorrect")
					flash := session.Flashes()
					session.Save()
					c.HTML(
						http.StatusOK,
						"signlog.html",
						gin.H{
							"title": "Sign Up / Log In",
							"flashes": flash,
							})
			}
}
func logOut(c *gin.Context) {
	// Clear the cookie
	c.SetCookie("name", "", -1, "", "", false, true)
	c.Redirect(
		303,
		"/signlog",
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
			"title":   "Restaurant page",
			"payload": restaurants,
			"token":   token,
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
			"title":   "Order",
			"payload": tok,
		},
	)
}

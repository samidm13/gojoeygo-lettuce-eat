package main

import  (
  "database/sql"
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

  

func registerNewUser(firstname string, lastname string, mail string, pass string) {

  sqlStatement := `
  INSERT INTO users (first_name, last_name, email, password)
  VALUES ($1, $2, $3, $4)`

  _, err := DB.Exec(sqlStatement, firstname, lastname, mail, pass)

  if err != nil {
		panic(err)
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
	psqlInfo := fmt.Sprintf("host=%s port=%d  "+
		" dbname=%s sslmode=disable",
		host, port, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	var (
		userCount int
		user_id   int
		email     string
		password  string
	)
	db.QueryRow("SELECT COUNT(user_id) AS userCount, user_id, email, password FROM users WHERE email=$1 GROUP BY user_id", remail).Scan(&userCount, &user_id, &email, &password)
	if password == rpassword {
		// token := strconv.FormatInt(rand.Int63(), 16)
		// c.SetCookie("token", token, 3600, "", "", false, true)
		// c.Set("is_logged_in", true)
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

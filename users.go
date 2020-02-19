package main

<<<<<<< HEAD
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
=======
import (
	_ "github.com/lib/pq"
)

var (
	userCount int
	user_id   int
	email     string
	password  string
)

func userSignUp(firstname string, lastname string, mail string, pass string) int {
	sqlStatement := `
>>>>>>> 7968f279252dfa3df3cc1da07471be6d7e8ab139
  INSERT INTO users (first_name, last_name, email, password)
  VALUES ($1, $2, $3, $4) RETURNING user_id`
	err := DB.QueryRow(sqlStatement, firstname, lastname, mail, pass).Scan(&user_id)
	if err != nil {
		panic(err)
	}
	return user_id
}
func userLogIn(remail string, rpassword string) (int, string) {
	DB.QueryRow("SELECT COUNT(user_id) AS userCount, user_id, email, password FROM users WHERE email=$1 GROUP BY user_id", remail).Scan(&userCount, &user_id, &email, &password)
	return user_id, password
}

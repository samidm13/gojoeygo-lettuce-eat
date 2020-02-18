package main

import (
  "database/sql"
	"fmt"
  "net/http"
  "github.com/gin-gonic/gin"
  _"github.com/lib/pq"

)

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

  psqlInfo := fmt.Sprintf("host=%s port=%d "+
    " dbname=%s sslmode=disable",
    host, port, dbname)

  db, err := sql.Open("postgres", psqlInfo)

  if err != nil {
    panic(err)
  }
  defer db.Close()

  sqlStatement := `
  INSERT INTO users (first_name, last_name, email, password)
  VALUES ($1, $2, $3, $4)`
	_, err = db.Exec(sqlStatement, firstname, lastname, mail, pass)
	if err != nil {
		panic(err)
}

c.Redirect(
  303,
  "/",
)
}

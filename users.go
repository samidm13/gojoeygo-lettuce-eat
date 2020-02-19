package main

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

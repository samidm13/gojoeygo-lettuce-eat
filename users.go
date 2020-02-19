package main

import(

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

func userLogIn(remail string, rpassword string) bool {
	var (
		userCount int
		user_id   int
		email     string
		password  string
	)
	DB.QueryRow("SELECT COUNT(user_id) AS userCount, user_id, email, password FROM users WHERE email=$1 GROUP BY user_id", remail).Scan(&userCount, &user_id, &email, &password)
	if password == rpassword {
   return true
	} else {
return false
	}
}

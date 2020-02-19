package main

import   _"github.com/lib/pq"

func registerNewUser(firstname string, lastname string, mail string, pass string) {

  sqlStatement := `
  INSERT INTO users (first_name, last_name, email, password)
  VALUES ($1, $2, $3, $4)`

  _, err := DB.Exec(sqlStatement, firstname, lastname, mail, pass)

  if err != nil {
		panic(err)
  }
}

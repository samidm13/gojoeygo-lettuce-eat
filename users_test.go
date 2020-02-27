package main

import (

  "testing"

	_ "github.com/lib/pq"
)

func TestUserSignUp(t *testing.T) {

  emptyDB()

  actual := userSignUp("NameTest", "SurnameTest", "mail@test.com", "12345")

  if &actual == nil {
			t.Fail()
  }
}

func TestUserLogIn(t *testing.T) {

  emptyDB()

  expectedID := userSignUp("NameTest", "SurnameTest", "mail@test.com", "12345")

  actualID, actualPass := userLogIn("mail@test.com", "12345")

  if !CheckPasswordHash("12345", actualPass) || expectedID != actualID {
			t.Fail()
  }
}

package main

import (

  "testing"

	_ "github.com/lib/pq"
)

func TestGetOrders(t *testing.T) {

  emptyDB()

  populateDB()

  actual := getOrders(1)

  if len(actual) != 1 {
			t.Fail()
  }
}

func TestGetToken(t *testing.T) {

  emptyDB()

  populateDB()

  actual := getToken(4321)

  if len(actual) != 1 || actual[0].TokenID != 4321 {
			t.Fail()
  }
}

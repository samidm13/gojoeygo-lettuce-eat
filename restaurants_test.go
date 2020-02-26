package main

import (

  "testing"

	_ "github.com/lib/pq"
)

func TestGetAllRestaurants(t *testing.T) {

  emptyDB()

  populateDB()

  actual := getAllRestaurants()

  if len(actual) != 1 {
			t.Fail()
  }
}

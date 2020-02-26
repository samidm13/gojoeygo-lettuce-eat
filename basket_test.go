package main

import (

  "testing"

	_ "github.com/lib/pq"
)

func TestGetDish(t *testing.T) {

  emptyDB()

  populateDB()

  actual := getDish(1)

  expected := dish{"Fruits", "2.99"}

  if len(actual) != 1 || actual[0].dish_name != expected.dish_name || actual[0].dish_price != expected.dish_price{
			t.Fail()
  }
}

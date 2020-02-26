package main

import (

  "testing"
  "strconv"

	_ "github.com/lib/pq"
)

func TestDishesInBasket(t *testing.T) {

  emptyDB()

  populateDB()

  addBasket(1, "Fruits", 2.99, 4321, 1)
  addBasket(2, "Croissant", 1.99, 4321, 1)

  actual := dishesInBasket(4321, 1)
  expected := []dishbasket {
    {"Fruits", "2.99"},
    {"Croissant", "1.99"},
  }

  if len(actual) != 2 || actual[0].DishName != expected[0].DishName || actual[0].DishPrice != expected[0].DishPrice{
			t.Fail()
  }
}

func TestTotalPrice(t *testing.T) {

  emptyDB()

  populateDB()

  addBasket(1, "Fruits", 2.99, 4321, 1)
  addBasket(2, "Croissant", 1.99, 4321, 1)

  actual, _ := strconv.ParseFloat(totalPrice(4321, 1), 64)
  expected := 4.98

  if actual != expected {
    t.Fail()
  }
}

package main

import (

  "testing"
  "reflect"
  // "strconv"

	_ "github.com/lib/pq"
)

func TestDishesInOrder(t *testing.T) {

  emptyDB()

  populateDB()

  addBasket(1, "Fruits", 2.99, 4321, 1)
  addBasket(2, "Croissant", 1.99, 4321, 1)

  actual := dishesInOrder(4321)
  expected := []dishorder {
    {"Croissant", "1.99", "NameTest", "SurnameTest"},
    {"Fruits", "2.99", "NameTest", "SurnameTest"},
  }

  if !reflect.DeepEqual(actual, expected){
			t.Errorf("got %v want %v", actual, expected)
  }
}

func TestTotalOrderPrice(t *testing.T) {

  emptyDB()

  populateDB()

  addBasket(1, "Fruits", 2.99, 4321, 1)
  addBasket(2, "Croissant", 1.99, 4321, 1)

  actual := totalOrderPrice(4321)
  expected := "4.98"

  if actual != expected {
    t.Fail()
  }
}

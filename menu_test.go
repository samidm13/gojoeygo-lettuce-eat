package main

import (

  "testing"

	_ "github.com/lib/pq"
)

func TestDisplayMenu(t *testing.T) {

  emptyDB()

  populateDB()

  actual := displayMenu(4321)

  if len(actual) != 2 {
			t.Fail()
  }
}

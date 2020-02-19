package main

import (
	"database/sql"
	"fmt"

	_"github.com/lib/pq"
)

const (
	host   = "localhost"
	port   = 5432
	dbname = "lettuce_eat"
)

var DB *sql.DB
var err error


func main() {

	psqlInfo := fmt.Sprintf("host=%s port=%d "+
    " dbname=%s sslmode=disable",
    host, port, dbname)

  DB, err = sql.Open("postgres", psqlInfo)

  if err != nil {
    panic(err)
  }
  defer DB.Close()

	router := setupRouter()

	router.Run()
}

package main

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/lib/pq"
)

func emptyDB() {

  psqlInfo := fmt.Sprintf("host=%s port=%d "+
    " dbname=%s sslmode=disable",
    host, port, dbname)

  DB, err = sql.Open("postgres", psqlInfo)

  if err != nil {
    panic(err)
  }

  sqlStatement_1 := `TRUNCATE users CASCADE`
  sqlStatement_2 := `TRUNCATE restaurants CASCADE`
  sqlStatement_3 := `TRUNCATE dishes CASCADE`
  sqlStatement_4 := `TRUNCATE orders CASCADE`
  sqlStatement_5 := `TRUNCATE basket CASCADE`

  DB.Query(sqlStatement_1)
  DB.Query(sqlStatement_2)
  DB.Query(sqlStatement_3)
  DB.Query(sqlStatement_4)
  DB.Query(sqlStatement_5)
}

func populateDB() {

	psqlInfo := fmt.Sprintf("host=%s port=%d "+
    " dbname=%s sslmode=disable",
    host, port, dbname)

  DB, err = sql.Open("postgres", psqlInfo)

  if err != nil {
    panic(err)
  }

	sqlStatement_1 := `INSERT INTO users (user_id, first_name, last_name, email, password) VALUES (1, 'NameTest', 'SurnameTest', 'test@test.com', '12345')`
  sqlStatement_2 := `INSERT INTO restaurants (rest_id, rest_name, url) VALUES (1,'Test Restaurant', '/public/images/cherries.jpg')`
	sqlStatement_3 := `INSERT INTO orders (rest_id, token, address, user_id, order_time) VALUES (1, '4321', 'London', 1, $1)`
	sqlStatement_4 := `INSERT INTO dishes (dish_id, dish_name, dish_price, rest_id, url) VALUES (1, 'Fruits', 2.99, 1, '/public/images/cherries.jpg'), (2, 'Croissant', 1.99, 1, '/public/images/croissant.jpg')`

  DB.Exec(sqlStatement_1)
	DB.Exec(sqlStatement_2)
	DB.Exec(sqlStatement_3, time.Now())
	DB.Exec(sqlStatement_4)
}

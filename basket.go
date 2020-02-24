package main

import (
	_ "github.com/lib/pq"
)

type dish struct {
	dish_name  string
	dish_price string
}

func getDish(DishID int) []dish {
	sqlStatement := `SELECT dish_name, dish_price FROM dishes WHERE dish_id=$1`
	rows, err := DB.Query(sqlStatement, DishID)
	if err != nil {
		panic(err)
	}
	dishes := make([]dish, 0)
	for rows.Next() {
		var entry dish
		rows.Scan(&entry.dish_name, &entry.dish_price)
		dishes = append(dishes, entry)
	}
	return dishes
}

func addBasket(DishID int, DishName string, DishPrice float64, Token int, UserID int) {
	sqlStatement := `
	INSERT INTO basket (dish_id, dish_name, dish_price, token, user_id)
	VALUES ($1, $2, $3, $4, $5)`

	_, err := DB.Exec(sqlStatement, DishID, DishName, DishPrice, Token, UserID)

	if err != nil {
		panic(err)
	}
}

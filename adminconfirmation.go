package main

import (
	_ "github.com/lib/pq"
)

func totalOrderPrice(token int) string {
	var total string
	sqlStatement := `SELECT SUM(dish_price) as total FROM basket WHERE token=$1;`
	row := DB.QueryRow(sqlStatement, token)
	row.Scan(&total)
	return total
}

type dishorder struct {
	DishName  string
	DishPrice string
	FirstName string
	LastName  string
}

func dishesInOrder(token int) []dishorder {
	sqlStatement := `SELECT basket.dish_name, basket.dish_price, users.first_name, users.last_name FROM basket LEFT JOIN users ON basket.user_id = users.user_id WHERE basket.token=$1`
	rows, err := DB.Query(sqlStatement, token)
	if err != nil {
		panic(err)
	}
	dishes := make([]dishorder, 0)
	for rows.Next() {
		var entry dishorder
		rows.Scan(&entry.DishName, &entry.DishPrice, &entry.FirstName, &entry.LastName)
		dishes = append(dishes, entry)
	}
	return dishes
}

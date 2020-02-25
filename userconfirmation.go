package main

import (
	_ "github.com/lib/pq"
)

func totalPrice(token int, userID int) string {
	var total string
	sqlStatement := `SELECT SUM(dish_price) as total FROM basket WHERE token=$1 and user_id=$2;`
	row := DB.QueryRow(sqlStatement, token, userID)
	row.Scan(&total)
	return total
}

type dishbasket struct {
	DishName  string
	DishPrice string
}

func dishesInBasket(token int, userID int) []dishbasket {
	sqlStatement := `SELECT dish_name, dish_price FROM basket WHERE token=$1 and user_id=$2`
	rows, err := DB.Query(sqlStatement, token, userID)
	if err != nil {
		panic(err)
	}
	dishes := make([]dishbasket, 0)
	for rows.Next() {
		var entry dishbasket
		rows.Scan(&entry.DishName, &entry.DishPrice)
		dishes = append(dishes, entry)
	}
	return dishes
}

func orderTime(token int) string {
	var time string
	sqlStatement := `SELECT order_time FROM orders WHERE token=$1;`
	row := DB.QueryRow(sqlStatement, token)
	row.Scan(&time)
	return time
}

package main

import (
	_ "github.com/lib/pq"
)

type tokens struct {
	Token           int
	RestID          int
	OrderID         int
	DishName        string
	DishPrice       float64
	DishDescription string
	DishID          int
}

func displayMenu(token int) []tokens {
	rows, err := DB.Query("SELECT orders.token, orders.rest_id, orders.order_id, dishes.dish_name, dishes.dish_price, dishes.description, dishes.dish_id FROM orders LEFT JOIN dishes ON dishes.rest_id = orders.rest_id WHERE orders.token=$1", token)
	if err != nil {
		panic(err)
	}
	validTokens := make([]tokens, 0)
	for rows.Next() {
		var entry tokens
		rows.Scan(&entry.Token, &entry.RestID, &entry.OrderID, &entry.DishName, &entry.DishPrice, &entry.DishDescription, &entry.DishID)
		validTokens = append(validTokens, entry)
	}
	return validTokens
}

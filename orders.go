package main

import  (
 "fmt"
)


func orderPlacing(rest_ID int, tok int, address string, usID int ){
	sqlStatement := `
	INSERT INTO orders (rest_id, token, address, user_id)
	VALUES ($1, $2, $3, $4)`

	_, err := DB.Exec(sqlStatement, rest_ID, tok, address, usID)

	if err != nil {
		panic(err)
	}
}

type order struct {
	Token int
	RestID int
	address string
	userID int
}

func getOrders(usID int) []order {
	sqlStatement := `SELECT token, rest_id FROM orders WHERE user_id=$1`
	rows, err := DB.Query(sqlStatement, usID)
	if err != nil {
		panic(err)
	}

	orders := make([]order, 0)
	for rows.Next() {
		var entry order
		rows.Scan(&entry.Token, &entry.RestID)
		orders = append(orders, entry)

		fmt.Println(orders)
	}

	return orders
}

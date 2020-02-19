package main

import  (
 "fmt"
)


func orderPlacing(rest_ID int, tok int, address string){
	sqlStatement := `
	INSERT INTO orders (rest_id, token, address)
	VALUES ($1, $2, $3)`

	_, err := DB.Exec(sqlStatement, rest_ID, tok, address)

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

func getOrders() []order {
	sqlStatement := `SELECT token, rest_id FROM orders`
	rows, err := DB.Query(sqlStatement)
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

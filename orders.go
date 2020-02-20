package main

import  (
 "fmt"
)


func orderPlacing(rest_ID int, tok int, address string, usID int, time string){
	sqlStatement := `
	INSERT INTO orders (rest_id, token, address, user_id, order_time)
	VALUES ($1, $2, $3, $4, $5)`

	_, err := DB.Exec(sqlStatement, rest_ID, tok, address, usID, time)

	if err != nil {
		panic(err)
	}
}

type order struct {
	Token int
	RestID int
	Address string
	UserID int
  Ordertime string
}

func getOrders(usID int) []order {
	sqlStatement := `SELECT token, rest_id, order_time FROM orders WHERE user_id=$1`
	rows, err := DB.Query(sqlStatement, usID)
	if err != nil {
		panic(err)
	}

	orders := make([]order, 0)
	for rows.Next() {
		var entry order
		rows.Scan(&entry.Token, &entry.RestID, &entry.Ordertime)
		orders = append(orders, entry)

		fmt.Println(orders)
	}

	return orders
}

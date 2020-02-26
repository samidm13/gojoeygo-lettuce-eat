package main

import (
  _"github.com/lib/pq"
)
type restaurant struct {
  RestID int
  RestName string
  RestURL string
}
type menu struct {
  dishID int
  dishName string
  dishPrice int
  description string
}
func getAllRestaurants() []restaurant {
  sqlStatement := `SELECT rest_id, rest_name, url FROM restaurants`
  rows, err := DB.Query(sqlStatement)
  if err != nil {
    panic(err)
  }
  restaurants := make([]restaurant, 0)
  for rows.Next() {
    var entry restaurant
    rows.Scan(&entry.RestID, &entry.RestName, &entry.RestURL)
    restaurants = append(restaurants, entry)
  }
  return restaurants
}

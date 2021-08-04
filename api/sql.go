package main

import (
  "database/sql"
  "fmt"
  "encoding/json"
  _ "github.com/mattn/go-sqlite3"
)

type Product struct {
  Name     string `json:"Name"`
  Quantity int    `json:"quantity"`
}

func GetAllRows() []Product{
  // open connection
  db, err := sql.Open("sqlite3", "/opt/db/api.db")
  checkSQLError(err)

  // build query
  rows, err := db.Query("SELECT * FROM products")
  checkSQLError(err)

  var name string
  var quantity int
  var products []Product

  // convert table rows to JSON
  for rows.Next() {
    err = rows.Scan(&name, &quantity)
    checkSQLError(err)

    fmt.Println(name)
    fmt.Println(quantity)

    products = append(products, Product{Name:name, Quantity:quantity})
  }

  rows.Close()

  return products
}


func checkSQLError(err error) {
  if err != nil {
    panic(err)
  }
}

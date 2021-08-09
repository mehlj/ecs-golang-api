package main

import (
  "database/sql"
  "fmt"
  _ "github.com/mattn/go-sqlite3"
)

func GetAllRows() []Product{
  // open connection
  db, err := sql.Open("sqlite3", "/opt/db/api.db")
  checkSQLError(err)

  // build statement
  rows, err := db.Query("SELECT * FROM products")
  checkSQLError(err)

  var name string
  var quantity int
  var products []Product

  // convert table rows to JSON
  for rows.Next() {
    err = rows.Scan(&name, &quantity)
    checkSQLError(err)

    products = append(products, Product{Name:name, Quantity:quantity})
  }

  return products
}

func InsertRow(product Product){
  // open connection
  db, err := sql.Open("sqlite3", "/opt/db/api.db")
  checkSQLError(err)

  // build statement
  stmt, err := db.Prepare("INSERT INTO products(name, quantity)  values(?,?)")
  checkSQLError(err)

  fmt.Println("insert name is", product.Name)

  // execute statement
  _, err = stmt.Exec(product.Name, product.Quantity)
  checkSQLError(err)
}

func RemoveRow(product Product){
  // open connection
  db, err := sql.Open("sqlite3", "/opt/db/api.db")
  checkSQLError(err)

  fmt.Println("delete name is", product.Name)
  // delete product
  _, err = db.Exec("DELETE FROM products WHERE name=?", product.Name)
  checkSQLError(err)
}


func checkSQLError(err error) {
  if err != nil {
    panic(err)
  }
}

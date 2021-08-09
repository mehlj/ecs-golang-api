package main

import (
  "database/sql"
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

  // execute statement
  _, err = stmt.Exec(product.Name, product.Quantity)
  checkSQLError(err)
}

func RemoveRow(product Product){
  // open connection
  db, err := sql.Open("sqlite3", "/opt/db/api.db")
  checkSQLError(err)

  // delete product
  _, err = db.Exec("DELETE FROM products WHERE name=?", product.Name)
  checkSQLError(err)
}

func UpdateRow(product Product){
  // open connection
  db, err := sql.Open("sqlite3", "/opt/db/api.db")
  checkSQLError(err)

  // update product
  _, err = db.Exec("UPDATE products SET quantity=? WHERE name=?", product.Quantity, product.Name)
  checkSQLError(err)
}

func QueryRow(k string) Product{
  // open connection
  db, err := sql.Open("sqlite3", "/opt/db/api.db")
  checkSQLError(err)

  // query product
  rows, err := db.Query("SELECT * FROM products WHERE name=?", k)
  checkSQLError(err)

  var name string
  var quantity int
  var p Product

  // convert table rows to JSON
  for rows.Next() {
    err = rows.Scan(&name, &quantity)
    checkSQLError(err)

    p = Product{Name:name, Quantity:quantity}
  }
  return p
}


func checkSQLError(err error) {
  if err != nil {
    panic(err)
  }
}

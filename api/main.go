package main

import (
  "fmt"
  "net/http"
  "log"
  "io/ioutil"
  "encoding/json"

  "github.com/gorilla/mux"
)

type Product struct {
  Name     string `json:"Name"`
  Quantity int    `json:"quantity"`
}

var products []Product

func DefaultHandler(w http.ResponseWriter, r *http.Request) {
  fmt.Fprintf(w, "Hello world\n")
}

func QueryAllProducts(w http.ResponseWriter, r *http.Request) {
  w.Header().Set("Content-Type", "application/json")
  json.NewEncoder(w).Encode(GetAllRows())
}

func CreateProduct(w http.ResponseWriter, r *http.Request) {
  w.Header().Set("Content-Type", "application/json")
  b, _ := ioutil.ReadAll(r.Body)

  var p Product
  json.Unmarshal(b, &p)        // convert JSON -> byte slice, store in Product p

  InsertRow(p)
  json.NewEncoder(w).Encode(p) // echo product back to the user
}

func DeleteProduct(w http.ResponseWriter, r *http.Request) {
  w.Header().Set("Content-Type", "application/json")
  b, _ := ioutil.ReadAll(r.Body)

  var p Product
  json.Unmarshal(b, &p)        // convert JSON -> byte slice, store in Product p

  RemoveRow(p)

  json.NewEncoder(w).Encode(p) // echo removed product back to the user
}

func UpdateProduct(w http.ResponseWriter, r *http.Request) {
  // mux variable input
  v := mux.Vars(r)
  k := v["name"]

  // read http body and store in Product p
  b, _ := ioutil.ReadAll(r.Body)
  var p Product
  json.Unmarshal(b, &p)

  for i, product := range products {
    if product.Name == k {
      products[i] = p
    }
  }
}

func QueryProduct(w http.ResponseWriter, r *http.Request) {
  w.Header().Set("Content-Type", "application/json")

  v := mux.Vars(r)
  k := v["name"]

  // Loop through all inventory, if product.Name == key, return JSON
  for _, product := range products {
    if product.Name == k {
      json.NewEncoder(w).Encode(product)
    }
  }
}


func main() {
  products = []Product{
    Product{Name:"oranges",  Quantity:25},
    Product{Name:"apples",   Quantity:53},
    Product{Name:"bananas",  Quantity:34},
  }

  r := mux.NewRouter().StrictSlash(true)

  r.HandleFunc("/", DefaultHandler)
  r.HandleFunc("/products", QueryAllProducts)
  r.HandleFunc("/product", CreateProduct).Methods("POST")
  r.HandleFunc("/product", DeleteProduct).Methods("DELETE")
  r.HandleFunc("/product/{name}", UpdateProduct).Methods("PUT")
  r.HandleFunc("/product/{name}", QueryProduct).Methods("GET")

  log.Fatal(http.ListenAndServe(":3000", r))
}

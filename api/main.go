package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Product struct {
	Name     string `json:"Name"`
	Quantity int    `json:"quantity"`
}

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
	json.Unmarshal(b, &p) // convert JSON -> byte slice, store in Product p

	InsertRow(p)
	json.NewEncoder(w).Encode(p) // echo product back to the user
}

func DeleteProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	b, _ := ioutil.ReadAll(r.Body)

	var p Product
	json.Unmarshal(b, &p) // convert JSON -> byte slice, store in Product p

	RemoveRow(p)
	json.NewEncoder(w).Encode(p) // echo removed product back to the user
}

func UpdateProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	b, _ := ioutil.ReadAll(r.Body)

	var p Product
	json.Unmarshal(b, &p) // convert JSON -> byte slice, store in Product p

	UpdateRow(p)
	json.NewEncoder(w).Encode(p) // echo updated product back to the user
}

// URL : /product?name=apple
func QueryProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	k := r.URL.Query().Get("name")

	p := QueryRow(k)
	json.NewEncoder(w).Encode(p)
}

func main() {
	r := mux.NewRouter().StrictSlash(true)

	r.HandleFunc("/", DefaultHandler)
	r.HandleFunc("/products", QueryAllProducts)
	r.HandleFunc("/product", CreateProduct).Methods("POST")
	r.HandleFunc("/product", DeleteProduct).Methods("DELETE")
	r.HandleFunc("/product", UpdateProduct).Methods("PUT")
	r.HandleFunc("/product", QueryProduct).Methods("GET")

	log.Fatal(http.ListenAndServe(":80", r))
}

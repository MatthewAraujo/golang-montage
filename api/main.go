package main

import (
	"encoding/json"
	"net/http"
)

type Customer struct {
	Name string `json:"name"`
	City string `json:"city"`
	Zip  string `json:"zip"`
}

func main() {

	//defines routes
	http.HandleFunc("/greet", greet)
	http.HandleFunc("/customers", getAllCustomers)

	http.ListenAndServe(":8080", nil)
}

func greet(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World!"))
}

func getAllCustomers(w http.ResponseWriter, r *http.Request) {
	customers := []Customer{
		{Name: "John Doe", City: "San Francisco", Zip: "94105"},
		{Name: "Jane Doe", City: "San Francisco", Zip: "94105"},
	}

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(customers)
}

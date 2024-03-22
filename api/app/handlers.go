package app

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"net/http"
)

type Customer struct {
	Name string `json:"name" xml:"name"`
	City string `json:"city" xml:"city"`
	Zip  string `json:"zip" xml:"zip"`
}

func greet(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World!"))
}

func getAllCustomers(w http.ResponseWriter, r *http.Request) {
	customers := []Customer{
		{Name: "John Doe", City: "San Francisco", Zip: "94105"},
		{Name: "Jane Doe", City: "San Francisco", Zip: "94105"},
	}

	if r.Header.Get("Content-Type") == "application/xml" {
		w.Header().Set("Content-Type", "application/xml")
		xml.NewEncoder(w).Encode(customers)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(customers)
}

func getCustomer(w http.ResponseWriter, r *http.Request) {
	customerId := r.PathValue("costumer_ID")

	customer := Customer{Name: "John Doe", City: "San Francisco", Zip: customerId}
	if r.Header.Get("Content-Type") == "application/xml" {
		w.Header().Set("Content-Type", "application/xml")
		xml.NewEncoder(w).Encode(customer)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(customer)
}

func createCustomer(w http.ResponseWriter, r *http.Request) {
	// var customer Customer
	// if r.Header.Get("Content-Type") == "application/xml" {
	// 	xml.NewDecoder(r.Body).Decode(&customer)
	// } else {
	// 	json.NewDecoder(r.Body).Decode(&customer)
	// }

	// w.Header().Set("Content-Type", "application/json")
	// json.NewEncoder(w).Encode(customer)
	fmt.Println("POST")
}

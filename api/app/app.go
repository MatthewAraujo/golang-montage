package app

import (
	"log"
	"net/http"
)

func Start() {

	router := http.NewServeMux()

	//defines routes
	router.HandleFunc("/greet", greet)
	router.HandleFunc("/customers", getAllCustomers)
	router.HandleFunc("/customer", createCustomer)
	router.HandleFunc("/customer/{costumer_ID}", getCustomer)

	log.Fatal(http.ListenAndServe(":8080", router))
}

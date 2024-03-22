package app

import (
	"log"
	"net/http"
)

func Start() {

	//defines routes
	http.HandleFunc("/greet", greet)
	http.HandleFunc("/customers", getAllCustomers)

	log.Fatal(http.ListenAndServe(":8080", nil))
}

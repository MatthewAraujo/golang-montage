package app

import (
	"log"
	"net/http"
)

func Start() {

	mux := http.NewServeMux()

	//defines routes
	mux.HandleFunc("/greet", greet)
	mux.HandleFunc("/customers", getAllCustomers)

	log.Fatal(http.ListenAndServe(":8080", mux))
}

package main

import "net/http"

func main() {

	//defines routes
	http.HandleFunc("/greet", greet)

	http.ListenAndServe(":8080", nil)
}

func greet(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World!"))
}

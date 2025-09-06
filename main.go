package main

import (
	"encoding/json"
	"net/http"

	"github.com/NeginSal/mini-ci-cd-go-demo.git/hello"
)

type response struct {
	Message string `json:"message"`
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name") 
	message := hello.Hello(name)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response{Message: message})
}

func main() {
	http.HandleFunc("/hello", helloHandler)
	http.ListenAndServe(":8080", nil)
}

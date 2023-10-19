package main

import (
	"fmt"
	"net/http"
	"github.com/gorilla/mux"
)

func HelloHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, "Hello, World!")
}

func HealthHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, "Healthy")
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/hello", HelloHandler)
	r.HandleFunc("/health", HealthHandler)

	http.ListenAndServe(":8080", r)
}

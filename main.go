package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	goMux := mux.NewRouter()
	goMux.HandleFunc("/", homeHandler)
	goMux.HandleFunc("/products", productsHandler)
	log.Fatal(http.ListenAndServe(":8080", goMux))
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Welcome to homepage")
}

func productsHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Welcome to products page")
}

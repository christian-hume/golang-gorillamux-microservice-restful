package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprint(w, `{"message": "Hola desde Gorilla Mux"}`)
	}).Methods("GET")

	fmt.Println("Servidor corriendo en http://localhost:8080 🚀")
	log.Fatal(http.ListenAndServe(":8080", router))
}
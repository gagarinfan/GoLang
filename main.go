package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	log.Println("Starting application")

	router := mux.NewRouter()

	router.HandleFunc("/people", getAllHandler).Methods("GET")
	router.HandleFunc("/people/{id}", getPersonByIDEndpoint).Methods("GET")
	router.HandleFunc("/people", createNewPersonHandler).Methods("POST")
	log.Fatal(http.ListenAndServe(":8000", router))

}

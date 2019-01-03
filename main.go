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
	log.Println("Creating table")
	check := checkIfTableExist()
	if check == false {
		err := createTable()
		if err != nil {
			log.Fatal("Table does not exist but can not create one. Caused by:", err)
		}
		log.Println("Starting router")
	} else {
		log.Println("Table already exist. Starting router")
	}
	router := mux.NewRouter()

	router.HandleFunc("/people", getAllHandler).Methods("GET")
	router.HandleFunc("/people/{id}", getPersonByIDHandler).Methods("GET")
	router.HandleFunc("/people/{id}", deletePersonByIDHandler).Methods("POST")
	router.HandleFunc("/people", createNewPersonHandler).Methods("POST")
	router.HandleFunc("/people/update/{id}", updatePersonByIDHandler).Methods("POST")
	log.Fatal(http.ListenAndServe(":8000", router))

}

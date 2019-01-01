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

	log.Fatal(http.ListenAndServe(":8000", router))

}

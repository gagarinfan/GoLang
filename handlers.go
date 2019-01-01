package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func getAllHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Gethering data")
	persons, _ := getAll()
	json.NewEncoder(w).Encode(persons)
}

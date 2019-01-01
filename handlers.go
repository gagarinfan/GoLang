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

func createNewPersonHandler(w http.ResponseWriter, r *http.Request) {
	var people People

	err := json.NewDecoder(r.Body).Decode(&people)
	if err != nil {
		log.Println("Request body does not match model !", err)
	} else {
		log.Println("Creating new person")
		id := createNewPerson(people.Name, people.Surname, people.Age, people.Hobby, people.Active)
		log.Println("Created new person with ID: ", id)
	}
}

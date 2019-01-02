package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func getAllHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Gathering data")
	persons, _ := getAll()
	json.NewEncoder(w).Encode(persons)
}

func createNewPersonHandler(w http.ResponseWriter, r *http.Request) {
	var people People

	err := json.NewDecoder(r.Body).Decode(&people)
	if err != nil {
		log.Println("Can not save new person. Caused by: ", err)
	} else {
		log.Println("Creating new person")
		id := createNewPerson(people.Name, people.Surname, people.Age, people.Hobby, people.Active)
		log.Println("Created new person with ID: ", id)
	}
}

func getPersonByIDHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])
	person, _ := getPersonByID(id)
	json.NewEncoder(w).Encode(person)

}

func deletePersonByIDHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])
	deletePersonByID(id)
	log.Println("Deleted person with ID: ", id)

}

func updatePersonByIDHandler(w http.ResponseWriter, r *http.Request) {
	var person People
	params := mux.Vars(r)
	err := json.NewDecoder(r.Body).Decode(&person)
	if err != nil {
		log.Println("Error. Caused by: ", err)
	}

	id, _ := strconv.Atoi(params["id"])
	updatePersonByID(id, person.Name, person.Surname, person.Age, person.Hobby, person.Active)
}

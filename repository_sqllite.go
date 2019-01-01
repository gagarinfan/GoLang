package main

import (
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

func getAll() (persons []People, err error) {
	db, err := gorm.Open("sqlite3", "/Users/michal/Desktop/Go/GoLang/gorm.db")
	if err != nil {
		log.Fatal("Can not connect to DB. Caused by ", err)
	}
	log.Println("Connected to DB")

	db.Find(&persons)
	defer db.Close()
	return persons, err
}

/*
func getPersonByID() {

}
*/
func createNewPerson(name string, surname string, age int, hobby string, active int) (id int) {
	db, err := gorm.Open("sqlite3", "/Users/michal/Desktop/Go/GoLang/gorm.db")
	if err != nil {
		log.Fatal("Can not connect to DB. Caused by ", err)
	}
	log.Println("Connected to DB")
	new_person := &People{
		Name:    name,
		Surname: surname,
		Age:     age,
		Hobby:   hobby,
		Active:  active,
	}
	err_create := db.Create(new_person).Error
	if err_create != nil {
		log.Fatal("Can not create new person. Caused by: ", err_create)
	}

	defer db.Close()
	return new_person.ID
}

/*
func updatePersonByID() {

}
func deletePersonByID() {

}
*/
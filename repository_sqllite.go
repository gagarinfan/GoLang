package main

import (
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

func checkIfTableExist() bool {
	var people People
	db, err := gorm.Open("sqlite3", "gorm.db")
	if err != nil {
		log.Fatal("Can not connect to DB. Caused by ", err)
	}
	exist := db.HasTable(&people)
	return exist
}

func createTable() (errcreated error) {
	db, err := gorm.Open("sqlite3", "gorm.db")
	if err != nil {
		log.Fatal("Can not connect to DB. Caused by ", err)
	}
	var people People
	errcreated = db.CreateTable(&people).Error
	if errcreated != nil {
		log.Println(errcreated)
	}
	return errcreated
}

func getAll() (persons []People, err error) {
	db, err := gorm.Open("sqlite3", "gorm.db")
	if err != nil {
		log.Fatal("Can not connect to DB. Caused by ", err)
	}
	log.Println("Connected to DB")

	db.Find(&persons)
	defer db.Close()
	return persons, err
}

func getPersonByID(id int) (person People, err error) {
	db, err := gorm.Open("sqlite3", "gorm.db")
	if err != nil {
		log.Fatal("Can not connect to DB. Caused by ", err)
	}
	db.Find(&person, "id = ?", id)
	defer db.Close()
	return person, err
}

func createNewPerson(name string, surname string, age int, hobby string, active int) (id int) {
	db, err := gorm.Open("sqlite3", "gorm.db")
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

func updatePersonByID(id int, name string, surname string, age int, hobby string, active int) {
	db, err := gorm.Open("sqlite3", "gorm.db")
	if err != nil {
		log.Fatal("Can not connect to DB. Caused by ", err)
	}
	var person People
	person.ID = id
	new_data := &People{
		Name:    name,
		Surname: surname,
		Age:     age,
		Hobby:   hobby,
		Active:  active,
	}
	errupdate := db.Model(&person).Where("id = ?", id).Update(new_data).Error
	if errupdate != nil {
		log.Println("Can not create new person. Caused by: ", errupdate)
	} else {
		log.Println("Updated ID: ", id)
	}
	defer db.Close()
}

func deletePersonByID(id int) {
	db, err := gorm.Open("sqlite3", "gorm.db")
	if err != nil {
		log.Fatal("Can not connect to DB. Caused by ", err)
	}
	var person People
	person.ID = id
	db.Delete(&person)
}

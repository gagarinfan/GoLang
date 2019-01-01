package main

//People type
type People struct {
	//gorm.Model
	Name    string
	Surname string
	Age     int
	Hobby   string
	Active  int
	ID      int `gorm:"primary_key"`
}

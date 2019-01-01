package test

import (
	"fmt"
)

type Person struct {
	name    string
	surname string
	age     int
	hobby   string
	active  bool
	id      string
}

type People interface {
	GetName() string
}

func createUser(name, surname, hobby, id string, age int, active bool) *Person {
	return &Person{name: name, surname: surname, age: age, hobby: hobby, active: active, id: id}
}

func (p *Person) GetName() string {
	return (p.name + " " + p.surname)
}

func describe(p People) {
	fmt.Println("Hi", p.GetName())
}

func loadToMap(m map[string]Person, p Person) {
	m[p.id] = p
	fmt.Println(m)
}
func main() {

	m := make(map[string]Person)
	fmt.Println("Insert your data here:")
	var name, surname, hobby string
	var age int
	fmt.Scanln(&name)
	fmt.Scanln(&surname)
	fmt.Scanln(&age)
	fmt.Scanln(&hobby)
	active := true
	id := "100" //for now let it be hardcoded

	newMan := createUser(name, surname, hobby, id, age, active)
	describe(newMan)
	loadToMap(m, *newMan)
}

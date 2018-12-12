package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Person struct {
	name    string
	surname string
	age     int
	hobby   string
	active  bool
	id      int
}

type people interface {
	Describe() string
}

func createUser(name, surname, hobby string, age, id int, active bool) *Person {
	return &Person()
}

func getUsername(s string) {
	fmt.Printf("Hi %v, how are you?\n", s)
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Insert your data here:")
	name, _ := reader.ReadString('\n')
	name = strings.Replace(name, "\n", "", -1)
	getUsername(name)
}

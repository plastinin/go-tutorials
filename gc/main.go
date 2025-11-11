package main

import (
	"fmt"
	"time"
)

func main() {
	people := make([]Person, 0, 10_000_000)
	before := time.Now()
	for i := 0; i < 10_000_000; i++ {
		people = append(people, MakePerson("Fred", "Williamson", 25))
	}
	fmt.Println("time spend:", time.Since(before))
}

type Person struct {
	FirstName string
	LastName  string
	Age       int
}

func MakePerson(firstName, lastName string, age int) Person {
	return Person{
		FirstName: firstName,
		LastName:  lastName,
		Age:       age,
	}
}

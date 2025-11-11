package main

import "fmt"

type Person struct {
	FirstName string
	LastName  string
	Age       int
}

func MakePerson(FirstName string, LastName string, Age int) Person {
	return Person{
		FirstName: FirstName, LastName: LastName, Age: Age,
	}
}

func MakePersonPointer(FirstName string, LastName string, Age int) *Person {
	return &Person{
		FirstName: FirstName, LastName: LastName, Age: Age,
	}
}

func main() {
	P1 := MakePerson("Пластинин", "Артем", 26)
	P2 := MakePersonPointer("Иванов", "Евгений", 26)
	fmt.Println(P1)
	fmt.Println(*P2)
}

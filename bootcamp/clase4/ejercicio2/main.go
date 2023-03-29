package main

import "fmt"

type Person struct {
	ID          int
	Name        string
	DateOfBirth string
}

type Employee struct {
	ID       int
	Position string
	Person
}

func main() {
	p1 := Person{
		ID:          1,
		Name:        "Laura",
		DateOfBirth: "1998-01-01",
	}

	e1 := Employee{
		ID:       1,
		Position: "Developer",
		Person:   p1,
	}

	e1.PrintEmployed()
}

func (e Employee) PrintEmployed() {
	fmt.Println(e)
}

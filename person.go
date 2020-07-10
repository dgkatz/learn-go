package main

import "fmt"

type Person struct {
	firstName string
	lastName string
}

func (person Person) FullName() string {
	return fmt.Sprintf("%s %s", person.firstName, person.lastName)
}

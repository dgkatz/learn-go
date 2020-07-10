package main

import "fmt"

const greetingTemplate string = "Hello, %v %v for the %v%v time"

type Person struct {
	firstName string
	lastName string
}

func countSuffix(count int) (countSuffix string){
	lastDigit := count % 10
	if lastDigit == 1 && count != 11 {
		countSuffix = "st"
	} else if lastDigit == 2 && count != 12{
		countSuffix = "nd"
	} else if lastDigit == 3 && count != 13{
		countSuffix = "rd"
	} else {
		countSuffix = "th"
	}
	return
}

func createGreeting(person Person, count int) (greeting string){
	greeting = fmt.Sprintf(
		greetingTemplate,
		person.firstName,
		person.lastName,
		count,
		countSuffix(count))
	return
}

func printPerson(person Person, count int) {
	for i := 1; i <= count; i++ {
		fmt.Println(createGreeting(person, i))
	}
}

func main() {
	dan := Person{"Dan", "Katz"}
	printPerson(dan, 30)
}
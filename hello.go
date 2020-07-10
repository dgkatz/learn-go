package main

import "fmt"

const greetingTemplate string = "Hello, %v for the %v%v time"

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

func printName(name string, count int) {
	for i := 1; i <= count; i++ {
		greeting := fmt.Sprintf(greetingTemplate, name, i, countSuffix(i))
		fmt.Println(greeting)
	}
}

func main() {
	printName("Dan", 30)
}
package main

import (
	"fmt"
)

func main() {
	// longhand
	var fullName string = "Inw Ion"
	var age = 20
	const lasName = "lolipop"
	// shorthand
	fullName2 := "mokun"

	fmt.Println(fullName)
	fmt.Printf("test %s age = %d\n", fullName, age)
	fmt.Printf("lastName %s \n", lasName)
	fmt.Println("Hi go from ", fullName)
	fmt.Println("Hi go from jrBoy")

	fmt.Printf("Name %s\n", fullName2)
	fullName2 = "kohung"
	fmt.Printf("Name %s\n", fullName2)

	// calc grade
	var score int = 62
	var grade string
	
	if score >= 90 {
		grade = "A"	
	} else if score >= 80 {
		grade = "B"
	} else if score >= 70 {
		grade = "C"
	} else if score >= 60 {
		grade = "D"
	} else {
		grade = "F"
	}
	
	fmt.Printf("Your score = %d and Your grade = %s\n", score, grade)
	
	// DayofWeek 
	var dayOfWeek int = 7
	
	switch dayOfWeek {
	case 1:
		fmt.Println("Monday")
	case 2:
		fmt.Println("Tuesday")
	case 3:
		fmt.Println("Wednesday")
	case 4:
		fmt.Println("Thursday")
	case 5:
		fmt.Println("Friday")
	case 6:
		fmt.Println("Saturday")
	case 7:
		fmt.Println("Sunday")
	case 8:
		fmt.Println("Invalid Day")

	}
}
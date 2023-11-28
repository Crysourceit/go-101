package main

import (
	"fmt"
)

func sayHello() {
	fmt.Println("Hello World")
}

func sayHello2(name string) {
	fmt.Printf("Hello %s\n", name)
}

func add(a int, b int) int {
	return a + b
  }

// Define a struct type
type Student struct {
	Name    string
	Weight  int
	Height  int
	Grade   string
}
// Use With Map
type Student2 struct {
	Name    string
	Weight  int
	Height  int
	Grade   string
}

// Define a struct type
type Person struct {
	Name    string
	Age     int
	Address Address
}

// Another struct type used in Person
type Address struct {
	Street  string
	City    string
	ZipCode int
}

//--------------------------------------
// Define the Student struct
type StudentReceiver struct {
	Firstname string
	Lastname  string
}
// Method with a receiver of type Student
// This method returns the full name of the student
func (s StudentReceiver) FullName() string {
	return s.Firstname + " " + s.Lastname
}
// Define a struct type
type Rectangle struct {
	Length float64
	Width  float64
}
// Method with a receiver of type Rectangle
func (r Rectangle) Area() float64 {
	return r.Length * r.Width
}


//------------------------------------
// Speaker interface
type Speaker interface {
	Speak() string
}
// Dog struct
type Dog struct {
	Name string
}

// Dog's implementation of the Speaker interface
func (d Dog) Speak() string {
	return "Woof!"
}

// Person struct
type PersonInter struct {
	Name string
}

// Person's implementation of the Speaker interface
func (p PersonInter) Speak() string {
	return "Hello!"
}

// function that accepts Speaker interface
func makeSound(s Speaker) {
	fmt.Println(s.Speak())
}


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
	
	// For Loop
	i := 1
	for {
	  fmt.Printf("number: %d\n\n", i)
	  i++
	  if i >= 10 {
		break
	  }
	}
	
	// Do While Loop
	i2 := 1
	for {
	  fmt.Printf("number: %d\n\n", i2)
	  i2++
	  if i2 >= 10 {
		break
	  }
	}
	
	// While Loop
	i3 := 1
	for i3 < 10 {
		fmt.Printf("number: %d\n", i3)
		i3++
	}
	

	// Array
	var myArray [3]int // An array of 3 integers
	myArray[0] = 10    // Assign values
	myArray[1] = 20
	myArray[2] = 30
	fmt.Println(myArray) // Output: [10 20 30]
	
	// Use With Loop
	var myArray2 [3]int // An array of 3 integers
	myArray2[0] = 10    // Assign values
	myArray2[1] = 20
	myArray2[2] = 30
	// Looping through the array
	for i := 0; i < len(myArray2); i++ {
	  fmt.Println(myArray2[i])
	}
	
	// Slice
	mySlice := []int{10, 20, 30, 40, 50} // A slice of integers
	fmt.Println(mySlice)          // Output: [10 20 30 40 50]
	fmt.Println(len(mySlice))     // Length of the slice: 5
	fmt.Println(cap(mySlice))     // Capacity of the slice: 5
	// Slicing a slice
	subSlice := mySlice[1:3]      // Slice from index 1 to 2
	fmt.Println(subSlice)         // Output: [20 30]


	// Appending data to the slice
	var mySlice2 []int // Declared but not initialized
	// Appending data to the slice
	mySlice2 = append(mySlice2, 10)
	mySlice2 = append(mySlice2, 20, 30)
	fmt.Printf("Appending %d\n\n", mySlice2) // Output: [10 20 30]
	
	// Converting Array to Slice
	var myArray3 [3]int // An array of 3 integers
	myArray3[0] = 10    // Assign values
	myArray3[1] = 20
	myArray3[2] = 30
	// Converting array to slice
	mySlice3 := myArray3[:]
	// Resizing slice by appending new elements
	mySlice3 = append(mySlice3, 40, 50)
	fmt.Println(mySlice3) // Output will show a slice with 5 elements: [10 20 30 40 50]
	

	// Map
	myMap := make(map[string]int)
	// Add key-value pairs to the map
	myMap["apple"] = 5
	myMap["banana"] = 10
	myMap["orange"] = 8
	// Access and print a value for a key
	fmt.Println("Apples:", myMap["apple"])
	// Update the value for a key
	myMap["banana"] = 12
	// Delete a key-value pair
	delete(myMap, "orange")
	// Iterate over the map
	for key, value := range myMap {
	fmt.Printf("%s -> %d\n", key, value)
	}
	// Checking if a key exists
	val, ok := myMap["pear"]
	if ok {
	fmt.Println("Pear's value:", val)
	} else {
	fmt.Println("Pear not found inmap")
	}
	

	// Struct
	// Create an instance of the Student struct
	var student1 Student
	student1.Name = "Mikelopster"
	student1.Weight = 60
	student1.Height = 180
	student1.Grade = "F"
	// Print struct values
	fmt.Println(student1)
	
	// Use With Array
	// Create an array of Student structs
	var students [3]Student
	// Initialize the first student
	students[0] = Student{
		Name:   "Mikelopster",
		Weight: 60,
		Height: 180,
		Grade:  "F",
	}
	// Initialize the second student
	students[1] = Student{
		Name:   "Alice",
		Weight: 55,
		Height: 165,
		Grade:  "A",
	}
	// Initialize the third student
	students[2] = Student{
		Name:   "Bob",
		Weight: 68,
		Height: 175,
		Grade:  "B",
	}
	// Print array of structs
	fmt.Println(students)
	
	//Use With Map
	// Create a map with string keys and Student struct values
	students2 := make(map[string]Student2)
	// Add Student structs to the map
	students2["st01"] = Student2{Name: "Mikelopster", Weight: 60, Height: 180, Grade: "F"}
	students2["st02"] = Student2{Name: "Alice", Weight: 55, Height: 165, Grade: "A"}
	students2["st03"] = Student2{Name: "Bob", Weight: 68, Height: 175, Grade: "B"}
	// Print the map
	fmt.Println("Students Map:", students2)
	// Access and print a specific student by key
	fmt.Println("Student st01:", students2["st01"])
	

	// Use Struct together
	// Create an instance of the Person struct
	var person Person
	person.Name = "Alice"
	person.Age = 30
	person.Address = Address{
		Street:  "123 Main St",
		City:    "Gotham",
		ZipCode: 12345,
	}
	// Alternative way to initialize a struct
	bob := Person{
		Name: "Bob",
		Age:  25,
		Address: Address{
			Street:  "456 Elm St",
			City:    "Metropolis",
			ZipCode: 67890,
		},
	}
	// Print struct values
	fmt.Println(person)
	fmt.Println(bob)
	


	// Function
	sayHello()
	sayHello()
	sayHello()
	//parameter function 
	sayHello2("mike")
	sayHello2("mikelopster")
	sayHello2("tanitphon")
	//return function
	number1 := 3
	number2 := 5
	sumNumber := add(number1, number2)
	fmt.Println(sumNumber)
	

	//Receiver Method
	StudentRec := StudentReceiver{
		Firstname: "Mike",
		Lastname:  "Lopster",
	}
	// Call the FullName method on the Student instance
	fullNameRec := StudentRec.FullName()
	fmt.Println("Full Name of the student:", fullNameRec)
	// Receiver exp2
	rect := Rectangle{Length: 10, Width: 5}
	// Call the Area method on Rectangle instance
	area := rect.Area()
	fmt.Println("Area of rectangle:", area)


	// Interface
	dog := Dog{Name: "Buddy"}
	personIn := PersonInter{Name: "Alice"}

	makeSound(dog)
	makeSound(personIn)


}
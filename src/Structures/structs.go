package Structures

import (
	"fmt"
	"strconv"
)

// Define person struct
type Person struct {
	// Define properties of struct
	firstName 	string
	lastName 	string
	city 		string
	gender 		string
	age 		int
}

func Exec() {
	structs()
	valueReceivers()  // For methods that just do calculations and aren't changing things
	pointerReceivers()  // For methods where you are changing things
}

func structs() {
	// Init. person using Person struct
	person1 := Person {
		firstName: "Tim",
		lastName: "Mit",
		city: "Florida",
		gender: "M",
		age: 102,
	}
	//person1 := Person {firstName: "Patrick", lastName: "Archer", city: "Tempe", gender: "M", age: 22}
	//person1 := Person {
	//	"Patrick",
	//	"Archer",
	//	"Tempe",
	//	"M",
	//	22,
	//}

	fmt.Println(person1)
	fmt.Println(person1.age, person1.lastName)
	person1.age++
	fmt.Println("New age:", person1.age)
}

// Value receiver function -- acts a lot like "this.x" functionality
func (p Person) greeting() string {
	return "Hello, my name is " + p.firstName + " and I am " + strconv.Itoa(p.age) + " years old."
}

func valueReceivers() {
	person2 := Person {
		firstName: "Cheryl",
		lastName: "Star",
		city: "Tempe",
		gender: "F",
		age: 45,
	}
	fmt.Println(person2.greeting())
}

// Pointer receiver function
func (p *Person) hasBirthday() {
	p.age++
}

// Pointer receiver function
func (p *Person) getsMarried(newLastName string) {
	if p.gender == "M" {
		return
	} else if p.gender == "F" {
		p.lastName = newLastName
	} else {
		return
	}
}

func pointerReceivers() {
	person3 := Person {
		firstName: "Joe",
		lastName: "Shmoe",
		city: "Phoenix",
		gender: "M",
		age: 22,
	}
	fmt.Println(person3.age)
	person3.hasBirthday()
	fmt.Println(person3.age)

	person4 := Person {
		firstName: "Sam",
		lastName: "Smith",
		city: "Los Angeles",
		gender: "F",
		age: 65,
	}
	fmt.Println(person4)
	person4.getsMarried("McDonald")
	fmt.Println(person4)
}

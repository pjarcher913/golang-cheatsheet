package Structures

import (
	"fmt"
)

// Define person struct
type Person struct {
	// Define properties of struct
	firstName string
	lastName string
	city string
	gender string
	age int
}

func Exec() {
	structs()
	valueReceivers()  // For methods that just do calculations and aren't changing things
	pointerReceivers()  // For methods where you are changing things
}

func structs() {
	// Init. person using Person struct
	person1 := Person {
		firstName: "Patrick",
		lastName: "Archer",
		city: "Tempe",
		gender: "M",
		age: 22,
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
	return "Hello, my name is " + p.firstName + "and I am " + string(p.age)
}

func valueReceivers() {
	person2 := Person {
		firstName: "Patrick",
		lastName: "Archer",
		city: "Tempe",
		gender: "M",
		age: 22,
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
		firstName: "Patrick",
		lastName: "Archer",
		city: "Tempe",
		gender: "M",
		age: 22,
	}
	fmt.Println(person3.age)
	person3.hasBirthday()
	fmt.Println(person3.age)

	person4 := Person {
		firstName: "Marissa",
		lastName: "Archer",
		city: "Flagstaff",
		gender: "F",
		age: 65,
	}
	fmt.Println(person4)
	person4.getsMarried("Smith")
	fmt.Println(person4)
}

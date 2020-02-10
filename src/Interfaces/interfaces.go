/*
INTERFACES:
	Data types that represent a set of method signatures for structs.
	Define interfaces that have methods that will be used by other structs.
	Same concept as java interfaces.
*/

package Interfaces

import (
	"fmt"
	"math"
)

func Exec() {
	interfaces()
}

// Define interface
type Shape interface {
	// Put everything you want the interface to implement in here
	area() float64
}

// Use interface with different structs and methods
type Circle struct {
	x, y, radius float64
}
type Rectangle struct {
	w, h float64
}
func (c Circle) area() float64 {
	return math.Pi * c.radius * c.radius
}
func (r Rectangle) area() float64 {
	return r.w * r.h
}
func getArea(s Shape) float64 {
	return s.area()
}

func interfaces() {
	circle := Circle{x: 0, y: 0, radius: 5}
	rectangle := Rectangle{w: 10, h: 5}
	fmt.Printf("Circle Area: %f\n", getArea(circle))
	fmt.Printf("Rectangle Area: %f\n", getArea(rectangle))
}

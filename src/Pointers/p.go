package Pointers

import (
	"fmt"
)

func Exec() {
	pointers()
}

func pointers() {
	a := 5
	b := &a  // assign "b" to pointer at memory address of "a"
	fmt.Println(a, b)
	/*
	OUTPUT:
		5 0xc000012078
	NOTE:
		typeof(&a) == *int
	*/
	fmt.Println(a, *b)
	/*
	OUTPUT:
		5 5
	*/

	// Change value w/ pointer
	*b = 10
	// "b" is set to mem. addr. of "a", so we actually changed a's value
	// (which "b" is set to the value at the address of "a")
	fmt.Println(a)
}

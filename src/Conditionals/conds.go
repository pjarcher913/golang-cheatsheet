package Conditionals

import (
	"fmt"
)

func Exec() {
	ifStatements()
	switches()
}

func ifStatements() {
	x := 5
	y := 10

	if (x < y) {
		fmt.Printf("x=%d is less than y=%d\n", x, y)
	} else if x > y {
		fmt.Printf("x=%d is greater than y=%d\n", x, y)
	} else {
		fmt.Printf("x=%d is equal to y=%d\n", x, y)
	}
}

func switches() {
	color := "purple"
	switch color {
	case "red":
		fmt.Println("Color is red.")
	case "blue":
		fmt.Println("Color is blue.")
	default:
		fmt.Println("Color is not red or blue.")
	}
}

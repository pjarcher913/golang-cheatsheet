package Loops

import (
	"fmt"
)

func Exec() {
	forLoops_LONG()
	forLoops_SHORT()
	fizzBuzz()
}

func forLoops_LONG() {
	i := 1
	for i <= 3 {
		fmt.Println(i)
		//i = i + 1
		i++
	}
}

func forLoops_SHORT() {
	for i := 1; i <= 3; i++ {
		fmt.Println(i)
	}
}

func fizzBuzz() {
	for i := 1; i <= 15; i++ {
		if i % 15 == 0 {
			fmt.Printf("FizzBuzz (%d)\n", i)
		} else if i % 3 == 0 {
			fmt.Printf("Fizz (%d)\n", i)
		} else if i % 5 == 0 {
			fmt.Printf("Buzz (%d)\n", i)
		}
	}
}

package Closures

import (
	"fmt"
)

func Exec() {
	sum := adder()
	for i := 0; i < 5; i++ {
		fmt.Println(sum(i))
	}
}

func adder() func(int) int{
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}
}

package Functions

import "fmt"

func Exec() {
	funcs()
}

func funcs() {
	fmt.Println(greeting("Jerry"))
	fmt.Println(getSum(71, 22))
}

// func <FUNCTION_NAME>(<PARAMS>) <RETURN_TYPE> { }
func greeting(name string) string {
	return "Hello " + name
}

func getSum(num1 int, num2 int) int {
	return num1 + num2
}

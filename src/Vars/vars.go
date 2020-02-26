package Vars

import(
	"fmt"
)

const TEST_STR string = "THIS IS A TEST STRING 123!"

func Exec() {
	vars()
	printGlobalStrVar()
}

func vars() {
	/*
		Main var types:
			string
			bool
			int int8 int 16 int32 int64
			uint uint8 uint16 uint32 uint64
			byte (alias for uint8)
			rune (alias for int32)
			float32 float64
			complex64 complex128
	*/
	// STRING
	var name_1 string = "Terry_1"
	name_2 := "Terry_2"		// := is shorthand for declaring and initializing a var
	name_3, name_4 := "Terry_3", "Terry_4"
	fmt.Println("name_1: ", name_1,
		"\nname_2: ", name_2,
		"\nname_3: ", name_3,
		"\nname_4: ", name_4)

	// INT
	var age_1 int32 = 22
	age_2 := 22
	fmt.Println("age_1: ", age_1, "\nage_2: ", age_2)
	fmt.Println("Var type of age_1: %T", age_1)
	fmt.Println("Var type of age_2: %T", age_2)

	// BOOL
	var isCool_1 bool = true
	isCool_2 := true
	const isCool_3 = true	// const works the same way as other languages
	fmt.Println("isCool_1: ", isCool_1, "\nisCool_2: ", isCool_2, "\nisCool_3: ", isCool_3)

	// TODO: Examples of uint, byte, rune, float32/64, complex64/128 vars
	// UINT
	//...

	// BYTE
	//...

	// RUNE
	//...

	// FLOAT
	//...

	// COMPLEX
	//...
}

func printGlobalStrVar() {
	fmt.Println(TEST_STR)
}

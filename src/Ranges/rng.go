package Ranges

import (
	"fmt"
)

func Exec() {
	rng_with_slice()
	rng_with_map()
}

func rng_with_slice() {
	ids := []int{33, 76, 54, 23, 11, 2}

	// Loop thru ids using range
	for i, id := range ids {
		fmt.Printf("%d - ID: %d\n", i, id)
	}
	/*
	OUTPUT:
		0 - ID: 33
		1 - ID: 76
		2 - ID: 54
		3 - ID: 23
		4 - ID: 11
		5 - ID: 2
	*/

	// Loop thru ids using range w/out index
	for _, id := range ids {
		fmt.Printf("ID: %d\n", id)
	}
	/*
		OUTPUT:
			ID: 33
			ID: 76
			ID: 54
			ID: 23
			ID: 11
			ID: 2
	*/
}

func rng_with_map() {
	emails := map[string]string{
		"Bob": "bob@gmail.com",
		"Sally": "sally@gmail.com",
	}

	// for <KEY>, <VALUE> := range <MAP>
	for k, v := range emails {
		fmt.Printf("%s: %s\n", k, v)
	}
	/*
	OUTPUT:
		Bob: bob@gmail.com
		Sally: sally@gmail.com
	*/
}

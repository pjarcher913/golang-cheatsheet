package Arrays_Slices

import (
	"fmt"
)

func Exec() {
	arrays()
	slices()
}

func arrays() {
	var fruitArr [2]string
	fruitArr[0] = "Apple"
	fruitArr[1] = "Orange"
	fmt.Println("fruitArr: ", fruitArr, "\nfruitArr[0]: ", fruitArr[0], "\nfruitArr[1]: ", fruitArr[1], "\nlen(fruitArr): ", len(fruitArr))
}

func slices() {
	//fruitArr := [2]string{"Apple", "Orange"}
	fruitArr := []string{"Apple", "Orange", "Grape"}
	fmt.Println("New fruitArr: ", fruitArr)
	fmt.Println("fruitArr[0:2] slice: ", fruitArr[0:2])
	fmt.Println("fruitArr[1:3] slice: ", fruitArr[1:3])
}

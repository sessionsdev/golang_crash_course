package main

import (
	"fmt"
)

func main() {
	// Arrays
	var fruitArr [2]string

	// Assign values
	fruitArr[0] = "Apple"
	fruitArr[1] = "Orange"

	fmt.Println(fruitArr)

	// Declare and assign
	petArr := [2]string{"Dog", "Cat"}
	fmt.Println(petArr)

	fruitSlice := []string{"Apple", "Orange", "Grape", "Cherry"}
	fmt.Println(fruitSlice)

	fmt.Println(len(fruitSlice))
	fmt.Println(fruitSlice[1:3])
}

package main

import (
	"fmt"
)

func main() {
	a := 5

	// assigning a pointer to a
	b := &a

	fmt.Println(a, b)
	fmt.Printf("%T, %T\n", a, b)

	// User * to read val from mem address
	fmt.Println(*b)

	// Change val with pointer
	*b = 10
	fmt.Println(a)
}

package main

import "fmt"

func main() {

	// typing
	var name string = "Jonny"
	var age int = 31

	// constant variable
	const isCool bool = true

	//shorthand
	email := "jonny@sessionsdev.com"
	size := 1.3

	fmt.Println(name, age, isCool, email, size)
	fmt.Printf("%T\n", name)

}

// TYPES //
// string
// bool
// int
// int int8 int16 int32 int64
// uint uint8 uint16 uint 32 uint64 uintptr
// byte - alias for uint8
// rune - alias for int32
// float32 float64
// complex64 complex128

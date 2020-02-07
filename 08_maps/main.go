package main

import (
	"fmt"
)

func main() {
	// // define map
	// emails := make(map[string]string)

	// //Assign kv
	// emails["jon"] = "jon@test.com"
	// emails["jonny"] = "jonny@test.com"
	// emails["jonnytsunami"] = "jonnytsunami@test.com"

	// Declare and add kv
	emails := map[string]string{
		"jon":          "jon@test.com",
		"jonny":        "jonny@test.com",
		"jonnytsunami": "jonnytsunami@test.com",
	}

	fmt.Println(emails)
	fmt.Println(len(emails))
	fmt.Println(emails["jon"])

	delete(emails, "jon")
	fmt.Println(emails)
	fmt.Println(len(emails))
}

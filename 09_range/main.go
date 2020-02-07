package main

import "fmt"

func main() {

	// create slice of ids
	ids := []int{33, 21, 12, 21, 674, 3, 24, 141, 23}

	// loop through with range
	for i, id := range ids {
		fmt.Printf("%d - ID: %d\n", i, id)
	}

	// without using index
	for _, id := range ids {
		fmt.Printf("ID: %d\n", id)
	}

	// range to sum numbers
	sum := 0
	for _, id := range ids {
		sum += id
	}
	fmt.Println("sum: ", sum)

	emails := map[string]string{
		"jon":          "jon@test.com",
		"jonny":        "jonny@test.com",
		"jonnytsunami": "jonnytsunami@test.com",
	}

	for k, v := range emails {
		fmt.Printf("%s: %s\n", k, v)
	}

	for k := range emails {
		fmt.Println("Name: " + k)
	}

}

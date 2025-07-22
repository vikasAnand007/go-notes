package main

import "fmt"

func main() {
	// simple if-else

	age := 11
	if age >= 18 {
		fmt.Println("adult")
	} else if age >= 12 {
		fmt.Println("teenager")
	} else {
		fmt.Println("kid")
	}

	// declare variable inside if

	if age2 := 25; age2 >= 18 {
		fmt.Println("adult", age2)
	} else if age2 >= 12 {
		fmt.Println("teenager", age2)
	} else {
		fmt.Println("kid", age2)
	}

	// fmt.Println(age2) // age2 is only accessable in the scope of above if-else construct
}

// NOTE: Go do not have ternary operator

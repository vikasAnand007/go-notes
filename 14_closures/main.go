package main

import "fmt"

// Closure in GO works like same as javascript

// this function returns a function with integer type return value
func closureExample() func() int {
	// this variable will be passed as closure along with the function because it is used in that function
	var num int = 0

	return func() int {
		num = num + 1
		return num + 1
	}
}

func main() {
	incrementFunction := closureExample()
	// since closure is used here, values logged here will be 2,3,4 because it is refering to num variable.
	// if closule were not used, then logged values will be 1,1,1
	fmt.Println(incrementFunction(), incrementFunction(), incrementFunction())
}

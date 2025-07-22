package main

import "fmt"

// for is is the only looping construct in GO. We use different variation of for loop to achieve different functionalities
func main() {
	// Clissic for loop
	fmt.Println("Simple for loop")
	for i := 0; i < 10; i++ {
		if i == 2 {
			continue
		}
		if i > 8 {
			break
		}
		fmt.Println(i)
	}

	// while loop
	fmt.Println("While loop")
	j := 0
	for j < 5 {
		fmt.Println(j)
		j = j + 1
	}

	// infinite loop
	// fmt.Println("infinite loop")
	// for {
	// 	fmt.Println("infinite")
	// }

	// itrating with help of range
	fmt.Println("itrating with help of range")
	for k := range 10 {
		fmt.Println("range", k)
	}
}

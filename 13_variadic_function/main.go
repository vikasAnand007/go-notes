package main

import "fmt"

// Variadic function can take any number of argumnet.
func sum(nums ...int) int {
	total := 0
	for _, val := range nums {
		total = total + val
	}
	return total
}

func main() {
	fmt.Println(sum(12, 3, 45, 6, 4, 3, 2))

	// OR
	nums := []int{3, 5, 4, 1, 0, 9}
	fmt.Println(sum(nums...))
}

package main

import "fmt"

func some(nums ...int) []int {
	newNums := make([]int, len(nums))
	for ele, val := range nums {
		newNums[ele] = val * 2
	}
	return newNums
}

func main() {
	vals := []int{1, 2, 3, 4, 5, 6, 7, 8}
	fmt.Println(some(vals...))
}

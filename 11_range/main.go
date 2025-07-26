package main

import "fmt"

func main() {
	// Use of range for itrating over slices
	slice1 := []int{10, 15, 12, 11}
	sum := 0
	for ind, elem := range slice1 {
		fmt.Println(ind, slice1[ind], elem)
		sum = sum + elem
	}
	fmt.Println("sum =", sum)

	// Use of range for itrating over maps
	map1 := map[string]string{"fname": "Vikas", "lname": "Anand"}
	for key, val := range map1 {
		fmt.Println(key, val, map1[key])
	}

	// Use of range for itrating over string
	str1 := "My name is vikas"
	for v1, v2 := range str1 {
		// v1 -> starting byte of rune (Will see later about this)
		// v2 -> unicode code of character
		fmt.Println(v1, v2, string(v2))
	}
}

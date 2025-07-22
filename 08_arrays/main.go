package main

import "fmt"

// MOTE: Array in go is NOT widely used. Because in array we need to give the length in advance. But in real world cases we are mostly un aware of the array length. Hence in GO we mostly use slices.

func main() {
	var strings [4]int = [4]int{1, 2, 3, 4}
	strings[2] = 20

	fmt.Println(strings, strings[2])
	fmt.Println("length of array is", len(strings))
	//  NOTE: unassigned places in an array is filled with zeroed value, which is different for different types.
	// string -> ""
	// int -> 0
	// float -> 0
	// bool -> false

	// shorthand
	arr2 := [3]string{"a", "b"}
	fmt.Println(arr2)

	// 2D array
	arr3 := [2][3]int{{11, 12, 13}, {21, 22, 23}}
	fmt.Println(arr3, arr3[1][2])
}

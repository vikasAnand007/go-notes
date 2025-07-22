package main

import "fmt"

func main() {
	// assigning variable with type - [not necessary assign value at same time]
	var str1 string = "vikas"
	var num1 int = 20
	var bool1 bool
	bool1 = false
	fmt.Println(str1, num1, bool1)

	// assigning variable type automatically infered - [must assign value at same time]
	var str2 = "vikas"
	var num2 = 20
	var bool2 = false
	fmt.Println(str2, num2, bool2)

	// assigning variable : shorthand - [must assign value at same time]
	str3 := "vikas"
	num3 := 20
	bool3 := false
	fmt.Println(str3, num3, bool3)
}

var str0 string = "vikas" // can use this syntex outside a function
var num0 = 20             // can use this syntex outside a function
// bool0 := "vikas"	// CANNOT use this syntex outside a function

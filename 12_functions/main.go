package main

import "fmt"

// simple
func add(num1 int, num2 int) int {
	return num1 + num2
}

// if multiple argument have same type
func subtract(num1, num2 int) int {
	return num1 - num2
}

// if no return value
func printInfo(fName, lName string, age int) {
	fmt.Println(fName, lName, "is", age, "years old.")
}

// Functions in go can return multiple values
func mulVal() (string, string, int, bool) {
	return "Vikas", "Anand", 27, true
}

// Functions in GO are first class citizen. They can be passed/receive as argument in another parameter. And can be assigned to a variable.
func processIt(fff func(a int) int) int {
	value := fff(5)
	return value
}

func main() {
	result := add(1, 2)
	fmt.Println(result)
	fmt.Println(subtract(1, 2))
	printInfo("Vikas", "Anand", 27)
	fname, _, age, isAdmin := mulVal()
	fmt.Println(fname, age, isAdmin)

	// Function as a first class citizen
	multiplyBy10 := func(arg1 int) int {
		return arg1 * 10
	}

	fmt.Println(processIt(multiplyBy10))
}

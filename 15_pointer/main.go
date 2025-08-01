package main

import "fmt"

// Pointers are the address of location of any variable stored in computer memory
// for example we have a variable v1
// then,
// for any variable v1, &v1 = gives pointer of that varaible
// for any pointer p1, *p1 = gives value stored on that pointer

// by default behaviour when variable (not a function) is passed in another function then a new copy of variable is created which do not share any reference with the original variable. Hence when its value is chaged, it has no effect in original variable.
func changeNum(num int) {
	num = 5
	fmt.Println("in changeNum function", num) // this prints 5
}

func changePointerNum(num *int) {
	*num = 6
	fmt.Println("in changePointerNum function", *num) // this prints 6
}

func main() {
	str := "vikas"
	fmt.Println("variable =", str)
	fmt.Println("pointer =", &str)

	// example
	num := 1
	// passing variable as argument
	changeNum(num)
	fmt.Println("in main function", num) // this prints 1

	// passing pointer as argument
	changePointerNum(&num)
	fmt.Println("in main function", num) // this prints 6

}

package main

import (
	"fmt"
	"time"
)

func main() {
	// simple switch
	switch i := 3; i {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	case 3:
		fmt.Println("three")
	default:
		fmt.Println("other")
	}

	// multiple condition switch
	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("Weekend")
	default:
		fmt.Println("Workday")
	}

	// type switch
	var someVariable interface{} = "vikas" // interface{} is like "any" type
	switch someVariable.(type) {
	case int:
		fmt.Println("integer")
	case string:
		fmt.Println("string")
	case bool:
		fmt.Println("boolean")
	default:
		fmt.Println("other")
	}
}

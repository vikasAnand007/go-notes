package main

import "fmt"

const name = "vikas"

// name = "anand" // invalid (cannot re assign)

// constant grouping
const (
	city = "Jaipur"
	pin  = 302033
)

func main() {
	const age = 20
	// age = 21 // invalid (cannot re assign)

	fmt.Println("Age of", name, "is", age, ".He lives in", city, pin)
}

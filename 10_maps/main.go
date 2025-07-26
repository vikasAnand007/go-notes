package main

import (
	"fmt"
	"maps"
)

// maps === hash / objects / dictionaries
func main() {
	// uninitialised map

	m1 := make(map[string]string)
	m1["name"] = "Vikas"
	m1["lastName"] = "Anand"
	fmt.Println("m1", m1)
	fmt.Println("m1", m1["name"])
	fmt.Println("m1", m1["lastName"])
	fmt.Println("m1", m1["city"]) // return zeroed value i.e, empty string ("")

	m2 := make(map[string]int)
	m2["value1"] = 100
	fmt.Println("m2", m2)
	fmt.Println("m2", m2["value1"])
	fmt.Println("m2", m2["value2"]) // return zeroed value i.e, empty string (0)

	// length of map
	fmt.Println(len(m1), len(m2))

	// delete key from map
	delete(m1, "lastName")
	fmt.Println("m1", m1)

	// clear map
	clear(m2)
	fmt.Println("m2", m2)

	// intialised slice
	m3 := map[string]bool{"isAdmin": false, "isVerified": true}
	fmt.Println("m3", m3)

	// Accessing a key-value in GO returns 2 things
	// first - actual value / zeroed value
	// second - boolean (is any actual value returned, false if zeroed value)
	m3r1, m3r2 := m3["isAdmin"]
	fmt.Println(m3r1, m3r2)

	m3r3, m3r4 := m3["isUser"]
	fmt.Println(m3r3, m3r4)

	// Comparing map
	m4 := map[string]int{"v1": 1, "v2": 2}
	m5 := map[string]int{"v1": 1, "v2": 2}
	m6 := map[string]int{"v1": 1, "v2": 2, "v3": 3}

	fmt.Println(maps.Equal(m4, m5), maps.Equal(m5, m6))
}

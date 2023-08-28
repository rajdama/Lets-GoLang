package main

import "fmt"

func main() {

	// Using make keyword to form a map from string to int
	ages := make(map[string]int)

	ages["John"] = 37
	ages["Mary"] = 35

	fmt.Println(ages)

	// Defining and instantiating map at same time
	ages2 := map[string]int{
		"John": 37,
		"Mary": 35,
	}
	fmt.Println(ages2)

	elem := ages2["John"] // Getting an element from map
	fmt.Println(elem)

	delete(ages2, "John") // Deleting an element from map
	fmt.Println(ages2)

	val, ok := ages2["Mary"] // Checking an element in map
	fmt.Println(val, ok)
}

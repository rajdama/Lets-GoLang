package main
import "fmt"

/*
1]A slice is a lightweight data structure that wraps around an underlying array. 
2]It represents a segment of that array, allowing you to work with a variable-length sequence of elements.
3]Slices provide a way to create, modify, and manipulate sequences of elements without having to worry about memory 
 allocation and management.
*/

func main()  {
	
	// Creates a slice of int with length 5 and capacity 10
	mySlice1 := make([]int, 5, 10) 
	fmt.Println(mySlice1)

	// Creates a slice of int with length 5 and capacity 5
	mySlice2 := make([]int, 5) 
	fmt.Println(mySlice2)

	// Creates array of size 5 with respective elements
	mySlice3 := []int{1, 2, 3, 4, 5}
	fmt.Println(mySlice3)
	
	// Slicing an Existing Array or Slice:
	array := [5]int{1, 2, 3, 4, 5}
	mySlice4 := array[1:4] // Creates a slice containing elements 2, 3, and 4
	fmt.Println(mySlice4)
}
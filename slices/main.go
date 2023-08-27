package main
import "fmt"

/*
1]A slice is a lightweight data structure that wraps around an underlying array. 
2]It represents a segment of that array, allowing you to work with a variable-length sequence of elements.
3]Slices provide a way to create, modify, and manipulate sequences of elements without having to worry about memory 
 allocation and management.
4]Slice is a reference of array
5]A variadic function in programming is a function that can accept a variable number of arguments of the same type
*/



func sum(nums ...int) int {				// A variadic function receives the variadic arguments as a slice.
	// nums is just a slice

	sum := 0
	for i:= 0; i < len (nums); i++{
	sum = sum + nums[i]
	}

	return sum
}

func mul(nums ...int) int {				// A variadic function receives the variadic arguments as a slice.
	// nums is just a slice

	mul := 1
	for i:= 0; i < len (nums); i++{
	mul = mul*nums[i]
	}

	return mul
}


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

	// Append function
	mySlice5 := []int{1, 2, 3, 4, 5}
	mySlice5 = append(mySlice5,8,9)
	fmt.Println(mySlice5)

	// Variadic slice arguments
	total := sum(1,2,3)
	fmt.Println(total)

	// Spread operator
	myNums := []int{3,4,5}
	mulResult := mul(myNums...)
	fmt.Println(mulResult)

	

}
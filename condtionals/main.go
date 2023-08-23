package main

import "fmt"

func main() {

	// one way to write if-else

	num1 := 3
	num2 := 2

	if num1 > 2 {
		fmt.Println("True")
	} else if num2 > 2 {
		fmt.Println("True")
	}else{
		fmt.Println("False")
	}

	// second way to write if-else (optimized) --> if INITIAL_STATEMENT; CONDITION{ / CODE GOES HERE / }

	if num3 := num1 + num2; num3 == 5 {
		fmt.Println("True")
	}
	
}
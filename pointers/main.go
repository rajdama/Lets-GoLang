package main
import "fmt"


func add (x, y *int) int {
	return *x + *y
}



func main()  {

	x := 7
	y := &x            // y contains memory address of x
	fmt.Println(y)

	z := *y			   // dereference operator for getting actual x value
	fmt.Println(z)

	sum := add(&x,&z)	// passing address in a function
	fmt.Println(sum)

	
}
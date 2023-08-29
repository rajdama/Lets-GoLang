package main
import "fmt"

func main()  {

	x := 7
	y := &x            // y contains memory address of x
	fmt.Println(y)

	z := *y			   // dereference operator for getting actual x value
	fmt.Println(z)
}
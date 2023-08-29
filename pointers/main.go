package main
import "fmt"

func main()  {

	x := 7
	y := &x            // y contains memory address of x
	fmt.Println(*y)
}
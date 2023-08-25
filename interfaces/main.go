package main
import "fmt"

/*
1]Interfaces are collections of method signatures. A type "implements" an
interface if it has all of the methods of the given interface defined on it.
2]A type implements an interface by implementing its methods. Unlike in
many other languages, there is no explicit declaration of intent, there is no
"implements" keyword.
*/


// Single interface
type shape interface {
	area() float64
	perimeter() float64
}

type rect struct {				//rect implements interface shape as it defines methods area and perimeter
	width, height float64
}
	
func (r rect) area() float64 {
	return r.width * r.height
}
	
func (r rect) perimeter() float64 {
	return 2*r.width + 2*r.height
}

// Multiple Interface
type expense interface {
	cost() float64
}

type printer interface {
	print()
}

type email struct{
	subscribed bool
	body string
}

func (e email) cost() float64 {
	if e.subscribed {
		return 10
	} else{
		return 0
	}
}

func (e email) print() {
	fmt.Println(e.body)
}

func main(){

	myRect := rect{}			// Implements interface shape
	myRect.width = 5
	myRect.height = 10
	fmt.Println(myRect.area())

	myemail := email{}			// Implements interface cost and print
	myemail.subscribed = true
	myemail.body = " Hello !"
	fmt.Println(myemail.cost())
	myemail.print()

}
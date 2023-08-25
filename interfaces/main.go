package main
import "fmt"

/*Interfaces are collections of method signatures. A type "implements" an
interface if it has all of the methods of the given interface defined on it.*/

type shape interface {
	area() float64
	perimeter() float64
}

type rect struct {					//rect implements interface shape as it defines methods area and perimeter
	width, height float64
}
	
func (r rect) area() float64 {
	return r.width * r.height
}
	
func (r rect) perimeter() float64 {
	return 2*r.width + 2*r.height
}

func main(){
	myRect := rect{}
	myRect.width = 5
	myRect.height = 10
	fmt.Println(myRect.area())
}
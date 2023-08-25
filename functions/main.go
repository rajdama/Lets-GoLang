package main
import "fmt"

// ------------Args data type-----Return type---//
func concat  (s1 string, s2 string)  string     {
return s1 + s2
}


// Here last argument's data type is every arguments data type
func add(x, y int) int {
	return x + y
}

// Returning multiple values
func getPoint1()  (int,int) {
	return 10, 4
}

// Returning multiple values with label
func getPoint2()  (x,y int) {
	x = 15
	y = 16
	return x,y
}

// Multiple Returns
func divide(dividend, divisor int) (int, error) {
	if divisor == 0 {
	return 0, errors.New("Can't divide by zero")
	}
	return dividend/divisor, nil
}

func main() {
	fmt.Println(concat("hello"," world"))
	fmt.Println(add(1,2))
	x, _ := getPoint1() // Ignores y
	fmt.Println(x)
}
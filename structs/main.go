package main
import "fmt"

/*
1]Use the type keyword to define a new type.
2]Provide a name for your struct (e.g., MyStruct in the example above).
3]Inside curly braces {}, list the fields of the struct. Each field has a name and a type separated by a space. You can have multiple fields of different types.
*/


// Simple struct
type car struct {
	Make string
	Model string
	Height int
	Width int
}

// Nested struct
type messageToSend struct {
	message string
	sender user
	recipient user

}
type user struct {
	name string
	number int
}

// Embedded struct (kinda inheritence)
type dog struct {
	name string
	age int
}

type animal struct {
	// "car" is embeddded, so the definition of a
	// "truck" now also additionally contains all
	// of the fields of the car struct
	dog
	legs int
}

// Methods on struct
type rect struct {
	width int
	height int
}
	
func (r rect) area() int {  // area has a receiver of (r rect)
	return r.width * r.height
}



func main() {

	myCar := car{}		    //All values initialized to default value i.e 0

	myCar.Make = "honda" 	//Assigning values
	myCar.Height = 5

	fmt.Println(myCar)  	// Prints simple struct

	msg := messageToSend{}
	msg.sender.name = "X"
	msg.recipient.name = "Y"

	fmt.Println(msg) 		// Prints nested struct

	student := struct {  	//Anonymous structs prevent you from re-using a struct definition you never intended to re-use
        name      string
        branch    string
    }{
        name:      "Raj",
        branch:    "CSE",
    }

	fmt.Println(student) 	// Prints anonymus struct

	myDog := animal{}

	myDog.name = "Buzo"
	myDog.age = 10
	myDog.legs = 4

	fmt.Println(myDog)		// Prints embedded struct

	r := rect{
		width: 5,
		height: 10,
		}

	fmt.Println(r.area())   // Prints method on struct


}
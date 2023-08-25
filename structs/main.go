package main
import "fmt"


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

	fmt.Println(myDog)


}
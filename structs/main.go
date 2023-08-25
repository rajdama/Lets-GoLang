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

func main() {
	myCar := car{}		 //All values initialized to default value i.e 0

	myCar.Make = "honda" //Assigning values
	myCar.Height = 5

	fmt.Println(myCar.Height)

	msg := messageToSend{}
	msg.sender.name = "X"
	msg.recipient.name = "Y"

	fmt.Println(msg.sender.name, msg.recipient.name )
}
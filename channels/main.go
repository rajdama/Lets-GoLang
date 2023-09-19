package main

import "fmt"

func main() {
	userch1 := make(chan string) //unbuffered channel i.e capacity is only 1

	go func() {
		userch1 <- "raj" //blocking
	}()

	user := <-userch1

	fmt.Println(user)

	userch2 := make(chan string, 2) //buffred channel i.e capacity is more than 1

	userch2 <- "ram" //non-blocking because channel not full

	user = <-userch2

	fmt.Println(user)

	userch2 <- "rohan" //non-blocking because channel not full
	userch2 <- "roy"
	// userch2 <- "sam" this will result in blocking

}

func sendMessage(msgch chan<- string) { // This specifies we are going to only write to channel msgch and not read anything

	msgch <- "hello"

}

func readMessage(msgch <-chan string) { // This specifies we are going to only read from channel msgch and not write anything

	msg := <-msgch
	fmt.Println(msg)

}

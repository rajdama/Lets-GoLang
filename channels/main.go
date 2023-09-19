package main

import "fmt"

func main() {
	userch1 := make(chan string) //buffered channel i.e capacity is only 1

	go func() {
		userch1 <- "raj" //blocking
	}()

	user := <-userch1

	fmt.Println(user)

	userch2 := make(chan string, 2) //unbuffred channel i.e capacity is more than 1

	userch2 <- "ram" //non-blocking because channel not full

	user = <-userch2

	fmt.Println(user)

	userch2 <- "rohan" //non-blocking because channel not full
	userch2 <- "roy"
	// userch2 <- "sam" this will result in blocking

}

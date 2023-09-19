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

	go func() {
		userch2 <- "ram" //non-blocking
	}()

	user = <-userch2

	fmt.Println(user)
}

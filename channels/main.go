package main

import "fmt"

func main() {
	userch := make(chan string) //buffered channel i.e capacity is only 1

	go func() {
		userch <- "raj" //blocking
	}()

	user := <-userch

	fmt.Println(user)
}

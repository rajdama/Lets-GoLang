package main

import "fmt"

func main() {
	userch := make(chan string)

	go func() {
		userch <- "raj"
	}()

	user := <-userch

	fmt.Println(user)
}

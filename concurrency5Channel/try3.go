package main

import (
	"fmt"
)

func main() {
	c := make(chan string, 2)
	c <- "hello"
	c <- "world"

	// will not block, untill the channel is full.
	msg := <-c
	fmt.Println(msg)

	// not a new variable, so := will not work, have to use =
	msg = <-c
	fmt.Println(msg)

	// but if we put three data, it will block, and it will not work.
}

package main

import (
	"fmt"
)

func main() {
	c := make(chan string, 2)
	c <- "hello"
	// will not block, untill the channel is full.
	msg := <-c
	fmt.Println(msg)
}

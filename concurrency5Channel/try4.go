package main

import (
	"fmt"
)

func main() {

	msg := <-c
	fmt.Println(msg)

	msg = <-c
	fmt.Println(msg)

	// Note: this will not work, as we have not defined the "c" yet.

	c := make(chan string, 2)
	c <- "hello"
	c <- "world"
}

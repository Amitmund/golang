// package declaration
package main

// import of module
import (
	"fmt"
)

// Creating a function called "count"
// which take two argument of "string" type, of "thing" and "channel c"
// whatever passed as an argument in the "count function", get in to channel "c"
// Once the loop is done or work is done, we need to close the channel "c"
func count(thing string, c chan string) {
	for i := 1; i <= 5; i++ {
		c <- thing
	}
	close(c)
}

// main function declaration.
// over here we also define the channel "c" with "make" function with argument "chan" and "type of channel"
func main() {
	c := make(chan string)
	go count("sheep", c)

	for msg := range c {
		fmt.Println(msg)
	}
}

// youtube link: https://www.youtube.com/watch?v=LvgVSSpwND8
// git: https://github.com/jakewright/tutorials/blob/master/go/02-go-concurrency/03-channels.go
//

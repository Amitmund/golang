package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	var count = 10

	for count > 0 {
		fmt.Println(count)
		time.Sleep(time.Second)
		stats := rand.Intn(10)
		fmt.Printf("Random Number: %v", stats)
		if stats == 7 {
			break
		}
		count--
	}

	if count == 0 {
		fmt.Println("Liftoff!")
	} else {
		fmt.Println("Launch failed.")
	}
}

// Actually I am not really able to understand the same.

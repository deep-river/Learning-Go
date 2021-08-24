package main

import "fmt"

func main() {
	// make a channel with buffer size of 2
	messages := make(chan string, 2)
	messages <- "buffered"
	messages <- "channel"

	fmt.Println(<-messages)
	fmt.Println(<-messages)
}
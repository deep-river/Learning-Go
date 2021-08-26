// This example also shows that it’s possible to close a non-empty channel but still have the remaining values be received.
package main

import "fmt"

func main() {
	queue := make(chan string, 2)
	queue <- "one"
	queue <- "two"
	close(queue)
	// Because we closed the channel above, the iteration terminates after receiving the 2 elements.
	for element := range queue {
		fmt.Println(element)
	}
}
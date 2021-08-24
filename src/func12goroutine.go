/*
goroutine: go协程
goroutine的栈内存可变，初始一般为 2KB ，随着需求可以扩大达到 1GB
*/
package main

import (
	"fmt"
	"time"
)

// function f()
func f(from string) {
	for i := 0; i < 3; i++ {
		fmt.Println(from, ": ", i)
	}
}

func main() {
	// direct call of f()
	f("direct call of f()")
	// call f() concurrently with goroutine
	go f("gorountine")
	// call an anonymous function with goroutine 
	go func(msg string) {
		fmt.Println(msg)
	}("goroutine with anonymous fucntion")

	time.Sleep(time.Second)
	fmt.Println("done")
}
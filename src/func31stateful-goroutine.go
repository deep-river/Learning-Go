/*
In the previous example we used explicit locking with mutexes to synchronize access to shared state across multiple goroutines. Another option is to use the built-in synchronization features of goroutines and channels to achieve the same result.

This channel-based approach aligns with Go’s ideas of sharing memory by communicating and having each piece of data owned by exactly 1 goroutine.
*/
package main

import (
	"fmt"
	"math/rand"
	"sync/atomic"
	"time"
)

/*
In this example our state will be owned by a single goroutine. This will guarantee that the data is never corrupted with concurrent access.

In order to read or write that state, other goroutines will send messages to the owning goroutine and receive corresponding replies.

These readOp and writeOp structs encapsulate those requests and a way for the owning goroutine to respond
*/
type readOp struct {
	key int
	resp chan int
}

type writeOp struct {
	key int 
	val int 
	resp chan bool 
}

func main() {
	var readOps uint64 
	var writeOps uint64

	reads := make(chan readOp)
	writes := make(chan writeOp)

	// the goroutine that owns the state, which is a map
	go func() {
		var state = make(map[int]int)
		for {
			select {
			case read := <-reads:
				read.resp <- state[read.key]
			case write := <-writes:
				state[write.key] = write.val
				write.resp <- true
			}
		}
	}()

	for r := 0; r < 100; r++ {
		go func() {
			for {
				read := readOp {
					key: rand.Intn(5),
					resp: make(chan int)}
				reads <- read
				<- read.resp
				atomic.AddUint64(&readOps, 1)
				time.Sleep(time.Millisecond)
			}
		}()
	}
	
	for w := 0; w < 10; w++ {
		go func() {
			for {
				write := writeOp{
					key: rand.Intn(5),
					val: rand.Intn(100),
					resp: make(chan bool)}
				writes <- write
				<-write.resp 
				atomic.AddUint64(&writeOps, 1)
				time.Sleep(time.Millisecond)
			}
		}()
	}

	time.Sleep(time.Second)

	readOpsFinal := atomic.LoadUint64(&readOps)
	fmt.Println("readOps: ", readOpsFinal)
	writeOpsFinal := atomic.LoadUint64(&writeOps)
	fmt.Println("writeOps: ", writeOpsFinal)
}
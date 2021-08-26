/*
In the previous example we saw how to manage simple counter state using atomic operations. For more complex state we can use a mutex to safely access data across multiple goroutines.
*/
package main

import (
	"fmt"
	"math/rand"
	"sync"
	"sync/atomic"
	"time"
)

func main() {
	var state = make(map[int]int)
	// This mutex will synchronize access to state
	var mutex = &sync.Mutex{}
	// We’ll keep track of how many read and write operations we do
	var readOps uint64
	var writeOps uint64

	// Here we start 100 goroutines to execute repeated reads against the state, once per millisecond in each goroutine.
	for r := 0; r < 100; r++ {
		go func() {
			total := 0
			for {
				key := rand.Intn(5)
				mutex.Lock()
				total += state[key]
				mutex.Unlock()
				atomic.AddUint64(&readOps, 1)
				// Wait a bit between reads.
				time.Sleep(time.Millisecond)
			}
		}()
	}

	// We’ll also start 10 goroutines to simulate writes, using the same pattern we did for reads
	for w := 0; w < 10; w++ {
		go func() {
			for {
				key := rand.Intn(5)
				val := rand.Intn(100)
				mutex.Lock()
				state[key] = val 
				mutex.Unlock()
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

	mutex.Lock()
	fmt.Println("state: ", state)
	mutex.Unlock()
}
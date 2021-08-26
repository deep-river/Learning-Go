/*
Timers represent a single event in the future.
*/
package main

import (
	"fmt"
	"time"
)

func main() {
	// You tell the timer how long you want to wait, and it provides a channel that will be notified at that time. This timer will wait 2 seconds.
	timer1 := time.NewTimer(2 * time.Second)
	// It's a channel, bro! 

	<-timer1.C 
	fmt.Println("Timer 1 fired")
	
	// One reason a timer may be useful is that you can cancel the timer before it fires. Here’s an example of that.
	timer2 := time.NewTimer(time.Second * 2)
	go func() {
		<-timer2.C 
		fmt.Println("Timer 2 fired")
	}()
	// timer2 stopped 
	stop2 := timer2.Stop()
	if stop2 {
		fmt.Println("Timer 2 stopped")
	}
	// If you just wanted to wait, you could have used time.Sleep.
	// Here we give the timer2 enough time to fire, if it ever was going to, to show it is in fact stopped.
	time.Sleep(2 * time.Second)
}
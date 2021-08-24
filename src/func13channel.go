package main

import "fmt"

func main() {
	// create new channel with [make(chan value-type)]
	messages := make(chan string)
	// send value into channel using [channel <- value]
	go func() { messages <- "ping" }()
	// receive value from the channel using [receiver <- channel]
	msg := <-messages
	/* 上面三行代码不能写成以下形式：
	messages := make(chan string)
	messages <- "ping"
	msg := <-messages
	因为channel是用来给不同goroutine通信的，以上代码在同一个协程内同时发送与接受，因而会造成死锁
	*/
	fmt.Println(msg)
	// 关闭channel
	defer close(messages)
}
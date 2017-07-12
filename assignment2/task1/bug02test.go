package main

import "fmt"
import "time"

var counter = 0
var ch = make(chan int)
// This program should go to 11, but sometimes it only prints 1 to 10.
func main() {
	for i:=1; i<=1000; i++ {
		loop()
	}
	for t:=1; t<=2; t++ {
		fmt.Println("t=", t)
		time.Sleep(time.Second)
	}
	close(ch)
	fmt.Println("counter: ",counter)
}

func loop() {
	go Print()
	for j := 1; j <= 11; j++ {
		ch <- j
	}
}

// Print prints all numbers sent on the channel.
// The function returns when the channel is closed.
func Print() {
	count := 0
	for range ch { // reads from channel until it's closed
		count += 1
	}
	if count == 11 {
		counter += 1
	} else {
		fmt.Println("fail")
	}
}

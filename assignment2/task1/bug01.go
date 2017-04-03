package main

import "fmt"

func main() {
	ch := make(chan string)
	// Send string inside go-routine instead of in main function
	go func() {
		ch <- "Hello world!"
	}()
	fmt.Println(<-ch)
}

package main

import "fmt"

// Add adds the numbers in a and
// sends the result on res.
func Add(a []int, res chan<- int) {
	sum := 0
	length := len(a)
	for i := 0; i < length; i++ {
		sum += a[i]
	}
	res <- sum
}

func main() {
	a := []int{7, 2, 8, -9, 4, 0}
	n := len(a)
	ch := make(chan int)
	go Add(a[:n/2], ch)
	go Add(a[n/2:], ch)

	x, y := <-ch, <-ch

	fmt.Println("x =", x)
	fmt.Println("y =", y)
	fmt.Println("sum =", x+y)
}

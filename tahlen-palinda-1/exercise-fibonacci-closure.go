package main

import "fmt"

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func() int {
	fib := []int{0, 0, 1}
	return func() int {
		fib[0] = fib[1]
		fib[1] = fib[2]
		fib[2] += fib[0]
		return fib[0]
	}
}

func Testfibonacci(t *testing.T) {
	res := []int{0, 1, 1, 2, 3, 5, 8, 13, 21, 34}
	f := fibonacci()
	test := 0
	for i := 0; i < 10; i++ {
		test := f()
		if test != res[i] {
			t.Error("wrong answer")
		}
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 20; i++ {
		fmt.Println(f())
	}
}

package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) float64 {
	z1 := 1.0
	var z2 float64
	delta := 1.0
	deltaLimit := math.Pow(10, -6)

	for delta > deltaLimit {
		z2 = z1 - (math.Pow(z1, 2)-x)/(z1*2) // Newton's method
		delta = math.Abs(z2 - z1)
		z1 = z2 // set z1 for next loop
	}
	return z1
}

func main() {
	x := float64(3)

	fmt.Println("Sqrt:     ", Sqrt(x))
	fmt.Println("math.Sqrt:", math.Sqrt(x))
}

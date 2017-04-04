package main

import (
	"golang.org/x/tour/pic"
	"math/rand"
)

func Pic(dx, dy int) [][]uint8 {
	slice := make([][]uint8, dy)
	for i := 0; i < dy; i++ {
		slice[i] = make([]uint8, dx)
		for j := 0; j < dx; j++ {
			slice[i][j] = uint8(rand.Uint32() / 4)
		}
	}
	return slice
}

func main() {
	pic.Show(Pic)
}

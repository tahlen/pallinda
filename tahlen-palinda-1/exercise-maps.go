package main

import (
	"golang.org/x/tour/wc"
	"strings"
)

func WordCount(s string) map[string]int {
	vs := strings.Fields(s)
	stringLength := len(vs)
	retMap := make(map[string]int)
	for i := 0; i < stringLength; i++ {
		retMap[vs[i]] += 1
	}
	return retMap
}

func main() {
	wc.Test(WordCount)
}

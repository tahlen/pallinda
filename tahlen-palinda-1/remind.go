package main

import (
	"fmt"
	"strconv"
	"time"
)

func Remind(text string, delay time.Duration) {
	for {
		currentTime := time.Now()
		hr := strconv.Itoa((currentTime.Hour()))
		if currentTime.Hour() < 10 {
			hr = "0" + hr
		}
		min := strconv.Itoa((currentTime.Minute()))
		if currentTime.Minute() < 10 {
			min = "0" + min
		}

		fmt.Printf("Klockan är " + hr + ":" + min)
		fmt.Println(": ", text)

		time.Sleep(delay)
	}
}

func main() {
	go Remind("Dags att äta", 3 * time.Hour)
	time.Sleep(1 * time.Second)
	go Remind("Dags att arbeta", 8 * time.Hour)
	time.Sleep(1 * time.Second)
	Remind("Dags att sova", 24 * time.Hour)
}

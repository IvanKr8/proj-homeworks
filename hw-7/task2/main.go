package main

import (
	"fmt"
	"math/rand"
	"time"
)

func GenerateRandomInt(ch1 chan int, ch2 chan int, done chan bool) {
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 5; i++ {
		ch1 <- rand.Intn(100)
	}
	close(ch1)
	minInt, maxInt := <-ch2, <-ch2
	fmt.Println()
	fmt.Println("Min:", minInt)
	fmt.Println("Max:", maxInt)
	done <- true
}

func minMaxFunc(ch chan int, ch2 chan int) {
	var minInt, maxInt int
	first := true
	fmt.Printf("Received from minMaxFunc: ")
	for val := range ch {
		fmt.Printf("%d ", val)
		if first {
			minInt = val
			maxInt = val
			first = false
		} else {
			if val < minInt {
				minInt = val
			}
			if val > maxInt {
				maxInt = val
			}
		}
	}
	ch2 <- minInt
	ch2 <- maxInt
	close(ch2)
}

func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)
	done := make(chan bool)

	go GenerateRandomInt(ch1, ch2, done)
	go minMaxFunc(ch1, ch2)

	<-done
}

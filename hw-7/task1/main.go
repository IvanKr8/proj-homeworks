package main

import (
	"fmt"
	"math/rand"
	"time"
)

func GenerateRandomInt(integers chan int) {
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 5; i++ {
		integers <- rand.Int()
	}
	close(integers)
}

func AverageInt(integers chan int, avInt chan int) {
	var CountInt int
	var IntSum int
	CountInt = 0
	IntSum = 0
	fmt.Printf("numbers: ")
	for val := range integers {
		fmt.Printf("%d ", val)
		CountInt++
		IntSum += val
	}
	fmt.Println()
	fmt.Println("---------------------")
	AverageInteger := IntSum / CountInt
	avInt <- AverageInteger
}

func PrintAverage(avInt chan int, done chan bool) {
	average := <-avInt
	fmt.Println("average:", average)
	done <- true
}

func main() {
	ch := make(chan int)
	avInt := make(chan int)
	done := make(chan bool)
	go GenerateRandomInt(ch)
	go AverageInt(ch, avInt)
	go PrintAverage(avInt, done)
	<-done
}

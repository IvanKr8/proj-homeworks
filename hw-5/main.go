package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

func SearchByWord(texts []string, word string) []string {
	index := make(map[string][]int)
	for i, text := range texts {
		words := strings.Fields(text)
		for _, word := range words {
			found := false
			word = strings.ToLower(word)
			if _, ok := index[word]; !ok {
				index[word] = make([]int, 0)
			}
			for _, idx := range index[word] {
				if idx == i+1 {
					found = true
					break
				}
			}
			if !found {
				index[word] = append(index[word], i+1)
			}
		}
	}
	word = strings.ToLower(word)
	lines, found := index[word]
	if !found {
		return []string{"Nothing found"}
	}
	result := make([]string, len(lines))
	for i, line := range lines {
		result[i] = fmt.Sprintf("You have this word in %d offer", line)
	}
	return result
}

func fileOpen(fileName string) []string {
	var texts []string
	Smallfile, err := os.Open(fileName)
	if err != nil {
		texts = append(texts, fmt.Sprintf("Error opening file: %v", err))
		return texts
	}

	scanner := bufio.NewScanner(Smallfile)
	for scanner.Scan() {
		texts = append(texts, scanner.Text())
	}
	return texts
}

func AverageCalc(array []float64) float64 {
	if array != nil {
		var numberSum float64
		for _, num := range array {
			numberSum = +num
		}
		averageNum := numberSum / float64(len(array))
		return averageNum
	} else {
		return 0
	}
}

func main() {
	smallTexts := fileOpen("small.txt")
	averageTexts := fileOpen("average.txt")
	bigTexts := fileOpen("big.txt")
	fmt.Println()
	word := "грати"
	startSearchSmall := time.Now()
	resultsSmall := SearchByWord(smallTexts, word)
	endSearchSmall := time.Now()

	startSearchAverage := time.Now()
	resultsAverage := SearchByWord(averageTexts, word)
	endSearchAverage := time.Now()

	startSearchBig := time.Now()
	resultsBig := SearchByWord(bigTexts, word)
	endSearchBig := time.Now()
	for _, result := range resultsSmall {
		fmt.Println("small:", result)
	}
	for _, result := range resultsAverage {
		fmt.Println("average:", result)
	}
	for _, result := range resultsBig {
		fmt.Println("big:", result)
	}
	smallTime := endSearchSmall.Sub(startSearchSmall).Seconds() * 1000
	averageTime := endSearchAverage.Sub(startSearchAverage).Seconds() * 1000
	bigTime := endSearchBig.Sub(startSearchBig).Seconds() * 1000
	var times []float64
	fmt.Printf("small: %.6f ms\n", smallTime)
	fmt.Printf("average: %.6f ms\n", averageTime)
	fmt.Printf("big: %.6f ms\n", bigTime)
	times = append(times, bigTime, smallTime, averageTime)
	fmt.Println(times)
	fmt.Println()
	fmt.Println(AverageCalc(times), "ms")
}

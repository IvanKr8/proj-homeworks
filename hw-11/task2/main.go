package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
)

var (
	wordReg = regexp.MustCompile(`(?m)^[бвгджзйклмнпрстфхцчшщБВГДЖЗЙКЛМНПРСТФХЦЧШЩ]\p{L}*[аеиоуюяієїґАЕИОУЮЯІЄЇҐ]$`)
)

func main() {
	readFile, err := os.Open("./file.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer readFile.Close()

	scanner := bufio.NewScanner(readFile)
	for scanner.Scan() {
		line := scanner.Text()
		words := strings.Fields(line)
		for _, word := range words {
			isValid := wordReg.MatchString(word)
			if isValid {
				fmt.Printf("%s ", word)
			}
		}
	}
}

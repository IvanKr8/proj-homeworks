package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
)

var (
	phoneReg = regexp.MustCompile(`(\+\d{1,2}\s?)?(\(\d{3}\)|\d{3})[\s\.-]?\d{3}[\s\.-]?\d{4}`)
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
		isValid := phoneReg.MatchString(scanner.Text())
		fmt.Printf("%t %s\n", isValid, scanner.Text())
	}
}

package main

import (
	"fmt"
	phoneRegPackage "github.com/IvanKr8/phoneReg"
)

func main() {
	var test []string = []string{
		"+380952434319",
		"1234567890",
		"(123) 456-7890",
		"(123)456-7890",
		"123-456-7890",
		"123.456.7890",
		"123 456 7890",
		"asddas",
		"131sadassg",
	}

	var test2 []string
	test2 = phoneRegPackage.PhoneReg(test)
	for _, val := range test2 {
		fmt.Println(val)
	}
}

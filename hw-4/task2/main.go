package main

import (
	"fmt"
	"sort"
)

type number struct {
	ID   int
	Name string
}

func RemoveDuplicates(NumStruct []number) []number {
	var numbers []number
	for _, obj := range NumStruct {
		found := false
		for _, i := range numbers {
			if i == obj {
				found = true
				break
			}
		}
		if !found {
			numbers = append(numbers, obj)
		}
	}
	return numbers
}

func main() {
	numbers := []number{
		{
			ID:   5,
			Name: "five",
		},
		{
			ID:   2,
			Name: "two",
		},
		{
			ID:   3,
			Name: "three",
		},
		{
			ID:   3,
			Name: "three",
		},
		{
			ID:   4,
			Name: "four",
		},
		{
			ID:   1,
			Name: "one",
		},
		{
			ID:   7,
			Name: "seven",
		},
		{
			ID:   7,
			Name: "seven",
		},
	}
	withoutDuplicates := RemoveDuplicates(numbers)
	fmt.Println("Without sort: ", withoutDuplicates)
	fmt.Println()
	sort.SliceStable(withoutDuplicates, func(i, j int) bool {
		return withoutDuplicates[i].ID < withoutDuplicates[j].ID
	})
	fmt.Println("With sort: ", withoutDuplicates)
}

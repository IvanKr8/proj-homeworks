package main

import "fmt"

type strring string

func typeTest(a strring) {
	fmt.Printf("%s", a)
}

func calculator(num1, num2 int) int {
	return num1 + num2
}

func main() {
	var (
		stringstr string
		integer   int
		float     float64
		hasBool   bool
	)
	fmt.Printf("%s %d %f %t \n", stringstr, integer, float, hasBool)

	// Біографія
	age := 15
	firstname := "Іван"
	lastname := "Крайч"
	description := "3 роки вчив front-end. Починаю вчити back.\nПочав вчити back з пайтону.\nGolang подобається своєю швидкістью, захистом, синтаксисом і тд.\n"
	stack := "HTML, CSS, Java Script, React.js, Bootstrap, jQuery"
	fmt.Printf("Моя біографія😀\n\n")
	fmt.Printf("Мене звати %s %s, я маю %d\n", firstname, lastname, age)
	fmt.Printf("%s\n", description)
	fmt.Printf("Мій стек в фронті: %s\n\n", stack)

	// Калькулятор додавання
	int1 := 19
	int2 := 5
	fmt.Println(calculator(int1, int2))
}

package main

import (
	"Users/user/Documents/Sites/projector/go-course/homeworks/hw-12/internal"
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	pm := internal.NewPasswordManager()
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Println("Please enter one of the following commands:")
		fmt.Println("- 'add [name] [password]' to add a new password")
		fmt.Println("- 'get [name]' to retrieve a password")
		fmt.Println("- 'list' to list all stored names")
		fmt.Print("Enter command: ")
		if !scanner.Scan() {
			break
		}
		input := scanner.Text()
		parts := strings.SplitN(input, " ", 2)
		if len(parts) != 2 {
			fmt.Println("Invalid input. Please enter two values separated by a space.")
			break
		}
		command := parts[0]
		switch command {
		case "add":
			nameAndPassword := strings.SplitN(parts[1], " ", 2)
			pm.AddPassword(nameAndPassword[0], nameAndPassword[1])
		case "get":
			password, ok := pm.GetPassword(parts[1])
			if ok {
				fmt.Println(parts[1], ":", password)
			} else {
				fmt.Println("not found")
			}
		case "list":
			for name := range pm.ListPassword() {
				fmt.Println(name)
			}
		}
	}
}

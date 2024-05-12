package main

import (
	"fmt"
	"time"
)

func main() {
	var startGame string
	fmt.Println("Welcome to my mini-game. Write a `start` to start passing it!")
	startGame = Input()
	for {
		switch startGame {
		case "start":
			fmt.Println("Loading...")
			time.Sleep(2 * time.Second)
			chooseCharacter()
			return
		default:
			fmt.Println("The game is not running. Good luck :D")
		}
	}
}

func Input() string {
	var text string
	fmt.Scanln(&text)
	return text
}

func SwitchHandler(c1 string, c2 string, c1f func(), c2f func()) {
	var text string
	text = Input()
	for {
		switch text {
		case c1:
			c1f()
			return
		case c2:
			c2f()
			return
		default:
			fmt.Println("I don't know this command. Please try again.")
			text = Input()
		}
	}
}

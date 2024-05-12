package transports

import "fmt"

type Train struct{}

func (t *Train) AcceptPassengers() {
	fmt.Println("Пасажирів залізають в потяг")
}

func (t *Train) DisembarkPassengers() {
	fmt.Println("Пасажирів висаджують з потягу")
}

package transports

import "fmt"

type Bus struct{}

func (b *Bus) AcceptPassengers() {
	fmt.Println("Пасажирів залізають в автобус")
}

func (b *Bus) DisembarkPassengers() {
	fmt.Println("Пасажирів висаджують з автобусу")
}

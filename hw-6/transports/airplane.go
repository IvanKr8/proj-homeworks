package transports

import "fmt"

type Airplane struct{}

func (a *Airplane) AcceptPassengers() {
	fmt.Println("Пасажирів залізають в літак")
}

func (a *Airplane) DisembarkPassengers() {
	fmt.Println("Пасажирів висаджують з літаку")
}

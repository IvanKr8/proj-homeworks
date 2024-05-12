package passenger

import (
	"Users/user/Documents/Sites/projector/go-course/lesson6/hw-6/traver"
	"fmt"
)

type Passenger struct{}

func (p *Passenger) Travel(route *traver.Route) {
	fmt.Println("Подорож починається!")
	route.ShowTransport()
	for _, t := range route.Transport {
		t.AcceptPassengers()
		t.DisembarkPassengers()
	}
	fmt.Println("Подорож закінчена!")
}

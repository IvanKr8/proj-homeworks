package main

import (
	passengerPackage "Users/user/Documents/Sites/projector/go-course/lesson6/hw-6/passenger"
	"Users/user/Documents/Sites/projector/go-course/lesson6/hw-6/transports"
	"Users/user/Documents/Sites/projector/go-course/lesson6/hw-6/traver"
)

func main() {
	bus := &transports.Bus{}
	train := &transports.Train{}
	airplane := &transports.Airplane{}

	route := &traver.Route{}
	route.AddTransport(bus)
	route.AddTransport(train)
	route.AddTransport(airplane)
	passenger := &passengerPackage.Passenger{}
	passenger.Travel(route)
}

package traver

import "fmt"

type Route struct {
	Transport []Transport
}

func (r *Route) AddTransport(Transport Transport) {
	r.Transport = append(r.Transport, Transport)
}

func (r *Route) ShowTransport() {
	fmt.Println("Транспорт на маршруті:")
	for _, t := range r.Transport {
		fmt.Printf("%t\n", t)
	}
}

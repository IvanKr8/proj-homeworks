package main

import (
	"Users/user/Documents/Sites/projector/go-course/homeworks/hw-14/internal/api"
	"Users/user/Documents/Sites/projector/go-course/homeworks/hw-14/internal/data/connect"
	"flag"
	"fmt"
	"net/http"
)

var (
	port = 8080
)

func main() {
	portFlag := flag.Int("port", port, "A port to run the service on")
	flag.Parse()
	dbConn := connect.DataBaseConnect()

	app := api.App{}
	router := app.Route(dbConn)

	if portFlag != nil {
		fmt.Println("PORT -> ", *portFlag)
	}
	err := http.ListenAndServe(fmt.Sprintf(":%d", *portFlag), router)
	if err != nil {
		fmt.Print("Error", err)
	}
}

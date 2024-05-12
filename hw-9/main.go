package main

import (
	"Users/user/Documents/Sites/projector/go-course/homeworks/hw-9/pages"
	"Users/user/Documents/Sites/projector/go-course/homeworks/hw-9/pages/login"
	"fmt"
	"log"
	"net/http"
)

const serverPort = 5000

func main() {

	http.HandleFunc("/", pages.MainPage)
	http.HandleFunc("/class-info", login.Authorize(pages.ClassInfo))
	http.HandleFunc("/children/", login.Authorize(pages.ChildrenPage))
	http.HandleFunc("/login", login.LoginPage)
	log.Println("Starting server on port ->", serverPort)

	err := http.ListenAndServe(fmt.Sprintf(":%d", serverPort), nil)
	log.Fatal(err)

}

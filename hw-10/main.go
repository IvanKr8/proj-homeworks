package main

import (
	"Users/user/Documents/Sites/projector/go-course/homeworks/hw-10/client"
	"github.com/gorilla/mux"
	"net/http"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/api/tasks", client.GetTasks).Methods(http.MethodGet)
	r.HandleFunc("/api/tasks", client.CreateTask).Methods(http.MethodPost)
	r.HandleFunc("/api/tasks/{id}", client.GetTask).Methods(http.MethodGet)
	r.HandleFunc("/api/tasks/{id}", client.EditTask).Methods(http.MethodPut)
	r.HandleFunc("/api/tasks/{id}", client.DeleteTask).Methods(http.MethodDelete)
	http.Handle("/", r)
	http.ListenAndServe(":8080", nil)
}

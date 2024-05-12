package client

import (
	"net/http"
)

func CreateTask(w http.ResponseWriter, r *http.Request) {

	title := "test"
	description := "test description"
	err := AddTask(title, description)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

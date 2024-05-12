package client

import (
	"Users/user/Documents/Sites/projector/go-course/homeworks/hw-10/jsonutil"
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

func DeleteTask(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	vars := mux.Vars(r)
	idStr := vars["id"]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	taskList, err := jsonutil.OpenJSONFile("tasks.json")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var task Tasks
	json.Unmarshal(taskList, &task)
	for _, val := range task.Tasks {
		if val.Id == id {
			e := RemoveTask(id)
			if e != nil {
				http.Error(w, e.Error(), http.StatusInternalServerError)
			}
		}
	}
}

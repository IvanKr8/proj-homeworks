package client

import (
	"Users/user/Documents/Sites/projector/go-course/homeworks/hw-10/jsonutil"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

func EditTask(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	idStr := r.URL.Query().Get("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		fmt.Println(err)
		return
	}
	taskList, err := jsonutil.OpenJSONFile("tasks.json")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	var status bool
	var task Tasks
	json.Unmarshal(taskList, &task)
	for _, val := range task.Tasks {
		if val.Id == id {
			status = val.Status
		}
	}

	newStatus := !bool(status)
	statusErr := UpdateTaskStatus(id, newStatus)
	if statusErr != nil {
		http.Error(w, statusErr.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(task)
}

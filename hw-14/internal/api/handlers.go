package api

import (
	"Users/user/Documents/Sites/projector/go-course/homeworks/hw-14/internal/entities"
	"Users/user/Documents/Sites/projector/go-course/homeworks/hw-14/internal/services"
	"Users/user/Documents/Sites/projector/go-course/homeworks/hw-14/internal/storage"
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"strconv"
	"time"
)

// GetTours - get all tours
func (app *App) GetTours(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	tourList, err := storage.GetToursList(app.db)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	writeResponse(w, tourList)
}

// GetTour - get definite tour
func (app *App) GetTour(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	tour, err := storage.GetTour(app.db, id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	writeResponse(w, tour)
}

// GetTourStatus - get all tours with any way status
func (app *App) GetTourStatus(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	timeNow := time.Now()

	status := r.URL.Query().Get("status")

	tourList, err := storage.GetToursList(app.db)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var filteredTours []entities.Tour

	for _, value := range tourList {
		startDate, _ := time.Parse("2006-01-02", value.StartDate)
		endDate, _ := time.Parse("2006-01-02", value.EndDate)

		tourStatus := services.GetTourStatus(timeNow, startDate, endDate)
		if status == tourStatus {
			value.Status = tourStatus
			filteredTours = append(filteredTours, value)
		}
	}
	writeResponse(w, entities.Tours{Tours: filteredTours})
}

// CreateTour - create a tour
func (app *App) CreateTour(w http.ResponseWriter, r *http.Request) {
	tourList, err := storage.GetToursList(app.db)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var tours entities.Tours
	tours.Tours = tourList

	newTask := entities.Tour{
		Id:          len(tours.Tours) + 1,
		Title:       "title",
		Description: "description",
		Price:       1500,
		Transport:   "Літак",
		StartDate:   "2024-08-15",
		EndDate:     "2024-09-15",
	}
	err = storage.SaveNewTour(app.db, newTask)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	writeResponse(w, newTask)
}

// DeleteTour - remove definite tour
func (app *App) DeleteTour(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = storage.RemoveTour(app.db, id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (app *App) BuyingTour(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	NewOrder := entities.Order{
		Name:  "test",
		Email: "test@test.com",
	}
	err = services.BuyTour(app.db, id, NewOrder)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func writeResponse(w http.ResponseWriter, data any) {
	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(data)
	if err != nil {
		log.Printf("writing response: %s", err)
	}
}

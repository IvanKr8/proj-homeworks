package api

import (
	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
	"net/http"
)

// App storage
type App struct {
	db *sqlx.DB
}

func (app *App) Route(DbConn *sqlx.DB) *mux.Router {
	r := mux.NewRouter()
	a := App{db: DbConn}

	r.HandleFunc("/api/tours", a.GetTourStatus).Methods(http.MethodGet).Queries("status", "{status}") // (api/tours?status={status})
	r.HandleFunc("/api/tours", a.GetTours).Methods(http.MethodGet)
	r.HandleFunc("/api/tour/{id}", a.GetTour).Methods(http.MethodGet)
	r.HandleFunc("/api/tour", a.CreateTour).Methods(http.MethodPost)
	r.HandleFunc("/api/tour/{id}", a.DeleteTour).Methods(http.MethodDelete)

	r.HandleFunc("/api/tour/{id}/buy-process", a.BuyingTour).Methods(http.MethodPost)

	return r
}

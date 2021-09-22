package main

import (
	"github.com/gorilla/mux"
)

func (app *application) routes() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/api/trips/", app.findTrip).Methods("GET")
	r.HandleFunc("/api/trips/{uuid}", app.requestTrip).Methods("POST")
	return r
}

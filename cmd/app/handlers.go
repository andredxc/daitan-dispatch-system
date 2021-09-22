package main

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

func (app *application) requestTrip(w http.ResponseWriter, r *http.Request) {

	var u TripRequest
	err := json.NewDecoder(r.Body).Decode(&u)
	trip, err := u.RequestTrip()
	if err != nil {
		app.serverError(w, err)
	}

	body, err := json.Marshal(trip)
	if err != nil {
		app.serverError(w, err)
	}

	app.info.Println("Trip created")

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(body)
}

func (app *application) findTrip(w http.ResponseWriter, r *http.Request) {

	var u Trip
	vars := mux.Vars(r)
	uuid := vars["uuid"]
	trip, err := u.FindTrip(uuid)
	if err != nil {
		app.serverError(w, err)
	}
	err = json.NewDecoder(r.Body).Decode(&u)
	body, err := json.Marshal(trip)
	if err != nil {
		app.serverError(w, err)
	}

	app.info.Println("Trip was found")

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(body)
}

package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"daitan-dispatch-system/cmd/app/models"
	"daitan-dispatch-system/cmd/app/services"

	"github.com/gorilla/mux"
)

func (app *application) requestTrip(w http.ResponseWriter, r *http.Request) {

	var u *models.TripRequest
	// TODO: Decode does not return errror when the parameters in the json do not correspond to the TripRequest struct
	err := json.NewDecoder(r.Body).Decode(&u)

	err = services.GetInstance().NewTripRequest(u)
	if err != nil {
		app.serverError(w, err)
	}

	// body, err := json.Marshal(trip)
	// if err != nil {
	// 	app.serverError(w, err)
	// }

	// app.info.Println("Trip created")

	// w.Header().Set("Content-Type", "application/json")
	// w.WriteHeader(http.StatusCreated)
	// w.Write(body)
}

func (app *application) findTrip(w http.ResponseWriter, r *http.Request) {

	var u models.Trip
	vars := mux.Vars(r)
	uuid := vars["uuid"]
	fmt.Printf("[findTrip] Read uuid=%s\n", uuid)
	trip, err := services.GetInstance().FindTrip(uuid)
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

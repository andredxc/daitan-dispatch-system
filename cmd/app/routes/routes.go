package routes

import (
	"fmt"
	"net/http"
	"strconv"
)

// func (app *application) routes() *mux.Router {
// 	r := mux.NewRouter()
// 	r.HandleFunc("/api/trips/", app.findTrip).Methods("GET")
// 	r.HandleFunc("/api/trips/{uuid}", app.requestTrip).Methods("POST")
// 	return r
// }

func RequestTrip(w http.ResponseWriter, req *http.Request) {

	var lat, lon float64
	var err error
	const LAT_PARAM = "lat"
	const LON_PARAM = "lon"

	keys := req.URL.Query()

	param := keys.Get(LAT_PARAM)
	if lat, err = strconv.ParseFloat(param, 64); err != nil {
		fmt.Printf("Error converting lat=%s to float\n", param)
		return
	}

	param = keys.Get(LON_PARAM)
	if lon, err = strconv.ParseFloat(param, 64); err != nil {
		fmt.Printf("Error converting lon=%s to float\n", param)
		return
	}

	fmt.Printf("Parameters read from the URL, latitude=%.3f; longitude=%.3f\n", lat, lon)
}

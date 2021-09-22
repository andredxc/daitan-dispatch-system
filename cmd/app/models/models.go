package models

type Status int

const (
	Received = 1
	OnTrip
	TripFinished
	Cancel
)

type Location struct {
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

type Car struct {
	PlateNumber string `json:"plateNumber"`
	Brand       string `json:"brand"`
	Model       string `json:"model"`
}

type Driver struct {
	Uuid    string  `json:"uuid"`
	Name    string  `json:"name"`
	Ranking float64 `json:"ranking"`
	Trips   int     `json:"trips"`
	Car     Car
}

type TripRequest struct {
	Datetime string `json:"datetime"`
	Location Location
	Uuid     string `json:"uuid"`
	Status   Status
}

type Trip struct {
	Location Location
	Uuid     string `json:"uuid"`
	Status   Status
	Driver   Driver
}

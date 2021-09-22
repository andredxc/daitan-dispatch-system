package main

import (
	"github.com/google/uuid"
	"sync"
)

type service struct {
}

var instance *service
var once sync.Once

func GetInstance() *service {
	once.Do(func() {
		instance = &service{}
	})

	return instance
}

func (t *TripRequest) RequestTrip() (Trip, error) {

	_uuid := uuid.New().String()
	driver, _ := t.Location.FindDriver()
	location := Location{
		Latitude:  -15.838399,
		Longitude: -48.010947,
	}

	return Trip{
		Uuid:     _uuid,
		Driver:   driver,
		Status:   2,
		Location: location,
	}, nil

}

func (trip *Trip) FindTrip(uuid string) (Trip, error) {

	driver := Driver{
		Name:    "Agatha Ayla",
		Ranking: 4.9,
		Trips:   1212,
		Car: Car{
			PlateNumber: "GJX-2343",
			Brand:       "Lamborghini",
			Model:       "aventador-svj",
		},
	}

	location := Location{
		Latitude:  -15.838399,
		Longitude: -48.010947,
	}

	return Trip{
		Uuid:     uuid,
		Driver:   driver,
		Status:   2,
		Location: location,
	}, nil
}

func (t *Location) FindDriver() (Driver, error) {

	return Driver{
		Name:    "Agatha Ayla",
		Ranking: 4.9,
		Trips:   1212,
		Car: Car{
			PlateNumber: "GJX-2343",
			Brand:       "Lamborghini",
			Model:       "aventador-svj",
		},
	}, nil
}

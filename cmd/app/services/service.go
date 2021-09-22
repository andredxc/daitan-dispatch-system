package services

import (
	"daitan-dispatch-system/cmd/app/models"
	"github.com/google/uuid"
	"sync"
)

type Service struct {
	models.TripRequest
	models.Trip
	models.Location
}

var instance *Service
var once sync.Once

func GetInstance() *Service {
	once.Do(func() {
		instance = &Service{}
	})

	return instance
}

func (t *Service) RequestTrip(request *models.TripRequest) (models.Trip, error) {
	userLocation := request.Location
	_uuid := uuid.New().String()
	driver, _ := t.FindDriver(&userLocation)
	location := models.Location{
		Latitude:  -15.838399,
		Longitude: -48.010947,
	}

	return models.Trip{
		Uuid:     _uuid,
		Driver:   driver,
		Status:   2,
		Location: location,
	}, nil

}

func (trip *Service) FindTrip(uuid string) (models.Trip, error) {

	driver := models.Driver{
		Name:    "Agatha Ayla",
		Ranking: 4.9,
		Trips:   1212,
		Car: models.Car{
			PlateNumber: "GJX-2343",
			Brand:       "Lamborghini",
			Model:       "aventador-svj",
		},
	}

	location := models.Location{
		Latitude:  -15.838399,
		Longitude: -48.010947,
	}

	return models.Trip{
		Uuid:     uuid,
		Driver:   driver,
		Status:   2,
		Location: location,
	}, nil
}

func (t *Service) FindDriver(location *models.Location) (models.Driver, error) {

	if location == nil {

	}

	return models.Driver{
		Uuid:    "ad716768-3a58-4a5a-aa40-cccda2d5b718",
		Name:    "Agatha Ayla",
		Ranking: 4.9,
		Trips:   1212,
		Car: models.Car{
			PlateNumber: "GJX-2343",
			Brand:       "Lamborghini",
			Model:       "aventador-svj",
		},
	}, nil
}
